package api

import (
	"strconv"
	"strings"
	"time"

	"github.com/okgotool/gocodegen/example/biz"
	"github.com/okgotool/gocodegen/example/dal/model"
	"github.com/okgotool/gocodegen/example/model/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gen/field"
)

var (
	SysRoleApi = &SysRoleApiType{}

	SysRoleInsertBatchSize = 100
)

type SysRoleApiType struct {
}

// Swagger doc refer: https://github.com/swaggo/swag
// @Summary 查询所有，支持分页、排序
// @Schemes
// @Description 查询所有，支持分页、排序
// @Tags SysRole
// @Param orderBy query string false "orderBy" default()
// @Param page query int false "page" default(1)
// @Param pageSize query int false "pageSize" default(10)
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessSysRoleArray "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数错误\"}"
// @Router /api/v1/sysroles/all [get]
// @Security ApiKeyAuth
func (c *SysRoleApiType) QueryAll(g *gin.Context) {
	page, pageSize, errMsgs := getPagerValuesFromQuery(g)

	orderBy := g.Query("orderBy")
	orderBys, err := biz.SysRoleService.GetOrderByExpr(orderBy)
	if err != nil {
		errMsgs["orderBy"] = err.Error()
	}

	if len(errMsgs) > 0 {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code: response.StatusBadRequest,
			Msg:  "参数错误",
			Data: errMsgs,
		})
		return
	}

	rs, total, err := biz.SysRoleService.QueryAll(nil, orderBys, page, pageSize)

	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code:  response.StatusInternalServerError,
			Msg:   "执行失败",
			Data: map[string]string{"error": err.Error()},
		})
	} else {
		g.JSON(response.OKCode, &response.SuccessSysRoleArray{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: rs,
			Total: total,
		})
	}
}

