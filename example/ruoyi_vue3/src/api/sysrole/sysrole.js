import request from '@/utils/request'

// Api详情：http://server_ip:port/swagger/index.html#/SysRole

// 查询列表
export function listSysRole(query) {
  return request({
    url: '/api/v1/sysroles',
    method: 'get',
    params: query
  })
}

// 按ID查询详细
export function getSysRole(id) {
  return request({
    url: '/api/v1/sysroles?id=' + id,
    method: 'get'
  })
}

// 新增
export function addSysRole(data) {
  return request({
    url: '/api/v1/sysrole',
    method: 'post',
    data: data
  })
}

// 修改
export function updateSysRole(data) {
  return request({
    url: '/api/v1/sysrole',
    method: 'put',
    data: data
  })
}

// 删除
export function delSysRole(id) {
  return request({
    url: '/api/v1/sysroles?id=' + id,
    method: 'delete'
  })
}
