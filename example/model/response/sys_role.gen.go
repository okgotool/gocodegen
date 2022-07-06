package response

import (
	"github.com/okgotool/gocodegen/example/dal/model"
)

type (
	SuccessSysRoleArray struct {
		Code StatusCode       `json:"code"`
		Msg  string           `json:"msg"`
		Data []*model.SysRole `json:"data"`
	}
)
