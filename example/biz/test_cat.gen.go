package biz

import (
	"errors"
	"strings"

	"github.com/okgotool/gocodegen/example/dal"
	"github.com/okgotool/gocodegen/example/dal/model"
	"github.com/okgotool/gocodegen/example/dal/query"
	"gorm.io/gen/field"
)

var (
	TestCatService = TestCatServiceType{}
	TestCatConn    = query.Use(dal.CreateDbConn("TestCat"))
	TestCatDao     = TestCatConn.TestCat
)

type TestCatServiceType struct {
}

// 解析orderby 字符串为表达式，orderBy格式： "id desc,name desc"
func (c *TestCatServiceType) GetOrderByExpr(orderBy string) ([]field.Expr, error) {
	orderBys := []field.Expr{}
	orderBy = strings.Trim(orderBy, " ")
	if len(orderBy) < 1 {
		return orderBys, nil
	}

	orderByStrs := strings.Split(orderBy, ",")
	for _, orderByStr := range orderByStrs {
		if len(orderByStr) > 0 {

			orderByStr = strings.ToLower(strings.Trim(orderByStr, " "))
			orderColStr := strings.Split(orderByStr, " ")[0]
			orderCol, ok := TestCatDao.GetFieldByName(strings.Trim(orderColStr, " "))
			if !ok {
				return orderBys, errors.New("排序字段未找到：" + orderColStr)
			}

			// 添加orderby表达式
			if strings.HasSuffix(orderByStr, " desc") { // 倒序
				orderBys = append(orderBys, orderCol.Desc())
			} else {
				orderBys = append(orderBys, orderCol)
			}
		}
	}
	
	// if not set, set default order by id desc
	if len(orderBy) < 1 {
		orderByIdDescy := TestCatDao.ID.Desc()
		orderBys = append(orderBys, orderByIdDescy)
	}

	return orderBys, nil
}

// 查询所有,limits<0 时无限制
func (c *TestCatServiceType) QueryAll(wheres []field.Expr, orderBys []field.Expr, page int, pageSize int) ([]*model.TestCat, int64, error) {
	if len(orderBys) < 1 {
		orderBys = []field.Expr{TestCatDao.ID.Desc()}
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	tx := TestCatDao.WithContext(dal.QueryCtx)
	for _, where := range wheres {
		tx = tx.Where(where)
	}

	// if have deleted column:
	// {QueryDeletedWhereCondition}

	for _, orderBy := range orderBys {
		tx = tx.Order(orderBy)
	}
	total, _ := tx.Count()
	rs, err := tx.Offset(offset).Limit(pageSize).Find()

	return rs, total, err
}

// 按id查询一个
func (c *TestCatServiceType) QueryById(id int64) (*model.TestCat, error) {
	rs, err := TestCatDao.WithContext(dal.QueryCtx).Where(TestCatDao.ID.Eq(id)).First()

	// 其它处理...

	return rs, err
}

// 批量创建，成功返回成功插入的行数
func (c *TestCatServiceType) CreateBatch(models []*model.TestCat, batchSize int) ([]*model.TestCat, error) {
	for _, model := range models {
		model.ID = 0
	}
	err := TestCatDao.WithContext(dal.QueryCtx).CreateInBatches(models, batchSize)

	return models, err

}

// 创建一个，成功返回的对象有记录ID
func (c *TestCatServiceType) Create(model *model.TestCat) (*model.TestCat, error) {
	model.ID = 0
	err := TestCatDao.WithContext(dal.QueryCtx).Create(model)

	return model, err

}

// 按条件更新一批指定的字段，成功返回影响的行数
func (c *TestCatServiceType) UpdateColumns(whereExpr field.Expr, columns map[string]interface{}) (int64, error) {
	if whereExpr == nil {
		return 0, errors.New("whereExpr is null!")
	}
	info, err := TestCatDao.WithContext(dal.QueryCtx).Where(whereExpr).Updates(columns)

	return info.RowsAffected, err

}

// 批量更新，成功返回成功的行数
func (c *TestCatServiceType) UpdateBatch(models []*model.TestCat) (int64, error) {
	for _, model := range models {
		if model.ID < 1 {
			return 0, errors.New("Some model ID not giving!")
		}
	}

	var rowsAffected int64 = 0
	err := TestCatConn.Transaction(func(tx *query.Query) error {
		for _, model := range models {
			info, err := tx.WithContext(dal.QueryCtx).TestCat.Where(TestCatDao.ID.Eq(model.ID)).Updates(model)
			if err != nil {
				rowsAffected = 0
				return err
			} else {
				rowsAffected = rowsAffected + info.RowsAffected
			}
		}
		return nil
	})

	return rowsAffected, err
}

// 更新一个，成功返回影响的行数
func (c *TestCatServiceType) Update(model *model.TestCat) (int64, error) {
	if model.ID < 1 {
		return 0, errors.New("ID not giving!")
	}
	info, err := TestCatDao.WithContext(dal.QueryCtx).Where(TestCatDao.ID.Eq(model.ID)).Updates(model)

	return info.RowsAffected, err
}

// 硬删除记录，成功返回影响的行数
func (c *TestCatServiceType) Delete(ids []int64) (int64, error) {
	if len(ids) < 1 {
		return 0, errors.New("ID not giving!")
	}
	info, err := TestCatDao.WithContext(dal.QueryCtx).Where(TestCatDao.ID.In(ids...)).Delete()

	return info.RowsAffected, err
}

