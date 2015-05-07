//go:generate ffjson $GOFILE
package sgo

const (
	ErrorParse 			= "JSON Parse Error"
	ErrorApi 				= "Square API Error"
)

type Error struct {
	ErrorType string
	ApiError *ApiError
	ParseError error
}

type ApiError struct {
	ErrorCode int
	ErrorType string						`json:"type"`
	Message string							`json:"message"`
}

func MakeError(errType string, api *ApiError, parse error) *Error {
	return &Error{
		ErrorType: errType,
		ApiError: api,
		ParseError: parse,
	}
}

func MakeApiError(code int, errType string, message string) *Error {
	return &Error{
		ErrorType: ErrorApi,
		ApiError: &ApiError{
			ErrorCode: code,
			ErrorType: errType,
			Message: message,
		},
		ParseError: nil,
	}
}