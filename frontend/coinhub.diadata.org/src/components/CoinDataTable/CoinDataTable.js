import { AtomSpinner } from 'epic-spinners';
import { EventBus } from '@/main';
import { ModelListSelect} from 'vue-search-select';
import shared from  '@/shared/shared';

let vm = null;
let allCoinData = [];
export default {
  components: {
      AtomSpinner,
      ModelListSelect
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
      selectedCurrency: '',
      selectedCoin:undefined,
      options: [],
      currencies: [],
      searcharray: [],
    };
  },
  created: function () {
    vm = this
  },
  async mounted() {
    try {
      const { currencyArray, coinsArray, searchArray } = shared.formatCoinData(await shared.fetchCoins());
      allCoinData = coinsArray;
      vm.currencies = currencyArray;
      vm.searcharray = searchArray;

      if(vm.currencies.length > 0) {
        if(localStorage.selectedCurrency) {
           vm.selectedCurrency = localStorage.selectedCurrency;
        }
        else {
          if(vm.currencies.includes("EUR")){
            vm.selectedCurrency = "EUR";
          }
          else{
            vm.selectedCurrency = "USD";
          }
        }
        
      }
      // load the currencies
      vm.switchCurrencies(vm.selectedCurrency);
       
    }
    catch (error) {
      console.log(error);
      vm.errored = true;
      vm.loading = false;
    }

  },
  methods: {
      switchCurrencies : function(selectedCurrency){
        allCoinData .forEach(function(coin,index){
            if(coin[selectedCurrency]) {
              vm.selectedCurrency = selectedCurrency;
              localStorage.selectedCurrency = selectedCurrency;
              vm.coindata = coin[selectedCurrency];
              vm.loading = false;
            }
        });
      },
      initSearch : () => {
        if(vm.options.length <=0){
          vm.options = vm.searcharray;
        }

        if(vm.selectedCoin){
          vm.$router.push({ name: 'coin-details', params: { coinSymbol: vm.selectedCoin}});
          vm.selectedCoin = undefined;
        }
      }
  },

};
