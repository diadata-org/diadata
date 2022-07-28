---
description: >-
  This example demonstrates the functionality of the Laminar chain's oracle
  pallet including the aggregation of price values.
---

# Polkadot Medianizer

### Laminar Chain

The Laminar chain is built using the [Substrate framework](https://substrate.dev). In our example we are going to setup a network consisting if two nodes, one is operated by Alice and one is operated by Bob.

#### Setup

First we download the Laminar chain [sources](https://github.com/laminar-protocol/laminar-chain) and bootstrap the network. After that we execute the following statements in the chain's directory.

```text
example$ git clone https://github.com/laminar-protocol/laminar-chain.git
example$ cd laminar-chain
example/laminar-chain$ make init
example/laminar-chain$ make build
```

`make init` will setup the environment by downloading and installing all necessary dependencies and the required toolchain.

`make build` builds the chain's code so we can execute it.

After that copy the `chainSpec.json` file from this repository into the chain's directory. More on the spec file will be explained a little bit later.

#### Cleanup

When running a node being part of the laminar chain network it creates and stores lots of data because it's constantly building new blocks and extending the actual blockchain. We are going to store all that data in a subdirectory of `/tmp` for each node \(its base directory\). If the base directory doesn't exist it will be created but it is necessary to purge all data of a specific node on subsequent launches:

```text
example/laminar-chain$ SKIP_WASM_BUILD= cargo run -- purge-chain --base-path /tmp/alice
```

This wipes the base directory of Alice's node. For wiping Bob's node's data adjust the base-path parameter accordingly. The action asks for confirmation which is given by answering with `y`.

**Bootstrap the network**

We are going to launch Alice's node first:

```text
example/laminar-chain$ SKIP_WASM_BUILD= cargo run -- --base-path /tmp/alice --chain=chainSpec.json --alice --port 30333 --ws-port 9945 --rpc-port 9933
```

* This creates and runs a node that stores its data in the specified base directory.
* We also pass the chain specification file into the node. We use this file to configure the scenario in which all participating parties will own a sufficient amount of gas in order to be able to "pay" for their transmitted transactions.
* The parameter `--alice` will make Alice the node's authority uses her predefined key pair.
* We specify three ports: One for the other peers to connect to the node, a websocket port for the off-chain worker to connect to and a port for incoming RPCs.

After that we fire up Bob's node and tell it to connect to Alice. So in another terminal session we are executing

```text
example/laminar-chain$ SKIP_WASM_BUILD= cargo run -- --base-path /tmp/bob --chain=chainSpec.json --bob --port 30334 --ws-port 9946 --rpc-port 9934 --bootnodes /ip4/127.0.0.1/tcp/30333/p2p/12D3KooWEyoppNCUx8Yx66oV9fJnriXwCcXwDDUA2kj6vnc6iDEp
```

The main difference is the `bootnodes` parameter which specifies the network and identity information of Alice's node.

#### Use the Laminar chain's oracle

Now the two nodes start working together and our laminar chain network is ready for playing with the oracle.

**About oracles**

A blockchain oracle is an interface for the chain to the real world and vice versa. By feeding off-chain data to the oracle they enter the chain's scope and can get processed by smart contracts. In this case our oracle is called an **inbound oracle**. For our example imagine a smart contract that operates on currency prices. It might trigger an action as soon as the price for a specific currency exceeds a specific value. By feeding the current prices from external sources periodically to the oracle we provide the smart contract with the required input data on-chain.

**Oracle pallet**

A Substrate chain's runtime consists of modules called pallets. In the case of Laminar one of these pallets is our [oracle](https://github.com/open-web3-stack/open-runtime-module-library/tree/master/oracle). It can be interacted with in two ways:

1. Feeding new key-value pairs to the oracle via signed transactions.
2. Retrieving key-value pairs from the oracle either by other modules on-chain or by off-chain workers via RPC.

The Laminar implementation configures the oracle pallet to use the type `CurrencyId` as key type. That means it can be one of the predefined values _FBTC_, _FETH_, _FEUR_, etc. The value uses the type `Price`.

In this example we are going to use an off-chain worker to feed currency prices and fetch them afterwards.

**Our off-chain worker**

The component we create to simulate new price reports and oracle queries lives off-chain and will be written in Typescript. It is completely separate from anything stored on the Laminar chain itself. We rather interact with the chain through specific calls made using the [Polkadot API for JavaScript](https://polkadot.js.org/docs/api). This API defines the technical aspects of how we can interact with the chain \(and specifically with the oracle\), i. e. how to build transactions and make RPCs. For the specific parameters regarding the Laminar chain we also make use of the [Laminar chain JS SDK](https://github.com/laminar-protocol/laminar-chain.js).

Let's start by switching the directory

```text
example/laminar-chain$ cd ..
example$ mkdir medianizer
example/medianizer$ touch package.json
```

and setting up the dependencies in the `package.json` file:

```text
{
  "name": "medianizer-example",
  "version": "0.1.0",
  "description": "",
  "main": "index.ts",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "start": "node --require ts-node/register index.ts"
  },
  "dependencies": {
    "@laminar/api": "^0.2.0-beta.144",
    "@polkadot/api": "^3.8.1"
  },
  "devDependencies": {
    "ts-node": "^9.1.1",
    "typescript": "^4.1.5"
  },
  "keywords": [],
  "author": "",
  "license": "ISC"
}
```

Next step is coding the off-chain worker in the file `index.ts`. We import the required types

```text
import { ApiPromise, WsProvider, Keyring } from "@polkadot/api";
import { options } from "@laminar/api";
```

and build the frame for the program:

```text
const run = async () => {
  // Everything is going to happen here in this block ...
};

export default run;

if (require.main === module) {
    run().then(() => process.exit(0))
}
```

Now we can run

```text
example/medianizer$ yarn install
example/medianizer$ yarn start
```

for fetching the dependencies and launching the program. However nothing is going on right now so let's write the actual code. From now on we put everything into the async block that we just prepared.

We start creating an API object:

```text
const wsProvider = new WsProvider('ws://localhost:9945');
const api = await ApiPromise.create({
    ...options({}),
    provider: wsProvider });
```

Please note that our API object uses a websocket to communicate with the node. And this websocket connects to port 9945 which we have specified above when firing up Alice's node. So all the requests are submitted to her node.

Next step is configuring the key material. We want to transmit signed transactions that contain the prices to be fed in the oracle. Therefore the off-chain worker needs to know the key pairs of Alice, Bob and Charlie \(yes, he's also going to take part in this example\). Usually the key pairs would be imported, e. g. from a file. But as for the chain nodes we use the predefined key material which is addressed by `//[Name]`. We derive those key pairs from the keyring we create specifying the key pair format SR25519.

```text
const keyring = new Keyring({ type: 'sr25519'});

const alice = keyring.addFromUri("//Alice");
const bob = keyring.addFromUri("//Bob");
const charlie = keyring.addFromUri("//Charlie");
```

Alice wants to trigger a transaction that feeds the oracle. If we start typing `api.tx.` the code completion shows us the available extrinsics but the oracle isn't one of them. That's because the default methods of the Substrate runtime are known at compile time but not the custom chain specific extrinsics that might be part of the runtime. Let's find out what's really available by querying the node:

```text
const transactions = await api.tx;
console.log(transactions);
```

When executing the program we get a list of much more actually available extrinsics - and one of them is `laminarOracle`. The list tells us even more about the oracle: the available query method `feedValues`. This is exactly what we needed to know.x

Now Alice reports a current BTC price of 1, i. e. she feeds the oracle with the corresponding key-value pair:

```text
await api.tx['laminarOracle'].feedValues([['FBTC', 1]] as any).signAndSend(alice);
```

This line consists of several interesting pieces:

1. We use the API object to trigger an extrinsic \(`tx` for transaction\). The available extrinsics depend on the used chain. In this case we address the oracle pallet of the laminar chain which is addressed by `laminarOracle`.
2. On that extrinsic we call the method `feedValues` in order to submit key-value pairs that should be fed into the oracle. In this case we just submit a single key-value pair for the Bitcoin price.
3. As this operation is a transaction that is going to be included in a block and costs Alice some gas it needs to be signed by her. This is what happens in the last part of the line.

When this line returns the transaction has been finalized and we can now query the oracle with the current Bitcoin price. But there's again the question of how to do that. As mentioned earlier the oracle offers retrieving the current values via RPC. So let's find out the available RPCs:

```text
const methods = await api.rpc.rpc.methods;
console.log(methods);
```

The output again gives us what we need. Two calls should catch our attention: `oracle_getAllValues` and `oracle_getValue`. That means we have the two functions `getAllValues` and `getValue` available for the RPC `oracle.`

Now we query the current BTC price:

```text
const result = await api.rpc['oracle'].getValue('Laminar', 'FBTC');
console.log(`Current aggregated BTC price: ${result}`);
```

As you can see we use an RPC as the query doesn't change the state of the chain and therefore the operation is free. We specify the correct RPC \(which is `oracle`\) and query the value to the key associated with the Bitcoin price \(which is `FBTC`\). The output is a pair of two values: The current price and the timestamp of the transaction that contained the feeding operation.

That's it! Now we can run the program again with `yarn start`. The oracle will be fed and queried and the output is the expected BTC price of 1.

**Data aggregation**

In the previous section we fed the oracle with one single data point and queried the oracle to retrieve a value. It isn't surprising at all that we read the value that we had written into the oracle just before. But it requires a closer look to what value is actually returned by `getValue`. The oracle pallet states that this method returns a "combined value" which indicates that there is some data aggregation happening.

In fact the way the prices of each currency are aggregated is entirely up to the chain implementation. Currently the Laminar chain doesn't implement any aggregation algorithm but the oracle pallet comes with a default implementation: It takes all fed values of a specific currency submitted within the last 3 days and returns the median value. If there aren't any values to calculate the median from it just returns the most recent value.

Let's try it out. Disable the lines of the previous section that interacted with the oracle \(via tx and RPC\) and add the following:

```text
const feedPrice = async (value: string, sender: any, completion: () => void) => {
    const unsub = await api.tx['laminarOracle'].feedValues([['FBTC', value]] as any).signAndSend(sender, (result) => {
        if (result.status.isInBlock) {
            console.debug(`Transaction included at blockHash ${result.status.asInBlock}`);
        } else if (result.status.isFinalized) {
            console.debug(`Transaction finalized at blockHash ${result.status.asFinalized}`);
            unsub();
            completion();
        }
    });
};
```

For more convenience we define a function `feedPrice` which takes a price value and a sender and triggers the transaction accordingly. We use the [subscription API](https://polkadot.js.org/docs/api/start/api.query.subs) of Polkadot JS in order to make sure that we wait for every transaction to be finalized in a block. Now our three users submit their values:

```text
await new Promise<void>((resolve) => { feedPrice("1", alice, resolve); });
await new Promise<void>((resolve) => { feedPrice("5", bob, resolve); });
await new Promise<void>((resolve) => { feedPrice("10", charlie, resolve); });

const result = await api.rpc['oracle'].getValue('Laminar', 'FBTC');
console.log(result);
```

When we query the oracle for the BTC price it returns the median value of 5 as expected.

**Important!** Remember to purge the chain status and restart the two nodes before executing the code of this section.

### TL;DR - Quick demo

```text
mkdir example
cd example
```

Download sources for the example and the laminar chain:

```text
git clone https://github.com/diadata-org/medianizer-example.git
git clone https://github.com/laminar-protocol/laminar-chain.git
```

Setup laminar chain:

```text
cd laminar-chain
make init
make build
cp ../medianizer-example/chain/chainSpec.json .
```

Run first chain node \(Answer with "y" when asked for confirmation before purging\):

```text
SKIP_WASM_BUILD= cargo run -- purge-chain --base-path /tmp/alice

SKIP_WASM_BUILD= cargo run -- --base-path /tmp/alice --chain=chainSpec.json --alice --port 30333 --ws-port 9945 --rpc-port 9933 --node-key 0000000000000000000000000000000000000000000000000000000000000001
```

Run second chain node \(e.g. in another terminal session\):

```text
SKIP_WASM_BUILD= cargo run -- purge-chain --base-path /tmp/bob

SKIP_WASM_BUILD= cargo run -- --base-path /tmp/bob --chain=chainSpec.json --bob --port 30334 --ws-port 9946 --rpc-port 9934 --bootnodes /ip4/127.0.0.1/tcp/30333/p2p/12D3KooWEyoppNCUx8Yx66oV9fJnriXwCcXwDDUA2kj6vnc6iDEp
```

After both nodes have started to communicate with each other run medianizer example \(in yet another terminal session\):

```text
cd ../medianizer-example
yarn install
yarn start
```

