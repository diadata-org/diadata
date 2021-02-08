---
description: >-
  How do I access information on crypto farming pools using the DIA oracle
  system?
---

# Farming Pools



The oracle contains information about pool rates and balances of crypto farming pools.   
You can execute an oracle call as follows:

1.  Access our [oracle smart contract](https://etherscan.io/address/0xD47FDf51D61c100C447E2D4747c7126F19fa23Ef).
2. Call `getCoinInfo(protocol_name)` with `protocol_name` being the name of the farming protocol such as  `Balancer`. You can use the "Read" section on Etherscan to execute this call.
3. The response of the call contains four values:
   1. Pool rate \(see [here](https://docs.diadata.org/documentation/methodology/digital-assets/return-rates-in-crypto-farming) for an explanation\).
   2. Pool balance \(measured in pool token\).
   3. The [UNIX timestamp](https://www.unixtimestamp.com/) of the last oracle update.
   4. Pool ID.

As of now, the following list of pools is available in this oracle:  
- BALANCER: largest WETH/WBTC pool \(by market cap\)  
- CVAULT: WETH pool  
- YFI \(yearn.finance\): WETH pool  
- SYNTHETIX: sETH pool  
- LOOPRING: LRC pool  
- CURVEFI: 3pool \(USDC, DAI, USDT\)  
- BARNBRIDGE: Stablecoin pool \(USDC, DAI, sUSD\)

_Remark_: Due to different mechanics of farming, the meaning of pool rate can differ between protocols. See [here](https://docs.diadata.org/documentation/methodology/digital-assets/return-rates-in-crypto-farming) for a detailed explanation. The emitted value "Pool rate" has the following meaning for the corresponding protocol:  
- Pool rate: _Balancer, CVault, YFI_  
- Total debt: _Synthetix_  
- Total reward: _Loopring_  
- Virtual price: _Curve Finance_  
- Fixed reward: _Barnbridge_  


