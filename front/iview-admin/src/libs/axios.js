import axios from 'axios'
import { getToken } from '@/libs/util'
import store from '@/store'
import router from '@/router'
import { Message, Spin } from 'view-design'

class HttpRequest {
  constructor (baseUrl = baseURL) {
    this.baseUrl = baseUrl
    this.queue = {}
  }
  getInsideConfig () {
    const config = {
      baseURL: this.baseUrl,
      headers: {
        //
      }
    }
    return config
  }
  destroy (url) {
    delete this.queue[url]
    if (!Object.keys(this.queue).length) {
      // Spin.hide()
    }
  }
  interceptors (instance, url) {
    // 请求拦截
    instance.interceptors.request.use(config => {
      // 添加全局的loading...
      if (!Object.keys(this.queue).length) {
        // Spin.show() // 不建议开启，因为界面不友好
      }
      this.queue[url] = true
      config.headers['Authorization'] = 'bearer ' + getToken()
      return config
    }, error => {
      return Promise.reject(error)
    })
    // 响应拦截
    instance.interceptors.response.use(res => {
      this.destroy(url)
      const { data, status } = res
      return { data, status }
    }, err => {
      this.destroy(url)
      console.log(err.response)
      let status = err.response.status
      let msg = err.response.data.msg
      let code = err.response.data.code
      if (code === '3001') store.commit('logout', [])
      switch (status) {
        // case 400: Message.error({message: (msg != null && msg !== '') ? msg : '操作失败,参数不对!', showClose: true}); break
        case 401:
          Message.warning({ content: msg || '当前会话已过期,请重新登录', closable: true })
          router.replace({ name: 'login' })
          break
        case 404: Message.error({ content: msg || '请求的资源不存在!', closable: true }); break
        case 403: Message.error({ content: msg || '权限不足,请联系管理员', closable: true }); break
        case 500: Message.error({ content: msg || '服务器错误', closable: true }); break
        case 504: Message.error({ content: msg || '网关连接超时', closable: true }); break
        default: Message.warning({ content: msg || '失败，未知原因', closable: true })
      }
      return Promise.reject(err)
    })
  }
  request (options) {
    const instance = axios.create()
    options = Object.assign(this.getInsideConfig(), options)
    this.interceptors(instance, options.url)
    return instance(options)
  }
}
export default HttpRequest
