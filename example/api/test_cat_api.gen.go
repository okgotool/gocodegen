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
	TestCatApi = &TestCatApiType{}

	TestCatInsertBatchSize = 100
)

type TestCatApiType struct {
}

// Swagger doc refer: https://github.com/swaggo/swag
// @Summary 查询所有，支持分页、排序
// @Schemes
// @Description 查询所有，支持分页、排序
// @Tags TestCat
// @Param orderBy query string false "orderBy" default()
// @Param page query int false "page" default(1)
// @Param pageSize query int false "pageSize" default(10)
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessTestCatArray "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数错误\"}"
// @Router /api/v1/testcats/all [get]
// @Security ApiKeyAuth
func (c *TestCatApiType) QueryAll(g *gin.Context) {
	page, pageSize, errMsgs := getPagerValuesFromQuery(g)

	orderBy := g.Query("orderBy")
	orderBys, err := biz.TestCatService.GetOrderByExpr(orderBy)
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

	rs, total, err := biz.TestCatService.QueryAll(nil, orderBys, page, pageSize)

	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code:  response.StatusInternalServerError,
			Msg:   "执行失败",
			Data: map[string]string{"error": err.Error()},
		})
	} else {
		g.JSON(response.OKCode, &response.SuccessTestCatArray{
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
// @Tags TestCat
// @Param id query int false "id" default()
// @Param catName query string false "CatName" default()
// @Param createdBy query string false "CreatedBy" default()
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
// @Success 200 {object} response.SuccessTestCatArray "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数错误\"}"
// @Router /api/v1/testcats [get]
// @Security ApiKeyAuth
func (c *TestCatApiType) QueryByCondition(g *gin.Context) {
	idStr := g.Query("id")

	whereConditions := []field.Expr{}
	page, pageSize, errMsgs := getPagerValuesFromQuery(g)

	orderBy := g.Query("orderBy")
	orderBys, err := biz.TestCatService.GetOrderByExpr(orderBy)
	if err != nil {
		errMsgs["orderBy"] = err.Error()
	}

	if len(idStr) > 0 && !strings.EqualFold(idStr, "0") { // 按id查询
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			errMsgs["id"] = err.Error()
		} else {
			whereConditions = append(whereConditions, biz.TestCatDao.ID.Eq(id))
		}
	} else { // 多条件查询

		// 其它查询条件处理：
		if len(g.Query("catName")) > 0 {
			queryValue := g.Query("catName")
			whereConditions = append(whereConditions, biz.TestCatDao.CatName.Eq(queryValue))
		}
		if len(g.Query("createdBy")) > 0 {
			queryValue := g.Query("createdBy")
			whereConditions = append(whereConditions, biz.TestCatDao.CreatedBy.Eq(queryValue))
		}

		// query data of createdAt between createdAtMin and createdAtMax:
		if len(g.Query("createdAtMin")) > 0 {
			createdAtMills, err := strconv.ParseInt(g.Query("createdAtMin"), 10, 64)
			if err == nil {
				createdAtMin := time.Unix(createdAtMills/1000, 0)
				whereConditions = append(whereConditions, biz.TestCatDao.CreatedAt.Gte(createdAtMin))
			}
		}
		if len(g.Query("createdAtMax")) > 0 {
			createdAtMills, err := strconv.ParseInt(g.Query("createdAt"), 10, 64)
			if err == nil {
				createdAtMax := time.Unix(createdAtMills/1000, 0)
				whereConditions = append(whereConditions, biz.TestCatDao.CreatedAt.Lt(createdAtMax))
			}
		}

		if len(g.Query("updatedBy")) > 0 {
			queryValue := g.Query("updatedBy")
			whereConditions = append(whereConditions, biz.TestCatDao.UpdatedBy.Eq(queryValue))
		}

		// query data of updatedAt between updatedAtMin and updatedAtMax:
		if len(g.Query("updatedAtMin")) > 0 {
			updatedAtMills, err := strconv.ParseInt(g.Query("updatedAtMin"), 10, 64)
			if err == nil {
				updatedAtMin := time.Unix(updatedAtMills/1000, 0)
				whereConditions = append(whereConditions, biz.TestCatDao.UpdatedAt.Gte(updatedAtMin))
			}
		}
		if len(g.Query("updatedAtMax")) > 0 {
			updatedAtMills, err := strconv.ParseInt(g.Query("updatedAt"), 10, 64)
			if err == nil {
				updatedAtMax := time.Unix(updatedAtMills/1000, 0)
				whereConditions = append(whereConditions, biz.TestCatDao.UpdatedAt.Lt(updatedAtMax))
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

	rs, total, err := biz.TestCatService.QueryAll(whereConditions, orderBys, page, pageSize)

	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code:  response.StatusInternalServerError,
			Msg:   "执行失败",
			Data: map[string]string{"error": err.Error()},
		})
	} else {
		g.JSON(response.OKCode, &response.SuccessTestCatArray{
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
// @Tags TestCat
// @Param body body []model.TestCat{} true "json format" default()
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessInt64Array "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数格式错误或为空\"}"
// @Router /api/v1/testcats [post]
// @Security ApiKeyAuth
func (c *TestCatApiType) CreateBatch(g *gin.Context) {
	newIds := []int64{}
	models := []*model.TestCat{}

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

	rs, err := biz.TestCatService.CreateBatch(models, TestCatInsertBatchSize)

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
// @Tags TestCat
// @Param body body model.TestCat{} true "json format" default()
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessInt64 "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数格式错误\"}"
// @Router /api/v1/testcat [post]
// @Security ApiKeyAuth
func (c *TestCatApiType) Create(g *gin.Context) {

	model := &model.TestCat{}

	err := g.ShouldBind(model)
	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code: response.StatusBadRequest,
			Msg:  "参数格式错误",
			Data: map[string]string{"error": err.Error()},
		})
		return
	}

	rs, err := biz.TestCatService.Create(model)

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
// @Tags TestCat
// @Param body body []model.TestCat{} true "json format" default()
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessInt64 "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数格式错误或为空\"}"
// @Router /api/v1/testcats [put]
// @Security ApiKeyAuth
func (c *TestCatApiType) UpdateBatch(g *gin.Context) {
	models := []*model.TestCat{}

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

	effectRows, err := biz.TestCatService.UpdateBatch(models)

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
// @Tags TestCat
// @Param body body model.TestCat{} true "json format" default()
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessInt64 "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数格式错误\"}"
// @Router /api/v1/testcat [put]
// @Security ApiKeyAuth
func (c *TestCatApiType) Update(g *gin.Context) {
	model := &model.TestCat{}

	err := g.ShouldBind(model)
	if err != nil {
		g.JSON(response.FailedCode, &response.FailedWithReason{
			Code: response.StatusBadRequest,
			Msg:  "参数格式错误",
			Data: map[string]string{"error": err.Error()},
		})
		return
	}

	effectRows, err := biz.TestCatService.Update(model)

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
// @Tags TestCat
// @Param id query string true "id，多个时逗号隔开" default()
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessInt64 "{\"code\":200,\"msg\":\"ok\"}"
// @Failure 210 {object} response.FailedWithReason "{\"code\":400,\"msg\":\"参数格式错误\"}"
// @Router /api/v1/testcats [delete]
// @Security ApiKeyAuth
func (c *TestCatApiType) Remove(g *gin.Context) {

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

	effectRows, err := biz.TestCatService.Delete(ids)

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

