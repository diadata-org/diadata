---
description: This page contains an overview on how to interact with the NEAR Request oracle
---

# NEAR Request Oracle

On the NEAR blockchain, DIA operates an oracle that can fetch quotations from the DIA API. The oracle follows a request/response pattern, i.e., an on-chain NEAR contract _requests_ data from the DIA API which is then promptly served by the oracle as a callback.

The latest oracle addresses are listed in [here](deployed-contracts.md).

### How to use the oracle

The oracle can be used by interfacing it with a client for quotations, that takes the callbacks and processes the data. An example deployment can be seen [here on the NEAR testnet](https://explorer.testnet.near.org/accounts/quote-test-client.dia-test.testnet). The corresponding source code is located [in this repository](https://github.com/diadata-org/dia-adapter).

Follow the instructions in its README to learn how to interact with the DIA oracle on the NEAR blockchain. There, the exact steps are described for getting a client up and running that is served quotations using the request oracle.

We will extend the NEAR oracle with additional data from our API in the future.

