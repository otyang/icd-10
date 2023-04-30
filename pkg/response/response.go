package response

import (
	"net/http"
	"runtime"
)

// Response is the response structure.
type Response struct {
	StatusCode        int     `json:"-"`
	Success           bool    `json:"success"`
	Message           string  `json:"message"`
	ErrorCode         *string `json:"errorCode,omitempty"`
	Data              any     `json:"data,omitempty"`
	RuntimeCallerInfo string  `json:"-"`
}

// Lets do this, so the structs satisfy's go error interface
func (a Response) Error() string {
	return a.Message
}

func NewError(statusCode int, msg string, errorCode *string) *Response {
	return &Response{
		StatusCode: statusCode,
		Success:    false,
		Message:    msg,
		ErrorCode:  errorCode,
	}
}

func NewSuccess(statusCode int, msg string, data any) *Response {
	if msg == "" {
		msg = "Request was succesful"
	}

	return &Response{
		StatusCode: statusCode,
		Success:    true,
		Message:    msg,
		Data:       data,
	}
}

func Ok(msg string, data any) *Response {
	if msg == "" {
		msg = "Request was succesful"
	}
	return NewSuccess(http.StatusOK, msg, data)
}

func Created(msg string, data any) *Response {
	if msg == "" {
		msg = "Request was succesful"
	}
	return NewSuccess(http.StatusCreated, msg, data)
}

func BadRequest(msg string, errorCode *string) *Response {
	if msg == "" {
		msg = "Your request is in a bad format"
	}

	return NewError(
		http.StatusBadRequest,
		msg,
		setErrCode(errorCode, "bad_request"),
	)
}

func Unauthorized(msg string, errorCode *string) *Response {
	if msg == "" {
		msg = "You are not authenticated to perform the requested action"
	}

	return NewError(
		http.StatusUnauthorized,
		msg,
		setErrCode(errorCode, "un_authorized"),
	)
}

func Forbidden(msg string, errorCode *string) *Response {
	if msg == "" {
		msg = "You are not authorized to perform the requested action"
	}

	return NewError(
		http.StatusForbidden,
		msg,
		setErrCode(errorCode, "forbidden"),
	)
}

func NotFound(msg string, errorCode *string) *Response {
	if msg == "" {
		msg = "The requested resource was not found"
	}

	return NewError(
		http.StatusNotFound,
		msg,
		setErrCode(errorCode, "not_found"),
	)
}

func Conflict(msg string, errorCode *string) *Response {
	if msg == "" {
		msg = "The requested resource was not found"
	}

	return NewError(
		http.StatusConflict,
		msg,
		setErrCode(errorCode, "conflict"),
	)
}

func InternalServerError(msg string, errorCode *string) *Response {
	if msg == "" {
		msg = "Something went wrong on our end."
	}

	resp := NewError(
		http.StatusInternalServerError,
		msg,
		setErrCode(errorCode, "internal_server_error"),
	)

	pc, file, lineNo, ok := runtime.Caller(1)
	resp.RuntimeCallerInfo = getRuntimeCallerInfo(pc, file, lineNo, ok)
	return resp
}

func setErrCode(errorCode *string, defaultCode string) *string {
	if errorCode == nil {
		errorCode = &defaultCode
	}
	return errorCode
}
