package notificationhubs

import (
	"fmt"
	"net/http"
)

// ErrorCode represents specific error types that can occur
type ErrorCode string

const (
	// ErrorCodeInvalidConnectionString indicates an invalid connection string
	ErrorCodeInvalidConnectionString ErrorCode = "INVALID_CONNECTION_STRING"
	// ErrorCodeAuthenticationFailed indicates authentication failure
	ErrorCodeAuthenticationFailed ErrorCode = "AUTHENTICATION_FAILED"
	// ErrorCodeUnauthorized indicates unauthorized access
	ErrorCodeUnauthorized ErrorCode = "UNAUTHORIZED"

	// ErrorCodeInvalidRequest indicates an invalid request
	ErrorCodeInvalidRequest ErrorCode = "INVALID_REQUEST"
	// ErrorCodeInvalidPayload indicates an invalid payload
	ErrorCodeInvalidPayload ErrorCode = "INVALID_PAYLOAD"
	// ErrorCodePayloadTooLarge indicates the payload is too large
	ErrorCodePayloadTooLarge ErrorCode = "PAYLOAD_TOO_LARGE"
	// ErrorCodeInvalidTags indicates invalid tags
	ErrorCodeInvalidTags ErrorCode = "INVALID_TAGS"

	// ErrorCodeServerError indicates a server error
	ErrorCodeServerError ErrorCode = "SERVER_ERROR"
	// ErrorCodeServiceUnavailable indicates the service is unavailable
	ErrorCodeServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"
	// ErrorCodeTimeout indicates a timeout occurred
	ErrorCodeTimeout ErrorCode = "TIMEOUT"

	// ErrorCodeRateLimited indicates rate limiting is active
	ErrorCodeRateLimited ErrorCode = "RATE_LIMITED"
	// ErrorCodeQuotaExceeded indicates quota has been exceeded
	ErrorCodeQuotaExceeded ErrorCode = "QUOTA_EXCEEDED"

	// ErrorCodeRegistrationNotFound indicates registration was not found
	ErrorCodeRegistrationNotFound ErrorCode = "REGISTRATION_NOT_FOUND"
	// ErrorCodeInvalidRegistration indicates invalid registration
	ErrorCodeInvalidRegistration ErrorCode = "INVALID_REGISTRATION"

	// ErrorCodeInstallationNotFound indicates installation was not found
	ErrorCodeInstallationNotFound ErrorCode = "INSTALLATION_NOT_FOUND"
	// ErrorCodeInvalidInstallation indicates invalid installation
	ErrorCodeInvalidInstallation ErrorCode = "INVALID_INSTALLATION"
)

// NotificationHubError represents an error from the notification hub service
type NotificationHubError struct {
	Code       ErrorCode
	Message    string
	Details    string
	StatusCode int
	RequestID  string
	Cause      error
}

// Error implements the error interface
func (e *NotificationHubError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("notification hub error [%s]: %s - %s", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("notification hub error [%s]: %s", e.Code, e.Message)
}

// Unwrap implements the error unwrapping interface for Go 1.13+
func (e *NotificationHubError) Unwrap() error {
	return e.Cause
}

// Is implements error comparison for Go 1.13+
func (e *NotificationHubError) Is(target error) bool {
	if t, ok := target.(*NotificationHubError); ok {
		return e.Code == t.Code
	}
	return false
}

// IsRetryable returns true if the error indicates the request can be retried
func (e *NotificationHubError) IsRetryable() bool {
	switch e.Code {
	case ErrorCodeServerError, ErrorCodeServiceUnavailable, ErrorCodeTimeout:
		return true
	case ErrorCodeRateLimited:
		return true // with backoff
	default:
		return false
	}
}

// IsAuthenticationError returns true if the error is related to authentication
func (e *NotificationHubError) IsAuthenticationError() bool {
	switch e.Code {
	case ErrorCodeInvalidConnectionString, ErrorCodeAuthenticationFailed, ErrorCodeUnauthorized:
		return true
	default:
		return false
	}
}

// NewError creates a new NotificationHubError
func NewError(code ErrorCode, message string) *NotificationHubError {
	return &NotificationHubError{
		Code:    code,
		Message: message,
	}
}

// NewErrorWithCause creates a new NotificationHubError with a cause
func NewErrorWithCause(code ErrorCode, message string, cause error) *NotificationHubError {
	return &NotificationHubError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

// NewErrorFromHTTPResponse creates an error from an HTTP response
func NewErrorFromHTTPResponse(resp *http.Response, body []byte) *NotificationHubError {
	err := &NotificationHubError{
		StatusCode: resp.StatusCode,
		RequestID:  resp.Header.Get("x-ms-request-id"),
	}

	switch resp.StatusCode {
	case http.StatusBadRequest:
		err.Code = ErrorCodeInvalidRequest
		err.Message = "Bad request"
	case http.StatusUnauthorized:
		err.Code = ErrorCodeUnauthorized
		err.Message = "Unauthorized"
	case http.StatusForbidden:
		err.Code = ErrorCodeAuthenticationFailed
		err.Message = "Authentication failed"
	case http.StatusNotFound:
		err.Code = ErrorCodeRegistrationNotFound
		err.Message = "Resource not found"
	case http.StatusRequestEntityTooLarge:
		err.Code = ErrorCodePayloadTooLarge
		err.Message = "Payload too large"
	case http.StatusTooManyRequests:
		err.Code = ErrorCodeRateLimited
		err.Message = "Rate limited"
	case http.StatusInternalServerError:
		err.Code = ErrorCodeServerError
		err.Message = "Internal server error"
	case http.StatusServiceUnavailable:
		err.Code = ErrorCodeServiceUnavailable
		err.Message = "Service unavailable"
	case http.StatusGatewayTimeout:
		err.Code = ErrorCodeTimeout
		err.Message = "Gateway timeout"
	default:
		err.Code = ErrorCodeServerError
		err.Message = fmt.Sprintf("HTTP %d", resp.StatusCode)
	}

	if len(body) > 0 {
		err.Details = string(body)
	}

	return err
}

// ValidationError represents input validation errors
type ValidationError struct {
	Field   string
	Message string
	Value   interface{}
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s': %s (value: %v)", e.Field, e.Message, e.Value)
}

// NewValidationError creates a new validation error
func NewValidationError(field, message string, value interface{}) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
		Value:   value,
	}
}

// MultiError represents multiple errors
type MultiError struct {
	Errors []error
}

// Error implements the error interface
func (e *MultiError) Error() string {
	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}
	return fmt.Sprintf("multiple errors occurred (%d errors)", len(e.Errors))
}

// Add adds an error to the multi-error
func (e *MultiError) Add(err error) {
	if err != nil {
		e.Errors = append(e.Errors, err)
	}
}

// HasErrors returns true if there are any errors
func (e *MultiError) HasErrors() bool {
	return len(e.Errors) > 0
}

// ToError returns the multi-error as a single error, or nil if no errors
func (e *MultiError) ToError() error {
	if !e.HasErrors() {
		return nil
	}
	return e
}

// NewMultiError creates a new multi-error
func NewMultiError() *MultiError {
	return &MultiError{
		Errors: make([]error, 0),
	}
}
 