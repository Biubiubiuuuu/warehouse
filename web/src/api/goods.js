import request from '@/plugin/axios'

export function QueryGoodsTypesByLimitOffset(pageSize, page) {
  return request({
    url: '/goodsType/queryGoodsTypesByLimitOffset',
    method: 'get',
    params: {
      pageSize,
      page
    }
  })
}

export function DeleteGoodsType(id) {
  return request({
    url: '/goodsType/deleteGoodsType',
    method: 'delete',
    params: {
      id
    }
  })
}

export function DeleteGoodsTypes(data) {
  return request({
    url: '/goodsType/deleteGoodsTypes',
    method: 'post',
    data
  })
}

export function AddGoodsType(data) {
  return request({
    url: '/goodsType/addGoodsType',
    method: 'post',
    data
  })
}

export function UpdateGoodsType(data) {
  return request({
    url: '/goodsType/updateGoodsType',
    method: 'put',
    data
  })
}