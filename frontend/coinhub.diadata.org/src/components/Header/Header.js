export default {
  name: 'Header',
  props: {},
  data() {
  	return {
  	  logo:'',
      selectedCurrency: ''
  	};
  },
  mounted() {
  	this.logo = require('@/assets/dia-logo.png');
  }
};
