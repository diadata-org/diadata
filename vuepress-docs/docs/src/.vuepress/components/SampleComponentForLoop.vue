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
      MyList: this.MyList.split(','),
      //MyList: MyInternalList.split(',')
      //bigCats: ['Mountain Lion', '	Clouded leopard', 'Eurasian lynx']
    }
  }
 
}
</script>
