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

	// Example device token (must be a valid hexadecimal string for APNS)
	// This is a sample token - in production, you would get this from the actual device
	deviceToken := "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

	// Create a registration with tags
	registration := notificationhubs.NewRegistration(
		deviceToken,                  // Device ID
		nil,                          // No expiration time
		notificationhubs.AppleFormat, // Notification format
		"",                           // Registration ID (empty for new registration)
		"user:123,location:US",       // Tags for targeting
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

	// Example of registering a device with a template
	templateRegistration := notificationhubs.NewTemplateRegistration(
		deviceToken,                      // Device ID
		nil,                              // No expiration time
		"",                               // Registration ID (empty for new registration)
		"user:123,location:US",           // Tags
		notificationhubs.ApplePlatform,   // Platform
		`{"aps":{"alert":"$(message)"}}`, // Template
	)

	// Register the device with template
	_, templateResult, err := hub.RegisterWithTemplate(ctx, *templateRegistration)
	if err != nil {
		log.Fatalf("Failed to register device with template: %v", err)
	}

	fmt.Printf("\nDevice registered with template successfully!\n")
	fmt.Printf("Registration ID: %s\n", templateResult.RegistrationContent.RegisteredDevice.RegistrationID)
	fmt.Printf("Template: %s\n", templateResult.RegistrationContent.RegisteredDevice.Template)
}
