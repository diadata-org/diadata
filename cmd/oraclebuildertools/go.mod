module oraclebuildertools

go 1.19

require (
	github.com/diadata-org/diadata v1.4.291-rc2
	github.com/sirupsen/logrus v1.9.3
)

replace (
	github.com/btcsuite/btcd/chaincfg/chainhash => github.com/btcsuite/btcd/chaincfg/chainhash v0.23.4
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/tendermint v0.34.28 => github.com/tendermint/tendermint v0.34.24
	github.com/diadata-org/diadata => /home/eric/GolandProjects/Dia/diadata
)