<template>
  <div>
    <li v-for="InternalArrayItem in MyInternalList" :key="InternalArrayItem">
      {{ InternalArrayItem }}
    </li>
  </div>
</template>
<script>
export default {
  props: ['MyList'],
  data: function() {
    return {
      MyInternalList: ['asdfasdf','qwerqwer'],
      //MyList: this.MyList.split(','),
      //MyList: MyInternalList.split(',')
      //bigCats: ['Mountain Lion', '	Clouded leopard', 'Eurasian lynx']
    }
  }
 
}
</script>
