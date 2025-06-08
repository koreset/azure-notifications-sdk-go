package notificationhubs_test

import (
	"testing"

	. "github.com/koreset/azure-notifications-sdk-go"
)

func TestNotificationFormat_GetContentType(t *testing.T) {
	var (
		testCases = []struct {
			format   NotificationFormat
			expected string
		}{
			{
				format:   Template,
				expected: "application/json",
			},
			{
				format:   FcmV1Format,
				expected: "application/json",
			},
			{
				format:   AppleFormat,
				expected: "application/json",
			},
			{
				format:   BaiduFormat,
				expected: "application/json",
			},
			{
				format:   KindleFormat,
				expected: "application/json",
			},
			{
				format:   WindowsFormat,
				expected: "application/xml",
			},
			{
				format:   WindowsPhoneFormat,
				expected: "application/xml",
			},
		}
	)

	for _, testCase := range testCases {
		obtained := testCase.format.GetContentType()
		if obtained != testCase.expected {
			t.Errorf("NotificationFormat '%s' GetContentType(). Expected '%s', got '%s'", testCase.format, testCase.expected, obtained)
		}
	}
}

func TestNotificationFormat_IsValid(t *testing.T) {
	var (
		testCases = []struct {
			format  NotificationFormat
			isValid bool
		}{
			{
				format:  Template,
				isValid: true,
			},
			{
				format:  FcmV1Format,
				isValid: true,
			},
			{
				format:  AppleFormat,
				isValid: true,
			},
			{
				format:  BaiduFormat,
				isValid: true,
			},
			{
				format:  KindleFormat,
				isValid: true,
			},
			{
				format:  WindowsFormat,
				isValid: true,
			},
			{
				format:  WindowsPhoneFormat,
				isValid: true,
			},
			{
				format:  NotificationFormat("wrong_format"),
				isValid: false,
			},
		}
	)

	for _, testCase := range testCases {
		obtained := testCase.format.IsValid()
		if obtained != testCase.isValid {
			t.Errorf("NotificationFormat '%s' isValid(). Expected '%t', got '%t'", testCase.format, testCase.isValid, obtained)
		}
	}
}

func TestTargetPlatform_IsValid(t *testing.T) {
	var (
		testCases = []struct {
			platform TargetPlatform
			isValid  bool
		}{
			{
				platform: AdmPlatform,
				isValid:  true,
			},
			{
				platform: AdmTemplatePlatform,
				isValid:  true,
			},
			{
				platform: ApplePlatform,
				isValid:  true,
			},
			{
				platform: AppleTemplatePlatform,
				isValid:  true,
			},
			{
				platform: BaiduPlatform,
				isValid:  true,
			},
			{
				platform: BaiduTemplatePlatform,
				isValid:  true,
			},
			{
				platform: FcmV1Platform,
				isValid:  true,
			},
			{
				platform: FcmV1TemplatePlatform,
				isValid:  true,
			},
			{
				platform: TemplatePlatform,
				isValid:  true,
			},
			{
				platform: WindowsphonePlatform,
				isValid:  true,
			},
			{
				platform: WindowsphoneTemplatePlatform,
				isValid:  true,
			},
			{
				platform: WindowsPlatform,
				isValid:  true,
			},
			{
				platform: WindowsTemplatePlatform,
				isValid:  true,
			},
			{
				platform: TargetPlatform("invalid_platform"),
				isValid:  false,
			},
		}
	)

	for _, testCase := range testCases {
		obtained := testCase.platform.IsValid()
		if obtained != testCase.isValid {
			t.Errorf("TargetPlatform '%s' isValid(). Expected '%t', got '%t'", testCase.platform, testCase.isValid, obtained)
		}
	}
}
