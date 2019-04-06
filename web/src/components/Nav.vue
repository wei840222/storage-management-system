<template lang="pug">
  el-menu.el-menu-demo(default-active="1" mode="horizontal")
    el-submenu(index="1")
      template(slot="title") New
      el-menu-item(index="minio" @click="minio.dialogFormVisible = true") minio
      el-menu-item(index="mysql" @click="mysql.dialogFormVisible = true") mysql
      el-menu-item(index="mongodb" @click="mongodb.dialogFormVisible = true") mongodb
    
    el-dialog(title="Create minio storage" :visible.sync="minio.dialogFormVisible")
      el-form(:model="minio.config" label-position="left" label-width="80px")
        el-form-item(label="accessKey")
          el-input(v-model="minio.config.accessKey")
        el-form-item(label="secretKey")
          el-input(v-model="minio.config.secretKey")
        el-form-item(label="size")
          el-select(v-model="minio.config['persistence.size']" placeholder="size")
            el-option(v-for="s in ['2G','4G','8G','20G','40G']" :key="s" :label="s" :value="s")
      .dialog-footer(slot="footer")
        el-button(type="primary" :loading="minio.creating" @click="create(minio)") Create
        el-button(@click="minio.dialogFormVisible = false") Cancel

    el-dialog(title="Create mysql storage" :visible.sync="mysql.dialogFormVisible")
      el-form(:model="mysql.config" label-position="left" label-width="80px")
        el-form-item(label="username")
          el-input(v-model="mysql.config.mysqlUser")
        el-form-item(label="password")
          el-input(v-model="mysql.config.mysqlPassword")
        el-form-item(label="database")
          el-input(v-model="mysql.config.mysqlDatabase")
        el-form-item(label="size")
          el-select(v-model="mysql.config['persistence.size']" placeholder="size")
            el-option(v-for="s in ['2G','4G','8G','20G','40G']" :key="s" :label="s" :value="s")
      .dialog-footer(slot="footer")
        el-button(type="primary" :loading="mysql.creating" @click="create(mysql)") Create
        el-button(@click="mysql.dialogFormVisible = false") Cancel

    el-dialog(title="Create mongodb storage" :visible.sync="mongodb.dialogFormVisible")
      el-form(:model="mongodb.config" label-position="left" label-width="80px")
        el-form-item(label="username")
          el-input(v-model="mongodb.config.mongodbUsername")
        el-form-item(label="password")
          el-input(v-model="mongodb.config.mongodbPassword")
        el-form-item(label="database")
          el-input(v-model="mongodb.config.mongodbDatabase")
        el-form-item(label="size")
          el-select(v-model="mongodb.config['persistence.size']" placeholder="size")
            el-option(v-for="s in ['2G','4G','8G','20G','40G']" :key="s" :label="s" :value="s")
      .dialog-footer(slot="footer")
        el-button(type="primary" :loading="mongodb.creating" @click="create(mongodb)") Create
        el-button(@click="mongodb.dialogFormVisible = false") Cancel
</template>

<script>
export default {
  data() {
    return {
      minio: {
        dialogFormVisible: false,
        creating: false,
        chartName: "stable/minio",
        config: {
          accessKey: "",
          secretKey: "",
          "persistence.size": "2G",
          "service.type": "NodePort",
          "service.nodePort": "",
          "resources.limits.memory": "512Mi"
        }
      },
      mysql: {
        dialogFormVisible: false,
        creating: false,
        chartName: "stable/mysql",
        config: {
          mysqlUser: "",
          mysqlPassword: "",
          mysqlDatabase: "database",
          "persistence.size": "2G",
          "service.type": "NodePort",
          "service.nodePort": "",
          "resources.limits.memory": "512Mi"
        }
      },
      mongodb: {
        dialogFormVisible: false,
        creating: false,
        chartName: "stable/mongodb",
        config: {
          mongodbUsername: "",
          mongodbPassword: "",
          mongodbDatabase: "database",
          "persistence.size": "2G",
          "service.type": "NodePort",
          "service.nodePort": "",
          "resources.limits.memory": "512Mi"
        }
      }
    };
  },
  methods: {
    async create(data) {
      data.creating = true;
      const res = await this.$axios.post("/storage", data);
      if (res.data.code === 201) {
        res.data.data.persistentVolumeClaim = {
          size: 0,
          capacity: 0
        };
        this.$store.dispatch("getStorageList");
        this.$notify({
          title: "Created",
          message: `Create ${data.chartName.split("/")[1]} success!`,
          type: "success"
        });
      } else {
        this.$notify({
          title: "Failed",
          message: `Create ${data.chartName.split("/")[1]} failed!`,
          type: "error"
        });
      }
      data.dialogFormVisible = false;
      data.creating = false;
    }
  }
};
</script>
