# DIA Technical Structure

DIA is setup as a hybrid system with off-chain components for storing and processing large amounts of data and on-chain components providing data sources for financial smart contracts. Currently, Ethereum is used as smart contract platform because of its widespread use and technical quality, but the DIA system is not limited to one blockchain technology.

### Centralized Backend

For collecting financial data, we use a centralized backend that runs collectors for all kinds of financial data. All collected data is processed by a database setup consisting of a stream-oriented kafka instance and a key-value store for faster access of certain intermediate results. These collectors are separated in different classes:

* Exchange Scrapers: These collect data from exchanges for cryptocurrencies and other assets \(like stocks, futures, and rare earth metals\).

  Every exchange scraper is run independently and is structured around the idea of `pairs` that are used to indicate the pair of assets that is used in a trade.

  For example, a trade from Bitcoin to US-Dollar is stored under `BTCUSD`.

* Blockchain Scrapers: These scrapers are used to determine attributes of blockchains. The most prominent attribute is the circulating supply of a cryptocurrency, for which various scrapers already exist. Depending on the expected update rate of a blockchain, these scrapers update the backend periodically about their data.

  Blockchain scrapers are run as docker containers that can be managed independently.

  Often, these collect Gigabytes of data over time \(they are dealing with ever-growing blockchains\) and thus should be managed in a separate way not directly baked in the database executable.

* Quotation Scrapers: These are used to collect official quotations from central trusted providers.

  Apart from daily exchange rates from the European Central Bank \(ECB\) against various international currencies we collect several interbank overnight interest rates such as SOFR and €STR.

