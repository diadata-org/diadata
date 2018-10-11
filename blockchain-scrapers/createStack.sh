#!/bin/bash
#set -x

function usage {
    echo "Usage: ./createStack.sh [-bdr] <service|all>"
    echo "  Each flag corresponds to an operation and takes a single argument. Multiple operations can be specified in the same command and will be run in that order. The supported flags/operations are:"
    echo ""
    echo "  -b      builds an image for a service (builds everything when run with 'all')"
    echo "  -d      deploys a service in a stack with its name (deploys every service in a stack named 'all' when run with 'all')"
    echo "  -r      removes a service (or every service if run with 'all')"
    echo "  -h      display this usage guide"
    exit 1
}

function error {
	RED='\033[0;31m'
	NC='\033[0m'
	echo -e "${RED}Error: ${NC}$1" 1>&3
}

function deploy {
	echo "Deploying stack '$1'..." 1>&3

	case $1 in
		# deploy every service in a stack called 'all'
		all)
			find $blockchain_dir -name docker-compose.*.yml -exec docker stack deploy -c {} $1 \;
			echo "Finished deployment" 1>&3
			;;

		# deploy a service in a stack with its name
		*)
			(docker stack deploy -c $(find $blockchain_dir -name docker-compose.$1.yml) $1 &&
				echo "Finished deployment" 1>&3) || error "Can't deploy '$1' (might not exist) "
			;;
	esac
}

function build {
	echo "Building service '$1'..." 1>&3

	case $1 in
		# build every service
		all) 
			find $blockchain_dir -name Dockerfile* -exec sh -c "docker build -f {} -t $DOCKER_HUB_LOGIN/${STACKNAME}_$(sed -n -e 's/^.*-//p' {}):latest $GOPATH" \;
			echo "Finished build" 1>&3			
			;;

		# build a particular service
		*) 
			(docker build -f $(find $blockchain_dir -name Dockerfile-$1) -t "$DOCKER_HUB_LOGIN/${STACKNAME}_$1:latest" $GOPATH && 
				docker push $DOCKER_HUB_LOGIN/${STACKNAME}_$1 && echo "Finished build" 1>&3) || error "Can't build '$1' (might not exist)"
			;;
	esac
}

function remove {
	echo "Removing stack '$1'..." 1>&3

	# remove a stack
	if [[ ! -z $(docker stack ls | grep "$1") ]]; then
		docker stack rm $1
		echo "Finished removal" 1>&3
	else
		error "Stack isn't deployed"
	fi
}

# if one of the commands is help, display only that
if [[ $@ == *"-h"* ]]; then
	usage
	exit 1
fi

# create necessary volumes 
blockchain_dir=$GOPATH/src/github.com/diadata-org/diadata/blockchain-scrapers/blockchains
export STACKNAME=blockchain-scrapers

if [ ! -e $HOME/srv ]; then
	sudo mkdir -p $HOME/srv/bitcoin $HOME/srv/geth $HOME/srv/monero $HOME/srv/litecoin  \
		$HOME/srv/cardano $HOME/srv/bitcoin-cash $HOME/srv/neo $HOME/srv/dash $HOME/srv/dogecoin
	sudo chmod -R 777 $HOME/srv
fi

exec 3>&1 4>&2
# parse input (uncomment these to silence non-script output)
#{

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

			*)
				error "Unknown operation '$opt'"
		esac
	done
#} &> /dev/null