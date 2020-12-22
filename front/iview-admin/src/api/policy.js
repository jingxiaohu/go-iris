import { Get, Post, Put, Delete } from './zz_restful'

export const PolicyAPI = {
  create (role) {
    return Post('/policy', role)
  },
  put (role) {
    return Put('/policy', role)
  },
  delete (params) {
    return Delete('/policy', params)
  },
  table (params) {
    return Post('/policy/table', params)
  },
  all () {
    return Get('/policy/all')
  }
}
