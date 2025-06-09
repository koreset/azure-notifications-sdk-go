# Azure Notifications SDK for Go

A Go library for interacting with Microsoft Azure Notification Hubs, providing a simple and efficient way to send push notifications to various platforms.

## Features

- Support for multiple notification platforms:
  - Apple Push Notification Service (APNS)
  - Firebase Cloud Messaging (FCM)
  - Windows Push Notification Service (WNS)
  - Baidu Push Notification Service
  - Amazon Device Messaging (ADM)
- Device registration and management
- Template-based notifications
- Tag-based targeting
- Scheduled notifications
- Custom headers support
- Multi-platform notifications

## Installation

```bash
go get github.com/koreset/azure-notifications-sdk-go
```

## Quick Start

```go
package main

import (
    "log"
    "github.com/koreset/azure-notifications-sdk-go"
)

func main() {
    // Initialize the notification hub
    hub, err := notificationhubs.NewNotificationHub(
        "your-connection-string",
        "your-hub-path",
    )
    if err != nil {
        log.Fatal(err)
    }

    // Create a notification
    notification, err := notificationhubs.NewNotification(
        notificationhubs.APNS,
        []byte(`{"aps":{"alert":{"title":"Hello","body":"World"}}}`),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Send the notification
    result, err := hub.Send(notification)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Notification sent successfully: %+v", result)
}
```

## Key Concepts

### Notification Hub

The `NotificationHub` is the main entry point for interacting with Azure Notification Hubs. It handles the connection and provides methods for sending notifications and managing registrations.

### Registrations

Device registrations are used to identify and target specific devices. You can register devices with:
- Tags for targeting
- Templates for consistent notification formats
- Platform-specific information

Example:
```go
registration := notificationhubs.NewRegistration(
    "device-id",
    nil, // expiration time
    notificationhubs.APNS,
    "registration-id",
    "tag1,tag2",
)
```

### Templates

Templates allow you to define reusable notification formats with placeholders. This is useful for maintaining consistent notification structures across your application.

Example template for APNS:
```json
{
    "aps": {
        "alert": {
            "title": "$(title)",
            "body": "$(message)"
        }
    }
}
```

### Tags

Tags are used to target specific devices or groups of devices. You can combine tags using logical operators:
- OR: `tag1 || tag2`
- AND: `tag1 && tag2`
- NOT: `!tag1`

Example: `(follows_RedSox || follows_Cardinals) && location_Boston`

## Examples

The repository includes several examples demonstrating different features:

- Basic notification sending
- Device registration
- Scheduled notifications
- Tag-based targeting
- Template notifications
- Custom headers
- Multi-platform notifications

See the [examples directory](examples/) for detailed examples.

## Error Handling

The SDK provides detailed error information through the `errors` package. Always handle errors appropriately in your application:

```go
if err != nil {
    switch {
    case errors.IsConnectionError(err):
        // Handle connection errors
    case errors.IsAuthenticationError(err):
        // Handle authentication errors
    default:
        // Handle other errors
    }
}
```

## Best Practices

1. Always use context with timeouts for API calls
2. Implement proper error handling and retry logic
3. Use tags for targeting specific devices
4. Use templates for consistent notification formats
5. Keep your connection string secure
6. Monitor notification delivery status
7. Handle platform-specific requirements appropriately

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms of the license included in the repository.

## Security

Please see [SECURITY.md](SECURITY.md) for security-related information. 