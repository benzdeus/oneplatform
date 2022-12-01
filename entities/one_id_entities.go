package entities

type OneIDService interface {
	// Register()
	LoginPWD(username, password string) interface{}
	// LoginSharedToken()
	// GetSharedToken()
	// RefreshToken()
	// LoginOauthWithSharedToken()
	// GetAccountData()
}

type ResponseOneIDLoginPWD struct {
	TokenType      string `json:"token_type"`
	ExpiresIn      int    `json:"expires_in"`
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	ExpirationDate string `json:"expiration_date"`
	AccountID      string `json:"account_id"`
	Result         string `json:"result"`
	UserName       string `json:"username"`
	ResponseOneIDFail
}

type ResponseOneIDFail struct {
	Result       string      `json:"result"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"`
	ResponseCode int         `json:"responseCode"`
}
