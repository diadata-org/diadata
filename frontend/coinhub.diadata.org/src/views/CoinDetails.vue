<template>
  <div>
    <Header/>
    <CoinDetails/>
    <Footer/>
  </div>
</template>

<script>
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
