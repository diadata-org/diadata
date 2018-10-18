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
	    		<div class="subtext1 w-100">
	    			<ul>
		    			<li>
		    				CoinHub provides you with crowd-curated open-source financial data about cryptocurrencies – verified, cleaned for outliers, transparent and free of charge accessible via API or Oracle.
		    			</li>
		    			<li>
		    				CoinHub is a dApp by <a href="https://diadata.org/">DIA</a>. DIA provides a decentralized creation, contribution and crowd-validation approach to financial data for a wide variety of use cases in the financial space.
		    			</li>
		    		</ul>
	    		</div>
	    		<div class="subtext2 w-100">
	    			Start building your own:
	    		</div>
	    	</b-row>
	    	<b-row class="link-row">
		    	<b-col>
		          <a id="first" href="https://github.com/diadata-org/diadata" target="_blank">DIA Methodology</a>
		        </b-col>
		        <b-col>
		          <a id="second" href="http://swagger.diadata.org/api/swagger/index.html" target="_blank">API Documentation</a>
		        </b-col>
		    </b-row>
	    	<b-row >
	    		<div class="line-separator w-100"></div>
	    	</b-row>
	    	<b-row class="search-currency-row">
	            <b-col cols="9" md="4">
	                <model-list-select 
	                              :list="options"
	                              option-value="value"
	                              option-text="text"
	                              v-model="selectedCoin"
	                              placeholder="Search for Digital Asset"
	                              @searchchange="initSearch">
	                </model-list-select>
	             
	            </b-col>
	            <b-col cols="2" md="2" offset-md="6" offset="1">
	                <b-dropdown id="dd-currency" 
	                          v-bind:text="selectedCurrency" 
	                          v-bind:class = "{ 'd-none' : currencies.length <= 0, 'd-inline-block' : currencies.length > 0 }">
		                <b-dropdown-item v-for="currency in currencies" @click="switchCurrencies(currency)">
		                  {{ currency }}
		                </b-dropdown-item>
	              	</b-dropdown>
	            </b-col>    
        	</b-row>
	    	<b-row>
				<b-table responsive striped :items="coindata" :fields="fields" id="coindata">
				    <template slot="coinName" slot-scope="data">
				      <b-img :src="data.item.coinImage" alt=" " width="16" height="16" />
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
				    	<b-link target="_blank" href="https://github.com/diadata-org/api-golang/blob/master/documentation/methodology/Oracles.md">
				    		<b-img :src="data.value" fluid alt="Oracle"  width ="20" height = "20"/>
				    	</b-link>
				    </template>
				</b-table>
			</b-row>
	    </div>
    </section>
</template>
<script src="./CoinDataTable.js"></script>

