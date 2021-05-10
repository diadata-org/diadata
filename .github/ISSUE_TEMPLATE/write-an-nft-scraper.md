---
name: Write an NFT Scraper
about: Write an NFT Scraper
title: ''
labels: ''
assignees: ''

---

We are looking for a Go developer who is experienced in writing applications interfacing with the Ethereum blockchain. Basic knowledge in NFTs is helpful but not necessary.
Your task is to write a scraper for the [...] NFTs. The basic structure of the scraping backend is already set up, including the go bindings of the Sorare contract. More precisely, you should implement the scraping in the file [...]
If available, the data should be fetched directly from the smart contract. Additional information such as player's attributes can be fetched from the opensea API. 
Metadata and attributes should go into an encoded json in the field "attributes" in the struct dia.NFT. Please adhere to the json standard as described in EIP721 where possible.
In case you have questions or suggestions, please don't hesitate to ask.
