<template>
  <header class="coin-data-header">
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
   
  </header>
</template>
<script src="./Header.js"></script>
<style src="./Header.css" scoped></style>
