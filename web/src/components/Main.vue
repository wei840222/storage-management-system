<template lang="pug">
  el-main#main
    el-collapse
      el-collapse-item(v-for="(item, index) in storageList" :title="`${item.releaseName}-${item.type}`" :name="index")
        div Status: {{ item.status }}
        div Endpoint
          a(:href="item.endpoint") {{ item.endpoint }}
</template>

<script>
export default {
  name: 'main',
  data() {
    return {
      storageList: []
    }
  },
  async mounted() {
    const res = await this.$axios.get('http://localhost:8080/storage')
    this.storageList = res.data.data
  }
}
</script>