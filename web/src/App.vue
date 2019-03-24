<template lang="pug">
  #app
    Nav
    Main
    Footer
</template>

<script>
import Nav from './components/Nav.vue'
import Main from './components/Main.vue'
import Footer from './components/Footer.vue'

export default {
  name: 'app',
  components: { Nav, Main, Footer }
}
</script>

<style lang="sass">
  #app
    font-family: 'Avenir', Helvetica, Arial, sans-serif
    -webkit-font-smoothing: antialiased
    -moz-osx-font-smoothing: grayscale
    text-align: center
    color: #2c3e50
</style>
