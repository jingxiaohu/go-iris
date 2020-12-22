import { RoleAPI } from '@/api/role'
// import {  } from '@/libs/util'

export default {
  state: {
  },
  mutations: {
  },
  getters: {
  },
  actions: {
    handleCreateRole ({ state, commit }, role) {
      return new Promise((resolve, reject) => {
        try {
          RoleAPI.create(role).then(res => resolve()).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    handleUpdateRole ({ state, commit }, role) {
      return new Promise((resolve, reject) => {
        try {
          RoleAPI.put(role).then(res => resolve()).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    handleDeleteRoles ({ state, commit }, params) {
      return new Promise((resolve, reject) => {
        try {
          RoleAPI.delete(params).then(res => resolve()).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    // 获取租户列表
    handleDomainList ({ state, commit }) {
      return new Promise((resolve, reject) => {
        try {
          RoleAPI.domainList().then(res => resolve(res.data.data)).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    // 获取所有用户组角色列表
    handleUserGroupRoleList ({ state, commit }) {
      return new Promise((resolve, reject) => {
        try {
          RoleAPI.userGroupRoleList().then(res => resolve(res.data.data)).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    // 报表
    handleRoleTable ({ state, commit }, params) {
      return new Promise((resolve, reject) => {
        try {
          RoleAPI.table(params).then(res => resolve(res.data.data)).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    handleRoleOfMenus ({ state, commit }, roleKey) {
      return new Promise((resolve, reject) => {
        try {
          RoleAPI.roleOfMenus(roleKey).then(res => resolve(res.data.data)).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    },
    handleRoleOfPolicys ({ state, commit }, tRole) {
      return new Promise((resolve, reject) => {
        try {
          RoleAPI.roleOfPolicys(tRole).then(res => resolve(res.data.data)).catch(err => reject(err))
        } catch (err) { reject(err) }
      })
    }
  }
}