// Swagger doc refer: https://github.com/swaggo/swag
// @Summary 条件查询，支持分页、排序
// @Schemes
// @Description 条件查询，支持分页、排序
// @Tags SysRole
// @Param id query int false "id" default()
// @Param roleName query string false "RoleName" default()
// @Param roleNameEn query string false "RoleNameEn" default()
// @Param status query string false "Status, 数字，多个时逗号隔开" default()
// @Param priority query string false "Priority, 数字，多个时逗号隔开" default()
// @Param comment query string false "Comment" default()
// @Param deleted query string false "Deleted, 数字，多个时逗号隔开" default()
// @Param lastmodifiedBy query string false "LastmodifiedBy" default()
// @Param lastmodifiedMin query int64 false "Lastmodified 起始时间, 毫秒数时间戳，查询大于等于lastmodifiedMin 该时间的数据" default()
// @Param lastmodifiedMax query int64 false "Lastmodified 结束时间, 毫秒数时间戳，查询小于lastmodifiedMax 该时间的数据" default()
// @Param createdBy query string false "CreatedBy" default()
// @Param createdMin query int64 false "Created 起始时间, 毫秒数时间戳，查询大于等于createdMin 该时间的数据" default()
// @Param createdMax query int64 false "Created 结束时间, 毫秒数时间戳，查询小于createdMax 该时间的数据" default()
// @Param createdAtMin query int64 false "CreatedAt 起始时间, 毫秒数时间戳，查询大于等于createdAtMin 该时间的数据" default()
// @Param createdAtMax query int64 false "CreatedAt 结束时间, 毫秒数时间戳，查询小于createdAtMax 该时间的数据" default()
// @Param updatedBy query string false "UpdatedBy" default()
// @Param updatedAtMin query int64 false "UpdatedAt 起始时间, 毫秒数时间戳，查询大于等于updatedAtMin 该时间的数据" default()
// @Param updatedAtMax query int64 false "UpdatedAt 结束时间, 毫秒数时间戳，查询小于updatedAtMax 该时间的数据" default()
// @Param orderBy query string false "orderBy" default()
// @Param page query int false "page" default(1)
// @Param pageSize query int false "pageSize" default(10)
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessSysRoleArray "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数错误\"}"
// @Router /api/v1/sysroles [get]
// @Security ApiKeyAuth
func (c *SysRoleApiType) QueryByCondition(g *gin.Context) {
	idStr := g.Query("id")

	whereConditions := []field.Expr{}
	page, pageSize, errMsgs := getPagerValuesFromQuery(g)

	orderBy := g.Query("orderBy")
	orderBys, err := biz.SysRoleService.GetOrderByExpr(orderBy)
	if err != nil {
		errMsgs["orderBy"] = err.Error()
	}

	if len(idStr) > 0 && !strings.EqualFold(idStr, "0") { // 按id查询
		id, err := parseInt64(idStr)
		if err != nil {
			errMsgs["id"] = err.Error()
		} else {
			whereConditions = append(whereConditions, biz.SysRoleDao.ID.Eq(id))
		}
	} else { // 多条件查询

		// 其它查询条件处理：
		if len(g.Query("roleName")) > 0 {
			queryValue := g.Query("roleName")
			whereConditions = append(whereConditions, biz.SysRoleDao.RoleName.Eq(queryValue))
		}
		if len(g.Query("roleNameEn")) > 0 {
			queryValue := g.Query("roleNameEn")
			whereConditions = append(whereConditions, biz.SysRoleDao.RoleNameEn.Eq(queryValue))
		}

		if len(g.Query("status")) > 0 {
			queryValues := []int32{}
			queryStrs := strings.Split(g.Query("status"), ",")
			for _, queryStr := range queryStrs {
				queryValue, err := parseInt64(queryStr)
				if err != nil {
				} else {
					queryValues = append(queryValues, int32(queryValue))
				}
			}
			if len(queryValues) > 0 {
				whereConditions = append(whereConditions, biz.SysRoleDao.Status.In(queryValues...))
			}
		}


		if len(g.Query("priority")) > 0 {
			queryValues := []int32{}
			queryStrs := strings.Split(g.Query("priority"), ",")
			for _, queryStr := range queryStrs {
				queryValue, err := parseInt64(queryStr)
				if err != nil {
				} else {
					queryValues = append(queryValues, int32(queryValue))
				}
			}
			if len(queryValues) > 0 {
				whereConditions = append(whereConditions, biz.SysRoleDao.Priority.In(queryValues...))
			}
		}

		if len(g.Query("comment")) > 0 {
			queryValue := g.Query("comment")
			whereConditions = append(whereConditions, biz.SysRoleDao.Comment.Eq(queryValue))
		}

		if len(g.Query("deleted")) > 0 {
			queryValues := []int32{}
			queryStrs := strings.Split(g.Query("deleted"), ",")
			for _, queryStr := range queryStrs {
				queryValue, err := parseInt64(queryStr)
				if err != nil {
				} else {
					queryValues = append(queryValues, int32(queryValue))
				}
			}
			if len(queryValues) > 0 {
				whereConditions = append(whereConditions, biz.SysRoleDao.Deleted.In(queryValues...))
			}
		}

		if len(g.Query("lastmodifiedBy")) > 0 {
			queryValue := g.Query("lastmodifiedBy")
			whereConditions = append(whereConditions, biz.SysRoleDao.LastmodifiedBy.Eq(queryValue))
		}

		// query data of lastmodified between lastmodifiedMin and lastmodifiedMax:
		if len(g.Query("lastmodifiedMin")) > 0 {
			lastmodifiedMills, err := strconv.ParseInt(g.Query("lastmodifiedMin"), 10, 64)
			if err == nil {
				lastmodifiedMin := time.Unix(lastmodifiedMills/1000, 0)
				whereConditions = append(whereConditions, biz.SysRoleDao.Lastmodified.Gte(lastmodifiedMin))
			}
		}
		if len(g.Query("lastmodifiedMax")) > 0 {
			lastmodifiedMills, err := strconv.ParseInt(g.Query("lastmodified"), 10, 64)
			if err == nil {
				lastmodifiedMax := time.Unix(lastmodifiedMills/1000, 0)
				whereConditions = append(whereConditions, biz.SysRoleDao.Lastmodified.Lt(lastmodifiedMax))
			}
		}

		if len(g.Query("createdBy")) > 0 {
			queryValue := g.Query("createdBy")
			whereConditions = append(whereConditions, biz.SysRoleDao.CreatedBy.Eq(queryValue))
		}

		// query data of created between createdMin and createdMax:
		if len(g.Query("createdMin")) > 0 {
			createdMills, err := strconv.ParseInt(g.Query("createdMin"), 10, 64)
			if err == nil {
				createdMin := time.Unix(createdMills/1000, 0)
				whereConditions = append(whereConditions, biz.SysRoleDao.Created.Gte(createdMin))
			}
		}
		if len(g.Query("createdMax")) > 0 {
			createdMills, err := strconv.ParseInt(g.Query("created"), 10, 64)
			if err == nil {
				createdMax := time.Unix(createdMills/1000, 0)
				whereConditions = append(whereConditions, biz.SysRoleDao.Created.Lt(createdMax))
			}
		}


		// query data of createdAt between createdAtMin and createdAtMax:
		if len(g.Query("createdAtMin")) > 0 {
			createdAtMills, err := strconv.ParseInt(g.Query("createdAtMin"), 10, 64)
			if err == nil {
				createdAtMin := time.Unix(createdAtMills/1000, 0)
				whereConditions = append(whereConditions, biz.SysRoleDao.CreatedAt.Gte(createdAtMin))
			}
		}
		if len(g.Query("createdAtMax")) > 0 {
			createdAtMills, err := strconv.ParseInt(g.Query("createdAt"), 10, 64)
			if err == nil {
				createdAtMax := time.Unix(createdAtMills/1000, 0)
				whereConditions = append(whereConditions, biz.SysRoleDao.CreatedAt.Lt(createdAtMax))
			}
		}

		if len(g.Query("updatedBy")) > 0 {
			queryValue := g.Query("updatedBy")
			whereConditions = append(whereConditions, biz.SysRoleDao.UpdatedBy.Eq(queryValue))
		}

		// query data of updatedAt between updatedAtMin and updatedAtMax:
		if len(g.Query("updatedAtMin")) > 0 {
			updatedAtMills, err := strconv.ParseInt(g.Query("updatedAtMin"), 10, 64)
			if err == nil {
				updatedAtMin := time.Unix(updatedAtMills/1000, 0)
				whereConditions = append(whereConditions, biz.SysRoleDao.UpdatedAt.Gte(updatedAtMin))
			}
		}
		if len(g.Query("updatedAtMax")) > 0 {
			updatedAtMills, err := strconv.ParseInt(g.Query("updatedAt"), 10, 64)
			if err == nil {
				updatedAtMax := time.Unix(updatedAtMills/1000, 0)
				whereConditions = append(whereConditions, biz.SysRoleDao.UpdatedAt.Lt(updatedAtMax))
			}
		}



	}

	if len(errMsgs) > 0 {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code: response.StatusBadRequest,
			Msg:  "参数错误",
			Data: errMsgs,
		})
		return
	}

	rs, total, err := biz.SysRoleService.QueryAll(whereConditions, orderBys, page, pageSize)

	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code:  response.StatusInternalServerError,
			Msg:   "执行失败",
			Data: map[string]string{"error": err.Error()},
		})
	} else {
		g.JSON(response.OKCode, &response.SuccessSysRoleArray{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: rs,
			Total: total, 
		})
	}
}

