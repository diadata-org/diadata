import axios from 'axios';
import router from '@/router';
import { AtomSpinner } from 'epic-spinners';
import sortBy from 'lodash/sortBy';
import { EventBus } from '@/main';
import shared from  '@/shared/shared';


export default {
  components: {
    AtomSpinner
  },
  name: 'CoinDetails',
  props: {},
  data() {
    return {
      exchange_fields: [
        { key: 'Name', label: 'Exchange', sortable: true },
        { key: 'PriceFormatted', label: 'Price', sortable: true },
        { key: 'Volume24', label: 'Volume (24h)', sortable: true },
        { key: 'TimeFormatted', label: 'Last Updated', sortable: true },
        { key: 'show_trades', label: 'Trades', sortable: false },
      ],
      exchanges: [],
      last_trade_fields: [
        { key: 'Pair', label: 'Pair', sortable: true },
        { key: 'PriceFormatted', label: 'Price', sortable: true },
        { key: 'Volume', label: 'Volume', sortable: true },
        { key: 'TimeFormatted', label: 'Last Updated', sortable: true },
        { key: 'EstimatedPrice', label: 'Estimated Price', sortable: true },
      ],
      loading: true,
      errored: false,
      coinDetails:{},
      coinSymbol: '',
      coindata:null,
      selectedCurrency:'',
    };
  },
  created() {
    this.coinSymbol = this.$route.params.coinSymbol;

    EventBus.$emit('hideSearchInput', true);
    EventBus.$on('currencyChange', this.formatPairData);
  },
  beforeDestroy: function () {
    EventBus.$off('currencyChange', this.formatPairData);
  },
  mounted() {
    if(this.$route.params.coinRank) {
      localStorage.rank = this.$route.params.coinRank;
    }

    this.fetchCoinDetails();
  
  },
  methods: {
  	formatPairData() {
      this.loading = true;
      if(localStorage.selectedCurrency) {
         this.selectedCurrency = localStorage.selectedCurrency;
      }
      else{
        this.selectedCurrency = "EUR";
      }

      let {Coin, Change, Exchanges } = this.coindata;
      // format the coin details
      const coinPrice = shared.calculateCurrencyFromRate(Coin.Price,Change.USD,this.selectedCurrency,"today");
      const coinPriceFormatted = shared.formatCurrency(coinPrice,this.selectedCurrency);
      const coinPriceYesterday = shared.calculateCurrencyFromRate(Coin.PriceYesterday,Change.USD,this.selectedCurrency,"yesterday");
      const change24 = (coinPrice  - coinPriceYesterday) / coinPriceYesterday * 100;
      const change24Formatted = shared.formatChange24(change24);
      const volume24Formatted = shared.formatMarketCapAndVolume24(shared.calculateCurrencyFromRate(Coin.VolumeYesterdayUSD,Change.USD,this.selectedCurrency,"yesterday"),this.selectedCurrency);
      const circulatingSupplyFormattedWithoutSymbol = shared.formatCirculatingSupply(Coin.CirculatingSupply, undefined);
      this.coinDetails = { 
          coinName: Coin.Name,
          coinSymbol: Coin.Symbol,
          coinPriceFormatted,
          change24,
          change24Formatted,
          rank: localStorage.rank,
          volume24Formatted,
          circulatingSupplyFormattedWithoutSymbol,
      };

      // format the exchanges
      Exchanges.forEach((exchange)=>{
        exchange.PriceFormatted = shared.formatCurrency(shared.calculateCurrencyFromRate(exchange.Price,Change.USD,this.selectedCurrency,"today"),this.selectedCurrency),
        exchange.Volume24 = shared.formatMarketCapAndVolume24(shared.calculateCurrencyFromRate(exchange.VolumeYesterdayUSD,Change.USD,this.selectedCurrency,"yesterday"),this.selectedCurrency),
        exchange.TimeFormatted = shared.formatDateTime(exchange.Time,"dddd, MMMM Do YYYY, h:mm:ss a");

        // format the last trades too
        exchange.LastTrades.forEach((lastTrade) => {
            lastTrade.PriceFormatted = shared.formatCurrency(shared.calculateCurrencyFromRate(lastTrade.Price,Change.USD,this.selectedCurrency,"today"),this.selectedCurrency);
            lastTrade.EstimatedPrice = shared.formatCurrency(shared.calculateCurrencyFromRate(lastTrade.EstimatedUSDPrice,Change.USD,this.selectedCurrency,"today"),this.selectedCurrency);
            lastTrade.TimeFormatted = shared.formatDateTime(lastTrade.Time,"dddd, MMMM Do YYYY, h:mm:ss a");
        });
      });

      Exchanges = sortBy(Exchanges, 'VolumeYesterdayUSD').reverse();

      this.exchanges = Exchanges;
      this.loading = false;
  	},

    async fetchCoinDetails() {
      try {
        const response = await axios.get(`https://api.diadata.org/v1/symbol/${this.coinSymbol.toUpperCase()}`);
        this.coindata = response.data;
        EventBus.$emit('currencyChange');
      }
      catch (error) {
        console.log(error);
        this.errored = true;
      }
    },

    async fetchCoinChartDetails() {

    }


  },
};
