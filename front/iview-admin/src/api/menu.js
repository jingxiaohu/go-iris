import { Post, Put, Delete, Get } from './zz_restful'

export const MenuAPI = {
  create (menu) {
    return Post('/menu', menu)
  },
  put (menu) {
    return Put('/menu', menu)
  },
  delete (params) {
    return Delete('/menu', params)
  },
  table (params) {
    return Post('/menu/table', params)
  },
  parentMenus () {
    return Get('/menu/parentMenus')
  },
  allMenu () {
    return Get('/menu/all')
  },
  // 获取自己的菜单id
  OwnMenu () {
    return Get('/menu/own')
  }
}
