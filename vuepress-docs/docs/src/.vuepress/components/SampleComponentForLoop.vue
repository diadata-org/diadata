<template>
  <div>
    <li v-for="InternalArrayKey in InternalArrayName" :key="InternalArrayKey">
      {{ InternalArrayKey }}
    </li>
  </div>
</template>
<script>
export default {
  props: ['MyList'],
  data: function() {
    var testArray = MyList.split(",");
    return {
      InternalArrayName: testArray
      //bigCats: ['Mountain Lion', '	Clouded leopard', 'Eurasian lynx']
    }
  }
}
</script>
