#!/bin/bash

if [  "$DOCKER_HUB_LOGIN" == "" ]
then
	echo missing env DOCKER_HUB_LOGIN variable
	exit
fi

REMOTE="n"
NOCACHE="--no-cache"
LINENUMBER=""
DEPLOY="n"
SLAVES="n"
WATCH="n"
DIRECTORY=`pwd`
MACHINE=""
STACKNAME=""
FILETMP=/tmp/$0.$$.tmp
FILETMP2=/tmp/services.txt.$$

selectLocalHost() {
	echo DOCKER_HOST:$DOCKER_HOST
	unset DOCKER_HOST
	unset DOCKER_TLS_VERIFY
	unset DOCKER_CERT_PATH
	echo switching to LocalHost
	env | grep DOCKER
}

checkMachine() {
	if [ ! -d ~/secrets/live/$1/ ]; then
		echo "no configuration for machine $1 in your ~/secrets directory"
		exit
	fi
}

selectMachine() {
	export DOCKER_HOST="tcp://$1:2376"
	export DOCKER_ID_USER="reg"
	export DOCKER_TLS_VERIFY="1"
	export DOCKER_CERT_PATH=~/secrets/docker-certs/$1/

	H=`echo $1 | cut -d"." -f2-3`

	#export DOCKER_HUB_LOGIN=registry.$H

	if [ -d "$DOCKER_CERT_PATH" ]; then
		rm ~/secrets/live/$1/ca.pem
		rm ~/secrets/live/$1/key.pem
		ln -s ~/secrets/live/$1/fullchain.pem ~/secrets/live/$1/ca.pem 2>&1 > /dev/null
		ln -s ~/secrets/live/$1/privkey.pem ~/secrets/live/$1/key.pem 2>&1 > /dev/null
	else
		echo cant find machine $1
		exit
	fi
}

selectMaster() {
	if [ "$REMOTE" != "n" ]
	then
		selectMachine $MACHINE
		echo switching to $DOCKER_HOST
	fi
}


selectSlave() {
	if [ "$REMOTE" != "n" ]
	then
		export DOCKER_HOST="tcp://TOFIX:2376"
		export DOCKER_ID_USER="TOFIX"
		export DOCKER_TLS_VERIFY="TOFIX"
		export DOCKER_CERT_PATH="TOFIX"
		echo switching to $DOCKER_HOST
	fi
}

watch_service() {
	docker service logs $1 -f
}

rebuild_service() {
	echo building $1 NOCACHE: $NOCACHE

	GITHUB_DIRECTORY=`pwd | sed 's/\.com.*/.com/'`
	echo cd $GITHUB_DIRECTORY
	cd $GITHUB_DIRECTORY
	selectLocalHost
	docker-compose -f $DIRECTORY/$COMPOSE_FILE build $NOCACHE $1

	if [ $? -eq 0 ]
	then

		IMAGE=${DOCKER_HUB_LOGIN}/${STACKNAME}_$1

		echo "$0 pushing $IMAGE"

		docker push $IMAGE
		if [ $? -ne 0 ]
		then
			echo error pushing image, press any key to continue
			read
		fi
		if [ "$REMOTE" != "n" ]
		then
			if [ "$SLAVES" != "n" ]
			then
				selectSlave
				echo $DOCKER_HOST
				docker pull $IMAGE
			fi
			selectMaster
			echo "$0 pulling image"
			docker pull $IMAGE
			if [ $? -ne 0 ]
			then
				echo error pulling image
				exit 0
			fi
		fi

		SERVICE=${STACKNAME}_$1
		echo service update on $DOCKER_HOST for $SERVICE
		RESTARTED=1
		while [ $RESTARTED -ne 0 ]
		do
			docker service update --with-registry-auth --force ${SERVICE} --image ${IMAGE}
			RESTARTED=$?
			if [ $RESTARTED -ne 0 ]
			then
				echo "$0 failed on update service ${SERVICE} trying again in 10 secs"
				echo "run this to get logs:"
				echo "docker service logs -f ${SERVICE}"
				sleep 10
			fi
		done

		if [ "$WATCH" == "y" ]
		then
			watch_service $SERVICE
		fi

	else
		exit
	fi

	cd -
}

