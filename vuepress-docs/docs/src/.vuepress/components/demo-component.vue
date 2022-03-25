<template>
  <p class="demo" style="border:solid blue 5px !important">
    {{ msg }}
    {{ text }}
  </p>
</template>

<script>
  export default {
    data() {
      return {
        msg: 'Test edited component, see demo-component.vue <demo-component>'
      }
    }
  }
</script>
