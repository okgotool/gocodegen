import request from '@/utils/request'

// 查询所有
// 参数：page-页码, pageSize-每页记录数, orderBy-排序
export function queryTestCatAll(query) {
  return request({
    url: '/api/v1/testcats/all',
    method: 'get',
    params: query
  })
}

// 按条件查询
// 参数：page-页码, pageSize-每页记录数, orderBy-排序
// id: int64,多个时，请用逗号隔开
// catName: string,
// createdBy: string,
// createdAt: time.Time,
// updatedBy: string,
// updatedAt: time.Time,
export function queryTestCatByCondition(query) {
  return request({
    url: '/api/v1/testcats',
    method: 'get',
    params: query
  })
}

// 批量创建
// 参数：对象数组
export function createTestCatBatch(data) {
  return request({
    url: '/api/v1/testcats',
    method: 'post',
    data
  })
}

// 创建一个
// 参数：对象
export function createTestCat(data) {
  return request({
    url: '/api/v1/testcat',
    method: 'post',
    data
  })
}

// 批量更新
// 参数：对象数组，对象id必须设置，需要修改的属性则设置值
export function updateTestCatBatch(data) {
  return request({
    url: '/api/v1/testcats',
    method: 'put',
    data
  })
}

// 更新一个
// 参数：对象，对象id必须设置，需要修改的属性则设置值
export function updateTestCat(data) {
  return request({
    url: '/api/v1/testcat',
    method: 'put',
    data
  })
}

// 删除，直接删除记录
// 参数：id，多个时逗号隔开
export function removeTestCat(query) {
  return request({
    url: '/api/v1/testcats',
    method: 'delete',
    params: query
  })
}
