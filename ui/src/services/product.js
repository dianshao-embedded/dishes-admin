import {request} from '@/utils/request'

/**
 * 示例表单
 */

export function productQuery(params) {
    return request('/api/v1/products', 'get', params)
}

export function productAdd(data) {
    return request('/api/v1/products', 'post', data)
}

export function productGet(id) {
    return request(`/api/v1/products/${id}`, 'get')
}

export function productDelete(id) {
    return request(`/api/v1/products/${id}`, 'delete')
}

export function productUpdate(id, data) {
    return request(`/api/v1/products/${id}`, 'put', data)
}

export function productImport(data) {
    return request('/api/v1/products/import', 'post', data)
}

export default {
    productQuery,
    productAdd,
    productGet,
    productDelete,
    productUpdate,
}
