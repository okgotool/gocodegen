package response

// common response structs:
type (
	// StatusCode : enum
	StatusCode int

	// Success :
	Success struct {
		Code StatusCode `json:"code"`
		Msg  string     `json:"msg"`
	}

	// SuccessUint :
	SuccessInt struct {
		Code StatusCode `json:"code"`
		Msg  string     `json:"msg"`
		Data int        `json:"data"`
	}

	// SuccessUint64 :
	SuccessInt64 struct {
		Code StatusCode `json:"code"`
		Msg  string     `json:"msg"`
		Data int64      `json:"data"`
	}
	// SuccessMap :
	SuccessMap struct {
		Code StatusCode        `json:"code"`
		Msg  string            `json:"msg"`
		Data map[string]string `json:"data"`
	}

	// SuccessMapInt :
	SuccessMapInt struct {
		Code StatusCode     `json:"code"`
		Msg  string         `json:"msg"`
		Data map[string]int `json:"data"`
	}

	// SuccessMapInt :
	SuccessMapInt64 struct {
		Code StatusCode       `json:"code"`
		Msg  string           `json:"msg"`
		Data map[string]int64 `json:"data"`
	}

	// SuccessMapArray :
	SuccessMapArray struct {
		Code StatusCode          `json:"code"`
		Msg  string              `json:"msg"`
		Data []map[string]string `json:"data"`
	}
	// ResponseIntMap :
	ResponseIntMap struct {
		Code StatusCode  `json:"code"`
		Msg  string      `json:"msg"`
		Data map[int]int `json:"data"`
	}

	// SuccessArray :
	SuccessArray struct {
		Code StatusCode `json:"code"`
		Msg  string     `json:"msg"`
		Data []string   `json:"data"`
	}

	// SuccessArray :
	SuccessInt64Array struct {
		Code StatusCode `json:"code"`
		Msg  string     `json:"msg"`
		Data []int64    `json:"data"`
	}

	// FailedWithReason :
	FailedWithReason struct {
		Code StatusCode        `json:"code"`
		Msg  string            `json:"msg"`
		Data map[string]string `json:"data"` // failed reason, one reason one map item
	}

	ResponseInterface struct {
		Code StatusCode  `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
)
