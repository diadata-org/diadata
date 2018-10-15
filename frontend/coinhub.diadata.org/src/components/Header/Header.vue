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
          <b-col sm="12" md="8" lg="5" xl="4">
              Decentrally sourced and maintained crypto-financial-data
          </b-col>
          <b-col sm="12" md="4" lg="5" xl="3">
            <b-navbar-nav class="search-form">
              <b-nav-form>
                <b-form-input size="md" class="search-input mr-lg-1" type="text" placeholder="Search for Digital Asset"/>
              </b-nav-form>
            </b-navbar-nav>
          </b-col>
          <b-col sm="12" md="12" lg="2" xl="2" offset-xl="3">
            <b-dropdown id="dd-currency" text="EUR">
              <b-dropdown-item>EUR</b-dropdown-item>
              <b-dropdown-item>USD</b-dropdown-item>
              <b-dropdown-item>NOK</b-dropdown-item>
            </b-dropdown>
          </b-col>
        </b-row>
      </div>
     
    </b-navbar>
  </header>
</template>
<script src="./Header.js"></script>

