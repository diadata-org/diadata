
const { ApiPromise, WsProvider } = require("@polkadot/api");
const ethers = require("ethers");
const bignumber = ethers.BigNumber;
const {createResponse,getPrice} = require("./utils")

// for exchange rate decoding
const FIXEDI128_SCALING_FACTOR = 18;

async function collateralCurrencies(api) {
    const collateralCurrencies =
      await api.query.vaultRegistry.systemCollateralCeiling.entries();
    const collateralCurrenciesmap = new Map();
    collateralCurrencies.map((collateralCurrencie) => {
      let key = collateralCurrencie[0].toHuman();
      let value = collateralCurrencie[1].toHuman();
      // TODO filter only ibtc as wrapped {"collateral":{"Token":"DOT"},"wrapped":{"Token":"IBTC"}}
      console.log("key",JSON.stringify(key[0].collateral));
      // console.log("value",JSON.stringify(value));
      
  
      collateralCurrenciesmap.set(JSON.stringify(key[0].collateral), value);
    });
    return collateralCurrenciesmap;
  }
  
  //https://github.com/DefiLlama/DefiLlama-Adapters/blob/f0bcc4e65ebe43ce00c77951376a2249f2931dc3/projects/interlay-collateral/api.js
  
  async function totalUserVaultCollateral(api) {
    const totalUserVaultCollaterals =
      await api.query.vaultRegistry.totalUserVaultCollateral.entries();
    const totalUserVaultCollateralsmap = new Map();
  
    totalUserVaultCollaterals.map((totalUserVaultCollateral) => {

       let key = totalUserVaultCollateral[0].toHuman();
      console.log("totalUserVaultCollaterals----key-----",key)

      let value = totalUserVaultCollateral[1].toHuman();
      console.log("totalUserVaultCollaterals-----value----",value)

      totalUserVaultCollateralsmap.set(JSON.stringify(key[0].collateral), value);
    });
    return totalUserVaultCollateralsmap;
  }
  
  async function oracleaggregrate(api) {
    const oracleaggregrators = await api.query.oracle.aggregate.entries();
  
    const oracleaggregratemap = new Map();
  
    oracleaggregrators.map((oracleaggregrate) => {
      let key = oracleaggregrate[0].toHuman();
      let value = oracleaggregrate[1].toHuman();
  
      if (key[0].ExchangeRate) {
        oracleaggregratemap.set(JSON.stringify(key[0].ExchangeRate), value);
      }
    });
    return oracleaggregratemap;
  }
  
  async function lending(api,oracleaggregatormap){
    const data = (await api.query.loans.markets.entries());

    for (let i = 0; i < data.length; i++) {
      const [underlyingId, marketData] = data[i];
  
      const isActive = marketData?.toJSON().state == "Active";
      console.log("---------")

      console.log("lending",isActive)
      console.log("underlyingId",i)
      console.log("---------")


      if (!isActive) {
        continue;
      }
  
      const lendTokenId = marketData.toJSON().lendTokenId;
      
      const [issuanceExchangeRate, totalIssuance, totalBorrows]  = await Promise.all([
        api.query.loans.exchangeRate(underlyingId.toHuman()[0]).then((rawExchangeRate) => 
          parseInt(rawExchangeRate)  / (10 ** FIXEDI128_SCALING_FACTOR)
        ),
        api.query.tokens.totalIssuance(lendTokenId),
        api.query.loans.totalBorrows(underlyingId.toHuman()[0]),
      ]);
  
      const totalTvl = (totalIssuance * issuanceExchangeRate) - totalBorrows;

      console.log("--------")

      let oac = oracleaggregatormap.get(JSON.stringify(underlyingId.toHuman()[0]))
      try{

      
      if(oac){

        const data = lendTokenId;
        const matchData = {};

        matchData.LendToken = data.lendToken.toString(); // converting 4 to string
         
        oracleaggregatormap.set(JSON.stringify(matchData),oac)

      }
    } catch(e){
        console.log("err",e)
        console.log(lendTokenId)
      }



      console.log("underlyingId",underlyingId.toHuman()[0])
      console.log("underlyingId o",oac)


      console.log("lendTokenId",lendTokenId)
      console.log("totalIssuance",totalIssuance.toString())
      console.log("issuanceExchangeRate",issuanceExchangeRate)



      console.log("totalTvl",totalTvl)
    }

return oracleaggregatormap
  }
  async function totalIssuance(api) {
    const totalIssuances = await api.query.tokens.totalIssuance.entries();
    const totalIssuancesemap = new Map();
  
    totalIssuances.map((totalIssuance) => {
      let key = totalIssuance[0].toHuman();
  
      let value = totalIssuance[1].toHuman();
      totalIssuancesemap.set(JSON.stringify(key[0]), value);
    });
    return totalIssuancesemap;
  }
  
  async function getInterlayValues(token) {
     let providerurl = "";
  
    switch (token) {
      case "IBTC":
        providerurl = process.env.INTERLAY_NODE_URL || "wss://interlay-rpc.dwellir.com";
        break;
      case "KBTC":
        providerurl = process.env.KITSUNGI_NODE_URL || "wss://kintsugi.api.onfinality.io/public-ws";
        break;
    }
    const wsProvider = new WsProvider(
       providerurl
    );

    console.log("getinterlay api")
    let api;
    try{
       api = await ApiPromise.create({
      provider: wsProvider,
      throwOnConnect: true,
      throwOnUnknown:true
    });
  }catch(e){
    console.log("throw getinterlay api")

throw e
  }

  
    let collateralCurrenciesmap = new Map();
    let totalUserVaultCollateralmap = new Map();
    let oracleaggregatormap = new Map();
    let totalIssuancesemap = new Map();
  
    collateralCurrenciesmap = await collateralCurrencies(api);
    console.log("collateralCurrencies", collateralCurrenciesmap);
  
    totalUserVaultCollateralmap = await totalUserVaultCollateral(api);
    console.log("totalUserVaultCollateral", totalUserVaultCollateralmap);
  
    oracleaggregatormap = await oracleaggregrate(api);

    oracleaggregatormap = await lending(api,oracleaggregatormap)

    console.log("oracleaggregrate", oracleaggregatormap);
  
    totalIssuancesemap = await totalIssuance(api);
    console.log("totalIssuance", totalIssuancesemap);
  
    /*
  
  total_backable = collateral_currencies.map(|currency| vaultRegistry.totalUserVaultCollateral[currency] / oracle.aggregate[currency]).sum();
  total_issued = tokens.totalIssuance[IBTC];
  assert!(total_issued < total_backable);
  
  */
  
    let total_backable = bignumber.from(0);

    let consideredTokens = [];
  
    for (let [collateralCurrency,value] of collateralCurrenciesmap) {
      console.log("collateralCurrency", collateralCurrency);

      let  collateralCurrencyString = totalUserVaultCollateralmap.get(collateralCurrency)

      console.log("collateralCurrencycollateralCurrencyString", collateralCurrency);


      let totalUserVaultCollateralcurrency;
      if(collateralCurrencyString){
        totalUserVaultCollateralcurrency = bignumber.from(
          totalUserVaultCollateralmap.get(collateralCurrency).replaceAll(",", "")
        );
      }else{
        console.log("collateralCurrency not found in totalUserVaultCollateralmap", collateralCurrency);
          
      }

      
  
     let oac = oracleaggregatormap.get(collateralCurrency)
     let oracleaggregatecurrency;

     if(oac){

      console.log("oac of",collateralCurrency)
 
      oracleaggregatecurrency = bignumber.from(
        oracleaggregatormap.get(collateralCurrency).replaceAll(",", "")
      );
     }else{
      console.log("oac not found for", collateralCurrency);

       continue
     }
     
        
     
      // console.log(
      //   "totalUserVaultCollateralcurrenct",
      //   totalUserVaultCollateralcurrency.toString()
      // );
  
      oracleaggregatecurrency = oracleaggregatecurrency.div(1e12);
      oracleaggregatecurrency = oracleaggregatecurrency.div(1e6); //TODO while doing 1e18 its crashing somehow.
  
      let currentcurrencydecimal = 10;
      let btcdecimal = 8;
      // TODO  oracleaggregatecurrency/currentcurrencydecimal ie  (10^10 / 10^8)
      oracleaggregatecurrency = oracleaggregatecurrency.div(1e2);
  
      // totalUserVaultCollateralcurrency =
      //   totalUserVaultCollateralcurrency.div(1e10);

  
      if(totalUserVaultCollateralcurrency){


     
        let by = totalUserVaultCollateralcurrency.div(oracleaggregatecurrency).toString()
        consideredTokens.push({collateralCurrency,by})


         total_backable = total_backable.add(
          by
        );
       

      }

      console.log("total_backable",total_backable.toString())

      
    }
  
 
    let t = {}
    t.Token = token
    let total_issued = bignumber.from(
      totalIssuancesemap.get(JSON.stringify(t)).replaceAll(",", "")
    );

     // total_issued = total_issued.div(1e8);
  
    await api.disconnect();
    total_backable = total_backable.div(1e2);


    return {
      total_backable: total_backable.toString(),
      total_issued: total_issued.toString(),
      decimal: 8,
      token: token,
      consideredTokens,
    };
  
    // return total_backable,total_issued;
  
    // console.log("total_backable", total_backable.toString());
    // console.log("total_issued", total_issued.toString());
  }

  
 module.exports = {
    totalIssuance:totalIssuance,
    oracleaggregrate:oracleaggregrate,
    totalUserVaultCollateral:totalUserVaultCollateral,
    collateralCurrencies:collateralCurrencies,
    getInterlayValues:getInterlayValues,
 }