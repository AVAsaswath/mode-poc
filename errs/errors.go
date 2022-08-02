package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
	ErrorId int    `json:"errorId"`
	//Status  string `json:"status"`
}

func (e AppError) Error() string {
	//TODO implement me
	panic("implement me")
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
		ErrorId: 0,
		//Status:  "Failed",
	}
}

func SuccessResponse(message string) *AppError {

	return &AppError{
		Code:    http.StatusOK,
		Message: message,
		ErrorId: 0,
	}
}

func Authentication(message string) *AppError {
	return &AppError{
		ErrorId: 0,
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusOK,
	}
}

func SecretValidation(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusOK,
	}
}

func EmptyResponse(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusOK,
	}
}
