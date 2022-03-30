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
    },
    parseAndReturnAssetList( assetList ) {
      var tmpList = assetList.split(',');
      var returnList = new Array();
      for ( var i=0; i<tmpList.length; i++ ) {
        returnList.push( this.getContentByAssetId( tmpList[i] ) );
      }
      return returnList;
    }
  },
  data: function() {
    return {
      MyInternalList: this.parseAndReturnAssetList( this.AssetList ),
    }
  }
 
}
</script>
