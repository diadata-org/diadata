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
            <b-col md="12" lg="5" xl="4" class="d-inline-block" >
      
               <b-nav-form class="search-form" v-on:submit.prevent>
                    <b-form-input size="md" class="search-input" type="text" placeholder="Search for Digital Asset"></b-form-input>
                </b-nav-form>
             
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
          <b-col v-bind:class = "{ 'd-none' : showSearch === true, 'd-inline-block' : showSearch === false }">&nbsp;</b-col>
          <b-col v-bind:class = "{ 'd-none' : showSearch === true, 'd-inline-block' : showSearch === false }">
            <b-link href="/">
              <b-img :src="logo" fluid alt="logo" width="150" />
            </b-link>
          </b-col>
          <b-col cols="12" v-bind:class = "{ 'd-none' : showSearch === false, 'd-inline-block' : showSearch === true }">
            <b-navbar-nav class="search-form">
                <b-nav-form  v-on:submit.prevent>
                  <b-input-group>
                    <b-form-input size="md" class="search-input" type="text" placeholder="Search for Digital Asset"></b-form-input>
                    <b-input-group-append>
                      <button class="hide-search" @click.stop="showHideSearch">
                        <font-awesome-icon icon="times-circle" />
                      </button>
                    </b-input-group-append>
                  </b-input-group>
                </b-nav-form>
            </b-navbar-nav>
          </b-col>
          <b-col class="btn-search text-right" v-bind:class = "{ 'd-none' : showSearch === true, 'd-inline-block' : showSearch === false }">
            <b-link @click.stop="showHideSearch">
              <font-awesome-icon icon="search" />
            </b-link>
          </b-col>
        </b-row>
        <b-row class="header-components">
          <b-dropdown id="dd-currency" text="EUR">
            <b-dropdown-item>EUR</b-dropdown-item>
            <b-dropdown-item>USD</b-dropdown-item>
            <b-dropdown-item>NOK</b-dropdown-item>
          </b-dropdown>
        </b-row>
      </b-navbar>
    </header>
  </div>
</template>
<script src="./Header.js"></script>

