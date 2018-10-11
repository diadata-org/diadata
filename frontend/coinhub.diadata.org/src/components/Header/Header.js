export default {
  name: 'Header',
  props: {},
  data() {
  	return {
  	  logo:''
  	};
  },
  mounted() {
  	this.logo = require('@/assets/dia-logo.png');
  }
};
