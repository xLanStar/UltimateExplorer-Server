package APIException

import "net/http"

type APIException struct {
	Code    int    `json:"-"`
	Type    int    `json:"type"`
	Msg     string `json:"msg"`
	Request string `json:"request"`
}

func (e *APIException) Error() string {
	return e.Msg
}

func newAPIException(code int, Type int, msg string) *APIException {
	return &APIException{
		Code: code,
		Type: Type,
		Msg:  msg,
	}
}

const (
	ERROR       = 1
	WARN        = 2
	INFORMATION = 4
)

func UnknownError(message string) *APIException {
	return newAPIException(http.StatusForbidden, ERROR, message)
}

func NewError(message string) *APIException {
	return newAPIException(http.StatusBadRequest, ERROR, message)
}

func NewWarn(message string) *APIException {
	return newAPIException(http.StatusBadRequest, WARN, message)
}

func NewInformation(message string) *APIException {
	return newAPIException(http.StatusBadRequest, INFORMATION, message)
}
