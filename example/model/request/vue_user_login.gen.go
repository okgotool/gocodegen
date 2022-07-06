package request

type (
	VueUserLogin struct {
		UserName string `json:"username"`
		Password string `json:"password"`
		Code     string `json:"code"`
		UUID     string `json:"uuid"`
	}
)
