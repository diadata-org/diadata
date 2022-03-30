<template>
  <div>
    <li v-for="bigCat in bigCats" :key="bigCat">
      {{ bigCat }}
    </li>
  </div>
</template>
<script>
export default {
  data: function() {
    return {
      bigCats: ['Mountain Lion', '	Clouded leopard', 'Eurasian lynx']
    }
  }
}
</script>
