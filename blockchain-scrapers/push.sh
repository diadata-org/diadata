#!/bin/bash
#set -x

function usage {
    echo "Usage: ./push.sh [-bdr] <service|all>"
    echo "  -b      builds an image for a service (builds everything when run with 'all'"
    echo "  -d      deploys a service in a stack with its name (deploys every service in a stack named 'all' when run with 'all'"
    echo " 	-r 		removes a service (or every service if run with 'all')"
    echo "  -h      display this usage guide"
    exit 1
}

function deploy {
	case $1 in
		# deploy every service in a stack called 'all'
		all) 		
			find $blockchain_dir -name docker-compose-*.yml -exec docker stack deploy -c {} $1 \;
			;;
		# deploy a service in it's individual stack
		btc|bch|ada|ltc|xmr|eth|erc20|xrp|usdt|trx|xlm|neo) 
			find $blockchain_dir -name docker-compose-$1.yml -exec docker stack deploy -c {} $1 \;
			;;
		*)
			echo "Error: unknown service '"$1"'"
	esac
}

function build {
	case $1 in
		# build every service
		all) 
			find $blockchain_dir -name Dockerfile* -exec sh -c "docker build -f {} -t $DOCKER_HUB_LOGIN/blockchain-scrapers_$(sed -n -e 's/^.*-//p' {}) $GOPATH" \;
			;;
		# build a single service
		btc|bch|ada|ltc|xmr|eth|erc20|xrp|usdt|trx|xlm|neo) 
			find $blockchain_dir -name Dockerfile-$1 -exec sh -c "docker build -f {} -t $DOCKER_HUB_LOGIN/blockchain-scrapers_$1 $GOPATH" \;
			;;
		*)
			echo "Error: unknown service '"$1"'"
	esac
}

function remove {
	case $1 in
		# build every service
		all) 
			find $blockchain_dir -name Dockerfile* -exec docker service rm $(docker service ls -q)
			;;
		# build a single service
		btc|bch|ada|ltc|xmr|eth|erc20|xrp|usdt|trx|xlm|neo) 
			docker stack rm $1
			;;
		*)
			echo "Error: unknown service '"$1"'"
	esac
}

if [[ $@ == *"-h"* ]] 
then
	usage
fi

blockchain_dir=$GOPATH/src/github.com/diadata-org/api-golang/blockchain-scrapers/blockchains
sudo mkdir -p $HOME/srv/bitcoin $HOME/srv/geth $HOME/srv/monero $HOME/srv/litecoin $HOME/srv/cardano $HOME/srv/bitcoin-cash $HOME/srv/neo 
#sudo chown -R $USER:$USER $HOME/srv
sudo chmod -R 777 $HOME/srv

while getopts "d:b:r:" opt; do
	case $opt in
		d)
			deploy $OPTARG
			;;

		b)
			build $OPTARG
			;;

		r)
			remove $OPTARG
			;;
	esac
done