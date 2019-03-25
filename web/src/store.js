import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    storageList: []
  },
  mutations: {
    storageList(state, storageList) {
      state.storageList = storageList
    },
    add2storageList(state, storage) {
      state.storageList.push(storage)
    }
  },
  actions: {
    async getStorageList(context) {
      const res = await Vue.axios.get("http://localhost:8080/storage");
      context.commit('storageList', res.data.data)
    },
    async updateStorageList(context) {
      context.state.storageList.filter(item => item.status != 'deleting').forEach(async item => {
        const res = await Vue.axios.get(`http://localhost:8080/storage/${item.releaseName}`);
        item = res.data.data
      })
    }
  }
})
