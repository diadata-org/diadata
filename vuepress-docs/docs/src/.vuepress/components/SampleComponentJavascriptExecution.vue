<template>
  <div>
  <h1>SampleComponentJavascriptExecution.vue</h1>
  <h2>{{ headlineText }}</h2>
    <button @click="testAlert">{{ displayText }}</button>
  </div>
</template>

<script>
export default {
  props: ['displayText','headlineText'],
  methods: {
    testAlert() {
      alert("TODO: check how to pass parameter to javascript if needed");
    }
  }
}
</script>

<style scoped>
button {
  background-color: #4CAF50;
}
</style>
