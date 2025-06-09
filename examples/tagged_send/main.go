package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	notificationhubs "github.com/koreset/azure-notifications-sdk-go"
)

func main() {
	// Get connection string and hub path from environment variables
	connectionString := os.Getenv("AZURE_NOTIFICATION_HUB_CONNECTION_STRING")
	hubPath := os.Getenv("AZURE_NOTIFICATION_HUB_PATH")

	if connectionString == "" || hubPath == "" {
		log.Fatal("Please set AZURE_NOTIFICATION_HUB_CONNECTION_STRING and AZURE_NOTIFICATION_HUB_PATH environment variables")
	}

	// Create a new notification hub client
	hub := notificationhubs.NewNotificationHub(connectionString, hubPath)

	// Create a notification payload for Apple Push Notification Service (APNS)
	payload := map[string]interface{}{
		"aps": map[string]interface{}{
			"alert": map[string]string{
				"title": "Tagged Notification",
				"body":  "This notification was sent to devices with specific tags",
			},
			"sound": "default",
		},
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal payload: %v", err)
	}

	// Create a new notification
	notification, err := notificationhubs.NewNotification(notificationhubs.AppleFormat, payloadBytes)
	if err != nil {
		log.Fatalf("Failed to create notification: %v", err)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Example 1: Send to devices with a single tag
	singleTag := "premium_users"
	_, singleTagTelemetry, err := hub.Send(ctx, notification, &singleTag)
	if err != nil {
		log.Fatalf("Failed to send notification to single tag: %v", err)
	}

	fmt.Printf("Notification sent to single tag successfully!\n")
	fmt.Printf("Message ID: %s\n", singleTagTelemetry.NotificationMessageID)

	// Example 2: Send to devices with multiple tags (AND condition)
	multipleTags := "premium_users,ios_users"
	_, multipleTagsTelemetry, err := hub.Send(ctx, notification, &multipleTags)
	if err != nil {
		log.Fatalf("Failed to send notification to multiple tags: %v", err)
	}

	fmt.Printf("\nNotification sent to multiple tags successfully!\n")
	fmt.Printf("Message ID: %s\n", multipleTagsTelemetry.NotificationMessageID)

	// Example 3: Send to devices with tag expressions
	tagExpression := "(premium_users || vip_users) && ios_users"
	_, expressionTelemetry, err := hub.Send(ctx, notification, &tagExpression)
	if err != nil {
		log.Fatalf("Failed to send notification with tag expression: %v", err)
	}

	fmt.Printf("\nNotification sent with tag expression successfully!\n")
	fmt.Printf("Message ID: %s\n", expressionTelemetry.NotificationMessageID)
}
