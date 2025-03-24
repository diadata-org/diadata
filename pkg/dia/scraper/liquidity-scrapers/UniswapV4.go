package liquidityscrapers

import (
	"context"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	uniswapcontractv4 "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswapv4"
	models "github.com/diadata-org/diadata/pkg/model"
	"golang.org/x/crypto/sha3"

	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapV4Scraper struct {
	RestClient      *ethclient.Client
	WsClient        *ethclient.Client
	relDB           *models.RelDB
	datastore       *models.DB
	poolChannel     chan dia.Pool
	doneChannel     chan bool
	blockchain      string
	startBlock      uint64
	factoryContract string
	exchangeName    string
	chunksBlockSize uint64
	waitTime        int
}

const (
	// https://github.com/Uniswap/v4-core/blob/main/src/libraries/StateLibrary.sol
	pools_slot_offset  = uint8(6)
	liquidity_offset   = uint8(3)
	ticks_offset       = uint8(4)
	tick_bitmap_offset = uint8(5)
	univ4_start_block  = uint64(21688329)
)

// NewUniswapV4Scraper returns a new UniswapV4Scraper.
func NewUniswapV4Scraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *UniswapV4Scraper {
	log.Info("NewUniswapV4Scraper ", exchange.Name)
	log.Info("NewUniswapV4Scraper Address ", exchange.Contract)

	var uls *UniswapV4Scraper

	switch exchange.Name {
	case dia.UniswapExchangeV4:
		uls = makeUniswapV4Scraper(exchange, "", "", relDB, datastore, "200", uint64(21688329))
	}
	// err := uls.watchLiquidityChange()
	// if err != nil {
	// 	log.Fatal("watchLiquidityChange: ", err)
	// }

	go func() {
		uls.fetchPools()
	}()
	return uls
}

// makeUniswapV4Scraper returns a uniswap scraper as used in NewUniswapV4Scraper.
func makeUniswapV4Scraper(exchange dia.Exchange, restDial string, wsDial string, relDB *models.RelDB, datastore *models.DB, waitMilliseconds string, startBlock uint64) *UniswapV4Scraper {
	var (
		restClient  *ethclient.Client
		wsClient    *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		uls         *UniswapV4Scraper
	)

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", wsDial))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	var waitTime int
	waitTimeString := utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", waitMilliseconds)
	waitTime, err = strconv.Atoi(waitTimeString)
	if err != nil {
		log.Error("could not parse wait time: ", err)
		waitTime = 500
	}

	uls = &UniswapV4Scraper{
		WsClient:        wsClient,
		RestClient:      restClient,
		relDB:           relDB,
		datastore:       datastore,
		poolChannel:     poolChannel,
		doneChannel:     doneChannel,
		blockchain:      exchange.BlockChain.Name,
		startBlock:      startBlock,
		factoryContract: exchange.Contract,
		exchangeName:    exchange.Name,
		waitTime:        waitTime,
	}
	blockSize := utils.Getenv("CHUNKS_BLOCK_SIZE", "10000")
	uls.chunksBlockSize, err = strconv.ParseUint(blockSize, 10, 64)
	if err != nil {
		log.Error("Parse CHUNKS_BLOCK_SIZE: ", err)
		uls.chunksBlockSize = uint64(10000)
	}

	return uls
}

// fetchPools fetches all registered pools from on-chain and sends them into the pool channel.
func (uls *UniswapV4Scraper) fetchPools() {

	// filter from contract created https://etherscan.io/tx/0x1e20cd6d47d7021ae7e437792823517eeadd835df09dde17ab45afd7a5df4603

	log.Info("get pool creations from address: ", uls.factoryContract)
	poolsCount := 0
	// contract, err := uniswapcontractv4.NewUniswapV4Filterer(common.HexToAddress(uls.factoryContract), uls.WsClient)
	contract, err := uniswapcontractv4.NewPoolmanagerFilterer(common.HexToAddress(uls.factoryContract), uls.WsClient)
	if err != nil {
		log.Error(err)
	}

	// Iterate over chunks of blocks.
	currentBlockNumber, err := uls.RestClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("Get current block number: ", err)
	}
	var startblock, endblock uint64
	startblock = uls.startBlock
	endblock = uls.startBlock + uls.chunksBlockSize

	for endblock < currentBlockNumber+uls.chunksBlockSize {

		time.Sleep(time.Duration(uls.waitTime) * time.Millisecond)

		poolCreated, err := contract.FilterInitialize(
			&bind.FilterOpts{
				Start: startblock,
				End:   &endblock,
			},
			[][32]byte{},
			[]common.Address{},
			[]common.Address{},
		)
		if err != nil {
			log.Error("filter pool created: ", err)
			startblock = endblock
			endblock = startblock + uls.chunksBlockSize
			continue
		}

		for poolCreated.Next() {
			poolsCount++
			var (
				pool   dia.Pool
				asset0 dia.Asset
				asset1 dia.Asset
			)
			log.Info("pools count: ", poolsCount)

			asset0, err = uls.relDB.GetAsset(poolCreated.Event.Currency0.Hex(), uls.blockchain)
			if err != nil {
				asset0, err = ethhelper.ETHAddressToAsset(poolCreated.Event.Currency0, uls.RestClient, uls.blockchain)
				if err != nil {
					log.Warn("cannot fetch asset from address ", poolCreated.Event.Currency0.Hex())
					continue
				}
			}
			asset1, err = uls.relDB.GetAsset(poolCreated.Event.Currency1.Hex(), uls.blockchain)
			if err != nil {
				asset1, err = ethhelper.ETHAddressToAsset(poolCreated.Event.Currency1, uls.RestClient, uls.blockchain)
				if err != nil {
					log.Warn("cannot fetch asset from address ", poolCreated.Event.Currency1.Hex())
					continue
				}
			}

			if !(poolCreated.Event.Currency0 == common.Address{}) && !(poolCreated.Event.Currency1 == common.Address{}) {
				log.Infof("%s-%s : %s -- %s", asset0.Symbol, asset1.Symbol, poolCreated.Event.Currency0.Hex(), poolCreated.Event.Currency1.Hex())
				l, err := uls.getLiquidity(poolCreated.Event.Id)
				if err != nil {
					log.Error("getLiquidity: ", err)
				} else {
					if l.String() != "0" {
						log.Info("liquidity: ", l.String())
					}
				}

				// Testing -----------------------------------------------------------------------------
				tickInfoSlot := getTickInfoSlot(poolCreated.Event.Id, big.NewInt(0))
				log.Info("tickInfoSlot ---------------- : ", "0x"+hex.EncodeToString(tickInfoSlot[:]))

				tickBitmap, err := uls.getTickBitmap(poolCreated.Event.Id, int16(0))
				if err != nil {
					log.Error("getTickBitmap: ", err)
				} else {
					log.Info("tickBitmap: ", tickBitmap.String())
				}

				bitmapSlot := computeBitmapSlot(poolCreated.Event.Id)
				log.Warn("bitmapSlot: ", "0x"+hex.EncodeToString(bitmapSlot[:]))
				// Testing -----------------------------------------------------------------------------

			} else {
				continue
			}

			pool.Exchange = dia.Exchange{Name: uls.exchangeName}
			pool.Blockchain = dia.BlockChain{Name: uls.blockchain}
			pool.Address = hex.EncodeToString(poolCreated.Event.Id[:])

			if !(poolCreated.Event.Hooks == common.Address{}) {
				log.Info("hooks: ", poolCreated.Event.Hooks.Hex())
			}

			balance0Big, err := ethhelper.GetBalanceOf(common.HexToAddress(asset0.Address), common.HexToAddress(pool.Address), uls.RestClient)
			if err != nil {
				// log.Error("GetBalanceOf: ", err)
			}
			balance1Big, err := ethhelper.GetBalanceOf(common.HexToAddress(asset1.Address), common.HexToAddress(pool.Address), uls.RestClient)
			if err != nil {
				// log.Error("GetBalanceOf: ", err)
			}
			balance0, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance0Big), new(big.Float).SetFloat64(math.Pow10(int(asset0.Decimals)))).Float64()
			balance1, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance1Big), new(big.Float).SetFloat64(math.Pow10(int(asset1.Decimals)))).Float64()

			pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{Asset: asset0, Volume: balance0, Index: uint8(0)})
			pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{Asset: asset1, Volume: balance1, Index: uint8(1)})

			// Determine USD liquidity
			if balance0 > GLOBAL_NATIVE_LIQUIDITY_THRESHOLD && balance1 > GLOBAL_NATIVE_LIQUIDITY_THRESHOLD {
				uls.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
			}

			pool.Time = time.Now()

			uls.poolChannel <- pool

		}
		startblock = endblock
		endblock = startblock + uls.chunksBlockSize

	}
	uls.doneChannel <- true
}

