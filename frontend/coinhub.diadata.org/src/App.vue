<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style>

body {
  background-color: #FFFFFF;
  width: 100%;
  display: flex;
  flex-direction: column;
}

body, html {
  margin: 0;
  height: 100%;
}

table td {
  font-size: 14px;
}

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
   flex: 1 0 auto;
}


</style>
