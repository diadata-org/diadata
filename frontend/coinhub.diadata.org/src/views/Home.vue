<template>
  <div>
    <Header/>
    <CoinDataTable/>
    <Footer/>
  </div>
</template>

<script>
import Header from '@/components/Header/Header.vue';
import CoinDataTable from '@/components/CoinDataTable/CoinDataTable.vue';
import Footer from '@/components/Footer/Footer.vue';

export default {
  name: 'home',
  components: {
    Header,
    CoinDataTable,
    Footer,
  },
};
</script>
