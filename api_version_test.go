package notificationhubs_test

import (
	"testing"

	. "github.com/koreset/azure-notifications-sdk-go"
)

func TestGetAPIVersionForOperation(t *testing.T) {
	testCases := []struct {
		name          string
		operationType string
		expected      string
	}{
		{
			name:          "Send operations use latest API 2016-07",
			operationType: "send",
			expected:      "2016-07",
		},
		{
			name:          "Registration operations use latest API 2016-07",
			operationType: "registration",
			expected:      "2016-07",
		},
		{
			name:          "Telemetry operations use latest API 2016-07",
			operationType: "telemetry",
			expected:      "2016-07",
		},
		{
			name:          "Message telemetry operations use latest API 2016-07",
			operationType: "message-telemetry",
			expected:      "2016-07",
		},
		{
			name:          "PNS error details operations use latest API 2016-07",
			operationType: "pns-error-details",
			expected:      "2016-07",
		},
		{
			name:          "All operations use latest API 2016-07",
			operationType: "unknown-operation",
			expected:      "2016-07",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetAPIVersionForOperation(tc.operationType)
			if result != tc.expected {
				t.Errorf("GetAPIVersionForOperation(%q) = %q; want %q", tc.operationType, result, tc.expected)
			}
		})
	}
}

func TestAPIVersionConstants(t *testing.T) {
	if DefaultAPIVersion != "2016-07" {
		t.Errorf("DefaultAPIVersion = %q; want %q", DefaultAPIVersion, "2016-07")
	}

	if LatestAPIVersion != "2016-07" {
		t.Errorf("LatestAPIVersion = %q; want %q", LatestAPIVersion, "2016-07")
	}

	if LegacyAPIVersion != "2015-01" {
		t.Errorf("LegacyAPIVersion = %q; want %q", LegacyAPIVersion, "2015-01")
	}
}
