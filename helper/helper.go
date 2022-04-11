package helper

import "github.com/go-playground/validator/v10"

//1. set cetakan response
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

//2. set cetakan meta
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

//3. set fungsi APIResponse()
func APIResponse(message string, code int, status string, data interface{}) Response {
	//-> set objek cetakan meta
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	//-> set objek cetakan response
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

//4. set fungsi FormatError()
func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
