import {request} from '@/utils/request'

/**
 * 示例表单
 */

export function firmwareQuery(params) {
    return request('/api/v1/firmwares', 'get', params)
}

export function firmwareAdd(data) {
    return request('/api/v1/firmwares', 'post', data)
}

export function firmwareGet(id) {
    return request(`/api/v1/firmwares/${id}`, 'get')
}

export function firmwareDelete(id) {
    return request(`/api/v1/firmwares/${id}`, 'delete')
}

export function firmwareUpdate(id, data) {
    return request(`/api/v1/firmwares/${id}`, 'put', data)
}

export function firmwareImport(data) {
    return request('/api/v1/firmwares/import', 'post', data)
}

export default {
    firmwareQuery,
    firmwareAdd,
    firmwareGet,
    firmwareDelete,
    firmwareUpdate,
}
