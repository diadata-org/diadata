import shared from  '@/shared/shared';
import { AtomSpinner } from 'epic-spinners';

let coinData = {};

export default {
  components: {
      AtomSpinner
  },
  name: 'CoinDataTable',
  props: {},
  data() {
    return {
      fields: [{ key: 'rank', label: '#', sortable: true },
        { key: 'coinName', label: 'Name', sortable: true },
        { key: 'coinPrice', label: 'Price', sortable: true },
        { key: 'change24', label: 'Change (24h)', sortable: true },
        { key: 'priceGraph', label: 'Price Graph (7d)', sortable: false },
        { key: 'volume24', label: 'Volume (24h)', sortable: true },
        { key: 'marketCap', label: 'Market Cap', sortable: true },
        { key: 'circulatingSupply', label: 'Circulating Supply', sortable: true },
        { key: 'oracle', label: 'Oracle', sortable: false },
      ],
      coindata: [],
      loading: true,
      errored: false,
      selectedCurrency: ''
    };
  },
  async mounted() {
    try {
       const { Coins, Change } = await shared.fetchCoins();
       coinData = shared.formatCoinData(Coins, Change);
       const { coinDataUSD, coinDataEUR } = coinData;
       this.coindata = coinDataUSD;
       this.selectedCurrency = "USD";
       this.$nextTick( () => this.loading = false);
    }
    catch (error) {
      console.log(error);
      this.errored = true;
    }
  },
  methods: {
      switchCurrencies : function(currency){
        const { coinDataUSD, coinDataEUR } = coinData;
        if(currency === 'EUR'){
          this.coindata = coinDataEUR;
        }

        if(currency === 'USD'){
          this.coindata = coinDataUSD;
        }

        this.selectedCurrency = currency;

      }
  },

};
