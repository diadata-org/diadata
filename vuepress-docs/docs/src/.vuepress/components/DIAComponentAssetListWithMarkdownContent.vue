<template>
  <div>
    <li v-for="InternalArrayItem in MyInternalList" :key="InternalArrayItem">
      {{ InternalArrayItem }}
    </li>
  </div>
</template>
<script>
export default {
  props: ['AssetList'],
  data: function() {
    return {
      MyInternalList: this.AssetList.split(','),
    }
  }
 
}
</script>
