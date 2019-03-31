<template lang="pug">
  el-table(:data="$store.state.storageList")
    el-table-column(type="index" width="50")
    el-table-column(label="ID" width="220" prop="id" sortable)
    el-table-column(label="ReleaseName" width="150" prop="releaseName" sortable)
    el-table-column(label="Type" width="100" prop="chartName" sortable)
      template(slot-scope="scope")
        p {{ scope.row.chartName.split('/')[1] }}
    el-table-column(label="Size/Capacity" prop="persistentVolumeClaim.size" sortable)
      template(slot-scope="scope")
        el-popover(placement="top-start" trigger="hover" :content="getSizeString(scope.row.persistentVolumeClaim.size)+'/'+getSizeString(scope.row.persistentVolumeClaim.capacity)")
          el-progress(slot="reference" :text-inside="true" :stroke-width="18" color="#8e71c7" :percentage="getPercentage(scope.row.persistentVolumeClaim.size, scope.row.persistentVolumeClaim.capacity)")
    el-table-column(label="Endpoint" prop="endpoint.port")
      template(slot-scope="scope")
        a(v-if="scope.row.chartName === 'stable/minio'" :href="`http://${scope.row.endpoint.host}:${scope.row.endpoint.port}`" target="_blank")
          el-tag(size="mini" type="primary") http://{{ scope.row.endpoint.host }}:{{ scope.row.endpoint.port }}
        el-tag(v-else size="mini" type="primary") {{ scope.row.endpoint.host }}:{{ scope.row.endpoint.port }}
    el-table-column(label="Status" width="110" prop="status")
      template(slot-scope="scope")
        el-tag(:type="scope.row.status === 'running' ? 'success' : 'warning'" disable-transitions) {{ scope.row.status }}
    el-table-column(label="Operation")
      template(slot-scope="scope")
        el-button(size="mini" type="primary" @click="showMonitor(scope.row)") Monit
        el-button(size="mini" type="danger" :loading="scope.row.status === 'deleting'" @click="deleteStorage(scope.row)") Delete
        el-dialog(:visible.sync="monitorShow" :fullscreen="true" @close="closeMonitor")
          iframe(:src="prometheusUrl.cpu" width="100%" height="180" frameborder="0")
          iframe(:src="prometheusUrl.memory" width="100%" height="180" frameborder="0")
          iframe(:src="prometheusUrl.network" width="100%" height="180" frameborder="0")
</template>

<script>
export default {
  name: "main",
  data() {
    return {
      monitorShow: false,
      prometheusUrl: {
        cpu: "",
        memory: "",
        network: ""
      }
    };
  },
  async mounted() {
    await this.$store.dispatch("getStorageList");
    setInterval(() => this.$store.dispatch("getStorageList"), 10000);
  },
  methods: {
    async deleteStorage(data) {
      data.status = "deleting";
      const res = await this.$axios.delete(
        `http://localhost:8080/storage/${data.releaseName}`
      );
      if (res.data.code === 200) {
        await this.$store.dispatch("getStorageList");
        this.$notify({
          title: "Deleted",
          message: `Delete ${data.chartName.split("/")[1]} success!`,
          type: "success"
        });
      }
    },
    getSizeString(size) {
      const sizeMi = size / 1024 / 1024;
      if (isNaN(size)) {
        return 0 + "Mi";
      } else if (sizeMi < 1000) {
        return Math.round(sizeMi * 100) / 100 + "Mi";
      } else {
        return Math.round((sizeMi / 1024) * 100) / 100 + "Gi";
      }
    },
    getPercentage(size, capacity) {
      if (isNaN(size / capacity) || capacity < size) {
        return 0;
      } else return Math.round((size / capacity) * 100);
    },
    showMonitor(data) {
      this.prometheusUrl.cpu = data.prometheusUrl.CPU;
      this.prometheusUrl.memory = data.prometheusUrl.Memory;
      this.prometheusUrl.network = data.prometheusUrl.Network;
      this.monitorShow = true;
    },
    closeMonitor() {
      this.prometheusUrl = {
        cpu: "",
        memory: "",
        network: ""
      };
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
