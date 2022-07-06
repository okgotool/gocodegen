package api

import "github.com/gin-gonic/gin"

func AddApis(api *gin.RouterGroup) {
	v1 := api.Group("v1")

	addSysRoleApis(v1)
	addTestCatApis(v1)
	// {otherApisPlaceHolder}

}


// addSysRoleApis
func addSysRoleApis(v1 *gin.RouterGroup) {
    // 查询：
    v1.GET("/sysroles/all", SysRoleApi.QueryAll)
    v1.GET("/sysroles", SysRoleApi.QueryByCondition)
 
    // 创建:
    v1.POST("/sysroles", SysRoleApi.CreateBatch)
    v1.POST("/sysrole", SysRoleApi.Create)
 
    // 修改:
    v1.PUT("/sysroles", SysRoleApi.UpdateBatch)
    v1.PUT("/sysrole", SysRoleApi.Update)
 
    // 删除:
    v1.DELETE("/sysroles", SysRoleApi.Remove)
    
	// 软删除
	v1.DELETE("/sysroles/soft", SysRoleApi.SoftRemove)
	// 清除软删除记录
	v1.DELETE("/sysroles/clean", SysRoleApi.Clean)

}

// addTestCatApis
func addTestCatApis(v1 *gin.RouterGroup) {
    // 查询：
    v1.GET("/testcats/all", TestCatApi.QueryAll)
    v1.GET("/testcats", TestCatApi.QueryByCondition)
 
    // 创建:
    v1.POST("/testcats", TestCatApi.CreateBatch)
    v1.POST("/testcat", TestCatApi.Create)
 
    // 修改:
    v1.PUT("/testcats", TestCatApi.UpdateBatch)
    v1.PUT("/testcat", TestCatApi.Update)
 
    // 删除:
    v1.DELETE("/testcats", TestCatApi.Remove)

}