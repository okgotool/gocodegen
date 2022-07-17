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
	SysRoleService = SysRoleServiceType{}
	SysRoleConn    = query.Use(dal.CreateDbConn("SysRole"))
	SysRoleDao     = SysRoleConn.SysRole
)

type SysRoleServiceType struct {
}

// 解析orderby 字符串为表达式，orderBy格式： "id desc,name desc"
func (c *SysRoleServiceType) GetOrderByExpr(orderBy string) ([]field.Expr, error) {
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
			orderCol, ok := SysRoleDao.GetFieldByName(strings.Trim(orderColStr, " "))
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
		orderByIdDescy := SysRoleDao.ID.Desc()
		orderBys = append(orderBys, orderByIdDescy)
	}

	return orderBys, nil
}

// 查询所有,limits<0 时无限制
func (c *SysRoleServiceType) QueryAll(wheres []field.Expr, orderBys []field.Expr, page int, pageSize int) ([]*model.SysRole, int64, error) {
	if len(orderBys) < 1 {
		orderBys = []field.Expr{SysRoleDao.ID.Desc()}
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	tx := SysRoleDao.WithContext(dal.QueryCtx)
	for _, where := range wheres {
		tx = tx.Where(where)
	}

	// if have deleted column:
	tx = tx.Where(SysRoleDao.Deleted.Eq(0))

	for _, orderBy := range orderBys {
		tx = tx.Order(orderBy)
	}
	total, _ := tx.Count()
	rs, err := tx.Offset(offset).Limit(pageSize).Find()

	return rs, total, err
}

// 按id查询一个
func (c *SysRoleServiceType) QueryById(id int64) (*model.SysRole, error) {
	rs, err := SysRoleDao.WithContext(dal.QueryCtx).Where(SysRoleDao.ID.Eq(id)).First()

	// 其它处理...

	return rs, err
}

// 批量创建，成功返回成功插入的行数
func (c *SysRoleServiceType) CreateBatch(models []*model.SysRole, batchSize int) ([]*model.SysRole, error) {
	for _, model := range models {
		model.ID = 0
	}
	err := SysRoleDao.WithContext(dal.QueryCtx).CreateInBatches(models, batchSize)

	return models, err

}

// 创建一个，成功返回的对象有记录ID
func (c *SysRoleServiceType) Create(model *model.SysRole) (*model.SysRole, error) {
	model.ID = 0
	err := SysRoleDao.WithContext(dal.QueryCtx).Create(model)

	return model, err

}

// 按条件更新一批指定的字段，成功返回影响的行数
func (c *SysRoleServiceType) UpdateColumns(whereExpr field.Expr, columns map[string]interface{}) (int64, error) {
	if whereExpr == nil {
		return 0, errors.New("whereExpr is null!")
	}
	info, err := SysRoleDao.WithContext(dal.QueryCtx).Where(whereExpr).Updates(columns)

	return info.RowsAffected, err

}

// 批量更新，成功返回成功的行数
func (c *SysRoleServiceType) UpdateBatch(models []*model.SysRole) (int64, error) {
	for _, model := range models {
		if model.ID < 1 {
			return 0, errors.New("Some model ID not giving!")
		}
	}

	var rowsAffected int64 = 0
	err := SysRoleConn.Transaction(func(tx *query.Query) error {
		for _, model := range models {
			info, err := tx.WithContext(dal.QueryCtx).SysRole.Where(SysRoleDao.ID.Eq(model.ID)).Updates(model)
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
func (c *SysRoleServiceType) Update(model *model.SysRole) (int64, error) {
	if model.ID < 1 {
		return 0, errors.New("ID not giving!")
	}
	info, err := SysRoleDao.WithContext(dal.QueryCtx).Where(SysRoleDao.ID.Eq(model.ID)).Updates(model)

	return info.RowsAffected, err
}

// 硬删除记录，成功返回影响的行数
func (c *SysRoleServiceType) Delete(ids []int64) (int64, error) {
	if len(ids) < 1 {
		return 0, errors.New("ID not giving!")
	}
	info, err := SysRoleDao.WithContext(dal.QueryCtx).Where(SysRoleDao.ID.In(ids...)).Delete()

	return info.RowsAffected, err
}


// 软删除，设置deleted=1，成功返回影响的行数, 有deleted字段时才有该方法
func (c *SysRoleServiceType) SoftDelete(ids []int64) (int64, error) {
	if len(ids) < 1 {
		return 0, errors.New("ID not giving!")
	}
	info, err := SysRoleDao.WithContext(dal.QueryCtx).Where(SysRoleDao.ID.In(ids...)).UpdateColumn(SysRoleDao.Deleted, true)

	return info.RowsAffected, err
}

// 清除已软删除的记录，成功返回影响的行数, 有deleted字段时才有该方法
func (c *SysRoleServiceType) Clean() (int64, error) {
	info, err := SysRoleDao.WithContext(dal.QueryCtx).Where(SysRoleDao.Deleted).Delete()

	return info.RowsAffected, err
}
