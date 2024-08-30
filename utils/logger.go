package utils

import "log"

// LogError logs an error message.
func LogError(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
