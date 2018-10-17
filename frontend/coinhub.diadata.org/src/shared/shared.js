import axios from 'axios';
import numeral from 'numeral';
import moment from 'moment';
import sortBy from 'lodash/sortBy';
import getSymbolFromCurrency from 'currency-symbol-map';

const lang = navigator.language;
//numeral.locale(lang);

export default {
	fetchCoins: async function(){ 
		let res = null;
		let coins = [];

		try{
		 res = await axios.get('https://api.diadata.org/v1/coins');
		 coins = res.data;
		}
		catch(error) {
		  console.log(error);
		}

		return coins;
	  
	},
	formatCoinData: function(coins, change) {
      let coinsArray = [];
      let searchArray = [];
      let currencyArray = [];
      // init the other indexes with empty arrays for storing other currencies
      change.USD.forEach((el,index) => {
         coinsArray[index] = [];
      });
      // last index is for storing usd
      coinsArray[change.USD.length] = [];
      currencyArray[change.USD.length] = "USD";

      for (let i = 0; i < coins.length; i++) {

        const coin = coins[i];
        let coinImage = '';
        try {
    		  coinImage = require(`@/assets/icons/${coin.Symbol.toLowerCase()}.png`);
    		}
    		catch (e) {
    		  coinImage = require('@/assets/icons/crypto.png');
    		}
        const coinSymbol = coin.Symbol;
        const coinName = coin.Name;
        // Price graph
        const priceGraph = `https://api.diadata.org/v1/chart/${coin.Symbol}`;
        // coin price
        // USD
        const coinPriceUSD = coin.Price;
        const coinPriceUSDFormatted =  this.formatCurrency(coinPriceUSD, "USD");
        const coinPriceYesterdayUSD = coin.PriceYesterday;
        // circulating supply
        const circulatingSupply = coin.CirculatingSupply;
        const circulatingSupplyFormattedWithoutSymbol = this.formatCirculatingSupply(circulatingSupply, undefined);
        const circulatingSupplyFormatted = this.formatCirculatingSupply(circulatingSupply, coinSymbol);
        // change 24
        // USD
        let change24USD = (coinPriceUSD  - coinPriceYesterdayUSD ) / coinPriceYesterdayUSD * 100;
        const change24USDFormatted = this.formatChange24(change24USD);
        change24USD = change24USD !== Number.POSITIVE_INFINITY ? change24USD : Number.NEGATIVE_INFINITY ;
         // volume24
        // USD
        const volume24USD = coin.VolumeYesterdayUSD;
        const volume24USDFormatted = this.formatMarketCapAndVolume24(volume24USD , "USD");
         // market cap
        // USD
     
        const marketCapUSD = coinPriceUSD * circulatingSupply;
        const marketCapUSDFormatted = this.formatMarketCapAndVolume24(marketCapUSD , "USD");
        
        const oracle = require('@/assets/icons/oracle_icon.png');
        
        coinsArray[change.USD.length].push({coinImage, 
                       coinSymbol, 
                       coinName,
                       coinPrice:coinPriceUSD, 
                       coinPriceFormatted:coinPriceUSDFormatted,
                       change24:change24USD, 
                       change24Formatted:change24USDFormatted,
                       priceGraph, 
                       volume24:volume24USD, 
                       volume24Formatted:volume24USDFormatted,
                       marketCap:marketCapUSD, 
                       marketCapFormatted:marketCapUSDFormatted, 
                       circulatingSupply, 
                       circulatingSupplyFormatted,
                       circulatingSupplyFormattedWithoutSymbol, 
                       oracle});

        // calculate the values for the other currencies as well
        for( let j = 0; j < change.USD.length; j++){
          // populate the currency array
          const currencySwiftCode = change.USD[j].Symbol.toUpperCase();
          currencyArray[j] = currencySwiftCode;
                  
          // coin price
          const coinPriceOtherCurrency = this.calculateCurrencyFromRate(coin.Price,change.USD,currencySwiftCode,"today");
          const coinPriceOtherCurrencyFormatted = this.formatCurrency(coinPriceOtherCurrency, currencySwiftCode);
          const coinPriceOtherCurrencyYesterDay = this.calculateCurrencyFromRate(coin.PriceYesterday,change.USD,currencySwiftCode,"yesterday");
          // change 24
          let change24OtherCurrency = (coinPriceOtherCurrency  - coinPriceOtherCurrencyYesterDay ) / coinPriceOtherCurrencyYesterDay * 100;
          const change24OtherCurrencyFormatted = this.formatChange24(change24OtherCurrency);
          change24OtherCurrency = change24OtherCurrency !== Number.POSITIVE_INFINITY ? change24OtherCurrency : Number.NEGATIVE_INFINITY ;
           //  volume 24
          const volume24OtherCurrency  = this.calculateCurrencyFromRate(coin.VolumeYesterdayUSD,change.USD,currencySwiftCode,"yesterday");
          const volume24OtherCurrencyFormatted = this.formatMarketCapAndVolume24(volume24OtherCurrency, currencySwiftCode);
          //  marketCap
          const marketCapOtherCurrency = coinPriceOtherCurrency * circulatingSupply;
          const marketCapOtherCurrencyFormatted = this.formatMarketCapAndVolume24(marketCapOtherCurrency, currencySwiftCode);
          

          //  add the currency to the coins array 
          coinsArray[j].push({coinImage, 
                 coinSymbol, 
                 coinName,
                 coinPrice:coinPriceOtherCurrency, 
                 coinPriceFormatted:coinPriceOtherCurrencyFormatted,
                 change24:change24OtherCurrency, 
                 change24Formatted:change24OtherCurrencyFormatted,
                 priceGraph, 
                 volume24:volume24OtherCurrency, 
                 volume24Formatted:volume24OtherCurrencyFormatted,
                 marketCap:marketCapOtherCurrency, 
                 marketCapFormatted:marketCapOtherCurrencyFormatted, 
                 circulatingSupply, 
                 circulatingSupplyFormatted,
                 circulatingSupplyFormattedWithoutSymbol, 
                 oracle});
        }
       
      }
      // reorder the arrays by market cap and add the rank field
      coinsArray.forEach((el,i) => {
        // re-order by market cap
        coinsArray[i] = sortBy(coinsArray[i], 'marketCap').reverse();
        // add rank
        coinsArray[i].forEach((coin,j) => {
          const rank = (j + 1);
          coin.rank = rank;
          if(i === 0){
            //populate the search array
            searchArray.push( { value: coin.coinSymbol, text: coin.coinSymbol + ' : ' + coin.coinName, });
          }    
        });

        let coinsObj = {};
        const key = currencyArray[i];
        const value  = coinsArray[i];
        coinsObj[key] = value;
        // add currency key
        coinsArray[i] = coinsObj;
      });

      
      return {coinsArray, currencyArray, searchArray};
  },
  calculateCurrencyFromRate : function(currencyValue, rateArray, currencySwiftCode, rateOption) {
    return currencyValue / this.getRate(rateArray, currencySwiftCode, rateOption);
  },
  formatCurrency : function(currency,currencySwiftCode) {
    const symbol = this.getCurrencySymbol(currencySwiftCode);
    return currency < 1 ? symbol.concat(numeral(currency).format('0.0[0000]')) : symbol.concat(numeral(currency).format('0,0.00'));
  },

  formatChange24 : (change24) => {
    return change24 !== Number.POSITIVE_INFINITY ? numeral(change24).format('0,0.00').concat('%') : 'N/A';
  },

  formatMarketCapAndVolume24: function(param,currencySwiftCode) {
    const symbol = this.getCurrencySymbol(currencySwiftCode);
    return symbol.concat(numeral(param).format('0,0'));
  },

  formatCirculatingSupply: function(circulatingSupply,coinSymbol) {
    const cs = coinSymbol !== undefined ? numeral(circulatingSupply).format('0,0').concat(` (${coinSymbol})`) : numeral(circulatingSupply).format('0,0');
    return cs;
  },

  formatDateTime : (dateTime,dateTimeFormat) => {
      return moment(dateTime).format(dateTimeFormat);
  },

  getCurrencySymbol: (currencySwiftCode) => {
    let symbol = getSymbolFromCurrency(currencySwiftCode);
    if(symbol == undefined){
      symbol = currencySwiftCode;
    }
    return symbol;
  },
  getRate : (rateArray, currencySwiftCode, rateOption) => {
    let rate = 1;
    const rateObj = rateArray.filter((obj) => obj.Symbol === currencySwiftCode)[0];

    if(rateObj != undefined){
      rate = rateOption === "today" ? rateObj.Rate : rateObj.RateYesterday;
    }
    
    return rate;
  },


}
