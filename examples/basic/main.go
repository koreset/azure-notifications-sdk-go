package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	notificationhubs "github.com/koreset/azure-notifications-sdk-go"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

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

	// Create a notification payload for Android (Firebase Cloud Messaging)
	//payload := map[string]interface{}{
	//	"data": map[string]string{
	//		"message": "This is a test notification",
	//	},
	//	"notification": map[string]string{
	//		"title": "Hello!",
	//		"body":  "This is a test notification",
	//	},
	//}

	payload := map[string]interface{}{
		"message": map[string]interface{}{
			"data": map[string]string{
				"message": "This is a direct notification",
				"userId":  "12345",
				"msisdn":  "1234567890",
			},
			"notification": map[string]string{
				"title": "Direct Notification",
				"body":  "This is a direct notification that was sent from the backend push",
			},
		},
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal payload: %v", err)
	}

	// Create a new notification
	notification, err := notificationhubs.NewNotification(notificationhubs.FcmV1Format, payloadBytes)
	if err != nil {
		log.Fatalf("Failed to create notification: %v", err)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Send the notification
	bytes, telemetry, err := hub.Send(ctx, notification, nil)

	// view the response bytes if needed
	fmt.Println("Response bytes:")
	fmt.Println(string(bytes))

	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	}

	fmt.Printf("Notification sent successfully! Message ID: %s\n", telemetry.NotificationMessageID)
}
