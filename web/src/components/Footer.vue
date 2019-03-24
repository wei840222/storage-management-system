<template lang="pug">
  #footer
    .text-muted
      .container
        p.float-right
        a(href="#nav") Back to top
        p Album example is © Bootstrap, but please download and customize it for yourself!
        p New to Bootstrap? <a href="https://getbootstrap.com/">Visit the homepage</a> or read our <a href="/docs/4.3/getting-started/introduction/">getting started guide</a>. 
</template>

<script>
export default {
  name: 'Nav',
  props: {
    msg: String
  }
}
</script>

<style scoped lang="sass">
  #footer
    padding-top: 3rem
    padding-bottom: 3rem
</style>
