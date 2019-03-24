<template lang="pug">
  el-main#main
    el-table(:data="storageList")
      el-table-column(type="index" width="50")
      el-table-column(label="ID" prop="id" sortable)
      el-table-column(label="ReleaseName" prop="releaseName" sortable)
      el-table-column(label="Type" prop="type" sortable)
      el-table-column(label="Endpoint" prop="endpoint")
        template(slot-scope="scope")
          a(:href="scope.row.endpoint" target="_blank")
            el-tag(size="mini" type="primary") {{ scope.row.endpoint }}
      el-table-column(label="Status" prop="status")
        template(slot-scope="scope")
          el-tag(:type="scope.row.status === 'active' ? 'success' : 'warning'" disable-transitions) {{ scope.row.status }}
      el-table-column(label="Operation")
        template(slot-scope="scope")
          el-button(size="mini" type="primary") Monit
          el-button(size="mini" type="danger" @click="deleteStorage(scope.row.releaseName)") Delete
</template>

<script>
export default {
  name: "main",
  data() {
    return {
      storageList: []
    };
  },
  async mounted() {
    const res = await this.$axios.get("http://localhost:8080/storage");
    this.storageList = res.data.data;
    setInterval(async () => {
      const res = await this.$axios.get("http://localhost:8080/storage");
      this.storageList = res.data.data;
    }, 5000);
  },
  methods: {
    async deleteStorage(releaseName) {
      const res = await this.$axios.delete(
        `http://localhost:8080/storage/${releaseName}`
      );
      if (res.data.code === 200) {
        const res = await this.$axios.get("http://localhost:8080/storage");
        this.storageList = res.data.data;
      }
    }
  }
};
</script>

<style>
.demo-table-expand {
  font-size: 0;
}
.demo-table-expand label {
  width: 90px;
  color: #99a9bf;
}
.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}
</style>
