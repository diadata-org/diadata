<style src="./Header.css" scoped></style>
<template>
  <header>
    <b-navbar type="dark" variant="info">
      <b-row>
        <b-col>
          <b-link href="/">
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
    </b-navbar>
  </header>
</template>
<script src="./Header.js"></script>

