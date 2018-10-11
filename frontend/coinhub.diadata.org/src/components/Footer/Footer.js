import moment from 'moment';

export default {
  name: 'Footer',
  props: {},
  data() {
    return {
      currentDate: null,
    };
  },
  mounted() {
  	this.currentDate = moment().format('dddd, MMMM Do YYYY');
  }
};
