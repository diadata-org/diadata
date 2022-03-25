<template>
  <div>
  <h1>{{ headlineText }}</h1>
    <button @click="testAlert">{{ displayText }}</button>
  </div>
</template>

<script>
export default {
  props: ['displayText','headlineText'],
  methods: {
    testAlert() {
      alert('Test Alert');
    }
  }
}
</script>

<style scoped>
button {
  background-color: #4CAF50;
}
</style>
