<template>
  <div>
    <li v-for="InternalArrayItem in MyList" :key="InternalArrayItem">
      {{ InternalArrayItem }}
    </li>
  </div>
</template>
<script>
export default {
  props: ['MyList'],
  data: function() {
    return {
      MyList: ['a','b','c']
      //bigCats: ['Mountain Lion', '	Clouded leopard', 'Eurasian lynx']
    }
  }
}
</script>
