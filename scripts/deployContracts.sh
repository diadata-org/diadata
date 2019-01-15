#! /usr/bin/env bash

# -----------------------------------------------------------------------
# Script configuration parameters
# -----------------------------------------------------------------------
set -e
NETWORK='development'
# Since we are running deterministic flag accounts[0] should be here
OWNER='0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1'
# Project name
PROJECT='DIA'
# Project version
VERSION='0.1.0'
echo 'Using zos version:'
npx zos --version


# -----------------------------------------------------------------------
# Project setup and first implementation of an upgradeable DIDRegistry
# -----------------------------------------------------------------------

# Clean up
rm -f zos.*
rm -rf build
# Compile
npx truffle compile

# Initialize zOS project
# NOTE: Creates a zos.json file that keeps track of the project's details
npx zos init $PROJECT $VERSION -v
# Register contracts in the project as an upgradeable contract.
# NOTE: here we need to add the rest of dia contracts
npx zos add DIAToken -v
npx zos add Dispute -v

# Deploy all implementations in the specified network.
# NOTE: Creates another zos.<network_name>.json file, specific to the network used, which keeps track of deployed addresses, etc.
npx zos push --network $NETWORK --skip-compile
# Request a proxy for the upgradeably contracts.
# Here we run initialize which replace contract constructors
# NOTE: A dapp could now use the address of the proxy specified in zos.<network_name>.json
# instance=MyContract.at(proxyAddress)
dt=$(npx zos create DIAToken --network $NETWORK --init initialize --args $OWNER -v)
npx zos create Dispute --network $NETWORK --init initialize --args $dt -v

#Copy JSON abis to web environment
if [ $# -eq 0 ]; then
    echo "No folder provided. ABIs will not be copied"
    exit 1
else
    cp build/contracts/NFT.json $1/
    cp build/contracts/NFTStore.json $1/
    cp zos.$NETWORK.json $1/
fi
