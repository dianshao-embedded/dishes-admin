import {request} from '@/utils/request'

/**
 * 示例表单
 */

export function upgradeQuery(params) {
    return request('/api/v1/upgrades', 'get', params)
}

export function upgradeAdd(data) {
    return request('/api/v1/upgrades', 'post', data)
}

export function upgradeGet(id) {
    return request(`/api/v1/upgrades/${id}`, 'get')
}

export function upgradeDelete(id) {
    return request(`/api/v1/upgrades/${id}`, 'delete')
}

export function upgradeUpdate(id, data) {
    return request(`/api/v1/upgrades/${id}`, 'put', data)
}

export function upgradeImport(data) {
    return request('/api/v1/upgrades/import', 'post', data)
}

export default {
    upgradeQuery,
    upgradeAdd,
    upgradeGet,
    upgradeDelete,
    upgradeUpdate,
}
