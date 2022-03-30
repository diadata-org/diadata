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
    //var MyListTEST = MyList.split(",");
    var MyListTEST = ['a','b','c'];
    return {
      MyList: ['a','b','c'];
      //bigCats: ['Mountain Lion', '	Clouded leopard', 'Eurasian lynx']
    }
  }
}
</script>
