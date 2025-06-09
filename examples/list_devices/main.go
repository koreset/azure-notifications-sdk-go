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
	hub, err := notificationhubs.NewNotificationHub(connectionString, hubPath)
	if err != nil {
		log.Fatalf("Failed to create notification hub: %v", err)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get all registrations
	_, registrations, err := hub.Registrations(ctx)
	if err != nil {
		log.Fatalf("Failed to list registrations: %v", err)
	}

	// Group registrations by platform
	platformGroups := make(map[string][]notificationhubs.RegistrationResult)
	for _, reg := range registrations.Entries {
		if reg.RegistrationContent == nil || reg.RegistrationContent.RegisteredDevice == nil {
			continue
		}

		platform := "Unknown"

		// Determine platform based on registration type
		if reg.RegistrationContent.AppleRegistrationDescription != nil {
			platform = "iOS"
		} else if reg.RegistrationContent.FcmV1RegistrationDescription != nil {
			platform = "Android"
		} else if reg.RegistrationContent.AppleTemplateRegistrationDescription != nil {
			platform = "iOS (Template)"
		} else if reg.RegistrationContent.FcmV1TemplateRegistrationDescription != nil {
			platform = "Android (Template)"
		}

		platformGroups[platform] = append(platformGroups[platform], reg)
	}

	// Print summary
	fmt.Printf("Total registered devices: %d\n\n", len(registrations.Entries))

	// Print details by platform
	for platform, devices := range platformGroups {
		fmt.Printf("=== %s Devices (%d) ===\n", platform, len(devices))
		for _, reg := range devices {
			device := reg.RegistrationContent.RegisteredDevice
			fmt.Printf("\nDevice ID: %s\n", device.DeviceID)
			fmt.Printf("Registration ID: %s\n", device.RegistrationID)

			// Print tags if available
			if device.TagsString != nil && *device.TagsString != "" {
				fmt.Printf("Tags: %s\n", *device.TagsString)
			}

			// Print template if available
			if device.Template != "" {
				fmt.Printf("Template: %s\n", device.Template)
			}

			// Print expiration time if available
			if device.ExpirationTimeString != nil && *device.ExpirationTimeString != "" {
				fmt.Printf("Expiration Time: %s\n", *device.ExpirationTimeString)
			}

			fmt.Println("---")
		}
		fmt.Println()
	}

	// Example: Get details for a specific registration
	if len(registrations.Entries) > 0 {
		firstReg := registrations.Entries[0]
		device := firstReg.RegistrationContent.RegisteredDevice

		fmt.Printf("\nGetting details for registration ID: %s\n", device.RegistrationID)
		_, regDetails, err := hub.Registration(ctx, device.RegistrationID)
		if err != nil {
			log.Printf("Failed to get registration details: %v", err)
		} else {
			fmt.Printf("Registration details retrieved successfully\n")
			if regDetails.RegistrationContent != nil && regDetails.RegistrationContent.RegisteredDevice != nil {
				details := regDetails.RegistrationContent.RegisteredDevice
				fmt.Printf("Device ID: %s\n", details.DeviceID)
				fmt.Printf("ETag: %s\n", details.ETag)
				if details.TagsString != nil {
					fmt.Printf("Tags: %s\n", *details.TagsString)
				}
			}
		}
	}
}
