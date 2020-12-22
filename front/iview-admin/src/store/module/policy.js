import { PolicyAPI } from '@/api/policy'
// import {  } from '@/libs/util'

export default {
  state: {
  },
  mutations: {
  },
  getters: {
  },
  actions: {
    handleCreatePolicy ({ state, commit }, role) {
      return new Promise((resolve, reject) => {
        try {
          PolicyAPI.create(role).then(res => resolve()).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    handleUpdatePolicy ({ state, commit }, role) {
      return new Promise((resolve, reject) => {
        try {
          PolicyAPI.put(role).then(res => resolve()).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    handleDeletePolicys ({ state, commit }, params) {
      return new Promise((resolve, reject) => {
        try {
          PolicyAPI.delete(params).then(res => resolve()).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    // 报表
    handlePolicyTable ({ state, commit }, params) {
      return new Promise((resolve, reject) => {
        try {
          PolicyAPI.table(params).then(res => resolve(res.data.data)).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    handlePolicyAll ({ state, commit }) {
      return new Promise((resolve, reject) => {
        try {
          PolicyAPI.all().then(res => resolve(res.data.data)).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    }
  }
}
