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
                <model-list-select 
                              :list="options"
                              option-value="value"
                              option-text="text"
                              v-model="selectedCoin"
                              placeholder="Search for Digital Asset"
                              @searchchange="initSearch"
                              v-bind:class = "{ 'd-none' : hideSearchInput === true, 'd-inline-block' : hideSearchInput === false }">
                </model-list-select>
             
            </b-col>
            <b-col md="12" lg="1" offset-lg="6" xl="1" offset-xl="7">
              <b-dropdown id="dd-currency" 
                          v-bind:text="selectedCurrency" 
                          v-bind:class = "{ 'd-none' : currencies.length <= 0, 'd-inline-block' : currencies.length > 0 }">
                <b-dropdown-item v-for="currency in currencies" @click="switchCurrencies(currency)">
                  {{ currency }}
                </b-dropdown-item>
                
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
            <b-navbar-nav class="search-form d-inline-block">
                <model-list-select 
                        :list="options"
                        option-value="value"
                        option-text="text"
                        v-model="selectedCoin"
                        placeholder="Search for Digital Asset"
                        @searchchange="initSearch"
                        class = "d-inline-block">
                </model-list-select>
                <button class="hide-search" @click.stop="showHideSearch">
                    <font-awesome-icon icon="times-circle" />
                </button>
            </b-navbar-nav>
          </b-col>
          <b-col class="btn-search text-right" v-bind:class = "{ 'd-none' : showSearch === true, 'd-inline-block' : showSearch === false }">
            <b-link @click.stop="showHideSearch"
                    v-bind:class = "{ 'd-none' : hideSearchInput === true, 'd-inline-block' : hideSearchInput === false }">
              <font-awesome-icon icon="search" />
            </b-link>
            <div v-bind:class = "{ 'd-none' : hideSearchInput === true, 'd-inline-block' : hideSearchInput === false }">&nbsp;</div>
          </b-col>
        </b-row>
        <b-row class="header-components">
          <b-dropdown id="dd-currency" 
                      v-bind:text="selectedCurrency" 
                      v-bind:class = "{ 'd-none' : currencies.length <= 0, 'd-inline-block' : currencies.length > 0 }">
                <b-dropdown-item v-for="currency in currencies" @click="switchCurrencies(currency)">
                  {{ currency }}
                </b-dropdown-item>
          </b-dropdown>
        </b-row>
      </b-navbar>
    </header>
  </div>
</template>
<script src="./Header.js"></script>

