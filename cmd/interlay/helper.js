

async function collateralCurrencies(api) {
    const collateralCurrencies =
      await api.query.vaultRegistry.systemCollateralCeiling.entries();
    const collateralCurrenciesmap = new Map();
    collateralCurrencies.map((collateralCurrencie) => {
      let key = collateralCurrencie[0].toHuman();
      let value = collateralCurrencie[1].toHuman();
      // TODO filter only ibtc as wrapped {"collateral":{"Token":"DOT"},"wrapped":{"Token":"IBTC"}}
      // console.log("key",key[0].collateral.Token);
      // console.log("value",JSON.stringify(value));
  
      collateralCurrenciesmap.set(key[0].collateral.Token, value);
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
      let value = totalUserVaultCollateral[1].toHuman();
  
      totalUserVaultCollateralsmap.set(key[0].collateral.Token, value);
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
        console.log("key,", key);
        console.log("value,", value);
        oracleaggregratemap.set(key[0].ExchangeRate.Token, value);
      }
    });
    return oracleaggregratemap;
  }
  
  async function totalIssuance(api) {
    const totalIssuances = await api.query.tokens.totalIssuance.entries();
    const totalIssuancesemap = new Map();
  
    totalIssuances.map((totalIssuance) => {
      let key = totalIssuance[0].toHuman();
  
      let value = totalIssuance[1].toHuman();
  
      console.log("key", key);
      console.log("value", value);
      totalIssuancesemap.set(key[0].Token, value);
    });
    return totalIssuancesemap;
  }

 module.exports = {
    totalIssuance:totalIssuance,
    oracleaggregrate:oracleaggregrate,
    totalUserVaultCollateral:totalUserVaultCollateral,
    collateralCurrencies:collateralCurrencies
 }