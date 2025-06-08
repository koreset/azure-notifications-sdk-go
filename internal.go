package notificationhubs

// Internal constants
const (
	apiVersionParam = "api-version"

	// Latest API version for Azure Notification Hubs service operations
	// 2016-07 is the latest available and provides enhanced features like PNS error details
	// while maintaining full backward compatibility with 2015-01
	apiVersionValue = "2016-07"

	// Legacy API version (maintained for reference)
	legacyAPIVersionValue = "2015-01"

	// Current telemetry API version (same as latest service API)
	telemetryAPIVersionValue = "2016-07"

	directParam = "direct"
)

// API version helpers
const (
	// Latest API version for all notification hub operations
	// Using 2016-07 provides enhanced features including PNS error details
	LatestAPIVersion = apiVersionValue

	// Legacy API version (for backward compatibility reference)
	LegacyAPIVersion = legacyAPIVersionValue

	// Default API version (same as latest)
	DefaultAPIVersion = LatestAPIVersion
)

// GetAPIVersionForOperation returns the latest API version for all operations
// All operations now use 2016-07 for enhanced features and better error reporting
func GetAPIVersionForOperation(operationType string) string {
	// All operations use the latest API version for consistency and enhanced features
	return LatestAPIVersion
}

// Internal constants continued
const (
	// for connection string parsing
	schemeServiceBus  = "sb"
	schemeDefault     = "https"
	paramEndpoint     = "Endpoint="
	paramSaasKeyName  = "SharedAccessKeyName="
	paramSaasKeyValue = "SharedAccessKey="

	// Http methods
	deleteMethod = "DELETE"
	getMethod    = "GET"
	postMethod   = "POST"
	putMethod    = "PUT"
	patchMethod  = "PATCH"

	// appleRegXMLString is the XML string for registering an iOS device
	// Replace {{Tags}} and {{DeviceID}} with the correct values
	appleRegXMLString string = `<?xml version="1.0" encoding="utf-8"?>
<entry xmlns="http://www.w3.org/2005/Atom">
  <content type="application/xml">
    <AppleRegistrationDescription xmlns:i="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://schemas.microsoft.com/netservices/2010/10/servicebus/connect">
      <Tags>{{Tags}}</Tags>
      <DeviceToken>{{DeviceID}}</DeviceToken>
    </AppleRegistrationDescription>
  </content>
</entry>`

	// appleTemplateRegXMLString is the XML string for registering an iOS device
	// Replace {{Tags}}, {{DeviceID}} and {{Template}} with the correct values
	appleTemplateRegXMLString string = `<?xml version="1.0" encoding="utf-8"?>
<entry xmlns="http://www.w3.org/2005/Atom">
  <content type="application/xml">
    <AppleTemplateRegistrationDescription xmlns:i="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://schemas.microsoft.com/netservices/2010/10/servicebus/connect">
      <Tags>{{Tags}}</Tags>
      <DeviceToken>{{DeviceID}}</DeviceToken>
      <BodyTemplate><![CDATA[{{Template}}]]></BodyTemplate>
    </AppleTemplateRegistrationDescription>
  </content>
</entry>`

	// fcmV1RegXMLString is the XML string for registering an FCM v1 device
	// Replace {{Tags}} and {{DeviceID}} with the correct values
	fcmV1RegXMLString string = `<?xml version="1.0" encoding="utf-8"?>
<entry xmlns="http://www.w3.org/2005/Atom">
  <content type="application/xml">
    <FcmV1RegistrationDescription xmlns:i="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://schemas.microsoft.com/netservices/2010/10/servicebus/connect">
      <Tags>{{Tags}}</Tags>
      <FcmV1RegistrationId>{{DeviceID}}</FcmV1RegistrationId>
    </FcmV1RegistrationDescription>
  </content>
</entry>`

	// fcmV1TemplateRegXMLString is the XML string for registering an FCM v1 device with template
	// Replace {{Tags}}, {{DeviceID}} and {{Template}} with the correct values
	fcmV1TemplateRegXMLString string = `<?xml version="1.0" encoding="utf-8"?>
<entry xmlns="http://www.w3.org/2005/Atom">
  <content type="application/xml">
    <FcmV1TemplateRegistrationDescription xmlns:i="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://schemas.microsoft.com/netservices/2010/10/servicebus/connect">
      <Tags>{{Tags}}</Tags>
      <FcmV1RegistrationId>{{DeviceID}}</FcmV1RegistrationId>
      <BodyTemplate><![CDATA[{{Template}}]]></BodyTemplate>
    </FcmV1TemplateRegistrationDescription>
  </content>
</entry>`
)
