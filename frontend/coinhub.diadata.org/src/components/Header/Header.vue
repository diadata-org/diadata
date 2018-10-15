<style src="./Header.css" scoped></style>
<template>
  <div>
    <header class="lg-screens">
      <b-navbar type="dark" variant="info">
        <b-row>
          <b-col>
            <b-link href="https://diadata.org" target="_blank">
              <b-img :src="logo" fluid alt="logo" width="150" />
            </b-link>
          </b-col>
          <b-col>
            <div class="w-100 text-right">
              <b-navbar-brand href="/">CoinHub by DIA</b-navbar-brand>
              <br>
              <div class="sub-header">
                  Decentrally sourced and maintained crypto-financial-data
              </div>
            </div>
            
          </b-col>
        </b-row>
        <b-row class="header-components">
            <b-col md="12" lg="5" xl="4" >
              <b-navbar-nav class="search-form">
                <b-nav-form>
                  <b-form-input size="md" class="search-input" type="text" placeholder="Search for Digital Asset"/>
                </b-nav-form>
              </b-navbar-nav>
            </b-col>
            <b-col md="12" lg="1" offset-lg="6" xl="1" offset-xl="7">
              <b-dropdown id="dd-currency" text="EUR">
                <b-dropdown-item>EUR</b-dropdown-item>
                <b-dropdown-item>USD</b-dropdown-item>
                <b-dropdown-item>NOK</b-dropdown-item>
              </b-dropdown>
            </b-col>    
        </b-row>
      </b-navbar>
    </header>
    <header class="sm-screens">
      <b-navbar type="dark" variant="info">
        <b-row>
          <b-col>&nbsp;</b-col>
          <b-col>
            <b-link href="/">
              <b-img :src="logo" fluid alt="logo" width="150" />
            </b-link>
            <b-navbar-nav class="search-form d-none">
                <b-nav-form>
                  <b-form-input size="md" class="search-input" type="text" placeholder="Search for Digital Asset"/>
                </b-nav-form>
            </b-navbar-nav>
          </b-col>
          <b-col class="text-right">
            <b-link>
              <font-awesome-icon icon="search" />
            </b-link>
          </b-col>
        </b-row>
        <b-row class="header-components">
            <b-col md="12" lg="1" offset-lg="6" xl="1" offset-xl="7">
              <b-dropdown id="dd-currency" text="EUR">
                <b-dropdown-item>EUR</b-dropdown-item>
                <b-dropdown-item>USD</b-dropdown-item>
                <b-dropdown-item>NOK</b-dropdown-item>
              </b-dropdown>
            </b-col>    
        </b-row>
      </b-navbar>
    </header>
  </div>
</template>
<script src="./Header.js"></script>

