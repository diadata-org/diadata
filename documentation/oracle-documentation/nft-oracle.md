---
description: How do I access NFT pricing data using the DIA oracle system?
---

# NFT Oracle

The oracle contains information about NFT collections' prices (see [sources](https://docs.diadata.org/documentation/methodology/digital-assets/nft-data-collection)). It can hold up to six values.\
For the current deployed demo version you can access a collection's latest floor price and the current circulating supply as well as the timestamp of the last update.



1. Access the corresponding [oracle smart contract](https://ropsten.etherscan.io/address/0x93263e599b63fc8602cd24f8a05355723ac0609f#readContract).
2. Call `getValue(collection_id)` with `collection_id` being the string `Blockchain-Address` , for instance `Ethereum-0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D`  for Bored Apes Yacht Club. You can use the "Read" section on Etherscan to execute this call.
3. The current version of the DIA demo oracle contains the following values:
   1. The current floor price in the blockchain's native currency with a fix-comma notation of eight decimals (see [API endpoint](https://docs.diadata.org/documentation/api-1/api-endpoints#nft-floor-price)).
   2. The 30-day moving average of the floor price (see [API endpoint](https://docs.diadata.org/documentation/api-1/api-endpoints#nft-moving-average-of-floor-price)).
   3. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.

<details>

<summary>See the list of collections available on the oracle</summary>

* Bored Ape Yacht Club - 0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d
* CryptoPunks - 0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb
* Moonbirds - 0x23581767a106ae21c074b2276d25e5c3e136a68b
* Doodles - 0x8a90cab2b38dba80c64b7734e58ee1db38b8992e
* Otherdeed - 0x34d85c9cdeb23fa97cb08333b511ac86e1c4e258

</details>
