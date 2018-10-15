export default {
  name: 'Header',
  props: {},
  data() {
  	return {
  	  logo:'',
      selectedCurrency: '',
      showSearch: false,
  	};
  },
  mounted() {
  	this.logo = require('@/assets/logo.svg');
  },
  methods: {
    showHideSearch : function() {
      this.showSearch = !this.showSearch;
    }
  }
};
