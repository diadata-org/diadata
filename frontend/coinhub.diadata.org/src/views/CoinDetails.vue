<template>
  <div class="coin-details">
    <Header/>
    <CoinDetails/>
    <Footer/>
  </div>
</template>

<script>
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import Header from '@/components/Header/Header.vue';
import CoinDetails from '@/components/CoinDetails/CoinDetails.vue';
import Footer from '@/components/Footer/Footer.vue';

export default {
  name: 'home',
  components: {
    Header,
    CoinDetails,
    Footer,
  },
};
</script>
