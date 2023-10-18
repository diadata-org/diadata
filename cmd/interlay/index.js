const { ApiPromise, WsProvider } = require("@polkadot/api");
const { getInterlayValues } = require("./interlayhelper");
const { getBiFrostValues } = require("./bifrosthelper");

const { cronstart } = require("./cron");
let cron = require("node-cron");

const { tokenkey, redis, allowedTokens, pricekey } = require("./utils");

let cache = redis();

cronstart();

cron.schedule("10 * * * *", () => {
  console.log("running a task every minute");
  cronstart();
});

let express = require("express");
let app = express();

app.get("/customer/interlay/state/:token", async function (req, res) {
  // kbtc, ibtc
  const allowedTokens = ["IBTC", "KBTC", "DOT"];

  let token = req.params.token;
  token = token.toUpperCase();

  if (allowedTokens.includes(token)) {
    // let values = await getInterlayValues(token);
    let values = JSON.parse(await cache.get("interlayraw" + token));

    res.send(values);
  } else {
    res.send({ err: "invalid token use, IBTC or KBTC" });
  }
});

app.get("/customer/bifrost/state/:token", async function (req, res) {
  // kbtc, ibtc
  const allowedTokens = ["KSM", "DOT"];

  let token = req.params.token;
  token = token.toUpperCase();

  if (allowedTokens.includes(token)) {
    let values = await getBiFrostValues(token);
    res.send(values);
  } else {
    res.send({ err: "invalid token use, KSM or DOT" });
  }
});

let values = {
  Token: "",
  FairPrice: null,
  Collateralratio: {
    IssuedToken: "",
    LockedToken: "",
    Ratio: "",
  },
  BaseAssetSymbol: "",
  BaseAssetPrice: "",
  Issuer: "",
};

function findTokenByVTokenAndIssuer(vtoken, issuer) {
  const lowerCaseIssuer = issuer.toLowerCase();  
  const lowerCaseVtoken = vtoken.toLowerCase();

  for (let i = 0; i < allowedTokens.length; i++) {
    if (
      allowedTokens[i].vtoken.toLowerCase() === lowerCaseVtoken &&
      allowedTokens[i].issuer.toLowerCase() === lowerCaseIssuer
    ) {
      return allowedTokens[i];
    }
  }
  return null;
}
app.get("/xlsd/:issuer/:vtoken", async function (req, res) {
  let issuer = req.params.issuer;
  let vtoken = req.params.vtoken;

  let token = findTokenByVTokenAndIssuer(vtoken, issuer);

  if (token) {
    let cacheddata = JSON.parse(
      await cache.get(tokenkey(token.source, token.vtoken))
    );

    let tokenvalues = await createXResponse(cacheddata, token);
    res.send(tokenvalues);
  } else {
    res.send({ err: "invalid token and issuer" });
  }
});

async function createXResponse(cacheddata, token) {
  let tokenvalues = JSON.parse(JSON.stringify(values));
  let baseassetprice;
  baseassetprice = await cache.get(pricekey(token.token));

  if (cacheddata && cacheddata.collateral_ratio) {
    tokenvalues.Collateralratio = {
      IssuedToken: cacheddata.collateral_ratio.issued_token,
      LockedToken: cacheddata.collateral_ratio.locked_token,
      Ratio: cacheddata.collateral_ratio.ratio,
    };
  }
  if (cacheddata && cacheddata.fair_price) {
    tokenvalues.FairPrice = cacheddata.fair_price;
  }
  if (cacheddata && cacheddata.decimal) {
    tokenvalues.decimal = cacheddata.decimal;
  }
  // if (cacheddata && cacheddata.ratio) {
  //   tokenvalues.ratio = cacheddata.ratio;
  // }
  tokenvalues.BaseAssetPrice = parseFloat(baseassetprice);
  tokenvalues.Token = token.vtoken;
  tokenvalues.BaseAssetSymbol = token.token;
  tokenvalues.Issuer = token.issuer;

  return tokenvalues;
}

app.get("/xlsd", async function (req, res) {
  // kbtc, ibtc

  let response = [];

  for (const token of allowedTokens) {
    let cacheddata = JSON.parse(
      await cache.get(tokenkey(token.source, token.vtoken))
    );

    let tokenvalues = await createXResponse(cacheddata, token);

    response.push(tokenvalues);
  }
  res.send(response);
});

let server = app.listen(8081, function () {
  var host = server.address().address;
  var port = server.address().port;
  console.log(" App listening at http://%s:%s", host, port);
});
