package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	ws "github.com/gorilla/websocket"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-libipfs/files"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/path"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"
	log "github.com/sirupsen/logrus"

	"github.com/ipfs/kubo/config"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	"github.com/ipfs/kubo/core/node/libp2p"
	"github.com/ipfs/kubo/plugin/loader" // This package is needed so that all the preloaded plugins are loaded automatically
	"github.com/ipfs/kubo/repo/fsrepo"
	"github.com/libp2p/go-libp2p/core/peer"
)

var (
	binancewsBaseString = "wss://stream.binance.com:9443/ws/"
	inputBasePath       = "data/"
	inputPathBinance    = inputBasePath + "trades_Binance_BTC_USDT/"
)

type BinanceTrade struct {
	Symbol    string `json:"s"`
	Price     string `json:"p"`
	Volume    string `json:"q"`
	Timestamp int64  `json:"T"`
	TradeID   int    `json:"t"`
}

/// ------ Setting up the IPFS Repo

func setupPlugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

func createTempRepo() (string, error) {
	repoPath, err := os.MkdirTemp("", "ipfs-shell")
	if err != nil {
		return "", fmt.Errorf("failed to get temp dir: %s", err)
	}

	// Create a config with default options and a 2048 bit key
	cfg, err := config.Init(io.Discard, 2048)
	if err != nil {
		return "", err
	}

	// When creating the repository, you can define custom settings on the repository, such as enabling experimental
	// features (See experimental-features.md) or customizing the gateway endpoint.
	// To do such things, you should modify the variable `cfg`. For example:
	if *flagExp {
		// cfg.Ipns.UsePubsub
		// cfg.Pubsub

		// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-filestore
		cfg.Experimental.FilestoreEnabled = true
		// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-urlstore
		cfg.Experimental.UrlstoreEnabled = true
		// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#ipfs-p2p
		cfg.Experimental.Libp2pStreamMounting = true
		// https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md#p2p-http-proxy
		cfg.Experimental.P2pHttpProxy = true
		// See also: https://github.com/ipfs/kubo/blob/master/docs/config.md
		// And: https://github.com/ipfs/kubo/blob/master/docs/experimental-features.md
	}

	// Create the repo with the config
	err = fsrepo.Init(repoPath, cfg)
	if err != nil {
		return "", fmt.Errorf("failed to init ephemeral node: %s", err)
	}

	return repoPath, nil
}

/// ------ Spawning the node

// Creates an IPFS node and returns its coreAPI
func createNode(ctx context.Context, repoPath string) (*core.IpfsNode, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	// Construct the node

	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
		Repo: repo,
	}

	return core.NewNode(ctx, nodeOptions)
}

var loadPluginsOnce sync.Once

// Spawns a node to be used just for this run (i.e. creates a tmp repo)
func spawnEphemeral(ctx context.Context) (icore.CoreAPI, *core.IpfsNode, error) {
	var onceErr error
	loadPluginsOnce.Do(func() {
		onceErr = setupPlugins("")
	})
	if onceErr != nil {
		return nil, nil, onceErr
	}

	// Create a Temporary Repo
	repoPath, err := createTempRepo()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create temp repo: %s", err)
	}

	node, err := createNode(ctx, repoPath)
	if err != nil {
		return nil, nil, err
	}

	api, err := coreapi.NewCoreAPI(node)

	return api, node, err
}

func connectToPeers(ctx context.Context, ipfs icore.CoreAPI, peers []string) error {
	var wg sync.WaitGroup
	peerInfos := make(map[peer.ID]*peer.AddrInfo, len(peers))
	for _, addrStr := range peers {
		addr, err := ma.NewMultiaddr(addrStr)
		if err != nil {
			return err
		}
		pii, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			return err
		}
		pi, ok := peerInfos[pii.ID]
		if !ok {
			pi = &peer.AddrInfo{ID: pii.ID}
			peerInfos[pi.ID] = pi
		}
		pi.Addrs = append(pi.Addrs, pii.Addrs...)
	}

	wg.Add(len(peerInfos))
	for _, peerInfo := range peerInfos {
		go func(peerInfo *peer.AddrInfo) {
			defer wg.Done()
			err := ipfs.Swarm().Connect(ctx, *peerInfo)
			if err != nil {
				log.Printf("failed to connect to %s: %s", peerInfo.ID, err)
			}
		}(peerInfo)
	}
	wg.Wait()
	return nil
}

func getUnixfsNode(path string) (files.Node, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}

/// -------

func init() {

}

var flagExp = flag.Bool("experimental", false, "enable experimental features")

