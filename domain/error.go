package domain

type ErrorFormat struct {
	Code    int
	Message string
}

type ResponseFormat struct {
	Code int         `json:"code" example:"200"`
	Data interface{} `json:"data"`
}

var (
	ErrorUnauthorized        = ErrorFormat{Code: 401, Message: "Unauthorized"}
	ErrorForbidden           = ErrorFormat{Code: 403, Message: "Forbidden"}
	ErrorBadRequest          = ErrorFormat{Code: 400, Message: "bad request"}
	ErrorServer              = ErrorFormat{Code: 500, Message: "Server Error"}
	ErrorUnknowInternalError = ErrorFormat{Code: 404, Message: "Unknow Internal Error"}
	ErrUserAlreadyExists     = ErrorFormat{Code: 4001, Message: "User already exists"}
	ErrUserNotFound          = ErrorFormat{Code: 4002, Message: "User not found"}
)
