import request from '@/plugin/axios'

export function QueryUserByLimitOffset(pageSize, page) {
  return request({
    url: '/users/queryUserByLimitOffset',
    method: 'get',
    params: {
      pageSize,
      page
    }
  })
}