<style src="./Header.css" scoped></style>
<template>
  <header>
    <b-navbar type="dark" variant="info">
      <b-navbar-brand href="/">CoinHub by DIA</b-navbar-brand>
	    <b-navbar-nav class="ml-auto logo">
	        <b-link href="https://diadata.org" target="_blank">
	    		<b-img :src="logo" fluid alt="logo" width="150" />
	    	</b-link>
	    </b-navbar-nav>
      <div class="w-100 sub-header text-left">
        <b-row>
          <b-col sm="12" md="8" lg="8">
              Decentrally sourced and maintained crypto-financial-data
          </b-col>
          <b-col sm="12" md="4" lg="4">
            <b-navbar-nav class="search-form">
              <b-nav-form>
                <b-form-input size="sm" class="search-input mr-lg-5" type="text" placeholder="Search for Digital Asset"/>
              </b-nav-form>
            </b-navbar-nav>
          </b-col>
        </b-row>
      </div>
     
    </b-navbar>
  </header>
</template>
<script src="./Header.js"></script>

