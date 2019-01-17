#! /usr/bin/env bash
# abort script on error
set -e
# -----------------------------------------------------------------------
# Script configuration
# -----------------------------------------------------------------------
# Project name
PROJECT='DIA'
# Project version
VERSION='0.1.0'
# Network to deploy
if [ -z $1 ]
then
    echo 'No network provided'
    NETWORK='development'
else
    echo 'Network provided'
    NETWORK=$1
fi
# -----------------------------------------------------------------------
# Network configuration
# -----------------------------------------------------------------------
case $NETWORK in
    'development' )
        echo 'Using network development'
        # Since is development is safe to erase all configs
        rm -f zos.dev*
        rm -f zos.json
        rm -rf build
        # Owner is the owner passed to ownables contracts
        # use accounts[0] for simplicity
        OWNER='0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1'
        # Admin is the account used to deploy and manage upgrades.
        # use accounts[9] for simplicity
        ADMIN='0x1df62f291b2e969fb0849d99d9ce41e2f137006e'
    ;;
    
    'rinkeby' )
        echo 'Using network rinkeby'
        # Owner is the owner passed to ownables contracts
        OWNER='0x05349268a998355280e55693b06B0823aB7C98cE'
        # Admin is the account used to deploy and manage upgrades.
        ADMIN='0x11aA0b53cbd2df11fa9e2D3CeA08d6F4EBd65e19'
    ;;
    * )
        echo 'network ['$NETWORK'] not implemented'
        echo 'please add network configuration before continuing'
        exit
    ;;
esac
# -----------------------------------------------------------------------
# Project setup using zOS
# -----------------------------------------------------------------------
# Set zos session (network, admin, timeout)
npx zos session --network $NETWORK --from $ADMIN --expires 36000
# Compile contracts
npx truffle compile
# Initialize zOS project
# NOTE: Creates a zos.json file that keeps track of the project's details
npx zos init $PROJECT $VERSION -v
# Register contracts in the project as an upgradeable contract.
npx zos add DIAToken -v --skip-compile
npx zos add Dispute -v --skip-compile
# Deploy all implementations in the specified network.
# NOTE: Creates another zos.<network_name>.json file, specific to the network used, which keeps track of deployed addresses, etc.
npx zos push --skip-compile  -v
# Request a proxy for the upgradeably contracts.
# Here we run initialize which replace contract constructors
# NOTE: A dapp could now use the address of the proxy specified in zos.<network_name>.json
# instance=MyContract.at(proxyAddress)
dt=$(npx zos create DIAToken --network $NETWORK --init initialize --args $OWNER -v)
npx zos create Dispute --network $NETWORK --init initialize --args $dt -v