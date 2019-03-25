<template lang="pug">
  el-table(:data="$store.state.storageList")
    el-table-column(type="index" width="50")
    el-table-column(label="ID" prop="id" sortable)
    el-table-column(label="ReleaseName" prop="releaseName" sortable)
    el-table-column(label="Type" prop="type" sortable)
    el-table-column(label="Size/Capacity" prop="volumeSize" sortable)
      template(slot-scope="scope")
        el-popover(placement="top-start" width="150" trigger="hover" :content="getSizeString(scope.row.volumeSize) + '/' +scope.row.volumeCapacity")
          el-progress(slot="reference" :text-inside="true" :stroke-width="18" color="#8e71c7" :percentage="Math.ceil(scope.row.volumeSize / (Number(scope.row.volumeCapacity.split('Gi')[0])*1024*1024*1024) * 100)")
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
        el-button(size="mini" type="danger" :loading="scope.row.status === 'deleting'" @click="deleteStorage(scope.row)") Delete
</template>

<script>
export default {
  name: "main",
  async mounted() {
    await this.$store.dispatch("getStorageList");
    setInterval(() => this.$store.dispatch("updateStorageList"), 5000);
  },
  methods: {
    async deleteStorage(data) {
      data.status = "deleting";
      const res = await this.$axios.delete(
        `http://localhost:8080/storage/${data.releaseName}`
      );
      if (res.data.code === 200) {
        await this.$store.dispatch("getStorageList");
      }
    },
    getSizeString(size) {
      const sizeMi = size / 1024 / 1024;
      if (sizeMi < 1000) {
        return Math.round(sizeMi * 100) / 100 + "Mi";
      } else {
        return Math.round((sizeMi / 1024) * 100) / 100 + "Gi";
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
