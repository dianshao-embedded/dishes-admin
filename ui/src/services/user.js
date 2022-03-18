import {request, removeAuthorization} from '@/utils/request'

/**
 * 登录服务
 * @param name 账户名
 * @param password 账户密码
 * @returns {Promise<AxiosResponse<T>>}
 */
export async function login(data) {
  return request('/api/v1/pub/login', 'post', data)
}

export async function getUser() {
  return request('/api/v1/pub/current/user', 'get')
}

/**
 * 退出登录
 */
export function logout() {
  localStorage.removeItem(process.env.VUE_APP_ROUTES_KEY)
  localStorage.removeItem(process.env.VUE_APP_PERMISSIONS_KEY)
  localStorage.removeItem(process.env.VUE_APP_ROLES_KEY)
  removeAuthorization()

  return request(`/api/v1/pub/login/exit`, 'post')
}

export function userQuery() {
  return request('/api/v1/users', 'get')
}

export function userAdd(data) {
  return request('/api/v1/users', 'post', data)
}

export function userGet(id) {
  return request(`/api/v1/users/${id}`, 'get')
}

export function userDelete(id) {
  console.log('delete')
  return request(`/api/v1/users/${id}`, 'delete')
}

export function userDisable(id) {
  return request(`/api/v1/users/${id}/disable`, 'patch')
}

export function userEnable(id) {
  return request(`/api/v1/users/${id}/enable`, 'patch')
}

export function userUpdate(id, data) {
  return request(`/api/v1/users/${id}`, 'put', data)
}

export function getRoles() {
  return request(`/api/v1/roles.select`, 'get')
}

export default {
  login,
  logout,
  userQuery,
  userAdd,
  userGet,
  userDisable,
  userEnable,
  userUpdate,
  getRoles,
}
