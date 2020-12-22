import { Get, Post, Put, Delete } from './zz_restful'

export const UserAPI = {
  login (user) {
    return Post('/user/login', user)
  },
  getUserBaseInfo (id) {
    return Get('/user/userBaseInfo', { id: id })
  },
  logout () {
    return Get('/user/logout')
  },
  create (user) {
    return Post('/user', user)
  },
  put (user) {
    return Put('/user', user)
  },
  delete (params) {
    return Delete('/user', params)
  },
  table (params) {
    return Post('/user/table', params)
  }
}
// ---------------
export const login = ({ userName, password }) => {
  const data = {
    userName,
    password
  }
  return axios.request({
    url: 'user/login',
    data,
    method: 'post'
  })
}

export const getUserInfo = (token) => {
  return axios.request({
    url: 'get_info',
    params: {
      token
    },
    method: 'get'
  })
}

export const logout = () => {
  return axios.request({
    url: 'users/logout',
    method: 'post'
  })
}

export const getUnreadCount = () => {
  return axios.request({
    url: 'message/count',
    method: 'get'
  })
}

export const getMessage = () => {
  return axios.request({
    url: 'message/init',
    method: 'get'
  })
}

export const getContentByMsgId = msg_id => {
  return axios.request({
    url: 'message/content',
    method: 'get',
    params: {
      msg_id
    }
  })
}

export const hasRead = msg_id => {
  return axios.request({
    url: 'message/has_read',
    method: 'post',
    data: {
      msg_id
    }
  })
}

export const removeReaded = msg_id => {
  return axios.request({
    url: 'message/remove_readed',
    method: 'post',
    data: {
      msg_id
    }
  })
}

export const restoreTrash = msg_id => {
  return axios.request({
    url: 'message/restore',
    method: 'post',
    data: {
      msg_id
    }
  })
}
