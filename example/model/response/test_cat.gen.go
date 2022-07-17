package response

import (
	"github.com/okgotool/gocodegen/example/dal/model"
)

type (
	SuccessTestCatArray struct {
		Code StatusCode       `json:"code"`
		Msg  string           `json:"msg"`
		Data []*model.TestCat `json:"data"`
		Total int64            `json:"total"`
	}
)
