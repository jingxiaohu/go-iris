import { MenuAPI } from '@/api/menu'
// import {  } from '@/libs/util'

export default {
  state: {
  },
  mutations: {
  },
  getters: {
  },
  actions: {
    handleAllMenu ({ state, commit }) {
      return new Promise((resolve, reject) => {
        try {
          MenuAPI.allMenu().then(res => resolve(res.data.data)).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    handleCreateMenu ({ state, commit }, menu) {
      return new Promise((resolve, reject) => {
        try {
          MenuAPI.create(menu).then(res => resolve()).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    handleUpdateMenu ({ state, commit }, menu) {
      return new Promise((resolve, reject) => {
        try {
          MenuAPI.put(menu).then(res => resolve()).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    handleDeleteMenus ({ state, commit }, params) {
      return new Promise((resolve, reject) => {
        try {
          MenuAPI.delete(params).then(res => resolve()).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    // 报表
    handleMenuTable ({ state, commit }, params) {
      return new Promise((resolve, reject) => {
        try {
          MenuAPI.table(params).then(res => resolve(res.data.data)).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    handleParentMenus ({ state, commit }) {
      return new Promise((resolve, reject) => {
        try {
          MenuAPI.parentMenus().then(res => resolve(res.data.data)).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    }
  }
}
