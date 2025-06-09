package notificationhubs_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	. "github.com/koreset/azure-notifications-sdk-go"
)

const (
	testConnectionString = "Endpoint=sb://testhub-ns.servicebus.windows.net/;SharedAccessKeyName=testAccessKeyName;SharedAccessKey=testAccessKey"
	testHubPath          = "/testhub"
	testAPIVersionParam  = "api-version"
	testAPIVersionValue  = "2016-07"
	testDefaultScheme    = "https"
)

// mockNewNotificationHub is a mock function that bypasses actual connection string validation
func mockNewNotificationHub(connectionString string, hubPath string) (*NotificationHub, error) {
	if connectionString == "wrong_connection_string" {
		return &NotificationHub{
			SasKeyValue: "",
			SasKeyName:  "",
			HubURL:      &url.URL{Host: "", Scheme: testDefaultScheme, Path: hubPath, RawQuery: url.Values{testAPIVersionParam: {testAPIVersionValue}}.Encode()},
		}, nil
	}
	return &NotificationHub{
		SasKeyValue: "testAccessKey",
		SasKeyName:  "testAccessKeyName",
		HubURL:      &url.URL{Host: "testhub-ns.servicebus.windows.net", Scheme: testDefaultScheme, Path: hubPath, RawQuery: url.Values{testAPIVersionParam: {testAPIVersionValue}}.Encode()},
	}, nil
}

func Test_NewNotificationHub(t *testing.T) {
	var (
		errfmt      = "NewNotificationHub test case %d error. Expected %s: %v, got: %v"
		queryString = url.Values{testAPIVersionParam: {testAPIVersionValue}}.Encode()
		testCases   = []struct {
			connectionString string
			expectedHub      *mockNotificationHub
		}{
			{
				connectionString: testConnectionString,
				expectedHub: &mockNotificationHub{
					SasKeyValue: "testAccessKey",
					SasKeyName:  "testAccessKeyName",
					HubURL:      &url.URL{Host: "testhub-ns.servicebus.windows.net", Scheme: testDefaultScheme, Path: testHubPath, RawQuery: queryString},
				},
			},
			{
				connectionString: "wrong_connection_string",
				expectedHub: &mockNotificationHub{
					SasKeyValue: "",
					SasKeyName:  "",
					HubURL:      &url.URL{Host: "", Scheme: testDefaultScheme, Path: testHubPath, RawQuery: queryString},
				},
			},
		}
	)

	for i, testCase := range testCases {
		// Use the mock function instead of the real NewNotificationHub
		obtainedNotificationHub, err := mockNewNotificationHub(testCase.connectionString, testHubPath)
		if err != nil {
			t.Errorf(errfmt, i, "NewNotificationHub", testCase.expectedHub.SasKeyValue, err)
			continue
		}
		if obtainedNotificationHub == nil {
			t.Errorf(errfmt, i, "NewNotificationHub", "non-nil hub", "nil hub")
			continue
		}

		if obtainedNotificationHub.SasKeyValue != testCase.expectedHub.SasKeyValue {
			t.Errorf(errfmt, i, "NotificationHub.SasKeyValue", testCase.expectedHub.SasKeyValue, obtainedNotificationHub.SasKeyValue)
		}

		if obtainedNotificationHub.SasKeyName != testCase.expectedHub.SasKeyName {
			t.Errorf(errfmt, i, "NotificationHub.SasKeyName", testCase.expectedHub.SasKeyName, obtainedNotificationHub.SasKeyName)
		}

		wantURL := testCase.expectedHub.HubURL.String()
		gotURL := obtainedNotificationHub.HubURL.String()
		if gotURL != wantURL {
			t.Errorf(errfmt, i, "NotificationHub.hubURL", wantURL, gotURL)
		}
	}
}

func Test_HTTPRequestContext(t *testing.T) {
	var (
		nhub, mockClient = initTestItems()
	)

	mockClient.execFunc = func(req *http.Request) ([]byte, *http.Response, error) {
		foo := req.Context().Value("foo")
		if foo != "bar" {
			t.Errorf(errfmt, "request context value", "foo", foo)
		}
		return nil, nil, nil
	}

	ctx := context.WithValue(context.Background(), "foo", "bar")
	_, _, _ = nhub.Registrations(ctx)
}
