<template>
  <div class="DIAComponentAssetListWithMarkdownContent">
    <slot />
    <li v-for="InternalArrayItem in MyInternalList" :key="InternalArrayItem">
      {{ InternalArrayItem }}
    </li>
  </div>
</template>
<script>
export default {
  props: ['AssetList'],
  methods: {
    getContentByAssetId( assetId ) {
      return assetId;
    }
  },
  data: function() {
    return {
      MyInternalList: this.AssetList.split(','),
    }
  }
 
}
</script>
