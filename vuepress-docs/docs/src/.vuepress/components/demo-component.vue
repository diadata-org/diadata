<template>
  <p class="demo" style="border:solid red 5px !important">
    {{ msg }}
    {{ textattribute }}
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
