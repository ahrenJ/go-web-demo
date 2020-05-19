package responseutil

const (
	DefaultCode = 0
	ErrorCode   = 1
	DefaultMsg  = "success"
)

func Success(data interface{}, msg string) map[string]interface{} {
	var responseBody = make(map[string]interface{})
	responseBody["code"] = DefaultCode
	responseBody["msg"] = msg
	responseBody["data"] = data
	return responseBody
}

func SuccessWithMsg(msg string) map[string]interface{} {
	var responseBody = make(map[string]interface{})
	responseBody["code"] = DefaultCode
	responseBody["msg"] = msg
	return responseBody
}

func SuccessWithData(data interface{}) map[string]interface{} {
	var responseBody = make(map[string]interface{})
	responseBody["code"] = DefaultCode
	responseBody["msg"] = DefaultMsg
	responseBody["data"] = data
	return responseBody
}

func Error(msg string) map[string]interface{} {
	var responseBody = make(map[string]interface{})
	responseBody["code"] = ErrorCode
	responseBody["msg"] = msg
	return responseBody
}

func ErrorArgs() map[string]interface{} {
	var responseBody = make(map[string]interface{})
	responseBody["code"] = ErrorCode
	responseBody["msg"] = "invalid arguments"
	return responseBody
}
