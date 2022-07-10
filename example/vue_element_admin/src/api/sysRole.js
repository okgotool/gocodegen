import request from '@/utils/request'

// 查询所有
// 参数：page-页码, pageSize-每页记录数, orderBy-排序
export function querySysRoleAll(query) {
  return request({
    url: '/api/v1/sysroles/all',
    method: 'get',
    params: query
  })
}

// 按条件查询
// 参数：page-页码, pageSize-每页记录数, orderBy-排序
// id: int64,多个时，请用逗号隔开
// roleName: string,
// roleNameEn: string,
// status: int32,多个时，请用逗号隔开
// priority: int32,多个时，请用逗号隔开
// comment: string,
// deleted: int32,多个时，请用逗号隔开
// lastmodifiedBy: string,
// lastmodified: time.Time,
// createdBy: string,
// created: time.Time,
// createdAt: time.Time,
// updatedBy: string,
// updatedAt: time.Time,
export function querySysRoleByCondition(query) {
  return request({
    url: '/api/v1/sysroles',
    method: 'get',
    params: query
  })
}

// 批量创建
// 参数：对象数组
export function createSysRoleBatch(data) {
  return request({
    url: '/api/v1/sysroles',
    method: 'post',
    data
  })
}

// 创建一个
// 参数：对象
export function createSysRole(data) {
  return request({
    url: '/api/v1/sysrole',
    method: 'post',
    data
  })
}

// 批量更新
// 参数：对象数组，对象id必须设置，需要修改的属性则设置值
export function updateSysRoleBatch(data) {
  return request({
    url: '/api/v1/sysroles',
    method: 'put',
    data
  })
}

// 更新一个
// 参数：对象，对象id必须设置，需要修改的属性则设置值
export function updateSysRole(data) {
  return request({
    url: '/api/v1/sysrole',
    method: 'put',
    data
  })
}

// 删除，直接删除记录
// 参数：id，多个时逗号隔开
export function removeSysRole(query) {
  return request({
    url: '/api/v1/sysroles',
    method: 'delete',
    params: query
  })
}

// 软删除，设置deleted字段为1，只有有deleted字段才支持
// 参数：id，多个时逗号隔开
export function roftRemoveSysRole(query) {
  return request({
    url: '/api/v1/sysroles/soft',
    method: 'delete',
    params: query
  })
}

// 清除已被软删除的记录，即删除deleted字段为1的所有记录，只有有deleted字段才支持
export function cleanSysRole() {
  return request({
    url: '/api/v1/sysroles/clean',
    method: 'delete'
  })
}
