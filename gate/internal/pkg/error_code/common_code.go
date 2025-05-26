package error_code

var (
	Success                   = NewErrorData(0, "success")
	ServerError               = NewErrorData(10000000, "服務內部錯誤")
	InvalidParams             = NewErrorData(10000001, "參數錯誤")
	NotFound                  = NewErrorData(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewErrorData(10000003, "驗證失敗，token遺失")
	UnauthorizedTokenError    = NewErrorData(10000004, "驗證失敗，token錯誤")
	UnauthorizedTokenTimeout  = NewErrorData(10000005, "驗證失敗，token超時")
	UnauthorizedTokenGenerate = NewErrorData(10000006, "驗證失敗，token生成失敗")
	TooManyRequests           = NewErrorData(10000007, "請求過多")
)
