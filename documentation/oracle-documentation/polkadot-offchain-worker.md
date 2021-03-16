---
description: This page contains an overview over the DIA offchain worker for Polkadot
---

# Polkadot Offchain Worker

DIA offers a flexible oracle solution for Polkadot. Our offchain worker can be instantiated by any node operator. You can integrate it into your parachain and access the DIA API using the offchain worker.

### How it works

The DIA offchain worker is a component that can be ported to parachains and is located in [this repository](https://github.com/diadata-org/dia-substrate).

This offchain worker \(ocw\) gets data from an endpoint and writes an event as signed transaction for all local keys with subkey type `dia!`.

### Installation

To add the ocw \(offchain worker\) pallet to your node, add it to your runtime like this \(in this repository already done\):

1. Edit [`runtime/Cargo.toml`](https://github.com/diadata-org/dia-substrate/blob/dia/bin/node/runtime/Cargo.toml):
   * Add the following under `[dependencies]`:

     ```text
     pallet-dia-ocw = { version = "2.0.0", default-features = false, path = "../../../frame/dia-ocw" }
     ```

   * Add `"pallet-dia-ocw/std",` at `[features]`:

     ```text
     [features]
     std = [
         [...]
         "pallet-dia-ocw/std",
     ]
     ```
2. Edit [`runtime/src/lib.rs`](https://github.com/diadata-org/dia-substrate/blob/dia/bin/node/runtime/src/lib.rs) like this:
   * Add the following:

     ```text
     impl pallet_dia_ocw::Trait for Runtime {
         type Event = Event;
         type Call = Call;
         type AuthorityId = pallet_dia_ocw::crypto::TestAuthId;
     }
     ```

   * Insert `DIAOCW: pallet_dia_ocw::{Module, Call, Event<T>},` to `Runtime` enum:

     ```text
     construct_runtime!(
         pub enum Runtime where
             Block = Block,
             NodeBlock = node_primitives::Block,
             UncheckedExtrinsic = UncheckedExtrinsic
         {
             // ...
             DIAOCW: pallet_dia_ocw::{Module, Call, Event<T>},
         }
     );
     ```

### Usage

For each block, this offchain worker automatically adds a signed transaction. The signer account needs to pay the fees.

#### Local development mode

* Start the node and dev network by running `cargo run -- --dev --tmp`.
* Create an account or add a subkey to an existing account, e.g. the example account `Alice` via RPC:

  ```text
  curl http://localhost:9933 -H "Content-Type:application/json;charset=utf-8" -d \
    '{
      "jsonrpc":"2.0",
      "id":1,
      "method":"author_insertKey",
      "params": [
        "dia!",
        "bottom drive obey lake curtain smoke basket hold race lonely fit walk//Alice",
        "0xd43593c715fdd31c61141abd04a99fd6822c8558854ccde39a5684e7a56da27d"
      ]
    }'
  ```

