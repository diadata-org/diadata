import axios from 'axios';
import numeral from 'numeral';
const sortBy = require('lodash/sortBy');

export default {
	fetchCoins: async function(){ 
		let res = null;
		let coins = [];

		try{
		 res = await axios.get('https://api.diadata.org/v1/coins');
		 coins = res.data.Coins.reverse();
		}
		catch(error) {
		  console.log(error);
		}

		return coins;
	  
	},
	formatCoinData: function(coins) {
      const coinData = [];

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
        const coinPrice = coin.Price;
        const coinPriceFormatted = coinPrice < 1 ? '$'.concat(numeral(coin.Price).format('0.0[0000]')) : '$'.concat(numeral(coin.Price).format('0,0.00'));
        let change24 = (coin.Price - coin.PriceYesterday) / coin.PriceYesterday * 100;
        const change24Formatted = change24 !== Number.POSITIVE_INFINITY ? numeral(change24).format('0,0.00').concat('%') : 'N/A';
        change24 = change24 !== Number.POSITIVE_INFINITY ? change24 : Number.NEGATIVE_INFINITY ;
        const priceGraph = `https://api.diadata.org/v1/chart/${coin.Symbol}`;
        const volume24 = coin.VolumeYesterdayUSD;
        const volume24Formatted = '$'.concat(numeral(coin.VolumeYesterdayUSD).format('0,0'));
        const marketCapFormatted = '$'.concat(numeral(coin.Price * coin.CirculatingSupply).format('0,0'));
        const marketCap = coin.Price * coin.CirculatingSupply;
        const circulatingSupplyFormattedWithoutSymbol = numeral(coin.CirculatingSupply).format('0,0');
        const circulatingSupplyFormatted = numeral(coin.CirculatingSupply).format('0,0').concat(` (${coin.Symbol})`);
        const circulatingSupply = coin.CirculatingSupply;
        const oracle = require('@/assets/icons/oracle_icon.png');

        coinData.push({coinImage, 
                       coinSymbol, 
                       coinName,
                       coinPrice, 
                       coinPriceFormatted,
                       change24, 
                       change24Formatted,
                       priceGraph, 
                       volume24, 
                       volume24Formatted,
                       marketCap, 
                       marketCapFormatted, 
                       circulatingSupply, 
                       circulatingSupplyFormatted,
                       circulatingSupplyFormattedWithoutSymbol, 
                       oracle});
      }

      const descendingOrder = sortBy(coinData, 'marketCap').reverse();

      descendingOrder.forEach((coin,index) => {
        const rank = (index + 1);
        coin.rank = rank;
      });
      return descendingOrder;
  	},

}
