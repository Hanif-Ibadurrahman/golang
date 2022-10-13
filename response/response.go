package response

import "strings"

//Response is used for static shape json return
type Response struct {
	Status     bool        `json:"status"`
	StatusCode string      `json:"status_code"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors"`
	Data       interface{} `json:"data"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(statusCode string, message string, data interface{}) Response {
	res := Response{
		Status:     true,
		StatusCode: statusCode,
		Message:    message,
		Errors:     nil,
		Data:       data,
	}
	return res
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(statusCode string, message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:     false,
		StatusCode: statusCode,
		Message:    message,
		Errors:     splittedError,
		Data:       data,
	}
	return res
}
