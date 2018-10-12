<template>
  <header>
    <b-navbar type="dark" variant="info">
      <b-navbar-brand href="/">CoinHub by DIA</b-navbar-brand>
	    <b-navbar-nav class="ml-auto logo">
	        <b-link href="https://diadata.org" target="_blank">
	    		<b-img :src="logo" fluid alt="logo" width="200" />
	    	</b-link>
	    </b-navbar-nav>
      <div class="w-100 sub-header text-left">
        Decentrally sourced and maintained crypto-financial-data
      </div>
    </b-navbar>
    <b-row>
      <div class="exchange-buttons">
          <b-link @click="switchCurrencies('EUR')" 
                  v-bind:class = "{ 'selected': selectedCurrency ==='EUR', 'not-selected' : selectedCurrency !=='EUR'}"
                  id="btn-eur">
            EUR
          </b-link>
          <b-link @click="switchCurrencies('USD')" 
              v-bind:class = "{ 'selected': selectedCurrency ==='USD', 'not-selected' : selectedCurrency !=='USD'}"
              id="btn-usd">
          USD
        </b-link>
      </div>
    </b-row>
  </header>
</template>
<script src="./Header.js"></script>
<style src="./Header.css" scoped></style>
