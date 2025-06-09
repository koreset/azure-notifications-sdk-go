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
	hub, err := notificationhubs.NewNotificationHub(connectionString, hubPath)
	if err != nil {
		log.Fatalf("Failed to create notification hub: %v", err)
	}

	// Create a notification payload for Apple Push Notification Service (APNS)
	payload := map[string]interface{}{
		"aps": map[string]interface{}{
			"alert": map[string]string{
				"title": "Hello!",
				"body":  "This is a test notification",
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

	// Send the notification
	_, telemetry, err := hub.Send(ctx, notification, nil)
	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	}

	fmt.Printf("Notification sent successfully! Message ID: %s\n", telemetry.NotificationMessageID)
}
