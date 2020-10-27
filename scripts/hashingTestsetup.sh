#!/bin/bash

cd $GOPATH/src/github.com/diadata-org/diadata/cmd

cd services/merkleService
go run main.go

cd ../tradesBlockService
go run main.go

cd ../exchange-scrapers/collector
go run collector.go -exchange=Huobi

cd ../../ratescrapers
go run rate.go

cd ../merkle
go run masterMerkleTree.go