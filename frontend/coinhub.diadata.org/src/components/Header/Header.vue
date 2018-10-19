<style src="./Header.css" scoped></style>
<template>
  <header>
    <b-navbar type="dark" variant="info">
      <b-row>
        <b-col cols="12" md="6">
          <b-link href="/" class="logo">
            <b-img :src="logo"  fluid alt="logo" width="150" />
          </b-link>
        </b-col>
        <b-col cols="12" md="6">
          <div class="w-100 subheading">
            <b-navbar-brand href="/">CoinHub by DIA</b-navbar-brand>
            <br>
            <div class="sub-header">
                Decentrally sourced and maintained crypto-financial-data
            </div>
          </div>   
        </b-col>
      </b-row>
    </b-navbar>
  </header>
</template>
<script src="./Header.js"></script>

