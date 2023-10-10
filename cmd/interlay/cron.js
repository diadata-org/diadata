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
  getPrice,
  pricekey,
} = require("./utils");

let cache = redis();

async function cronstart() {
  for (const value of allowedTokens) {
    switch (value.source) {
      case "interlay":
        {
          let saved
          try{
            saved = await getInterlayValues(value.vtoken);
            cache.set("interlayraw"+ value.vtoken,JSON.stringify(saved))
          }catch(e){
            saved = 0
          }
          let btcprice = await getPrice("BTC");
          cache.set(
            tokenkey("interlay", value.vtoken),
            JSON.stringify({
              collateral_ratio: {
                issued_token: saved.total_issued / 1e8,
                locked_token: saved.total_backable / 1e8,
                // decimal: saved.decimal,
                ratio: saved.total_backable / saved.total_issued,
              },
              fair_price:
                (saved.total_backable/1e8) / (saved.total_issued/1e8) > 1
                  ? btcprice
                  : (btcprice * saved.total_backable) / saved.total_issued,
            })
          );
        }
        break;
      case "bifrost":
        {
          
          let saved = await getBiFrostValues(value.token);
          let btcprice = await getPrice(value.token);

          let decimal = saved.decimal
 
          let ratio = saved.total_backable/saved.total_issued;

          let fairprice =  ratio  *btcprice


          // let decimal = 1e12;
          if(value.token=="DOT"){
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
              fair_price:
              fairprice
            })
          );
        }
        break;
      case "stDOT":
        {
          let stDOTcollateral = await getValuestdot();
          cache.set(
            tokenkey("stDOT", value.vtoken),
            JSON.stringify(stDOTcollateral)
          );
        }
        break;
      case "astar":
        {
          let astarcollateral = await getValueAstar();
          cache.set(
            tokenkey("astar", value.vtoken),
            JSON.stringify(astarcollateral)
          );
        }
        break;
      case "stETH":
        {
          let stETHcollateral = await getValuesteth();
          cache.set(
            tokenkey("stETH", value.vtoken),
            JSON.stringify(stETHcollateral)
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
          JSON.stringify(rETHcollateral)
        );
      }
    }

    let baseAssetPrice = await getPrice(value.token);
    await cache.set(pricekey(value.token), baseAssetPrice);
  }

   
}

module.exports = {
  cronstart: cronstart,
};
