import axios from 'axios';
import numeral from 'numeral';
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
        const coinPriceUSDFormatted = coinPriceUSD < 1 ? '$'.concat(numeral(coinPriceUSD).format('0.0[0000]')) : '$'.concat(numeral(coinPriceUSD).format('0,0.00'));
        const coinPriceYesterdayUSD = coin.PriceYesterday;
        // circulating supply
        const circulatingSupply = coin.CirculatingSupply;
        const circulatingSupplyFormattedWithoutSymbol = numeral(circulatingSupply).format('0,0');
        const circulatingSupplyFormatted = numeral(circulatingSupply).format('0,0').concat(` (${coin.Symbol})`);
        // change 24
        // USD
        let change24USD = (coinPriceUSD  - coinPriceYesterdayUSD ) / coinPriceYesterdayUSD * 100;
        const change24USDFormatted = change24USD !== Number.POSITIVE_INFINITY ? numeral(change24USD).format('0,0.00').concat('%') : 'N/A';
        change24USD = change24USD !== Number.POSITIVE_INFINITY ? change24USD : Number.NEGATIVE_INFINITY ;
         // volume24
        // USD
        const volume24USD = coin.VolumeYesterdayUSD;
        const volume24USDFormatted = '$'.concat(numeral(volume24USD).format('0,0'));
         // market cap
        // USD
        const marketCapUSDFormatted = '$'.concat(numeral(coinPriceUSD * circulatingSupply ).format('0,0'));
        const marketCapUSD = coinPriceUSD * circulatingSupply;
        
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
          currencyArray[j] = change.USD[j].Symbol.toUpperCase();
          // get the currency symbol
          let symbol = getSymbolFromCurrency(change.USD[j].Symbol.toUpperCase());
          if(symbol == undefined){
            symbol = change.USD[j].Symbol.toUpperCase();
          }
         
          // coin price
          const coinPriceOtherCurrency = coin.Price / change.USD[j].Rate;
          const coinPriceOtherCurrencyFormatted = coinPriceOtherCurrency < 1 ? symbol.concat(numeral(coinPriceOtherCurrency).format('0.0[0000]')) : symbol.concat(numeral(coinPriceOtherCurrency).format('0,0.00'));
          const coinPriceOtherCurrencyYesterDay = coin.PriceYesterday / change.USD[j].RateYesterday;
          // change 24
          let change24OtherCurrency = (coinPriceOtherCurrency  - coinPriceOtherCurrencyYesterDay ) / coinPriceOtherCurrencyYesterDay * 100;
          const change24OtherCurrencyFormatted = change24OtherCurrency !== Number.POSITIVE_INFINITY ? numeral(change24OtherCurrency).format('0,0.00').concat('%') : 'N/A';
          change24OtherCurrency = change24OtherCurrency !== Number.POSITIVE_INFINITY ? change24OtherCurrency : Number.NEGATIVE_INFINITY ;
           //  volume 24
          const volume24OtherCurrency  = coin.VolumeYesterdayUSD / change.USD[j].RateYesterday;
          const volume24OtherCurrencyFormatted = symbol.concat(numeral(volume24OtherCurrency).format('0,0'));
          //  marketCap
          const marketCapOtherCurrencyFormatted = symbol.concat(numeral(coinPriceOtherCurrency * circulatingSupply).format('0,0'));
          const marketCapOtherCurrency = coinPriceOtherCurrency * circulatingSupply;

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

  formatCurrent : (currency,currencySymbol) => {

  }

}
