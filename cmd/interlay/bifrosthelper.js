const { ApiPromise, WsProvider } = require("@polkadot/api");
const ethers = require("ethers");
const bignumber = ethers.BigNumber;


async function tokenPool(api, token) {
    const tokenPoolMap = new Map();
    const tokenPoolEntries = await api.query.vtokenMinting.tokenPool.entries();
  
    tokenPoolEntries.forEach((tokenPool) => {
      let key = tokenPool[0].toHuman();
      let value = tokenPool[1].toHuman();
      if (key[0].Token) {
        tokenPoolMap.set(key[0].Token, value);
      }
    });
  
    return tokenPoolMap.get(token);
  }

async function tokenIssuance(api, token) {
    const tokenIssuanceMap = new Map();
    const totalIssuance = await api.query.tokens.totalIssuance.entries();
  
    totalIssuance.forEach((totalIssuance) => {
      let key = totalIssuance[0].toHuman();
      let value = totalIssuance[1].toHuman();
  
      if (key[0].Token) {
        tokenIssuanceMap.set(key[0].Token, value);
      }
    });
  
    return tokenIssuanceMap.get(token);
  }

  async function vTokenIssuance(api, token) {
    const totalIssuance = await api.query.tokens.totalIssuance({vToken:token});
  
    return totalIssuance;
  }

  async function getBiFrostValues(token) {
    let providerurl = "";
  
    switch (token) {
      case "KSM":
        providerurl = "wss://bifrost-parachain.api.onfinality.io/public-ws";
        break;
    }
  
    const wsProvider = new WsProvider(
      // "wss://interlay.api.onfinality.io/public-ws"
      providerurl
    );
    const api = await ApiPromise.create({
      provider: wsProvider,
    });
  
    let tokeninpool = await tokenPool(api, token);
    // let tokenIssuance = await bifrosttokenIssuance(api, token);
    let vtokenIssuance = await vTokenIssuance(api, token);
  
 
  
    return {
      total_backable: bignumber.from(tokeninpool.replaceAll(",", "")).toString(),
      total_issued: vtokenIssuance.toString(),
      decimal: 12,
      token: token,
      time: Date.now(),
    };
  }
  
  

module.exports = {
  tokenPool: tokenPool,
  bifrosttokenIssuance:tokenIssuance,
  vTokenIssuance:vTokenIssuance,
  getBiFrostValues:getBiFrostValues,
};

// {
//   "total_backable": "250465499708802373",
//   "total_issued": "223466021908860328",
//   "decimal": 12,
//   "token": "KSM"
//   }