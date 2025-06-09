package main

import (
	"context"
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

	// Example FCM token (this would normally come from your Android app)
	// Format: A valid FCM token is a long string that looks like:
	// "dPdnr8HkqU8:APA91bHPRi7cRrfahB8H-2QJh..."
	fcmToken := "dPdnr8HkqU8:APA91bHPRi7cRrfahB8H-2QJhK9mN0pL3vX7yZ4wA1bC5dE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP9qR3tU6vW9xY2zB5cE8fG2hJ4kM7nP"

	// Create a registration with the FCM token
	registration := notificationhubs.NewRegistration(
		fcmToken,                     // device ID (FCM token)
		nil,                          // no expiration time
		notificationhubs.FcmV1Format, // notification format for FCM
		"",                           // registration ID (will be generated by the hub)
		"android,premium",            // tags for targeting this device
	)

	// Register the device
	_, result, err := hub.Register(ctx, *registration)
	if err != nil {
		log.Fatalf("Failed to register device: %v", err)
	}

	fmt.Printf("Device registered successfully!\n")
	fmt.Printf("Registration ID: %s\n", result.RegistrationContent.RegisteredDevice.RegistrationID)
	fmt.Printf("Device ID: %s\n", result.RegistrationContent.RegisteredDevice.DeviceID)
	fmt.Printf("Tags: %v\n", result.RegistrationContent.RegisteredDevice.Tags)

	// Example: Create a template registration for the same device
	template := `{
		"data": {
			"message": "$(message)",
			"title": "$(title)",
			"priority": "$(priority)",
			"custom_key": "$(custom_key)"
		}
	}`

	templateRegistration := notificationhubs.NewTemplateRegistration(
		fcmToken, // device ID (FCM token)
		nil,      // no expiration time
		result.RegistrationContent.RegisteredDevice.RegistrationID, // use the registration ID from previous registration
		"android,premium",              // tags
		notificationhubs.FcmV1Platform, // platform
		template,                       // template
	)

	// Register the template
	_, templateResult, err := hub.RegisterWithTemplate(ctx, *templateRegistration)
	if err != nil {
		log.Fatalf("Failed to register template: %v", err)
	}

	fmt.Printf("\nTemplate registered successfully!\n")
	fmt.Printf("Registration ID: %s\n", templateResult.RegistrationContent.RegisteredDevice.RegistrationID)
	fmt.Printf("Template: %s\n", templateResult.RegistrationContent.RegisteredDevice.Template)

	// Example: List all registrations for this device
	_, registrations, err := hub.Registrations(ctx)
	if err != nil {
		log.Fatalf("Failed to list registrations: %v", err)
	}

	fmt.Printf("\nRegistrations for device:\n")
	for _, reg := range registrations.Entries {
		if reg.RegistrationContent.RegisteredDevice.DeviceID == fcmToken {
			fmt.Printf("- Registration ID: %s\n", reg.RegistrationContent.RegisteredDevice.RegistrationID)
			fmt.Printf("  Platform: %s\n", reg.RegistrationContent.RegisteredDevice.DeviceToken)
			fmt.Printf("  Tags: %v\n", reg.RegistrationContent.RegisteredDevice.Tags)
			if reg.RegistrationContent.RegisteredDevice.Template != "" {
				fmt.Printf("  Template: %s\n", reg.RegistrationContent.RegisteredDevice.Template)
			}
			fmt.Println()
		}
	}
}
