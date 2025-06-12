package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	notificationhubs "github.com/koreset/azure-notifications-sdk-go"
)

func main() {

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
	//		"message": "This is a direct notification",
	//	},
	//	"notification": map[string]string{
	//		"title": "Direct Notification",
	//		"body":  "This notification was sent directly to your device",
	//	},
	//}

	payload := map[string]interface{}{
		"message": map[string]interface{}{
			"notification": map[string]string{
				"body": "This is a direct notification",
			},
		},
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)

	fmt.Println(string(payloadBytes))
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

	// Example device handles (in production, these would be actual device registration IDs)
	deviceHandles := []string{
		"device-registration-id-1",
		"device-registration-id-2",
	}

	// Send notification to multiple devices
	_, telemetry, err := hub.SendDirectBatch(ctx, notification, deviceHandles...)
	if err != nil {
		log.Fatalf("Failed to send direct batch notification: %v", err)
	}

	fmt.Printf("Direct batch notification sent successfully!\n")
	fmt.Printf("Message ID: %s\n", telemetry.NotificationMessageID)

	// Example of sending to a single device
	singleDeviceHandle := "device-registration-id-3"
	_, singleTelemetry, err := hub.SendDirect(ctx, notification, singleDeviceHandle)
	if err != nil {
		log.Fatalf("Failed to send direct notification: %v", err)
	}

	fmt.Printf("\nSingle direct notification sent successfully!\n")
	fmt.Printf("Message ID: %s\n", singleTelemetry.NotificationMessageID)
}