func (uls *UniswapV4Scraper) watchLiquidityChange() error {
	filterer, errFilterer := uniswapcontractv4.NewPoolmanagerFilterer(common.HexToAddress(uls.factoryContract), uls.RestClient)
	if errFilterer != nil {
		log.Error("NewPoolmanagerFilterer: ", errFilterer)
	}

	startBlock := univ4_start_block
	endBlock := startBlock + 5000
	currentBlock, err := uls.RestClient.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	for endBlock < currentBlock {
		log.Infof("startBlock -- endBlock: %v --- %v", startBlock, endBlock)
		liquidityModification, err := filterer.FilterModifyLiquidity(
			&bind.FilterOpts{
				Start: startBlock,
				End:   &endBlock,
			},
			[][32]byte{},
			nil,
		)
		if err != nil {
			return err
		}

		for liquidityModification.Next() {
			log.Info("ID: ", hex.EncodeToString(liquidityModification.Event.Id[:]))
			log.Info("liquidity delta: ", liquidityModification.Event.LiquidityDelta.String())
		}
		startBlock = endBlock
		endBlock = startBlock + 5000
	}
	return nil
}

func (uls *UniswapV4Scraper) getLiquidity(poolID [32]byte) (*big.Int, error) {
	// Compute storage slot (Example: Use keccak256(poolID) to determine slot)
	// This is a placeholder. Replace it with the correct formula once confirmed.
	l, e := uniswapcontractv4.NewPoolmanagerCaller(common.HexToAddress(uls.factoryContract), uls.RestClient)
	if e != nil {
		log.Error("NewPoolmanagerCaller: ", e)
	}

	stateSlot := getPoolStateSlot(poolID)
	stateSlotInt := new(big.Int).SetBytes(stateSlot[:])
	// Add liquidity offset.
	liquiditySlotInt := new(big.Int).Add(stateSlotInt, big.NewInt(int64(liquidity_offset)))
	// Convert back to bytes array.
	liquiditySlot := bigInt2Byte(liquiditySlotInt)
	log.Info("liquiditySlot: ", "0x"+hex.EncodeToString(liquiditySlot[:]))

	resultExtsload, err := l.Extsload(&bind.CallOpts{}, liquiditySlot)
	if err != nil {
		return big.NewInt(0), err
	}

	// Convert result to a big integer (liquidity amount)
	liquidity := new(big.Int).SetBytes(resultExtsload[:])
	return liquidity, nil
}

