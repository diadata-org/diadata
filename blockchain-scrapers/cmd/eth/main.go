package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/diadata-org/api-golang/dia"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
)

//FYI: docker run --rm -it -v /tmp/eth:/root/.ethereum -p 8545:8545 -p 30303:30303 ethereum/client-go:alpine --rpc --rpcaddr "0.0.0.0" --syncmode "light" --cache 1024 --rpc

var (
	//catchupTimelimit defines the time-limit for initially catching up to the latest block on start-up
	catchupTimelimit = time.Minute * 10
	//supplyTimelimit defines the time-limit for each loop
	supplyTimelimit  = time.Minute
	ethconf          = params.MainnetChainConfig
	big8             = big.NewInt(8)
	big32            = big.NewInt(32)
	premine          = big.NewInt(0)
	checkpointBlock  = big.NewInt(0)
	checkpointSupply = big.NewInt(0)
	one              = big.NewInt(1)
	ninenine         = big.NewInt(99)
	timer            = time.Tick(time.Second * 10)
	rpcEndpoint      = flag.String("rpc", "http://geth:8545", "geth RPC endpoint")
	dev              = flag.Bool("dev", false, "Dev mode - prints to stdout instead of sending to dia")
)

func init() {
	flag.Parse()
	//Premined ETH during genesis
	premine.SetString("72009990499480000000000000", 10)
	//Checkpoints - Hardcoded - Starting point, even if DB/persistant layer is unavailable
	//Might be good idea to update these while performing mantainence on this code.
	checkpointBlock = big.NewInt(6228502)
	checkpointSupply.SetString("29595286937500000000000000", 10)
	// Override the above from some database
	blockStorage, blockSupply, err := fetchSupplyFromStorage()
	if err != nil {
		log.Printf("Error fetching from storage: %v\n", err)
	} else if blockStorage.Cmp(checkpointBlock) > 0 {
		//The block from storage newer than hard-coded value
		checkpointBlock.Set(blockStorage)
		checkpointSupply.Set(blockSupply)
	} else {
		log.Printf("Storage provided older(%v) block than hardcoded(%v)\n", blockStorage, checkpointBlock)
	}
}

//fetchSupplyFromStorage fetches the last stored supply from some sort of persistant storage
func fetchSupplyFromStorage() (*big.Int, *big.Int, error) {
	//TODO: Implement fetching from persistant database
	//IMPORTANT: The block and supply should come from a single checkpoint and MUST be saved atomically
	return nil, nil, fmt.Errorf("TODO: Not Implemented")
}

//storeSupplyToStorage stores the checkpoint into persistant storage so that subsiquent application startup will be fast
func storeSupplyToStorage(block, supply *big.Int) error {
	//TODO: Implement storing to persistant database
	//IMPORTANT: The underlying database MUST store the block and supply atomically, and not one by one.
	//Store them in single field/document for no-sql. Store then in single transaction for SQL.
	return fmt.Errorf("TODO: Not Implemented")
}

//Basically prefetch nth block forward than the one we are currently fetching.
//Ignoring any error, low timeout, in the hopes that when we eventually request it, geth has it in cache
func prefetch(ch chan int64, n int64) {
	cl, err := ethclient.Dial(*rpcEndpoint)
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}
	defer cl.Close()
	for blkn := range ch {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		cl.BlockByNumber(ctx, big.NewInt(blkn+n))
		cancel()
	}
}

//getcheckpoint returns the next blocknumber whose supply has not yet been accounted for
//If returned supply is 10,000 ETH at block 10,000 then this function should return 10001, 10000000000000000000000
func getcheckpoint() (blockNum, supply *big.Int) {
	supply = big.NewInt(0)
	blockNum = big.NewInt(0)
	supply.Set(checkpointSupply)
	blockNum.Set(checkpointBlock)
	return
}

func setcheckpoint(blockNum, supply *big.Int) {
	log.Println("Checkpoint: ", blockNum, supply)
	checkpointSupply.Set(supply)
	checkpointBlock.Set(blockNum)
	//Send to some sort of database, async
	go func(b, s *big.Int) {
		err := storeSupplyToStorage(b, s)
		//Ignoring error here because inability to store a checkpoint should not crash the application.
		if err != nil {
			log.Printf("Error storing: %v\n", err)
		}
	}(new(big.Int).Set(blockNum), new(big.Int).Set(supply))
}

