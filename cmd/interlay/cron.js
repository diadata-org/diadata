const { getInterlayValues } = require("./interlayhelper");

const { getBiFrostValues } = require("./bifrosthelper");

const { exchangeRate, getValues: getValuecbeth } = require("./cbETHhelper");

const { getValues: getValuesteth } = require("./stETHhelper");

const { getValues: getValuereth } = require("./rETHhelper");

const { getValues: getValuestdot } = require("./stdothelper");

const { getValues: getValueAstar } = require("./nastrhelper");

const {
  tokenkey,
  redis,
  allowedTokens,
  getPriceFromLumina,
  pricekey,
  cacheTTLSeconds,
} = require("./utils");

let cache = redis();

async function cronstart() {
  for (const value of allowedTokens) {

    console.log("allowedTokens", value);
    switch (value.source) {
      case "interlay":
        {
          let saved
          try {
            saved = await getInterlayValues(value.vtoken);
            if (saved) {
              cache.set("interlayraw" + value.vtoken, JSON.stringify(saved), { EX: cacheTTLSeconds })
            } else {
              continue
            }
          } catch (e) {
            console.log("interlay cron error", value.vtoken, e.message);
            continue;
          }
          let btcprice;
          try {
            btcprice = await getPriceFromLumina("BTC");
          } catch (e) {
            console.log("lumina price error", "BTC", e.message);
            continue;
          }


          cache.set(
            tokenkey("interlay", value.vtoken),
            JSON.stringify({
              collateral_ratio: {
                issued_token: saved.total_issued / 1e8,
                locked_token: saved.total_backable / 1e8,
                // decimal: saved.decimal,
                ratio: saved.total_backable / saved.total_issued,
              },
              timestamp: Date.now(),
              fair_price:
                (saved.total_backable / 1e8) / (saved.total_issued / 1e8) > 1
                  ? btcprice
                  : (btcprice * saved.total_backable) / saved.total_issued,
            }),
            { EX: cacheTTLSeconds }
          );
        }
        break;
      case "bifrost":
        {
          let saved;
          try {
            saved = await getBiFrostValues(value.token);
          } catch (e) {
            console.log("bifrost cron error", value.token, e.message);
            continue;
          }
          let btcprice;
          try {
            btcprice = await getPriceFromLumina(value.token);
          } catch (e) {
            console.log("lumina price error", value.token, e.message);
            continue;
          }


          let ratio = saved.total_backable / saved.total_issued;

          let fairprice = ratio * btcprice

          let decimal = Math.pow(10, saved.decimal);

          if (value.token == "DOT") {
            decimal = 1e10
          }




          cache.set(
            tokenkey("bifrost", value.vtoken),
            JSON.stringify({
              collateral_ratio: {
                issued_token: saved.total_issued / decimal,
                locked_token: saved.total_backable / decimal,
                ratio: saved.total_backable / saved.total_issued,
                // decimal: saved.decimal,
              },
              fair_price: fairprice,
              timestamp: Date.now()
            }),
            { EX: cacheTTLSeconds }
          );
        }
        break;
      case "stDOT":
        {
          try {
            let stDOTcollateral = await getValuestdot();
            cache.set(
              tokenkey("stDOT", value.vtoken),
              JSON.stringify(stDOTcollateral),
              { EX: cacheTTLSeconds }
            );
          } catch (e) {
            console.log("stDOT cron error", e.message);
            continue;
          }
        }
        break;
      case "astar":
        {
          let astarcollateral = await getValueAstar();
          cache.set(
            tokenkey("astar", value.vtoken),
            JSON.stringify(astarcollateral),
            { EX: cacheTTLSeconds }
          );
        }
        break;
      case "stETH":
        {
          let stETHcollateral = await getValuesteth();
          cache.set(
            tokenkey("stETH", value.vtoken),
            JSON.stringify(stETHcollateral),
            { EX: cacheTTLSeconds }
          );
        }
        break;
      // case "cbETH": {
      //   let cbETHcollateral = await getValuecbeth();
      //   cache.set(
      //     tokenkey("cbETH", value.vtoken),
      //     JSON.stringify(cbETHcollateral)
      //   );
      // }
      //   break;
      case "rETH": {
        let rETHcollateral = await getValuereth();
        cache.set(
          tokenkey("rETH", value.vtoken),
          JSON.stringify(rETHcollateral),
          { EX: cacheTTLSeconds }
        );
      }
    }

    try {
      let baseAssetPrice = await getPriceFromLumina(value.token);
      await cache.set(pricekey(value.token), baseAssetPrice, { EX: cacheTTLSeconds });
    } catch (e) {
      console.log("lumina price error", value.token, e.message);
    }
  }


}

module.exports = {
  cronstart: cronstart,
};
