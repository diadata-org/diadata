---
description: >-
  How do I access information on crypto farming pools using the DIA oracle
  system?
---

# Farming Pools



The oracle exists on various blockchains and contains information about pool rates and balances of crypto farming pools.   
You can execute an oracle call as follows:

1.  Access the corresponding [oracle smart contract](https://docs.diadata.org/documentation/oracle-documentation/deployed-contracts).
2. Call `getCoinInfo(protocol_name)` with `protocol_name` being the name of the farming protocol such as  `Balancer`. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains four values:
   1. Pool rate \(see [here](https://docs.diadata.org/documentation/methodology/digital-assets/return-rates-in-crypto-farming) for an explanation\).
   2. Pool balance \(measured in pool token\).
   3. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.
   4. Pool ID.

As of now, the following list of pools is available in this oracle:

**Ethereum Mainnet:**

* YFI \(yearn.finance\): WETH pool

**Binance Smart Chain:**

* BARNBRIDGE: Stablecoin pool \(USDC, DAI, sUSD\)
* CURVEFI: 3pool \(USDC, DAI, USDT\)
* LOOPRING: LRC pool
* SYNTHETIX: sETH pool
* YFI \(yearn.finance\): WETH pool

**Polygon \(Matic\):**

* BARNBRIDGE: Stablecoin pool \(USDC, DAI, sUSD\)
* CURVEFI: 3pool \(USDC, DAI, USDT\)
* LOOPRING: LRC pool
* SYNTHETIX: sETH pool
* YFI \(yearn.finance\): WETH pool

**Moonbeam:**

* BARNBRIDGE: Stablecoin pool \(USDC, DAI, sUSD\)
* CURVEFI: 3pool \(USDC, DAI, USDT\)
* LOOPRING: LRC pool
* SYNTHETIX: sETH pool
* YFI \(yearn.finance\): WETH pool

_Remark_: Due to different mechanics of farming, the meaning of pool rate can differ between protocols. See [here](https://docs.diadata.org/documentation/methodology/digital-assets/return-rates-in-crypto-farming) for a detailed explanation. The emitted value "Pool rate" has the following meaning for the corresponding protocol:  
- Pool rate: _YFI_  
- Total debt: _Synthetix_  
- Total reward: _Loopring_  
- Virtual price: _Curve Finance_  
- Fixed reward: _Barnbridge_  


