<style src="./Header.css" scoped></style>
<template>
  <header>
    <b-navbar type="dark" variant="info">
        <b-col cols="12" md="2" class="logo">
          <b-link href="/" class="subheading-wrap logo">
            <b-img :src="logo"  fluid alt="logo" width="150" />
          </b-link>
        </b-col>
        <b-col cols="12" md="3" class="subheading-wrap">
          <div class="w-100 subheading">
              <div class="w-100">
                  <b-navbar-brand href="https://docs.diadata.org/documentation/methodology/digital-assets/exchangeprices" target="_blank">DIA Methodology</b-navbar-brand>
              </div>
          </div>
        </b-col>
        <b-col cols="12" md="3" class="subheading-wrap">
          <div class="w-100 subheading">
            <b-navbar-brand href="https://docs.diadata.org/documentation/api-1/api#get-v-1-chartpointsallexchanges" target="_blank">API Documentation</b-navbar-brand>
          </div>
        </b-col>
        <b-col cols="12" md="4" class="subheading-wrap">
          <div class="w-100 subheading">
            <b-navbar-brand href="/">CoinHub by DIA</b-navbar-brand>
            <br>
            <div class="sub-header">
              Decentrally sourced and maintained crypto-financial-data
            </div>
          </div>
        </b-col>
    </b-navbar>
  </header>
</template>
<script src="./Header.js"></script>