// Swagger doc refer: https://github.com/swaggo/swag
// @Summary 批量创建
// @Schemes
// @Description 批量创建，返回新创建的记录的ID列表
// @Tags SysRole
// @Param body body []model.SysRole{} true "json format" default()
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessInt64Array "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数格式错误或为空\"}"
// @Router /api/v1/sysroles [post]
// @Security ApiKeyAuth
func (c *SysRoleApiType) CreateBatch(g *gin.Context) {
	newIds := []int64{}
	models := []*model.SysRole{}

	err := g.ShouldBind(&models)
	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code: response.StatusBadRequest,
			Msg:  "参数格式错误或为空",
			Data: map[string]string{"error": err.Error()},
		})
		return
	} else if len(models) < 1 {
		g.JSON(response.OKCode, &response.SuccessInt64Array{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: newIds,
		})
		return
	}

	rs, err := biz.SysRoleService.CreateBatch(models, SysRoleInsertBatchSize)

	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code:  response.StatusInternalServerError,
			Msg:   "执行失败",
			Data: map[string]string{"error": err.Error()},
		})
	} else {
        for _, model := range rs {
            if model.ID > 0 {
                newIds = append(newIds, model.ID)
            }
        }

        // 返回新创建的记录的ID列表
		g.JSON(response.OKCode, &response.SuccessInt64Array{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: newIds,
		})
	}
}

