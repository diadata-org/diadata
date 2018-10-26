import axios from 'axios';
import numeral from 'numeral';
import moment from 'moment';
import sortBy from 'lodash/sortBy';
import getSymbolFromCurrency from 'currency-symbol-map';

const lang = navigator.language;
let currencyArray = [];
//numeral.locale(lang);

 function getApi() {
     var host = window.location.hostname;
     console.log("hello"+host)
     if (host === 'coinhub.diadata.org') {
           return 'https://api.diadata.org'
        } else {
          return 'http://api.diadata.org'
        }
      }

export default {
	fetchCoins: async function(){ 
		let res = null;
		let coins = [];

		try{
		 res = await axios.get(getApi()+'/v1/coins');
		 coins = res.data;
		}
		catch(error) {
		  console.log(error);
		}

		return coins;
	  
	},
	formatCoinData: function(coindata) {
      let change = coindata.Change !== undefined && coindata.Change !== null ? coindata.Change : null;
      let coins = coindata.Coins !== undefined && coindata.Coins !== null ? coindata.Coins : [];
      let coinsList = coindata.CompleteCoinList !== undefined && coindata.CompleteCoinList !== null ? coindata.CompleteCoinList : null;

      let coinsArray = [];
      let searchArray = [];
      currencyArray = [];
      let USDindex = 0;

      if(change != undefined && change != null) {

        if(change.USD != undefined && change.USD != null) {
            // init the other indexes with empty arrays for storing other currencies
          change.USD.forEach((el,index) => {
             coinsArray[index] = [];
          });
          // last index is for storing usd
          USDindex = change.USD.length;
          
        }
      }
   
      coinsArray[USDindex] = [];
      currencyArray[USDindex] = "USD";

      for (let i = 0; i < coins.length; i++) {

        const coin = coins[i];
        let coinImage = '';
        try {
    		  coinImage = require(`cryptocurrency-icons/32/color/${coin.Symbol.toLowerCase()}.png`);
    		}
    		catch (e) {
          console.log(e);
    		  coinImage = require('@/assets/icons/crypto.png');
    		}
        const coinSymbol = coin.Symbol;
        const coinName = coin.Name;
        // Price graph
        const priceGraph = getApi()+`/v1/chart/${coin.Symbol.toUpperCase()}`;
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
        
        coinsArray[USDindex].push({coinImage, 
                       coinSymbol, 
                       coinName,
                       coinNameLowerCase:coinName.toLowerCase(),
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


        if(change != undefined && change != null) {
          if(change.USD != undefined && change.USD != null) {
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
        });
        let coinsObj = {};
        const key = currencyArray[i];
        const value  = coinsArray[i];
        coinsObj[key] = value;
        // add currency key
        coinsArray[i] = coinsObj;
      });

      coinsList.forEach((coin) =>{
          searchArray.push({ value: coin.Symbol, text: coin.Symbol + ' : ' + coin.Name, });
      });

      currencyArray = sortBy(currencyArray);
      return {coinsArray, currencyArray, searchArray};
  },
  calculateCurrencyFromRate : function(currencyValue, rateArray, currencySwiftCode, rateOption) {
    return currencyValue * this.getRate(rateArray, currencySwiftCode, rateOption);
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

  formatDateTime: (dateTime,dateTimeFormat) => {
      return moment(dateTime).format(dateTimeFormat);
  },

  getCurrencySymbol: (currencySwiftCode) => {
    let symbol = getSymbolFromCurrency(currencySwiftCode);
    if(symbol == undefined){
      symbol = currencySwiftCode;
    }
    return symbol;
  },
  getRate: (rateArray, currencySwiftCode, rateOption) => {
    let rate = 1;
    let rateObj = undefined;
 
    if(rateArray != undefined && rateArray != null) {
        rateObj = rateArray.filter((obj) => obj.Symbol === currencySwiftCode)[0];
    }

    if(rateObj != undefined){
      rate = rateOption === "today" ? rateObj.Rate : rateObj.RateYesterday;
    }
    
    return rate;
  },

  getCurrencies: (rateArray) => {
    currencyArray = [];
    rateArray.forEach((rate,index)=>{
      currencyArray[index] = rate.Symbol.toUpperCase();
    });
    currencyArray.push("USD");
    return sortBy(currencyArray);
  },
  getApi: getApi
}
