package utils

import "fmt"

// HandleError panics if an error is encountered.
func HandleError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Error: %v", err))
	}
}