// Retrieves the tick bitmap from the Uniswap V4 pool
func (uls *UniswapV4Scraper) getTickBitmap(poolID [32]byte, tick int16) (*big.Int, error) {

	l, e := uniswapcontractv4.NewPoolmanagerCaller(common.HexToAddress(uls.factoryContract), uls.RestClient)
	if e != nil {
		log.Error("NewPoolmanagerCaller: ", e)
	}

	tickBitmapMapping := computeBitmapSlot(poolID)
	slot := computeTickBitmapSlot(tick, tickBitmapMapping)
	log.Info("tickBitmapSlot: ", "0x"+hex.EncodeToString(slot[:]))

	// Call the extsload function to retrieve tickBitmap
	resultExtsload, err := l.Extsload(&bind.CallOpts{}, slot)
	if err != nil {
		return big.NewInt(0), err
	}

	// Convert the result to a uint256
	tickBitmap := new(big.Int).SetBytes(resultExtsload[:])

	return tickBitmap, nil
}

// GetActiveTicks retrieves all active ticks for a given pool.
func (uls *UniswapV4Scraper) getActiveTicks(poolID [32]byte) ([]int, error) {
	activeTicks := []int{}

	poolManager, e := uniswapcontractv4.NewPoolmanagerCaller(common.HexToAddress(uls.factoryContract), uls.RestClient)
	if e != nil {
		log.Error("NewPoolmanagerCaller: ", e)
	}

	// Get the storage slot for the tick bitmap mapping
	bitmapSlot := computeBitmapSlot(poolID)

	// Iterate over possible tick bitmap words (Uniswap typically has a range from -887272 to 887272)
	for i := -3500; i < 3500; i++ { // Adjust this range based on expected tick indexes
		slotHash := keccak256Hash(big.NewInt(int64(i)), bitmapSlot)

		// Query the storage slot using Extsload
		bitmap, err := poolManager.Extsload(nil, slotHash)
		if err != nil {
			return nil, err
		}

		// Convert the [32]byte response into a big integer
		bitmapValue := new(big.Int).SetBytes(bitmap[:])

		// Iterate over each bit in the 256-bit bitmap
		for j := 0; j < 256; j++ {
			if bitmapValue.Bit(j) == 1 {
				// Calculate the actual tick value
				tick := (i * 256) + j
				activeTicks = append(activeTicks, tick)
			}
		}
	}

	return activeTicks, nil
}

