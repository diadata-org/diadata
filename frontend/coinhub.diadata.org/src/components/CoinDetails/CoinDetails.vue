<style src="./CoinDetails.css" scoped></style>
<template>

    <section v-if="errored" class="loading-error">
		<b-alert variant="danger" show>We're sorry, we're not able to retrieve this information at the moment, please try back later</b-alert>
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
			        	<div class="text-methodology w-100">
				        	<p class="font-weight-bold"> Methodology to the price collection can be found
				        		<b-link target="_blank" href="https://github.com/diadata-org/diadata/blob/master/documentation/methodology/ExchangePrices.md">
				        			here
				        		</b-link>
				        	</p>
				        </div>
			    	</b-col>
			    </b-row>
			    <b-row>
			        <b-col cols="12" md="6">
			        	<br>
			        	<br>
			        	<div class="w-100 text-left">
				        	<h3> Supply:
				        		<b-link target="_blank" href="https://github.com/diadata-org/diadata/blob/master/documentation/methodology/SupplyNumbers.md">
				        			{{ coinDetails.circulatingSupplyFormattedWithoutSymbol }}
				        	    </b-link>
				        	</h3>
				        </div>
			        </b-col>
			        <b-col cols="12" md="6" class="dropdown-wrap">
		                <b-dropdown id="dd-algorithm"
		                          v-bind:text="selectedAlgorithmName"
		                          v-bind:class = "{ 'd-none' : algorithms.length <= 0, 'd-inline-block' : algorithms.length > 0 }">
			                <b-dropdown-item v-for="algorithm in algorithms" @click="switchAlgorithm(algorithm)">
			                  {{ algorithm.displayName }}
			                </b-dropdown-item>
		              	</b-dropdown>
		                <b-dropdown id="dd-exchange"
		                          v-bind:text="selectedExchange"
		                          v-bind:class = "{ 'd-none' : exchangeNames.length <= 0, 'd-inline-block' : exchangeNames.length > 0 }">
			                <b-dropdown-item @click="switchExchange('All')">
			                  All
			                </b-dropdown-item>
							<b-dropdown-item v-for="exchange in exchangeNames" @click="switchExchange(exchange)">
			                  {{ exchange }}
			                </b-dropdown-item>
		              	</b-dropdown>
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
			    	<highcharts  v-bind:class = "{ 'd-none' : showAllCharts === false, 'd-inline-block' : showAllCharts === true }"
			    				 class="coindata-charts" :constructor-type="'stockChart'" :options="chartAllOptions">
			    	</highcharts>
			    	<!-- <highcharts class="coindata-charts" :options="chartSimexOptions"></highcharts> -->
			    </b-row>
			     <b-row>
			     	<b-col>
			        	<div class="data-sources-banner">Data Sources</div>
			        	<b-table responsive striped :items="exchanges" :fields="exchange_fields" class="main">
			        		<template slot="show_trades" slot-scope="row">
							      <!-- we use @click.stop here to prevent emitting of a 'row-clicked' event  -->
							      <b-button size="sm" @click.stop="row.toggleDetails" class="mr-2" >
							       {{ row.detailsShowing ? 'Hide' : 'Show'}}  Last 10 Trades
							      </b-button>
							</template>
							<template slot="row-details" slot-scope="row">
     							<b-table responsive striped :items="row.item.LastTrades" :fields="last_trade_fields" class="details" >
     								<template slot="Order" slot-scope="data">
								       <div v-bind:class = "{ 'text-success': data.item.Volume >= 0, 'text-danger': data.item.Volume < 0 }">
								       	{{ data.item.Volume >= 0 ? 'BUY' : 'SELL' }}
								       </div>
								    </template>
								    <template slot="Volume" slot-scope="data">
								      {{ data.value < 0 ?  data.value * -1 : data.value}}
								    </template>
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
