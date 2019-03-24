<template lang="pug">
  #nav
    .navbar.navbar-dark.bg-dark.shadow-sm
      .container.d-flex.justify-content-between
        a.navbar-brand.d-flex.align-items-center(href="#")
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" aria-hidden="true" class="mr-2" viewBox="0 0 24 24" focusable="false"><path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"></path><circle cx="12" cy="13" r="4"></circle></svg>
          strong Storage
        button.navbar-toggler(type="button" data-toggle="collapse" data-target="#navbarHeader" aria-controls="navbarHeader" aria-expanded="false" aria-label="Toggle navigation")
          span.navbar-toggler-icon
    #navbarHeader.collapse.bg-dark
      .container
        .row.text-left
          .col-sm-8.col-md-7.py-4
            h4.text-white About
            p.text-muted Add some information about the album below, the author, or any other background context. Make it a few sentences long so folks can pick up some informative tidbits. Then, link them off to some social networking sites or contact information.
          .col-sm-4.offset-md-1.py-4
            h4.text-white Contact
            ul.list-unstyled
              li
                a.text-white(href="#") Follow on Twitter
              li
                a.text-white(href="#") Like on Faceboo
              li
                a.text-white(href="#") Email me
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

</style>