// -------------------------------------------
// Helper functions for contract interaction
// -------------------------------------------

func getPoolStateSlot(poolID [32]byte) [32]byte {
	return shiftBytesArray(poolID, pools_slot_offset)
}

func getTickInfoSlot(poolID [32]byte, tick *big.Int) [32]byte {

	// Get slot for tick Bitmap.
	tickBitmapSlot := computeBitmapSlot(poolID)

	// Convert tick to [32]byte.
	tickBytes := make([]byte, 32)
	copy(tickBytes[32-len(tick.Bytes()):], tick.Bytes())

	// Concatenate tickBytes and tickBitmapSlot.
	encoded := append(tickBytes, tickBitmapSlot[:]...)

	// Compute Keccak-256 hash
	hash := sha3.NewLegacyKeccak256()
	hash.Write(encoded)

	// Convert tickInfoSlot to [32]byte
	var tickInfoSlot [32]byte
	copy(tickInfoSlot[:], hash.Sum(nil))

	return tickInfoSlot

}

func computeBitmapSlot(poolID [32]byte) [32]byte {
	stateSlot := getPoolStateSlot(poolID)
	stateSlotInt := new(big.Int).SetBytes(stateSlot[:])
	tickBitmapSlotInt := new(big.Int).Add(stateSlotInt, big.NewInt(int64(tick_bitmap_offset)))
	return bigInt2Byte(tickBitmapSlotInt)
}

// Computes keccak256(abi.encodePacked(int256(tick), tickBitmapMapping))
func computeTickBitmapSlot(tick int16, tickBitmapMapping [32]byte) [32]byte {

	// Encode int256(tick) into 32 bytes (big-endian)
	tickBigInt := big.NewInt(int64(tick))
	tickBytes := make([]byte, 32)
	copy(tickBytes[32-len(tickBigInt.Bytes()):], tickBigInt.Bytes()) // Right-pad to 32 bytes

	// Concatenate tickBytes and tickBitmapMapping
	encoded := append(tickBytes, tickBitmapMapping[:]...)

	// Compute Keccak-256 hash
	hash := sha3.NewLegacyKeccak256()
	hash.Write(encoded)
	hashed := hash.Sum(nil)

	// Convert result to [32]byte
	var result [32]byte
	copy(result[:], hashed)

	return result
}

func shiftBytesArray(array [32]byte, offset uint8) [32]byte {

	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(array[:])

	shiftedArray := [32]byte{}
	shiftedArray[31] = offset
	hasher.Write(shiftedArray[:])

	var slot [32]byte
	copy(slot[:], hasher.Sum(nil))

	return slot
}

// Decode uint128 (first 16 bytes) from bytes32
func decodeLiquidity(data [32]byte) *big.Int {
	liquidityBytes := data[:]
	liquidity := new(big.Int).SetBytes(liquidityBytes)
	return liquidity
}

func bigInt2Byte(i *big.Int) [32]byte {
	var slot [32]byte
	slotBytes := i.Bytes()
	copy(slot[32-len(slotBytes):], slotBytes)
	return slot
}

// Keccak256Hash simulates abi.encodePacked and keccak256 hash in Solidity:
//
//	keccak256(abi.encodePacked(int256(value), slot))
func keccak256Hash(value *big.Int, slot [32]byte) [32]byte {
	hasher := sha3.NewLegacyKeccak256()
	valueBytes := value.Bytes()
	hasher.Write(make([]byte, 32-len(valueBytes))) // Pad to 32 bytes
	hasher.Write(valueBytes)
	hasher.Write(slot[:])

	var result [32]byte
	copy(result[:], hasher.Sum(nil))
	return result
}

// Helper function to convert slot to hex string
func slotToHex(slot [32]byte) string {
	return fmt.Sprintf("%x", slot[:])
}

func (uls *UniswapV4Scraper) Pool() chan dia.Pool {
	return uls.poolChannel
}

func (uls *UniswapV4Scraper) Done() chan bool {
	return uls.doneChannel
}