clean() {
	docker volume rm $(docker volume ls -qf dangling=true) 2> /dev/null
	docker system prune -f
}

showHelp() {
	echo "-s <stack>" >&2
	ls deployments/docker-compose.*
	echo "-c to use cache" >&2
	echo "-d to deploy stack" >&2
	echo "-r <machine> for remote" >&2
	echo "-n <number> to build service <number> from stack" >&2
	echo "-a to build all services" >&2
}

main() {
	while getopts "n:hr:dcaws:" opt; do
		case $opt in
			s)
				export STACKNAME=$OPTARG
				export COMPOSE_FILE=`find . | grep docker-compose | grep "$STACKNAME\.yml"`
					if [ ! -f "$COMPOSE_FILE" ]
				then
				echo cant find stack $STACKNAME
				exit
				fi
			;;
			w)
				WATCH="y"
			;;
			c)
				echo using cache
				NOCACHE=""
			;;
			d)
				DEPLOY="y"
			;;
			a)
				LINENUMBER="a"
			;;
			n)
				echo "-n was triggered, Parameter: $OPTARG" >&2
				LINENUMBER=$OPTARG
			;;
			h)
				showHelp
				exit
			;;
			r)
				REMOTE="y"
				MACHINE=$OPTARG
				if [ "$MACHINE" == "" ]
				then
					echo you must select a machine
					exit
				fi
				checkMachine $MACHINE
			;;
			\?)
				echo "Invalid option: -$OPTARG" >&2
			;;
			:)
				echo "Option -$OPTARG requires an argument." >&2
				exit 1
			;;
		esac
	done

	if [ "$STACKNAME" == "" ]
	then
		echo missing -s option to choose stack
		showHelp
		exit
	fi


	if [ "$DEPLOY" = "y" ]
	then
			selectMaster
			if [ $STACKNAME == "dia" ]
			then
				docker stack deploy -c docker-compose.kafka.yml kafka
			fi
			echo deploying on stack  $DIRECTORY/$COMPOSE_FILE $STACKNAME
			docker stack deploy -c $DIRECTORY/$COMPOSE_FILE $STACKNAME
			exit
	fi

	NUMBER=0
	rm -f $FILETMP 2>/dev/null

	docker-compose -f $DIRECTORY/$COMPOSE_FILE config --services | sort -u > $FILETMP2

	if [ $? -ne 0 ]
	then
		echo error in docker compose file
		exit
	fi

	cat $FILETMP2 | grep -v visualizer | while read line; do echo "$NUMBER-$line" >> $FILETMP;NUMBER=`expr $NUMBER + 1`; done

	rm -f $FILETMP2

	cat $FILETMP
	echo "a-for all"

	if [ "$LINENUMBER" == "" ]
	then
		read LINENUMBER
	fi

	if [ "$LINENUMBER" == "a" ]
	then
		NUMBER=`wc -l $FILETMP | cut -f1 -d\ `
		for (( c=0; c<$NUMBER; c++ ))
		do
			RESULT=`cat $FILETMP | grep $c | cut -f2- -d\- `
			rebuild_service $RESULT
		done
	else
		RESULT=`cat $FILETMP | grep ${LINENUMBER}- | cut -f2- -d\- `
		rebuild_service $RESULT
	fi

	clean

	if [ "$SLAVES" != "n" ]
	then
		selectSlave
		clean
	fi

	selectLocalHost
	clean

	rm -f $FILETMP $FILETMP2

	echo "docker stack services $STACKNAME  "
	echo "run this to get logs:"
	echo "docker service logs -f ${SERVICE}"
}

main $*

echo bye bye
