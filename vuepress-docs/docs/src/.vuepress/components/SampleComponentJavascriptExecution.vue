<template>
  <div>
  <h1>SampleComponentJavascriptExecution.vue</h1>
  <h2>{{ headlineText }}</h2>
    <button @click="testAlert({{ headlineText }})">{{ displayText }}</button>
  </div>
</template>

<script>
export default {
  props: ['displayText','headlineText'],
  methods: {
    testAlert(arg) {
      alert(arg);
    }
  }
}
</script>

<style scoped>
button {
  background-color: #4CAF50;
}
</style>
