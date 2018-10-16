import { ModelListSelect} from 'vue-search-select';
import shared from  '@/shared/shared';
import { EventBus } from '@/main';

let vm = null;
export default {
  components: {
    ModelListSelect
  },
  name: 'Header',
  props: {},
  data() {
  	return {
  	  logo:'',
      selectedCurrency: '',
      selectedCoin:undefined,
      showSearch: false,
      options: [],
      currencies: [],
      coindata: [],
      searcharray: [],
      hideSearchInput: false,
      item: {
        value: '',
        text: ''
      }
  	};
  },
  created: function () {
    vm = this;
    EventBus.$on('hideSearchInput', vm.hideSrchInput);
  },
  beforeDestroy: function () {
    EventBus.$off('hideSearchInput', vm.hideSrchInput);
  },
  async mounted() {
    
    try {
      vm.logo = require('@/assets/logo.svg');
      const { Coins, Change } = await shared.fetchCoins();
      const coinData = shared.formatCoinData(Coins, Change);
      
      vm.currencies = coinData.currencyArray;
      vm.coindata = coinData.coinsArray;
      vm.searcharray = coinData.searchArray;

      if(vm.currencies.length > 0) {
        if(vm.currencies.includes("EUR")){
          vm.selectedCurrency = "EUR";
        }
        else{
          vm.selectedCurrency = this.currencies[0];
        }
      }
      // load the currencies
      vm.switchCurrencies(vm.selectedCurrency);
       
    }
    catch (error) {
      console.log(error);
      EventBus.$emit('coinDataError', error);
    }
  },
  methods: {
    showHideSearch : function() {
      vm.showSearch = !vm.showSearch;
    },
    switchCurrencies : function(selectedCurrency){
      vm.coindata.forEach(function(coin,index){
          if(coin[selectedCurrency]) {
            vm.selectedCurrency = selectedCurrency;
            EventBus.$emit('coinData', coin[selectedCurrency]);
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
    },
    hideSrchInput: (val) =>{
        vm.hideSearchInput = val;
    }
  }
};
