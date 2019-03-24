<template lang="pug">
  #main
    .jumbotron.text-center
      .container
          h1.jumbotron-heading Album example
          p.lead.text-muted Something short and leading about the collection below—its contents, the creator, etc. Make it short and sweet, but not too short so folks don’t simply skip over it entirely.
          p
            a.btn.btn-primary.my-2(href="#") Main call to action
            a.btn.btn-secondary.my-2(href="#") Secondary action
    .album.py-5.bg-light
      .container
        .row
          .col-md-4
            .card.mb-4.shadow-sm
              .logo
                img.logo-img(src="https://min.io/resources/img/logo.svg" alt="MinIO Logo")
              .card-body
                p.card-text This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.
                .d-flex.justify-content-between.align-items-center
                  .btn-group
                    button.btn.btn-sm.btn-outline-secondary(type="button") View
                    button.btn.btn-sm.btn-outline-secondary(type="button") Edit
                  small.text-muted 9 mins
</template>

<style lang="sass" scoped>
  .jumbotron
    padding-top: 3rem
    padding-bottom: 3rem
    margin-bottom: 0
    background-color: #fff

  .jumbotron p:last-child
    margin-bottom: 0

  .jumbotron-heading
    font-weight: 300

  .jumbotron .container
    max-width: 40rem

  .logo
    background-color: #55595c

  .logo-img
    width: 80%
    padding-top: 10%
    padding-bottom: 10%
</style>
