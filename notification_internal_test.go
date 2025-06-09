package notificationhubs

import (
	"testing"
)

func Test_isIosBackgroundNotification(t *testing.T) {
	tests := []struct {
		name     string
		payload  []byte
		expected bool
	}{
		{
			name:     "Valid background notification",
			payload:  []byte(`{"aps":{"content-available":1}}`),
			expected: true,
		},
		{
			name:     "Non-background notification",
			payload:  []byte(`{"aps":{"alert":{"title":"Test","body":"Message"}}}`),
			expected: false,
		},
		{
			name:     "Invalid JSON",
			payload:  []byte(`invalid json`),
			expected: false,
		},
		{
			name:     "Empty payload",
			payload:  []byte{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIosBackgroundNotification(tt.payload); got != tt.expected {
				t.Errorf("isIosBackgroundNotification() = %v, want %v", got, tt.expected)
			}
		})
	}
}
