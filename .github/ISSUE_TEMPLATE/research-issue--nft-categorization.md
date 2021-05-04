---
name: 'Research Issue: NFT Categorization'
about: Assign attributes to NFT classes
title: ''
labels: ''
assignees: ''

---

We are looking for someone experienced or at least interested in NFTs, not necessarily with a technical background.
Your task is to assign attributes to NFT classes such as Hashmasks ("Arts"), Cryptokitties ("Collectibles") or Ethereum Name Service ("Certificates and Licenses").
Each class should be assigned exactly one attribute from the list returned at the following API endpoint:
https://api.diadata.org/v1/NFTCategories
In case you think that this list is incomplete, please let us know.

Your submission consists of the json file downloaded from:
https://api.diadata.org/v1/NFTClasses/limit/offset
with the field "Category" filled out. From all we can tell so far, this is a manual task. Information on the asset class corresponding to the Address can be found on https://opensea.io (and certainly other platforms). Please spell the categories such as emitted in the above API endpoint, so we can machine read it.
If anything is unclear, do not hesitate to ask.
