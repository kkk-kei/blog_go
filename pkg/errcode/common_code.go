package errcode

var (
	Success                   = NewError(0, "success")
	ServerError               = NewError(10000000, "internal server error")
	InvalidParams             = NewError(10000001, "invalid parameters")
	NotFound                  = NewError(10000002, "not found")
	UnauthorizedAuthNotExist  = NewError(10000003, "authorized failed : can't find AppKey and AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "authorized token failed : token error")
	UnauthorizedTokenTimeout  = NewError(10000005, "authorized token failed : token timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "authorized token failed : can't generate token")
	TooManyRequests           = NewError(10000007, "too many requests")
)
