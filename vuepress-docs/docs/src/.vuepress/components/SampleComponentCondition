<template>
  <div class="code-block">
    <slot v-if="display"/>
  </div>
</template>

<script>
export default {
  name: 'code-block',
  props: {
    language: String
  },
  data() {
    return {
      display: false
    }
  },
  methods: {
    getCookie() {
      return document.cookie.replace(/(?:(?:^|.*;\s*)language\s*\=\s*([^;]*).*$)|^.*$/, "$1");
    },
  },
  mounted() {
    const cookieValue = this.getCookie();
    this.display = cookieValue === this.language;
  }
}
</script>
