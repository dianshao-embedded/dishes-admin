import {request} from '@/utils/request'

/**
 * 示例表单
 */

export function demoQuery() {
    return request('/api/v1/demos', 'get')
}

export function demoAdd(data) {
    return request('/api/v1/demos', 'post', data)
}

export function demoGet(id) {
    return request(`/api/v1/demos/${id}`, 'get')
}

export function demoDelete(id) {
    return request(`/api/v1/demos/${id}`, 'delete')
}

export function demoUpdate(id, data) {
    return request(`/api/v1/demos/${id}`, 'put', data)
}

export default {
    demoQuery,
    demoAdd,
    demoGet,
    demoDelete,
    demoUpdate,
}
