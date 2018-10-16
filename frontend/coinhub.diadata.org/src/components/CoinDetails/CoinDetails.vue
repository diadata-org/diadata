<style src="./CoinDetails.css" scoped></style>
<template>

    <section v-if="errored">
    <p>
      We're sorry, we're not able to retrieve this information at the moment, please try back later
    </p>
    </section>
    <section v-else class="coin-details">
	    <div v-if="loading" class="loading-data">
	     <atom-spinner
          :animation-duration="1000"
          :size="200"
          :color="'#564dfc'"
     	  />
	    </div>
	    <div v-else>
			<b-container fluid>
			    <b-row>
			        <b-col cols="12" md="6">
			        	<br>
			        	<br>
				        <div class="w-100 text-left">
				        	<h5> {{ coinDetails.coinName}} ({{coinDetails.coinSymbol}}) </h5>
				        </div>
				        <div class="w-100 text-left">
				        	<h2> {{ coinDetails.coinPriceFormatted}} </h2>
				        </div>
				        <div class="w-100 text-left">
				        	<h6 v-bind:class = "{ 'text-success': coinDetails.change24 >= 0, 'text-danger': coinDetails.change24 < 0 , 'text-normal' : coinDetails.change24Formatted === 'N/A'}">
			       				{{ coinDetails.change24Formatted }}
				            </h6> 
				        </div>
				        <div class="w-100 text-left">
				        	<span>Rank: {{ coinDetails.rank}}</span>;
				        	<span>24h Volume: {{ coinDetails.volume24Formatted}}</span>
				        </div>
			        </b-col>
			        <b-col 	cols="12" md="6">
			        	<br>
			        	<br>
			        	<div class="text-methodology w-100 text-right">
				        	<h6 class="font-weight-bold"> Methodology to the price collection can be found
				        		<b-link target="_blank" href="https://github.com/diadata-org/api-golang/blob/master/methodology/ExchangePrices.md">
				        			here
				        		</b-link>
				        	</h6>
				        </div>
			    	</b-col>
			    </b-row>
			    <b-row>
			        <b-col>
			        	<br>
			        	<br>
			        	<div class="w-100 text-left">
				        	<h3> Supply:
				        		<b-link target="_blank" href="https://github.com/diadata-org/api-golang/blob/master/methodology/SupplyNumbers.md">
				        			{{ coinDetails.circulatingSupplyFormattedWithoutSymbol }}
				        	    </b-link>
				        	</h3>
				        </div>
			        </b-col>
			    </b-row>
			     <b-row>
			     	<b-col>
			        	<div class="data-sources-banner font-weight-bold">Data Sources</div>
			        	<b-table responsive striped :items="exchanges" :fields="exchange_fields" >
			        		<template slot="show_trades" slot-scope="row">
							      <!-- we use @click.stop here to prevent emitting of a 'row-clicked' event  -->
							      <b-button size="sm" @click.stop="row.toggleDetails" class="mr-2">
							       {{ row.detailsShowing ? 'Hide' : 'Show'}}  Last 10 Trades
							      </b-button>
							</template>
							<template slot="row-details" slot-scope="row">
     							<b-table responsive striped :items="row.item.LastTrades" :fields="last_trade_fields" >
     								</b-table>
    						</template>
			        	</b-table>
			        </b-col>
			    </b-row>
			</b-container>
	    </div>

    </section>
</template>
<script src="./CoinDetails.js"></script>
