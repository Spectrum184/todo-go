package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "Invalid request", err.Error(), "INVALID_REQUEST")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong with server", err.Error(), "INTERNAL_SERVER")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot list %s", strings.ToLower(entity)), fmt.Sprintf("CANNOT_LIST_%s", strings.ToUpper(entity)))
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)), fmt.Sprintf("CANNOT_DELETE_%s", strings.ToUpper(entity)))
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot update %s", strings.ToLower(entity)), fmt.Sprintf("CANNOT_UPDATE_%s", strings.ToUpper(entity)))
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot create %s", strings.ToLower(entity)), fmt.Sprintf("CANNOT_CREATE_%s", strings.ToUpper(entity)))
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot get %s", strings.ToLower(entity)), fmt.Sprintf("CANNOT_GET_%s", strings.ToUpper(entity)))
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s deleted", strings.ToLower(entity)), fmt.Sprintf("%s_DELETED", strings.ToUpper(entity)))
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s is existed", strings.ToLower(entity)), fmt.Sprintf("%s_EXISTED", strings.ToUpper(entity)))
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s not found", strings.ToLower(entity)), fmt.Sprintf("%s_NOT_FOUND", strings.ToUpper(entity)))
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("You have no permission"), fmt.Sprintf("NO_PERMISSION"))
}

var RecordNotFound = errors.New("record not found")
