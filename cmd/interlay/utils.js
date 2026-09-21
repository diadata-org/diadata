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

const cacheTTLSeconds = parseInt(process.env.CACHE_TTL_SECONDS, 10) || 600;

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
  // {
  //   vtoken: "rETH",
  //   token: "ETH",
  //   source: "rETH",
  //   issuer: "RocketPool",
  // },
  // {
  //   vtoken: "stETH",
  //   token: "ETH",
  //   source: "stETH",
  //   issuer: "Lido",
  // },
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
  var providerurl = process.env.DIADATA_API || "https://api.diadata.org/v1/quotation/";

  let response = await fetch(providerurl + "/quotation/" + asset);
  let ethprice = await response.json();
  return ethprice.Price;
}

const { ethers } = require("ethers");

const luminaOracleAbi = [
  "function getValue(string memory key) external view returns (uint128, uint128)",
];

async function getPriceFromLumina(asset) {
  const luminaOracleAddress = process.env.LUMINA_ORACLE_ADDRESS || "0xBd82Fc0067CA5977b86c6EDB579f90DFcFb62D7d";
  const provider = new ethers.providers.JsonRpcProvider(process.env.DIA_RPC_URL);
  const oracle = new ethers.Contract(luminaOracleAddress, luminaOracleAbi, provider);

  const [price] = await oracle.getValue(asset + "/USD");

  return parseFloat(ethers.utils.formatUnits(price, 18));
}

module.exports = {
  collaterlaratio: collaterlaratio,
  tokenkey: tokenkey,
  redis: redis,
  createResponse: createResponse,
  allowedTokens: allowedTokens,
  getPrice: getPrice,
  getPriceFromLumina: getPriceFromLumina,
  pricekey: pricekey,
  cacheTTLSeconds: cacheTTLSeconds,
};
