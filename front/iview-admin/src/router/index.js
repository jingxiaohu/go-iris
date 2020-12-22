import Vue from 'vue'
import Router from 'vue-router'
import routes from './routers'
import store from '@/store'
import iView from 'iview'
import { setToken, getToken, canTurnTo, setTitle } from '@/libs/util'
import config from '@/config'
import { LoadMenu } from '@/api/routers'

const { homeName } = config

Vue.use(Router)
let router = new Router({
  routes
  /*
  // 前后端分离时可开启该模式
  base: '/admin/',
  mode: 'history',
   */
})
const LOGIN_PAGE_NAME = 'login'

const turnTo = (to, access, next) => {
  // console.log('localRoutes=>', router.options.routes, router.options.routes.length)
  // to.name === null是页面刷新的情况
  if (to.name === null || to.name === 'home') LoadMenu(router)
  if (to.name && (to.name.startsWith('add_') || to.name.startsWith('modify_'))) LoadMenu(router)
  next()
  // if (canTurnTo(to.name, access, routes)) { // 有权限，可访问
  //   RouterMenu.init(router)
  //   next()
  // } else next({ replace: true, name: 'error_401' }) // 无权限，重定向到401页面
}

router.beforeEach((to, from, next) => {
  iView.LoadingBar.start()
  const token = getToken()
  console.log('to name=', to.name)
  // console.log('token=', token)
  if (!token && to.name !== LOGIN_PAGE_NAME) {
    // 未登录且要跳转的页面不是登录页
    next({ name: LOGIN_PAGE_NAME })
  } else if (!token && to.name === LOGIN_PAGE_NAME) {
    // 未登陆且要跳转的页面是登录页
    next() // 跳转
  } else if (token && to.name === LOGIN_PAGE_NAME) {
    // 已登录且要跳转的页面是登录页
    next({ name: homeName })
  } else {
    console.log(1)
    if (store.state.user.hasGetInfo) {
      console.log(11)
      turnTo(to, store.state.user.access, next)
    } else {
      console.log(12)
      store.dispatch('getUserInfo').then(user => {
        // 拉取用户信息，通过用户权限和跳转的页面的name来判断是否有权限访问;access必须是一个数组，如：['super_admin'] ['super_admin', 'admin']
        turnTo(to, user.access, next)
      }).catch(() => {
        console.log(13)
        setToken('')
        next({ name: 'login' })
      })
    }
  }
})

router.afterEach(to => {
  setTitle(to, router.app)
  iView.LoadingBar.finish()
  window.scrollTo(0, 0)
})

export default router
