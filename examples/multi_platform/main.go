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

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Example 1: Send to Apple Push Notification Service (APNS)
	apnsPayload := map[string]interface{}{
		"aps": map[string]interface{}{
			"alert": map[string]string{
				"title": "iOS Notification",
				"body":  "This is a notification for iOS devices",
			},
			"sound": "default",
			"badge": 1,
		},
	}

	apnsPayloadBytes, err := json.Marshal(apnsPayload)
	if err != nil {
		log.Fatalf("Failed to marshal APNS payload: %v", err)
	}

	apnsNotification, err := notificationhubs.NewNotification(notificationhubs.AppleFormat, apnsPayloadBytes)
	if err != nil {
		log.Fatalf("Failed to create APNS notification: %v", err)
	}

	_, apnsTelemetry, err := hub.Send(ctx, apnsNotification, nil)
	if err != nil {
		log.Fatalf("Failed to send APNS notification: %v", err)
	}

	fmt.Printf("APNS notification sent successfully!\n")
	fmt.Printf("Message ID: %s\n", apnsTelemetry.NotificationMessageID)

	// Example 2: Send to Firebase Cloud Messaging (FCM)
	fcmPayload := map[string]interface{}{
		"data": map[string]string{
			"message": "This is a notification for Android devices",
		},
		"notification": map[string]string{
			"title": "Android Notification",
			"body":  "This is a notification for Android devices",
		},
	}

	fcmPayloadBytes, err := json.Marshal(fcmPayload)
	if err != nil {
		log.Fatalf("Failed to marshal FCM payload: %v", err)
	}

	fcmNotification, err := notificationhubs.NewNotification(notificationhubs.FcmV1Format, fcmPayloadBytes)
	if err != nil {
		log.Fatalf("Failed to create FCM notification: %v", err)
	}

	_, fcmTelemetry, err := hub.Send(ctx, fcmNotification, nil)
	if err != nil {
		log.Fatalf("Failed to send FCM notification: %v", err)
	}

	fmt.Printf("\nFCM notification sent successfully!\n")
	fmt.Printf("Message ID: %s\n", fcmTelemetry.NotificationMessageID)

	// Example 3: Send to Windows Notification Service (WNS)
	wnsPayload := `<?xml version="1.0" encoding="utf-8"?>
		<toast>
			<visual>
				<binding template="ToastText01">
					<text id="1">Windows Notification</text>
				</binding>
			</visual>
		</toast>`

	wnsNotification, err := notificationhubs.NewNotification(notificationhubs.WindowsFormat, []byte(wnsPayload))
	if err != nil {
		log.Fatalf("Failed to create WNS notification: %v", err)
	}

	_, wnsTelemetry, err := hub.Send(ctx, wnsNotification, nil)
	if err != nil {
		log.Fatalf("Failed to send WNS notification: %v", err)
	}

	fmt.Printf("\nWNS notification sent successfully!\n")
	fmt.Printf("Message ID: %s\n", wnsTelemetry.NotificationMessageID)
}
