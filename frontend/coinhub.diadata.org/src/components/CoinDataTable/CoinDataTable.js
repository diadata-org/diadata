import { AtomSpinner } from 'epic-spinners';
import { EventBus } from '@/main';

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
    };
  },
  created: function () {
    EventBus.$on('coinData', this.setCoinData);
    EventBus.$on('coinDataError', this.handleCoinDataError);
  },
  beforeDestroy: function () {
    EventBus.$off('coinData', this.setCoinData);
    EventBus.$off('coinDataError', this.handleCoinDataError);
  },
  methods: {
      setCoinData : function (coinData){
         this.coindata = coinData;
         this.loading = false;
      },
      handleCoinDataError : function (error) {
        this.errored = true;
        this.loading = false;
      }
  },

};
