package utils

type StatusCode int

const (
	StatusOK                  StatusCode = 200
	StatusCreated             StatusCode = 201
	StatusNoContent           StatusCode = 204
	StatusBadRequest          StatusCode = 400
	StatusUnauthorized        StatusCode = 401
	StatusForbidden           StatusCode = 403
	StatusNotFound            StatusCode = 404
	StatusInternalServerError StatusCode = 500
)

type APIError struct {
	Code    StatusCode
	Message string
}

func (e *APIError) Error() string {
	return e.Message
}

func NewError(message string, code StatusCode) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

type APIResponse struct {
	Error   *APIError   `json:"error,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Success bool        `json:"success"`
}

func ToAPIError(err error) *APIError {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr
	}
	return NewError(err.Error(), StatusInternalServerError)
}
