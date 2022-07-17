import request from '@/utils/request'

// Api详情：http://server_ip:port/swagger/index.html#/TestCat

// 查询列表
export function listTestCat(query) {
  return request({
    url: '/api/v1/testcats',
    method: 'get',
    params: query
  })
}

// 按ID查询详细
export function getTestCat(id) {
  return request({
    url: '/api/v1/testcats?id=' + id,
    method: 'get'
  })
}

// 新增
export function addTestCat(data) {
  return request({
    url: '/api/v1/testcat',
    method: 'post',
    data: data
  })
}

// 修改
export function updateTestCat(data) {
  return request({
    url: '/api/v1/testcat',
    method: 'put',
    data: data
  })
}

// 删除
export function delTestCat(id) {
  return request({
    url: '/api/v1/testcats?id=' + id,
    method: 'delete'
  })
}
