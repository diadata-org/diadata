---
description: This page contains an overview on how to interact with the NEAR Request oracle
---

# NEAR Request Oracle

On the NEAR blockchain, DIA operates an oracle that can fetch quotations from the DIA API. The oracle follows a request/response pattern, i.e., an on-chain NEAR contract _requests_ data from the DIA API which is then promptly served by the oracle as a callback.

The latest oracle addresses are listed [here](deployed-contracts.md).

### How to use the oracle

The oracle can be used by interfacing it with a client for quotations, that takes the callbacks and processes the data. An example deployment can be seen [here on the NEAR testnet](https://explorer.testnet.near.org/accounts/quote-test-client.dia-test.testnet). The corresponding source code is located [in this repository](https://github.com/diadata-org/dia-adapter).

Follow the instructions in its README to learn how to interact with the DIA oracle on the NEAR blockchain. There, the exact steps are described for getting a client up and running that is served quotations using the request oracle.

We will extend the NEAR oracle with additional data from our API in the future.

### Usage Example

You can test the oracle by querying a price quotation using code from the [example repository](https://github.com/diadata-org/dia-adapter). For that, checkout the repository and install dependencies using `npm install`\
After that, you can build the oracle consumer with `npm run build`and start the oracle consumer with `node dist/test/quote-make-request`

An example request and response for the asset DIA (can be configured in the source file) looks like this

`user@host% node dist/test/quote-make-request` \
`near.call quote-test-client.dia-test.testnet make_request`\
`near.view quote-test-client.dia-test.testnet get_callback_response`\
`result: {"request_id":"111","err":"","data":{"Symbol":"DIA","Name":"DIAData","Price":2.433626800388127,"PriceYesterday":2.542690797155571,"VolumeYesterdayUSD":583156.1339811679,"Source":"diadata.org","Time":"2021-03-12T14:10:18.357245071Z","ITIN":"undefined"}}`
