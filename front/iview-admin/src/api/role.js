import { Get, Post, Put, Delete } from './zz_restful'

export const RoleAPI = {
  create (role) {
    return Post('/role', role)
  },
  put (role) {
    return Put('/role', role)
  },
  delete (params) {
    return Delete('/role', params)
  },
  domainList () {
    return Get('/role/domainList')
  },
  userGroupRoleList () {
    return Get('/role/group')
  },
  table (params) {
    return Post('/role/table', params)
  },
  roleOfMenus (roleKey) {
    return Get('/role/roleOfMenus/' + roleKey)
  },
  roleOfPolicys (tRole) {
    return Get(`/role/roleOfPolicys/${tRole.v1}/${tRole.v2}`)
  },
}
