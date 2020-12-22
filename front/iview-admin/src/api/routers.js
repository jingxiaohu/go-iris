import { MenuAPI } from './menu'
import Main from '@/components/main'
import store from '@/store'

export const LoadMenu = async (router) => {
  let d1 = await MenuAPI.allMenu()
  let d2 = await MenuAPI.OwnMenu()
  let AllMenu = d1.data.data || []
  let OwnMenuIds = d2.data.data || []

  let localRoutes = router.options.routes
  console.log('*加载菜单。路由长度=', localRoutes.length)
  if (localRoutes.length > 2) {
    localRoutes.forEach(lr => {
      let pMenu = AllMenu.find(e => e.id === lr.id) // 用新的后端的菜单数据做对比，而不用route里的信息做对比
      if (pMenu && pMenu.meta) { // 必须这样修改嵌套对象vue才能生效
        let { access, beforeCloseName, hideInBread, hideInMenu, icon, notCache, showAlways } = pMenu.meta
        lr.meta.access = access
        lr.meta.beforeCloseName = beforeCloseName
        lr.meta.hideInBread = hideInBread
        lr.meta.hideInMenu = hideInMenu
        lr.meta.icon = icon
        lr.meta.notCache = notCache
        lr.meta.showAlways = showAlways
      }
      // console.log(lr)
      // if (pMenu) console.log('**********', pMenu.meta, OwnMenuIds.some(d => d === lr.id))
      // console.log(lr.meta.hideInMenu, !lr.meta.hideInMenu)
      // 父级
      if (lr.id && OwnMenuIds.some(d => d === lr.id) && !lr.meta.hideInMenu) lr.meta.hideInMenu = false // 菜单显示
      else lr.meta.hideInMenu = true // 隐藏
      // 子级
      if (lr.children && lr.children.length > 0) {
        lr.children.forEach(lrc => {
          if (pMenu && pMenu.children) {
            let cMenu = pMenu.children.find(e => e.id === lrc.id)
            if (cMenu && cMenu.meta) {
              let { access, beforeCloseName, hideInBread, hideInMenu, icon, notCache, showAlways } = cMenu.meta
              lrc.meta.access = access
              lrc.meta.beforeCloseName = beforeCloseName
              lrc.meta.hideInBread = hideInBread
              lrc.meta.hideInMenu = hideInMenu
              lrc.meta.icon = icon
              lrc.meta.notCache = notCache
              lrc.meta.showAlways = showAlways
            }
          }
          if (lrc.id && OwnMenuIds.some(d => d === lrc.id) && !lrc.meta.hideInMenu) {
            lrc.meta.hideInMenu = false
            lr.meta.hideInMenu = false
          } else lrc.meta.hideInMenu = true
        })
      }
    })
    return
  }
  InitRouterConfig(router, store, AllMenu || [])
}

// 动态添加路由
const InitRouterConfig = (router, store, menus) => {
  console.log('*初始化路由')
  let frs = BuildParentRoutes(menus)
  router.options.routes.push(...frs)
  router.addRoutes(router.options.routes)
  store.commit('setRoutes', menus)
}

/**
 * 构建路由父节点
 * https://router.vuejs.org/zh/api/#router-构建选项
 * @param menus
 * @returns {Array}
 */
const BuildParentRoutes = (menus) => {
  let parentRoutes = []
  menus.forEach(menu => {
    let {
      id,
      path,
      name,
      meta,
      children
    } = menu
    let pr = {
      id: id,
      path: '/' + path,
      component: Main,
      name: name,
      // icon: icon,
      meta: meta,
      children: children && BuildChildenRoutes(children)
    }
    parentRoutes.push(pr)
  })
  return parentRoutes
}

/**
 * 构建路由子节点
 * @param childs
 * @returns {Array}
 */
const BuildChildenRoutes = (childs) => {
  let result = []
  childs.forEach(child => {
    let {
      id,
      path,
      name,
      modular,
      component,
      meta,
      children
    } = child

    if (children && children instanceof Array) {
      children.forEach(ch => {
        if (ch && ch instanceof Array) {
          children = BuildChildenRoutes(ch)
        }
      })
    }

    let fcr = {
      id: id,
      path: path,
      component: resolve => require(['@/view' + modular + component + '.vue'], resolve),
      name: name,
      meta: meta,
      children: children
    }
    result.push(fcr)
  })
  return result
}
