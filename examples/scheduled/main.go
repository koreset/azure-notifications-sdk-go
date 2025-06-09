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
				"title": "Scheduled Notification",
				"body":  "This notification was scheduled to be delivered later",
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

	// Schedule the notification for 5 minutes in the future
	deliveryTime := time.Now().Add(5 * time.Minute)

	// Send the scheduled notification
	_, telemetry, err := hub.Schedule(ctx, notification, nil, deliveryTime)
	if err != nil {
		log.Fatalf("Failed to schedule notification: %v", err)
	}

	fmt.Printf("Notification scheduled successfully!\n")
	fmt.Printf("Message ID: %s\n", telemetry.NotificationMessageID)
	fmt.Printf("Scheduled for: %s\n", deliveryTime.Format(time.RFC3339))
}
