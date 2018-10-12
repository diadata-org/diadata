import axios from 'axios';
import numeral from 'numeral';
const sortBy = require('lodash/sortBy');

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
      let coinDataUSD = [];
      let coinDataEUR = [];

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

        // coin price
        // USD
        const coinPriceUSD = coin.Price;
        const coinPriceUSDFormatted = coinPriceUSD < 1 ? '$'.concat(numeral(coinPriceUSD).format('0.0[0000]')) : '$'.concat(numeral(coinPriceUSD).format('0,0.00'));
        const coinPriceYesterdayUSD = coin.PriceYesterday;
        // EUR
        const coinPriceEUR = coin.Price / change.EURUSD;
        const coinPriceEURFormatted = coinPriceEUR < 1 ? '€'.concat(numeral(coinPriceEUR).format('0.0[0000]')) : '€'.concat(numeral(coinPriceEUR).format('0,0.00'));
        const coinPriceYesterdayEUR = coin.PriceYesterday / change.EURUSDYesterday;

        // change 24
        // USD
        let change24USD = (coinPriceUSD  - coinPriceYesterdayUSD ) / coinPriceYesterdayUSD * 100;
        const change24USDFormatted = change24USD !== Number.POSITIVE_INFINITY ? numeral(change24USD).format('0,0.00').concat('%') : 'N/A';
        change24USD = change24USD !== Number.POSITIVE_INFINITY ? change24USD : Number.NEGATIVE_INFINITY ;
        //EUR
        let change24EUR = (coinPriceEUR  - coinPriceYesterdayEUR ) / coinPriceYesterdayEUR * 100;
        const change24EURFormatted = change24EUR !== Number.POSITIVE_INFINITY ? numeral(change24EUR).format('0,0.00').concat('%') : 'N/A';
        change24EUR = change24EUR !== Number.POSITIVE_INFINITY ? change24EUR : Number.NEGATIVE_INFINITY ;

        // Price graph
        const priceGraph = `https://api.diadata.org/v1/chart/${coin.Symbol}`;

        // volume24
        // USD
        const volume24USD = coin.VolumeYesterdayUSD;
        const volume24USDFormatted = '$'.concat(numeral(volume24USD).format('0,0'));
        // EUR
        const volume24EUR = coin.VolumeYesterdayUSD / change.EURUSDYesterday;
        const volume24EURFormatted = '€'.concat(numeral(volume24EUR).format('0,0'));

        // circulating supply
        const circulatingSupply = coin.CirculatingSupply;
        const circulatingSupplyFormattedWithoutSymbol = numeral(circulatingSupply).format('0,0');
        const circulatingSupplyFormatted = numeral(circulatingSupply).format('0,0').concat(` (${coin.Symbol})`);

        // market cap
        // USD
        const marketCapUSDFormatted = '$'.concat(numeral(coinPriceUSD * circulatingSupply ).format('0,0'));
        const marketCapUSD = coinPriceUSD * circulatingSupply;
        // EUR
        const marketCapEURFormatted = '€'.concat(numeral(coinPriceEUR * circulatingSupply ).format('0,0'));
        const marketCapEUR = coinPriceEUR * circulatingSupply;
      
        const oracle = require('@/assets/icons/oracle_icon.png');

        coinDataUSD.push({coinImage, 
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

        coinDataEUR.push({coinImage, 
                       coinSymbol, 
                       coinName,
                       coinPrice:coinPriceEUR, 
                       coinPriceFormatted:coinPriceEURFormatted,
                       change24:change24EUR, 
                       change24Formatted:change24EURFormatted,
                       priceGraph, 
                       volume24:volume24EUR, 
                       volume24Formatted:volume24EURFormatted,
                       marketCap:marketCapEUR, 
                       marketCapFormatted:marketCapEURFormatted, 
                       circulatingSupply, 
                       circulatingSupplyFormatted,
                       circulatingSupplyFormattedWithoutSymbol, 
                       oracle});
      }

      coinDataUSD = sortBy(coinDataUSD, 'marketCap').reverse();
      coinDataEUR = sortBy(coinDataEUR, 'marketCap').reverse();

      coinDataUSD.forEach((coin,index) => {
        const rank = (index + 1);
        coin.rank = rank;
      });

      coinDataEUR.forEach((coin,index) => {
        const rank = (index + 1);
        coin.rank = rank;
      });
      return {coinDataUSD,coinDataEUR};
  	},

}
