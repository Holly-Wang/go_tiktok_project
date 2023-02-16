package common

// Common Status
const (
	StatusOk = iota

	StausUnkown
)

const (
	LoginSuccess = iota
	LoginFailed
)

const (
	LoginSuccessMsg = "login success"
)

const (
	RegisterSucces = iota
	RegisterFailed
)

const (
	RegisterSueecssMsg = "register success"
)

const (
	TokenSuccess = iota
    TokenFailed 
)

const (
	TokenSuccessMsg = "token success"
    TokenFailedMsg ="token failed , please login again."
)

const (
	PublishSuccess = iota
    PublishFailed 
)

const (
	PublishSuccessMsg = "publish success"
    PublishFailedMsg="publish failed"
)

const (
	GetUserVideoSuccess = iota
    GetUserVideoFailed 
)

const (
	GetUserVideoSuccessMsg = "get user video success"
    GetUserVideoFailedMsg="get user video failed"
)