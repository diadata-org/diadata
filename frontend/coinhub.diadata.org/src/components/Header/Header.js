let vm = null;
export default {
  name: 'Header',
  props: {},
  data() {
  	return {
  	  logo:'',
  	};
  },
  created: function () {
    vm = this;
  },
  mounted() {
    vm.logo = require('@/assets/logo.svg');
  }
};
