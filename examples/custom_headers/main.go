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

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Example 1: Send to Apple Push Notification Service (APNS) with rich features
	apnsPayload := map[string]interface{}{
		"aps": map[string]interface{}{
			"alert": map[string]interface{}{
				"title":        "iOS Rich Notification",
				"body":         "This notification includes rich features",
				"subtitle":     "Optional subtitle",
				"launch-image": "launch.png",
			},
			"sound":             "custom.wav",
			"badge":             1,
			"content-available": 1,
			"mutable-content":   1,
			"category":          "MESSAGE_CATEGORY",
			"thread-id":         "thread-123",
		},
		"custom_data": map[string]interface{}{
			"key1": "value1",
			"key2": 123,
			"key3": true,
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

	fmt.Printf("APNS rich notification sent successfully!\n")
	fmt.Printf("Message ID: %s\n", apnsTelemetry.NotificationMessageID)

	// Example 2: Send to Firebase Cloud Messaging (FCM) with rich features
	fcmPayload := map[string]interface{}{
		"data": map[string]string{
			"message":      "This notification includes rich features",
			"priority":     "high",
			"collapse_key": "message-123",
			"time_to_live": "3600",
			"custom_key1":  "value1",
			"custom_key2":  "value2",
		},
		"notification": map[string]interface{}{
			"title":          "Android Rich Notification",
			"body":           "This notification includes rich features",
			"icon":           "notification_icon",
			"color":          "#FF0000",
			"sound":          "default",
			"tag":            "message-123",
			"click_action":   "OPEN_ACTIVITY",
			"body_loc_key":   "notification_body",
			"body_loc_args":  []string{"arg1", "arg2"},
			"title_loc_key":  "notification_title",
			"title_loc_args": []string{"arg1", "arg2"},
		},
		"android": map[string]interface{}{
			"priority": "high",
			"notification": map[string]interface{}{
				"channel_id":   "high_importance_channel",
				"click_action": "OPEN_ACTIVITY",
				"color":        "#FF0000",
				"icon":         "notification_icon",
				"sound":        "default",
				"tag":          "message-123",
			},
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

	fmt.Printf("\nFCM rich notification sent successfully!\n")
	fmt.Printf("Message ID: %s\n", fcmTelemetry.NotificationMessageID)

	// Example 3: Send to Windows Notification Service (WNS) with rich features
	wnsPayload := `<?xml version="1.0" encoding="utf-8"?>
		<toast>
			<visual>
				<binding template="ToastGeneric">
					<text id="1">Windows Rich Notification</text>
					<text id="2">This notification includes rich features</text>
					<image placement="appLogoOverride" src="logo.png"/>
					<image placement="hero" src="hero.png"/>
				</binding>
			</visual>
			<audio src="ms-winsoundevent:Notification.Default"/>
			<actions>
				<action content="View" arguments="view"/>
				<action content="Dismiss" arguments="dismiss"/>
			</actions>
		</toast>`

	wnsNotification, err := notificationhubs.NewNotification(notificationhubs.WindowsFormat, []byte(wnsPayload))
	if err != nil {
		log.Fatalf("Failed to create WNS notification: %v", err)
	}

	_, wnsTelemetry, err := hub.Send(ctx, wnsNotification, nil)
	if err != nil {
		log.Fatalf("Failed to send WNS notification: %v", err)
	}

	fmt.Printf("\nWNS rich notification sent successfully!\n")
	fmt.Printf("Message ID: %s\n", wnsTelemetry.NotificationMessageID)
}
