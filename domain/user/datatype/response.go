package datatype

type GeneralResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type GenerateTokenResponse struct {
	StatusCode int                       `json:"status_code"`
	Message    string                    `json:"message"`
	Data       GenerateTokenDataResponse `json:"data"`
}

type UserLoginResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type GenerateTokenDataResponse struct {
	Token string `json:"token"`
}
