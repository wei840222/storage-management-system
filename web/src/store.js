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
    }
  },
  actions: {
    async getStorageList(context) {
      const res = await Vue.axios.get("/storage");
      context.commit('storageList', res.data.data)
    }
  }
})
