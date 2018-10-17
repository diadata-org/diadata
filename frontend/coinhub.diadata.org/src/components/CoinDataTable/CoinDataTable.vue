<style src="./CoinDataTable.css" scoped></style>
<template>
  	<section v-if="errored" class="loading-error">
      	<b-alert variant="danger" show>We're sorry, we're not able to retrieve this information at the moment, please try back later</b-alert>
    </section>
    <section v-else>
	    <div v-if="loading" class="loading-data">
	     <atom-spinner
          :animation-duration="1000"
          :size="200"
          :color="'#564dfc'"
     	  />
	    </div>
	    <div v-else class="coin-data">
	    	<b-row>
				<b-table responsive striped :items="coindata" :fields="fields" id="coindata">
				    <template slot="coinName" slot-scope="data">
				      <b-img :src="data.item.coinImage" alt=" " />
				      <router-link :to="{ name: 'coin-details', params: { coinRank:data.item.rank, coinSymbol: data.item.coinSymbol }}">
				      	{{data.item.coinSymbol}} 
				      </router-link>
				      <br>
				      <span class="coin-name">{{data.item.coinName}} </span>
				    </template>
				    <template slot="coinPrice" slot-scope="data">
				      {{ data.item.coinPriceFormatted }}
				    </template>
					<template slot="change24" slot-scope="data">
				       <div v-bind:class = "{ 'text-success': data.value >= 0, 'text-danger': data.value < 0 , 'text-normal' : data.item.change24Formatted === 'N/A'}">
				       	{{ data.item.change24Formatted }}
				       </div>
				    </template>
					<template slot="priceGraph" slot-scope="data">
						<router-link :to="{ name: 'coin-details', params: { coinRank:data.item.rank, coinSymbol: data.item.coinSymbol }}">
							<b-img :src="data.value" alt="" />
						</router-link>
				    </template>
				    <template slot="volume24" slot-scope="data">
				      {{ data.item.volume24Formatted }}
				    </template>
				    <template slot="marketCap" slot-scope="data">
				      {{ data.item.marketCapFormatted }}
				    </template>
				    <template slot="circulatingSupply" slot-scope="data">
				      {{ data.item.circulatingSupplyFormatted }}
				    </template>
				    <template slot="oracle" slot-scope="data">
				    	<b-link target="_blank" href="https://github.com/diadata-org/api-golang/blob/master/methodology/Oracles.md">
				    		<b-img :src="data.value" fluid alt="Oracle"  width ="20" height = "20"/>
				    	</b-link>
				    </template>
				</b-table>
			</b-row>
	    </div>
    </section>
</template>
<script src="./CoinDataTable.js"></script>