//getblockReward returns total reward issued by a given block.
// Basically all new ETH generated as a result of this block being added into the blockchain
//https://github.com/ethereum/go-ethereum/blob/9bf6bb8f6388f0f8a6fc5603bb4034a8dc75395c/consensus/ethash/consensus.go#L564
//Per block reward calculation taken from upstream go-ethereum
func getblockReward(block *types.Block) *big.Int {
	totalReward := big.NewInt(0)
	blockReward := ethash.FrontierBlockReward
	if ethconf.IsByzantium(block.Header().Number) {
		blockReward = ethash.ByzantiumBlockReward
	}
	reward := new(big.Int).Set(blockReward)
	r := new(big.Int)
	for _, uncle := range block.Uncles() {
		r.Add(uncle.Number, big8)
		r.Sub(r, block.Header().Number)
		r.Mul(r, blockReward)
		r.Div(r, big8)
		//state.AddBalance(uncle.Coinbase, r)
		totalReward.Add(totalReward, r) //Uncle gets r

		r.Div(blockReward, big32)
		reward.Add(reward, r)
	}
	//state.AddBalance(header.Coinbase, reward)
	totalReward.Add(totalReward, reward) //Current blockminer gets rewarded

	return totalReward
}

//getSupply syncs the blockchain to the current block, storing checkpoints along the way
func getSupply(ctx context.Context) (*big.Int, error) {
	ch := make(chan int64)
	for i := 0; i < 10; i++ {
		go prefetch(ch, int64(50))
	}
	cl, err := ethclient.Dial(*rpcEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "error dialing")
	}
	chkblock := big.NewInt(0)
	chksupply := big.NewInt(0)
	block, err := cl.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching header ")
	}
	latestnum := block.Number
	//Restore current and supply from DB or some memory state
	current, supply := getcheckpoint()
	log.Println("Need to fetch: ", latestnum.Int64()-current.Int64())
	//supply := big.NewInt(0)
	//current := big.NewInt(0)
	//go prefetch(current, latestnum)
	log.Println("Fetching blocks")
	start := time.Now()
	count := 0
	for current.Cmp(latestnum) == -1 {
		select {
		case ch <- current.Int64():
		default:
		}

		tmpblk, err := cl.BlockByNumber(ctx, current)
		if err != nil {
			close(ch)
			cl.Close()
			return nil, errors.Wrap(err, fmt.Sprintf("Error fetching block %v", current))
		}
		supply.Add(supply, getblockReward(tmpblk))
		count++
		//log.Println(current, supply)
		//Regularly checkpoint current and supply to DB or some memory state
		mod := big.NewInt(0)
		mod.Mod(current, big.NewInt(100))
		if one.Cmp(mod) == 0 {
			//mod = 1
			chkblock.Add(current, one)
			chksupply.Set(supply)
		}
		if ninenine.Cmp(mod) == 0 {
			//mod = 99
			if chkblock.Cmp(big.NewInt(0)) != 0 {
				setcheckpoint(chkblock, chksupply)
				log.Println("Per block: ", time.Since(start)/time.Duration(count))
				start = time.Now()
				count = 0
			}
		}
		//Increment
		current = current.Add(current, one)
	}
	close(ch)
	log.Println(supply)
	log.Println(current)
	log.Println(supply.Add(supply, premine))
	return supply, nil
}

func weitoeth(wei *big.Int) float64 {
	f, _ := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether)).Float64()
	return f
}

func main() {
	//Init dia
	var config *dia.ConfigApi
	var client *dia.Client
	prev := big.NewInt(0)
	if !*dev {
		config = dia.GetConfigApi()
		if config == nil {
			panic("Couldnt load config")
		}
		client = dia.NewClient(config)
		if client == nil {
			panic("Couldnt load client")
		}
	}

	//Lets catch up to the head first, assuming checkpoint is old
	log.Println("Catching up to head")
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), catchupTimelimit)
	defer cancel()
	supply, err := getSupply(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total supply after initial catch-up is: %f ETH; took %v", weitoeth(supply), time.Since(start))

	//Actual loop
	for {
		log.Println("Fetching supply")
		start := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), supplyTimelimit)
		defer cancel()
		supply, err := getSupply(ctx)
		if err != nil {
			log.Println(err)
		} else {
			eth := weitoeth(supply)
			log.Printf("Total supply is: %f ETH; took %v", eth, time.Since(start))
			if prev.Cmp(supply) == 0 {
				log.Println("Skipping because its same as before")
			} else {
				prev.Set(supply)
				if !*dev {
					client.SendSupply(&dia.Supply{
						Symbol:            "ETH",
						CirculatingSupply: eth,
					})
				}
			}
		}

		<-timer
	}
}
