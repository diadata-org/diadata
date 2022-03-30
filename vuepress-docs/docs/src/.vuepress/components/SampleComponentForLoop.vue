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
  data:MyList.split(',')
    /**
  data: function() {
    return {
      //MyList: MyInternalList.split(',')
      //bigCats: ['Mountain Lion', '	Clouded leopard', 'Eurasian lynx']
    }
  }
  */
}
</script>
