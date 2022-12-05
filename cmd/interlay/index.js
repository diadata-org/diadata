const { ApiPromise, WsProvider } = require("@polkadot/api");
const { totalIssuance, oracleaggregrate,totalUserVaultCollateral,collateralCurrencies } = require("./helper");

const ethers = require("ethers");
const bignumber = ethers.BigNumber;

let express = require('express');
let app = express();


async function getValues() {
  const wsProvider = new WsProvider(
    "wss://interlay.api.onfinality.io/public-ws"
  );
  const api = await ApiPromise.create({
    provider: wsProvider,
  });

  let collateralCurrenciesmap = new Map();
  let totalUserVaultCollateralmap = new Map();
  let oracleaggregatormap = new Map();
  let totalIssuancesemap = new Map();

  collateralCurrenciesmap = await collateralCurrencies(api);
  console.log("collateralCurrencies", collateralCurrenciesmap);

  totalUserVaultCollateralmap = await totalUserVaultCollateral(api);
  console.log("totalUserVaultCollateral", totalUserVaultCollateralmap);

  oracleaggregatormap = await oracleaggregrate(api);
  console.log("oracleaggregrate", oracleaggregatormap);

  totalIssuancesemap = await totalIssuance(api);
  console.log("totalIssuance", totalIssuancesemap);

  /*

total_backable = collateral_currencies.map(|currency| vaultRegistry.totalUserVaultCollateral[currency] / oracle.aggregate[currency]).sum();
total_issued = tokens.totalIssuance[IBTC];
assert!(total_issued < total_backable);

*/

  let total_backable = bignumber.from(0);

  for (let [collateralCurrency, value] of collateralCurrenciesmap) {
    console.log("collateralCurrency", collateralCurrency);

    let totalUserVaultCollateralcurrency = bignumber.from(
      totalUserVaultCollateralmap.get(collateralCurrency).replaceAll(",", "")
    );
    let oracleaggregatecurrency = bignumber.from(
      oracleaggregatormap.get(collateralCurrency).replaceAll(",", "")
    );
    console.log(
      "totalUserVaultCollateralcurrenct",
      totalUserVaultCollateralcurrency.toString()
    );

    oracleaggregatecurrency = oracleaggregatecurrency.div(1e12);
    oracleaggregatecurrency = oracleaggregatecurrency.div(1e6); //TODO while doing 1e18 its crashing somehow.

    let currentcurrencydecimal = 10;
    let btcdecimal = 8;
    // TODO  oracleaggregatecurrency/currentcurrencydecimal ie  (10^10 / 10^8)
    oracleaggregatecurrency = oracleaggregatecurrency.div(1e2);

    // totalUserVaultCollateralcurrency =
    //   totalUserVaultCollateralcurrency.div(1e10);

    total_backable = total_backable.add(
      totalUserVaultCollateralcurrency.div(oracleaggregatecurrency)
    );
  }

  let total_issued = bignumber.from(
    totalIssuancesemap.get("IBTC").replaceAll(",", "")
  );
  // total_issued = total_issued.div(1e8);

  await api.disconnect();
  total_backable = total_backable.div(1e2);
  return {total_backable:total_backable.toString(),total_issued:total_issued.toString(), decimal:8}

  // return total_backable,total_issued;

  // console.log("total_backable", total_backable.toString());
  // console.log("total_issued", total_issued.toString());

}

app.get('/customer/interlay/state', async function (req, res) {
  let values =  await getValues();
  res.send(values);
})

let server = app.listen(8081, function () {
  var host = server.address().address
  var port = server.address().port
  console.log(" App listening at http://%s:%s", host, port)
})


