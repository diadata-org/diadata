<template>
  <div class="home">
    <Header/>
    <CoinDataTable/>
    <Footer/>
  </div>
</template>

<script>
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
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
