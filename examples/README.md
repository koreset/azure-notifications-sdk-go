# Azure Notification Hubs SDK Examples

This directory contains examples demonstrating how to use the Azure Notification Hubs SDK for Go.

## Prerequisites

1. An Azure account with an active subscription
2. An Azure Notification Hub instance
3. The connection string and hub path for your notification hub

## Environment Variables

Before running the examples, set the following environment variables:

```bash
export AZURE_NOTIFICATION_HUB_CONNECTION_STRING="your-connection-string"
export AZURE_NOTIFICATION_HUB_PATH="ntf-dsa-mmn-dev-san-1"
```

## Examples

### 1. Basic Send (basic/main.go)

This example demonstrates how to send a basic notification to all registered devices.

```bash
cd basic
go run main.go
```

### 2. Device Registration (registration/main.go)

This example shows how to:
- Register a device with tags
- Register a device with a template

```bash
cd registration
go run main.go
```

### 3. Scheduled Notifications (scheduled/main.go)

This example demonstrates how to schedule a notification for future delivery.

```bash
cd scheduled
go run main.go
```

## Notification Formats

The SDK supports various notification formats:

- Apple Push Notification Service (APNS)
- Firebase Cloud Messaging (FCM)
- Windows Push Notification Service (WNS)
- Baidu Push Notification Service
- Amazon Device Messaging (ADM)

## Tags

You can use tags to target specific devices or groups of devices. Tags can be combined using logical operators:

- OR: `tag1 || tag2`
- AND: `tag1 && tag2`
- NOT: `!tag1`

Example: `(follows_RedSox || follows_Cardinals) && location_Boston`

## Templates

Templates allow you to define a notification format that can be reused with different values. The template uses placeholders that are replaced with actual values when sending the notification.

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

## Error Handling

All examples include basic error handling. In a production environment, you should:

1. Handle specific error types
2. Implement retry logic for transient failures
3. Log errors appropriately
4. Monitor notification delivery status

## Best Practices

1. Always use context with timeouts for API calls
2. Handle errors appropriately
3. Use tags for targeting specific devices
4. Use templates for consistent notification formats
5. Monitor notification delivery status
6. Implement proper error handling and retry logic
7. Keep your connection string secure 