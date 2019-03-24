<template lang="pug">
  el-menu.el-menu-demo(default-active="1" mode="horizontal")
    el-submenu(index="1")
      template(slot="title") New
      el-menu-item(index="minio" @click="dialogFormVisible = true") minio
    
    el-dialog(title="Create minio storage" :visible.sync="dialogFormVisible")
      el-form(:model="form")
        el-form-item(label="accessKey" label-width="120px")
          el-input(v-model="form.accessKey" autocomplete="off")
        el-form-item(label="secretKey" label-width="120px")
          el-input(v-model="form.secretKey" autocomplete="off")
        el-form-item(label="size" label-width="120px")
          el-input(v-model="form['persistence.size']" autocomplete="off")
      .dialog-footer(slot="footer")
        el-button(type="primary" :loading="creating" @click="createMinio") Create
        el-button(@click="dialogFormVisible = false") Cancel
</template>

<script>
export default {
  data() {
    return {
      dialogFormVisible: false,
      creating: false,
      form: {
        accessKey: "",
        secretKey: "",
        "persistence.size": "",
        "service.type": "NodePort",
        "service.nodePort": "",
        "resources.limits.memory": "512Mi"
      }
    };
  },
  methods: {
    async createMinio() {
      this.creating = true
      const res = await this.$axios.post("http://localhost:8080/storage", {
        type: "minio",
        config: this.form
      });
      this.dialogFormVisible = false
    }
  }
};
</script>
