import axios from 'axios';
import numeral from 'numeral';
import router from '@/router';
import shared from  '@/shared/shared';


export default {
  name: 'CoinDetails',
  props: {},
  data() {
    return {
      fields: [],
      pairs: [{"exchange": "Kraken", "pair": "BTCUSD", "price": 1337.33, "volume": 12345}, 
              {"exchange": "Binance", "pair": "BTCETH", "price": 345.21, "volume": 7788}],
      loading: true,
      errored: false,
      coinDetails:{},
      coinSymbol: '',
    };
  },
  created() {
    this.coinDetails = this.$route.params.coinDetails;
    this.coinSymbol = this.$route.params.coinSymbol;
  },
  async mounted() {
    
    if(this.coinDetails == undefined){
      const coins = await shared.fetchCoins();
      const coinData = shared.formatCoinData(coins);
      this.coinDetails = coinData.filter((coin) => coin.coinSymbol === this.coinSymbol)[0];
    }
    axios
    .get(`https://api.diadata.org/v1/pairs/${this.coinSymbol}`)
    .then((response) => {
      this.formatPairData(response.data.Pairs);
    })
    .catch((error) => {
      console.log(error);
      // since the api doesn't return anything for now, don't show any error
      //this.errored = true;
      this.loading = false;
    });
  },
  methods: {
  	formatPairData(pairs) {

      if(pairs !== undefined){
        if (pairs.length >0 ) {
          this.pairs = pairs;
        }
      }
      
      this.loading = false;
  	},
  },
};
