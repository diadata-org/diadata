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

async function getPrice(asset) {
  let response = await fetch("https://api.diadata.org/v1/quotation/" + asset);
  let ethprice = await response.json();
  return ethprice.Price;
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