// Swagger doc refer: https://github.com/swaggo/swag
// @Summary 创建一个
// @Schemes
// @Description 创建一个，返回创建的记录的ID
// @Tags SysRole
// @Param body body model.SysRole{} true "json format" default()
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessInt64 "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数格式错误\"}"
// @Router /api/v1/sysrole [post]
// @Security ApiKeyAuth
func (c *SysRoleApiType) Create(g *gin.Context) {

	model := &model.SysRole{}

	err := g.ShouldBind(model)
	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code: response.StatusBadRequest,
			Msg:  "参数格式错误",
			Data: map[string]string{"error": err.Error()},
		})
		return
	}

	rs, err := biz.SysRoleService.Create(model)

	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code:  response.StatusInternalServerError,
			Msg:   "执行失败",
			Data: map[string]string{"error": err.Error()},
		})
	} else {
		g.JSON(response.OKCode, &response.SuccessInt64{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: rs.ID,
		})
	}
}

// Swagger doc refer: https://github.com/swaggo/swag
// @Summary 批量更新
// @Schemes
// @Description 批量更新，返回被修改的行数
// @Tags SysRole
// @Param body body []model.SysRole{} true "json format" default()
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessInt64 "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数格式错误或为空\"}"
// @Router /api/v1/sysroles [put]
// @Security ApiKeyAuth
func (c *SysRoleApiType) UpdateBatch(g *gin.Context) {
	models := []*model.SysRole{}

	err := g.ShouldBind(&models)
	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code: response.StatusBadRequest,
			Msg:  "参数格式错误或为空",
			Data: map[string]string{"error": err.Error()},
		})
		return
	} else if len(models) < 1 {
		g.JSON(response.OKCode, &response.SuccessInt64Array{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: []int64{},
		})
		return
	}

	effectRows, err := biz.SysRoleService.UpdateBatch(models)

	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code:  response.StatusInternalServerError,
			Msg:   "执行失败",
			Data: map[string]string{"error": err.Error()},
		})
	} else {
		g.JSON(response.OKCode, &response.SuccessInt64{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: effectRows,
		})
	}
}

// Swagger doc refer: https://github.com/swaggo/swag
// @Summary 更新一个
// @Schemes
// @Description 更新一个，返回被修改的行数
// @Tags SysRole
// @Param body body model.SysRole{} true "json format" default()
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessInt64 "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数格式错误\"}"
// @Router /api/v1/sysrole [put]
// @Security ApiKeyAuth
func (c *SysRoleApiType) Update(g *gin.Context) {
	model := &model.SysRole{}

	err := g.ShouldBind(model)
	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code: response.StatusBadRequest,
			Msg:  "参数格式错误",
			Data: map[string]string{"error": err.Error()},
		})
		return
	}

	effectRows, err := biz.SysRoleService.Update(model)

	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code:  response.StatusInternalServerError,
			Msg:   "执行失败",
			Data: map[string]string{"error": err.Error()},
		})
	} else {
		g.JSON(response.OKCode, &response.SuccessInt64{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: effectRows,
		})
	}
}

