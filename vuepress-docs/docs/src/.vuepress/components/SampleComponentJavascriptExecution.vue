<template>
  <div>
  <h1>SampleComponentJavascriptExecution.vue</h1>
  <h2>{{ headlineText }}</h2>
    <button @click="testAlert">{{ displayText }}</button>
  </div>
</template>

<script>
var myParam = {{ headlineText }};
export default {
  props: ['displayText','headlineText'],
  methods: {
    testAlert() {
      alert(myParam);
    }
  }
}
</script>

<style scoped>
button {
  background-color: #4CAF50;
}
</style>
