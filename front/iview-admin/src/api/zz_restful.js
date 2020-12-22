import axios from '@/libs/api.request'

export const Get = (suffixUrl, data) => {
  return axios.request({
    url: suffixUrl,
    params: data,
    method: 'GET'
  })
}

export const Post = (suffixUrl, data) => {
  return axios.request({
    url: suffixUrl,
    data: data,
    method: 'POST'
  })
}

export const Put = (suffixUrl, data) => {
  return axios.request({
    url: suffixUrl,
    data: data,
    method: 'PUT'
  })
}

export const Delete = (suffixUrl, data) => {
  return axios.request({
    url: suffixUrl,
    params: data,
    method: 'DELETE'
  })
}
