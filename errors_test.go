package notificationhubs_test

import (
	"net/http"
	"testing"

	. "github.com/koreset/azure-notifications-sdk-go"
)

func Test_NotificationHubError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *NotificationHubError
		expected string
	}{
		{
			name: "Error with details",
			err: &NotificationHubError{
				Code:    ErrorCodeInvalidRequest,
				Message: "Invalid request",
				Details: "Missing required field",
			},
			expected: "notification hub error [INVALID_REQUEST]: Invalid request - Missing required field",
		},
		{
			name: "Error without details",
			err: &NotificationHubError{
				Code:    ErrorCodeServerError,
				Message: "Internal server error",
			},
			expected: "notification hub error [SERVER_ERROR]: Internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.expected {
				t.Errorf("NotificationHubError.Error() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_NotificationHubError_IsRetryable(t *testing.T) {
	tests := []struct {
		name     string
		err      *NotificationHubError
		expected bool
	}{
		{
			name: "Server error is retryable",
			err: &NotificationHubError{
				Code: ErrorCodeServerError,
			},
			expected: true,
		},
		{
			name: "Service unavailable is retryable",
			err: &NotificationHubError{
				Code: ErrorCodeServiceUnavailable,
			},
			expected: true,
		},
		{
			name: "Timeout is retryable",
			err: &NotificationHubError{
				Code: ErrorCodeTimeout,
			},
			expected: true,
		},
		{
			name: "Rate limited is retryable",
			err: &NotificationHubError{
				Code: ErrorCodeRateLimited,
			},
			expected: true,
		},
		{
			name: "Invalid request is not retryable",
			err: &NotificationHubError{
				Code: ErrorCodeInvalidRequest,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.IsRetryable(); got != tt.expected {
				t.Errorf("NotificationHubError.IsRetryable() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_NotificationHubError_IsAuthenticationError(t *testing.T) {
	tests := []struct {
		name     string
		err      *NotificationHubError
		expected bool
	}{
		{
			name: "Invalid connection string is authentication error",
			err: &NotificationHubError{
				Code: ErrorCodeInvalidConnectionString,
			},
			expected: true,
		},
		{
			name: "Authentication failed is authentication error",
			err: &NotificationHubError{
				Code: ErrorCodeAuthenticationFailed,
			},
			expected: true,
		},
		{
			name: "Unauthorized is authentication error",
			err: &NotificationHubError{
				Code: ErrorCodeUnauthorized,
			},
			expected: true,
		},
		{
			name: "Server error is not authentication error",
			err: &NotificationHubError{
				Code: ErrorCodeServerError,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.IsAuthenticationError(); got != tt.expected {
				t.Errorf("NotificationHubError.IsAuthenticationError() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_NewErrorFromHTTPResponse(t *testing.T) {
	tests := []struct {
		name            string
		statusCode      int
		body            []byte
		expectedCode    ErrorCode
		expectedMsg     string
		expectedDetails string
	}{
		{
			name:            "Bad request",
			statusCode:      http.StatusBadRequest,
			body:            []byte("Invalid request"),
			expectedCode:    ErrorCodeInvalidRequest,
			expectedMsg:     "Bad request",
			expectedDetails: "Invalid request",
		},
		{
			name:            "Unauthorized",
			statusCode:      http.StatusUnauthorized,
			body:            []byte("Invalid token"),
			expectedCode:    ErrorCodeUnauthorized,
			expectedMsg:     "Unauthorized",
			expectedDetails: "Invalid token",
		},
		{
			name:            "Forbidden",
			statusCode:      http.StatusForbidden,
			body:            []byte("Access denied"),
			expectedCode:    ErrorCodeAuthenticationFailed,
			expectedMsg:     "Authentication failed",
			expectedDetails: "Access denied",
		},
		{
			name:            "Not found",
			statusCode:      http.StatusNotFound,
			body:            []byte("Resource not found"),
			expectedCode:    ErrorCodeRegistrationNotFound,
			expectedMsg:     "Resource not found",
			expectedDetails: "Resource not found",
		},
		{
			name:            "Payload too large",
			statusCode:      http.StatusRequestEntityTooLarge,
			body:            []byte("Payload exceeds limit"),
			expectedCode:    ErrorCodePayloadTooLarge,
			expectedMsg:     "Payload too large",
			expectedDetails: "Payload exceeds limit",
		},
		{
			name:            "Too many requests",
			statusCode:      http.StatusTooManyRequests,
			body:            []byte("Rate limit exceeded"),
			expectedCode:    ErrorCodeRateLimited,
			expectedMsg:     "Rate limited",
			expectedDetails: "Rate limit exceeded",
		},
		{
			name:            "Server error",
			statusCode:      http.StatusInternalServerError,
			body:            []byte("Internal error"),
			expectedCode:    ErrorCodeServerError,
			expectedMsg:     "Internal server error",
			expectedDetails: "Internal error",
		},
		{
			name:            "Service unavailable",
			statusCode:      http.StatusServiceUnavailable,
			body:            []byte("Service down"),
			expectedCode:    ErrorCodeServiceUnavailable,
			expectedMsg:     "Service unavailable",
			expectedDetails: "Service down",
		},
		{
			name:            "Gateway timeout",
			statusCode:      http.StatusGatewayTimeout,
			body:            []byte("Request timed out"),
			expectedCode:    ErrorCodeTimeout,
			expectedMsg:     "Gateway timeout",
			expectedDetails: "Request timed out",
		},
		{
			name:            "Unknown status code",
			statusCode:      418,
			body:            []byte("I'm a teapot"),
			expectedCode:    ErrorCodeServerError,
			expectedMsg:     "HTTP 418",
			expectedDetails: "I'm a teapot",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Header:     make(http.Header),
			}
			resp.Header.Set("x-ms-request-id", "test-request-id")

			err := NewErrorFromHTTPResponse(resp, tt.body)
			if err.Code != tt.expectedCode {
				t.Errorf("NewErrorFromHTTPResponse() code = %v, want %v", err.Code, tt.expectedCode)
			}
			if err.Message != tt.expectedMsg {
				t.Errorf("NewErrorFromHTTPResponse() message = %v, want %v", err.Message, tt.expectedMsg)
			}
			if err.Details != tt.expectedDetails {
				t.Errorf("NewErrorFromHTTPResponse() details = %v, want %v", err.Details, tt.expectedDetails)
			}
			if err.StatusCode != tt.statusCode {
				t.Errorf("NewErrorFromHTTPResponse() status code = %v, want %v", err.StatusCode, tt.statusCode)
			}
			if err.RequestID != "test-request-id" {
				t.Errorf("NewErrorFromHTTPResponse() request ID = %v, want %v", err.RequestID, "test-request-id")
			}
		})
	}
}

func Test_ValidationError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *ValidationError
		expected string
	}{
		{
			name: "Validation error with value",
			err: &ValidationError{
				Field:   "name",
				Message: "is required",
				Value:   "",
			},
			expected: "validation error for field 'name': is required (value: )",
		},
		{
			name: "Validation error with non-empty value",
			err: &ValidationError{
				Field:   "age",
				Message: "must be positive",
				Value:   -1,
			},
			expected: "validation error for field 'age': must be positive (value: -1)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.expected {
				t.Errorf("ValidationError.Error() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_MultiError(t *testing.T) {
	t.Run("Single error", func(t *testing.T) {
		multiErr := NewMultiError()
		multiErr.Add(NewError(ErrorCodeInvalidRequest, "Invalid request"))

		if !multiErr.HasErrors() {
			t.Error("MultiError.HasErrors() = false, want true")
		}

		expected := "notification hub error [INVALID_REQUEST]: Invalid request"
		if got := multiErr.Error(); got != expected {
			t.Errorf("MultiError.Error() = %v, want %v", got, expected)
		}

		if err := multiErr.ToError(); err == nil {
			t.Error("MultiError.ToError() = nil, want error")
		}
	})

	t.Run("Multiple errors", func(t *testing.T) {
		multiErr := NewMultiError()
		multiErr.Add(NewError(ErrorCodeInvalidRequest, "Invalid request"))
		multiErr.Add(NewError(ErrorCodeInvalidPayload, "Invalid payload"))

		if !multiErr.HasErrors() {
			t.Error("MultiError.HasErrors() = false, want true")
		}

		expected := "multiple errors occurred (2 errors)"
		if got := multiErr.Error(); got != expected {
			t.Errorf("MultiError.Error() = %v, want %v", got, expected)
		}

		if err := multiErr.ToError(); err == nil {
			t.Error("MultiError.ToError() = nil, want error")
		}
	})

	t.Run("No errors", func(t *testing.T) {
		multiErr := NewMultiError()

		if multiErr.HasErrors() {
			t.Error("MultiError.HasErrors() = true, want false")
		}

		if err := multiErr.ToError(); err != nil {
			t.Errorf("MultiError.ToError() = %v, want nil", err)
		}
	})
}