func main() {
	flag.Parse()

	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// chanTrade := make(chan dia.Trade)
	// go scrapeBinance("btcusdt", chanTrade, &wg)
	// defer wg.Wait()

	tradesBlock := createTradesblock(10)
	fileIdentifier := strconv.Itoa(int(tradesBlock.TradesBlockData.BeginTime.Unix())) + "-" + strconv.Itoa(int(tradesBlock.TradesBlockData.EndTime.Unix()))
	tradesBlockMarsh, err := tradesBlock.MarshalBinary()
	if err != nil {
		log.Error("marshal tradesblock")
	}

	filename := inputPathBinance + fileIdentifier + ".json"
	err = os.WriteFile(filename, tradesBlockMarsh, 0666)
	if err != nil {
		log.Error("write file: ", err)
	}
	log.Info("successfully wrote file")

	// time.Sleep(10 * time.Minute)

	/// Getting a IPFS node running

	fmt.Println("-- Getting an IPFS node running -- ")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Spawn a node using a temporary path, creating a temporary repo for the run
	fmt.Println("Spawning Kubo node on a temporary repo")
	ipfsA, _, err := spawnEphemeral(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to spawn ephemeral node: %w", err))
	}

	fmt.Println("IPFS node is running")

	/// Adding a tradesblock to IPFS

	fmt.Println("\n-- Adding and getting back files & directories --")

	// binanceFolder, err := getUnixfsNode(inputPathBinance)
	// if err != nil {
	// 	panic(fmt.Errorf("could not get File: %w", err))
	// }

	tradesblockFile, err := getUnixfsNode(filename)
	if err != nil {
		panic(fmt.Errorf("could not create node for File: %w", err))
	}

	cidTradesblock, err := ipfsA.Unixfs().Add(ctx, tradesblockFile)
	if err != nil {
		panic(fmt.Errorf("could not add File: %w", err))
	}

	fmt.Printf("Added tradesblock to IPFS with CID %s\n", cidTradesblock.String())

	// someDirectory, err := getUnixfsNode(inputBasePath)
	// if err != nil {
	// 	panic(fmt.Errorf("could not get File: %w", err))
	// }

	// cidDirectory, err := ipfsA.Unixfs().Add(ctx, someDirectory)
	// if err != nil {
	// 	panic(fmt.Errorf("could not add Directory: %w", err))
	// }

	// fmt.Printf("Added directory to IPFS with CID %s\n", cidDirectory.String())

	// Getting tradesblock from the IPFS Network

	mHashTBS, err := multihash.FromB58String(strings.Split(cidTradesblock.String(), "/")[2])
	if err != nil {
		log.Error("multihash: ", mHashTBS)
	}
	cidTradesblockRecovered, err := ipfsA.ResolvePath(ctx, path.IpfsPath(cid.NewCidV0(mHashTBS)))
	if err != nil {
		log.Error("resolve path: ", err)
	}

	log.Info("recovered cid: ", cidTradesblockRecovered.String())
	// cidTradesblock = path.NewResolvedPath(path.New("/ipfs/QmUsehujVVSHxJtneWW49fVbAjyJWsDJSnJ5EiXX82VEBi"), cid.NewCidV1())

	outputBasePath, err := os.MkdirTemp("", "tradesblock")
	if err != nil {
		panic(fmt.Errorf("could not create output dir (%w)", err))
	}
	fmt.Printf("output folder: %s\n", outputBasePath)
	outputPathFile := outputBasePath + strings.Split(cidTradesblock.String(), "/")[2]
	// outputPathDirectory := outputBasePath + strings.Split(cidDirectory.String(), "/")[2]

	rootNodeFile, err := ipfsA.Unixfs().Get(ctx, cidTradesblock)
	if err != nil {
		panic(fmt.Errorf("could not get file with CID: %w", err))
	}

	err = files.WriteTo(rootNodeFile, outputPathFile)
	if err != nil {
		panic(fmt.Errorf("could not write out the fetched CID: %w", err))
	}

	fmt.Printf("got file back from IPFS (IPFS path: %s) and wrote it to %s\n", cidTradesblock.String(), outputPathFile)

}

func scrapeBinance(market string, tradesChan chan dia.Trade, wg *sync.WaitGroup) error {
	defer wg.Done()
	log.Info("Entered Binance handler")

	conn, _, err := ws.DefaultDialer.Dial(binancewsBaseString+market+"@trade", nil)
	if err != nil {
		log.Error("dial: ", err)
		return err
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Errorln("read:", err)
			return err
		}
		var bt BinanceTrade
		err = json.Unmarshal(message, &bt)
		if err != nil {
			log.Error("unmarshal binance response: ", err)
		}

		var t dia.Trade
		t.BaseToken = dia.Asset{Symbol: "USDT", Blockchain: dia.ETHEREUM, Address: "0xdAC17F958D2ee523a2206206994597C13D831ec7"}
		t.QuoteToken = dia.Asset{Symbol: "BTC", Blockchain: dia.BITCOIN, Address: "0x0000000000000000000000000000000000000000"}
		t.Pair = market
		t.Price, err = strconv.ParseFloat(bt.Price, 64)
		if err != nil {
			log.Error("parse float: ", err)
		}
		t.Volume, err = strconv.ParseFloat(bt.Volume, 64)
		if err != nil {
			log.Error("parse float: ", err)
		}
		t.Time = time.Unix(0, bt.Timestamp*1e6)
		t.ForeignTradeID = strconv.Itoa(bt.TradeID)
		log.Info("got trade: ", t)
		// tradesChan <- t

	}
}

func createTradesblock(size int) (tradesBlock dia.TradesBlock) {
	for i := 0; i < size; i++ {
		t := dia.Trade{
			Symbol:     "BTC",
			Pair:       "BTC-USDT",
			BaseToken:  dia.Asset{Symbol: "USDT", Blockchain: dia.ETHEREUM, Address: "0xdAC17F958D2ee523a2206206994597C13D831ec7"},
			QuoteToken: dia.Asset{Symbol: "BTC", Blockchain: dia.BITCOIN, Address: "0x0000000000000000000000000000000000000000"},
			Price:      float64(28000) + float64(i/1000),
			Volume:     float64(1),
			Time:       time.Now(),
		}
		tradesBlock.TradesBlockData.Trades = append(tradesBlock.TradesBlockData.Trades, t)
	}
	tradesBlock.TradesBlockData.BeginTime = tradesBlock.TradesBlockData.Trades[0].Time
	tradesBlock.TradesBlockData.EndTime = tradesBlock.TradesBlockData.Trades[len(tradesBlock.TradesBlockData.Trades)-1].Time
	return
}
