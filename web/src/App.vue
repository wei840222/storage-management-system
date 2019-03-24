<template lang="pug">
  el-container#app
    el-header
      Nav
    el-main
      Main
</template>

<script>
import Nav from './components/Nav.vue'
import Main from './components/Main.vue'

export default {
  name: 'app',
  components: { Main, Nav }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
