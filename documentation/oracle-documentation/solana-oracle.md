---
description: This page contains an overview over the DIA Solana oracle.
---

# Solana Oracle

### Demo Application

A demo application has been deployed to the Solana blockchain. A simple frontend showing current asset prices is hosted as a static site and can be accessed here:

[Launch Demo Application](https://diadata-solanaoracle.netlify.app)

The application currently supports this set of asset prices.

| Asset Name | Query Command | Address (Devnet)                                                                                                                                |
| ---------- | ------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| Solana     | SOL           | [PD8AXFHHqVFqrywQYTrJfrDNY4oAENBKSHq6ctQXPfa](https://explorer.solana.com/address/PD8AXFHHqVFqrywQYTrJfrDNY4oAENBKSHq6ctQXPfa?cluster=devnet)   |
| Bitcoin    | BTC           | [FJpv98TrcWURFaGXVRnzwQ7gfdF2ZWzKYeRo6Y3Jim9Z](https://explorer.solana.com/address/FJpv98TrcWURFaGXVRnzwQ7gfdF2ZWzKYeRo6Y3Jim9Z?cluster=devnet) |
| Ethereum   | ETH           | [GzBueExjjLxHbUK5QYBUKVHxUh4vU4wztKdtQ1UyejT3](https://explorer.solana.com/address/GzBueExjjLxHbUK5QYBUKVHxUh4vU4wztKdtQ1UyejT3?cluster=devnet) |
| Diadata    | DIA           | [4xfZ1FhvuDqMLSUUMyLBKdJpEjfC6VQYsnZqAFV182BR](https://explorer.solana.com/address/4xfZ1FhvuDqMLSUUMyLBKdJpEjfC6VQYsnZqAFV182BR?cluster=devnet) |
| Circle USD | USDC          | [E6nLSeutmECcp1USAWUnuYJ3NRKNAtKyv1SBMNLPRv5V](https://explorer.solana.com/address/E6nLSeutmECcp1USAWUnuYJ3NRKNAtKyv1SBMNLPRv5V?cluster=devnet) |
| Cardano    | ADA           | [EieYNn9t56ma52Jt3UYbeMCVyzf3jt4bGXE55gQJFRJK](https://explorer.solana.com/address/EieYNn9t56ma52Jt3UYbeMCVyzf3jt4bGXE55gQJFRJK?cluster=devnet) |

### Solana Oracle Operations

The Solana blockchain uses a different address to store each oracle index (Read more about [Solana Program Derived Addresses](https://docs.solana.com/developing/programming-model/calling-between-programs)).&#x20;

Therefore, to read the price of an asset, it is a requirement to have the address that stores the price of that asset.&#x20;

The workflow presented below represents the configuration of the DIA Solana Oracle and how data should be supplied to it through the DIA Data Feeder.

### DIA Data Feeder

The feeder gets symbol price informaton available from DIA.

1. The feeder calls the API https://api.diadata.org/v1/quotation/\_symbol\_ with each asset obtained from the configured asset list.
2. For each symbol that the feeder calls to quote, it verifies that the data has been received correctly. If for some reason, the feeder does not obtain the data or the data is incorrect, it continues with the next symbol, and so on, until it ends with all the configured asset symbols.
3. Once the feeder verifies all the data, it creates an address for each symbol for the first time (creation phase). At the update phase, if the symbol was already stored before, a transaction is sent to the oracle updating the price. If new symbols appear, a new address is created for the symbol and a transaction is submitted loading the price and symbol for the first time and saved. Continue with the next symbol.
4. Once all the symbols have been traversed, the feeder returns to step 1 to call all the symbols that the API has.

There are several workflows associated with the DIA Data Feeder.

#### Create Data

1. The Feeder prepares the following data:
   * Symbol address: this data can be generated and stored at the moment of being supplied to the oracle or it can be created in a previous instance and stored in a database.
   * Authority: is the address that will be authorized in the future to update the asset data. Important: it can be the same address as the Symbol address.
   * Data asset: Price and symbol.
2. The feeder calls the contract with the data and the function "create\_coin\_info". This way, the asset of that Symbol address is initialized.

#### Update Data

1. Feeder prepares the data to call and update the asset:
   * Symbol address: address that stores the data.
   * Authority: previously authorized address, it must also be a signer for the transaction.
   * Data asset: Price.
2. Feeder calls the contract with the data and the function "update\_coin\_info", updating the price of the asset.

#### Oracle Asset Structure

In the current Oracle architecture, each asset has its own authority address; this address can either be different for each asset or the same for all the assets.

Assets are stored using the following structure:

```
#[account]
pub struct CoinInfo {
    price: u64,
    supply: u64,
    last_update_timestamp: u64,
    authority: Pubkey,
    symbol: String
}
```
