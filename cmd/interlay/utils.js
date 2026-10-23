function collaterlaratio(issuedtoken, lockedtoken) {
  return { issued_token: issuedtoken, locked_token: lockedtoken };
}

function createResponse(issuedtoken, lockedtoken, fairPrice, ratio) {
  return {
    collateral_ratio: {
      issued_token: issuedtoken,
      locked_token: lockedtoken,
      ratio: ratio,
    },
    fair_price: fairPrice,
    timestamp: Date.now(),

  };
}

function tokenkey(source, token) {
  return "LSD_" + token;
}

function pricekey(token) {
  return "BASE_PARICE_" + token;
}

let cache;

let isRedisStarted;
function redis() {
  const redis = require("redis");
  const redishost = process.env.REDISHOST;
  const redisport = process.env.REDISPORT;

  const redispassword = process.env.REDISPASSWORD;
  let redisurl = "redis://:" + redispassword + "@" + redishost + ":" + redisport;

 
  if (!isRedisStarted) {
    (async () => {
      cache = redis.createClient({
        url: redisurl,
      });

      // cache =  redis.createClient();

      cache.on("error", (error) => console.error(`Error : ${error}`));

      await cache.connect();
    })();
    isRedisStarted = true;
  }

  return cache;
}

const allowedTokens = [
  // {
  //   vtoken: "nASTR",
  //   token: "ASTR",
  //   source: "astar",
  //   issuer: "Algem",

  // },

  {
    vtoken: "vMOVR",
    token: "MOVR",
    source: "bifrost",
    issuer: "Bifrost",
  },
  {
    vtoken: "vBNC",
    token: "BNC",
    source: "bifrost",
    issuer: "Bifrost",
  },
  {
    vtoken: "vGLMR",
    token: "GLMR",
    source: "bifrost",
    issuer: "Bifrost",
  },
  {
    vtoken: "vASTR",
    token: "ASTR",
    source: "bifrost",
    issuer: "Bifrost",
  },
  {
    vtoken: "vFIL",
    token: "FIL",
    source: "bifrost",
    issuer: "Bifrost",
  },
  {
    vtoken: "IBTC",
    token: "BTC",
    source: "interlay",
    issuer: "Interlay",
  },
  {
    vtoken: "vKSM",
    token: "KSM",
    source: "bifrost",
    issuer: "Bifrost",
  },
  {
    vtoken: "vDOT",
    token: "DOT",
    source: "bifrost",
    issuer: "Bifrost",
  },
  {
    vtoken: "stDOT",
    token: "DOT",
    source: "stDOT",
    issuer: "Lido",
  },
  {
    vtoken: "rETH",
    token: "ETH",
    source: "rETH",
    issuer: "RocketPool",
  },
  {
    vtoken: "stETH",
    token: "ETH",
    source: "stETH",
    issuer: "Lido",
  },
  // {
  //   vtoken: "cbETH",
  //   token: "ETH",
  //   source: "cbETH",
  //   issuer: "Coinbase",

  // },
  {
    vtoken: "KBTC",
    token: "BTC",
    source: "interlay",
    issuer: "Interlay",
  },
];

const assetQuotationKeys = {
  MOVR: { blockchain: "Moonriver", address: "0x0000000000000000000000000000000000000000" },
  BNC: { blockchain: "Bifrost", address: "0" },
  GLMR: { blockchain: "Moonbeam", address: "0x0000000000000000000000000000000000000000" },
  ASTR: { blockchain: "Astar", address: "0x0000000000000000000000000000000000000000" },
  FIL: { blockchain: "Filecoin", address: "0x0000000000000000000000000000000000000000" },
  BTC: { blockchain: "Bitcoin", address: "0x0000000000000000000000000000000000000000" },
  KSM: { blockchain: "Kusama", address: "0x0000000000000000000000000000000000000000" },
  DOT: { blockchain: "Polkadot", address: "0x0000000000000000000000000000000000000000" },
  ETH: { blockchain: "Ethereum", address: "0x0000000000000000000000000000000000000000" },
};

async function getPrice(asset) {
  var providerurl = process.env.DIADATA_API || "https://api.diadata.org/v1/assetQuotation";

  const key = assetQuotationKeys[asset];
  if (!key) throw new Error("no assetQuotation mapping for " + asset);

  let response = await fetch(providerurl + "/" + key.blockchain + "/" + key.address);
  let quotation = await response.json();
  return quotation.Price;
}

module.exports = {
  collaterlaratio: collaterlaratio,
  tokenkey: tokenkey,
  redis: redis,
  createResponse: createResponse,
  allowedTokens: allowedTokens,
  getPrice: getPrice,
  pricekey: pricekey,
};
