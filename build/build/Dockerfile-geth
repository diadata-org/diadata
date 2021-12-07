FROM ethereum/client-go:stable

EXPOSE 8546 8545

CMD ["ethereum/client-go",'--rpcvhosts geth --rpc --rpcaddr "0.0.0.0" --cache 1024 --rpc --syncmode "fast" --ws --ws.addr "0.0.0.0" --ws.origins "*" --ws.port 8546 --http.api admin,db,eth,debug,miner,net,shh,txpool,personal,web3 --txpool.globalqueue 4096']