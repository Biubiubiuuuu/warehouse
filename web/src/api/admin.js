import request from '@/plugin/axios'

export function Login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function AdminsByLimitOffset(pageSize, page) {
  return request({
    url: '/queryAdmins',
    method: 'get',
    params: {
      pageSize,
      page
    }
  })
}

export function DeleteAdmin(id) {
  return request({
    url: '/deleteAdmin',
    method: 'delete',
    params: {
      id
    }
  })
}

export function DeleteBatchAdmin(data) {
  return request({
    url: '/deleteAdmins',
    method: 'post',
    data
  })
}

export function AddAdmin(data) {
  return request({
    url: '/addAdmin',
    method: 'post',
    data
  })
}

export function UpdateAdminPassword(data) {
  return request({
    url: '/updateAdminPass',
    method: 'post',
    data
  })
}