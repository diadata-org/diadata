<template>
  <div class="DIAComponentAssetListWithMarkdownContent">
    <slot />
    <li v-for="InternalArrayItem in MyInternalList" :key="InternalArrayItem">
      {{ InternalArrayItem }}
    </li>
  </div>
</template>
<script>
var assetData = new Array();
assetData['BTC'] = "bitcoin";
assetData['ETH'] = "ethereum";
assetData['DIA'] = "DIAdata";
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
