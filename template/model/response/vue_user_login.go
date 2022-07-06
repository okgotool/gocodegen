package response

type (
	VueLoginResponse struct {
		Code StatusCode            `json:"code"`
		Data *VueLoginResponseData `json:"data"`
	}

	VueLoginResponseData struct {
		Token string `json:"token"`
	}

	VueLoginFailedResponse struct {
		Code StatusCode `json:"code"`
		Msg  string     `json:"message"`
	}

	VueGetUserInfoResponse struct {
		Code StatusCode          `json:"code"`
		Data *VueGetUserInfoData `json:"data"`
	}

	VueGetUserInfoData struct {
		Name         string   `json:"name"`
		Roles        []string `json:"roles"`
		Introduction string   `json:"introduction"`
		Avatar       string   `json:"avatar"`
	}

	VueLogout struct {
		Code int64  `json:"code"`
		Data string `json:"data"`
	}
)