// Swagger doc refer: https://github.com/swaggo/swag
// @Summary 硬删除，删除记录，不可找回
// @Schemes
// @Description 硬删除，删除记录，不可找回，可以传多个id，逗号隔开，返回被删除的行数
// @Tags SysRole
// @Param id query string true "id，多个时逗号隔开" default()
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessInt64 "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数格式错误\"}"
// @Router /api/v1/sysroles [delete]
// @Security ApiKeyAuth
func (c *SysRoleApiType) Remove(g *gin.Context) {

	idStr := g.Query("id")
	idStrs := strings.Split(idStr, ",")
	ids := []int64{}

	for _, idStr := range idStrs {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			g.JSON(response.FailedCode, &response.FailedWithReason{
			Code: response.StatusBadRequest,
			Msg:  "参数格式错误",
			Data: map[string]string{"error": err.Error()},
		    })
			return
		}

		ids = append(ids, id)
	}
    if len(ids) < 1 {
        g.JSON(response.OKCode, &response.SuccessInt64{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: 0,
		})
        return
    }

	effectRows, err := biz.SysRoleService.Delete(ids)

	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code:  response.StatusInternalServerError,
			Msg:   "执行失败",
			Data: map[string]string{"error": err.Error()},
		})
	} else {
		g.JSON(response.OKCode, &response.SuccessInt64{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: effectRows,
		})
	}
}



// Swagger doc refer: https://github.com/swaggo/swag
// @Summary 软删除，更新deleted=1
// @Schemes
// @Description 软删除，更新deleted=1，返回被删除的行数
// @Tags SysRole
// @Param id query string true "id，多个时逗号隔开" default()
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessInt64 "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数错误\"}"
// @Router /api/v1/sysroles/soft [delete]
// @Security ApiKeyAuth
func (c *SysRoleApiType) SoftRemove(g *gin.Context) {

	idStr := g.Query("id")
	idStrs := strings.Split(idStr, ",")
	ids := []int64{}

	for _, idStr := range idStrs {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			g.JSON(response.FailedCode, &response.FailedWithReason{
			    Code: response.StatusBadRequest,
			    Msg:  "参数格式错误",
			    Data: map[string]string{"error": err.Error()},
		    })
			return
		}

		ids = append(ids, id)
	}
    if len(ids) < 1 {
        g.JSON(response.OKCode, &response.SuccessInt64{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: 0,
		})
        return
    }

	effectRows, err := biz.SysRoleService.SoftDelete(ids)

	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code:  response.StatusInternalServerError,
			Msg:   "执行失败",
			Data: map[string]string{"error": err.Error()},
		})
	} else {
		g.JSON(response.OKCode, &response.SuccessInt64{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: effectRows,
		})
	}
}

// Swagger doc refer: https://github.com/swaggo/swag
// @Summary 清除deleted=1的所有记录
// @Schemes
// @Description 清除deleted=1的所有记录，返回被清除的记录行数
// @Tags SysRole
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessInt64 "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":500,\"msg\":\"执行失败\"}"
// @Router /api/v1/sysroles/clean [delete]
// @Security ApiKeyAuth
func (c *SysRoleApiType) Clean(g *gin.Context) {

	effectRows, err := biz.SysRoleService.Clean()

	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code:  response.StatusInternalServerError,
			Msg:   "执行失败",
			Data: map[string]string{"error": err.Error()},
		})
	} else {
		g.JSON(response.OKCode, &response.SuccessInt64{
			Code: response.StatusOK,
			Msg:  "ok",
			Data: effectRows,
		})
	}
}
