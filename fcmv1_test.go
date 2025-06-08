package notificationhubs_test

import (
	"testing"

	. "github.com/koreset/azure-notifications-sdk-go"
)

func TestFcmV1Constants(t *testing.T) {
	// Test FCM v1 notification format constant
	if FcmV1Format != "fcmv1" {
		t.Errorf("Expected FcmV1Format to be 'fcmv1', got '%s'", FcmV1Format)
	}

	// Test FCM v1 target platform constants
	if FcmV1Platform != "fcmv1" {
		t.Errorf("Expected FcmV1Platform to be 'fcmv1', got '%s'", FcmV1Platform)
	}

	if FcmV1TemplatePlatform != "fcmv1template" {
		t.Errorf("Expected FcmV1TemplatePlatform to be 'fcmv1template', got '%s'", FcmV1TemplatePlatform)
	}

	// Test FCM v1 installation platform constant
	if FCMV1Platform != "fcmv1" {
		t.Errorf("Expected FCMV1Platform to be 'fcmv1', got '%s'", FCMV1Platform)
	}
}

func TestFcmV1Format_GetContentType(t *testing.T) {
	contentType := FcmV1Format.GetContentType()
	expected := "application/json"

	if contentType != expected {
		t.Errorf("FcmV1Format.GetContentType() = %s, expected %s", contentType, expected)
	}
}

func TestFcmV1Format_IsValid(t *testing.T) {
	if !FcmV1Format.IsValid() {
		t.Error("FcmV1Format should be valid")
	}
}

func TestFcmV1Platforms_IsValid(t *testing.T) {
	if !FcmV1Platform.IsValid() {
		t.Error("FcmV1Platform should be valid")
	}

	if !FcmV1TemplatePlatform.IsValid() {
		t.Error("FcmV1TemplatePlatform should be valid")
	}
}

func TestFcmV1Registration(t *testing.T) {
	// Test FCM v1 registration creation
	registration := Registration{
		DeviceID:           "fcmv1-device-123",
		NotificationFormat: FcmV1Format,
		RegistrationID:     "fcmv1-reg-123",
		Tags:               "tag1,tag2",
	}

	if registration.NotificationFormat != FcmV1Format {
		t.Errorf("Expected notification format to be FcmV1Format, got %s", registration.NotificationFormat)
	}
}

func TestFcmV1TemplateRegistration(t *testing.T) {
	// Test FCM v1 template registration creation
	templateReg := TemplateRegistration{
		DeviceID:       "fcmv1-device-123",
		RegistrationID: "fcmv1-template-reg-123",
		Platform:       FcmV1TemplatePlatform,
		Template:       `{"message":{"notification":{"title":"$(title)","body":"$(message)"}}}`,
		Tags:           "tag1,tag2",
	}

	if templateReg.Platform != FcmV1TemplatePlatform {
		t.Errorf("Expected platform to be FcmV1TemplatePlatform, got %s", templateReg.Platform)
	}
}

func TestFcmV1Installation(t *testing.T) {
	// Test FCM v1 installation creation
	installation := Installation{
		InstallationID: "fcmv1-installation-123",
		Platform:       FCMV1Platform,
		PushChannel:    "fcmv1_token_sample_here",
		Tags:           []string{"tag1", "tag2"},
		Templates: map[string]InstallationTemplate{
			"template1": {
				Body: `{"message":{"notification":{"title":"$(title)","body":"$(message)"}}}`,
				Tags: []string{"templateTag1"},
			},
		},
	}

	if installation.Platform != FCMV1Platform {
		t.Errorf("Expected platform to be FCMV1Platform, got %s", installation.Platform)
	}

	if installation.PushChannel != "fcmv1_token_sample_here" {
		t.Errorf("Expected push channel to be set correctly")
	}

	template, exists := installation.Templates["template1"]
	if !exists {
		t.Error("Expected template1 to exist")
	}

	expectedBody := `{"message":{"notification":{"title":"$(title)","body":"$(message)"}}}`
	if template.Body != expectedBody {
		t.Errorf("Expected template body to be FCM v1 format, got %s", template.Body)
	}
}

func TestFcmV1NotificationFormat(t *testing.T) {
	// Test that FCM v1 format is properly categorized as JSON
	jsonFormats := []NotificationFormat{
		Template,
		AppleFormat,
		FcmV1Format,
		KindleFormat,
		BaiduFormat,
	}

	for _, format := range jsonFormats {
		if format.GetContentType() != "application/json" {
			t.Errorf("Format %s should return application/json content type", format)
		}
	}
}
