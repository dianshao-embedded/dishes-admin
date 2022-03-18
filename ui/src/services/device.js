import {request} from '@/utils/request'

/**
 * 示例表单
 */

export function deviceQuery(params) {
    return request('/api/v1/devices', 'get', params)
}

export function deviceAdd(data) {
    return request('/api/v1/devices', 'post', data)
}

export function deviceGet(id) {
    return request(`/api/v1/devices/${id}`, 'get')
}

export function deviceDelete(id) {
    return request(`/api/v1/devices/${id}`, 'delete')
}

export function deviceUpdate(id, data) {
    return request(`/api/v1/devices/${id}`, 'put', data)
}

export function deviceImport(data) {
    return request('/api/v1/devices/import', 'post', data)
}

export default {
    deviceQuery,
    deviceAdd,
    deviceGet,
    deviceDelete,
    deviceUpdate,
}
