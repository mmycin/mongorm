// Package utils contains utilities for working with BSON and JSON data.
package utils

import (
    "encoding/json"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
)

// Json is a custom type that extends bson.M to add utility methods.
// bson.M is a map representation used in MongoDB to store BSON data in Go, where key-value pairs are stored as map[string]interface{}.
type Json bson.M

// PrintAsJSON formats and prints the Json object as an indented JSON string.
// This method uses json.MarshalIndent to pretty-print the JSON.
// It returns an error if the JSON marshalling fails.
func (j Json) PrintAsJSON() error {
    // Marshal the Json (bson.M) into an indented JSON format for readability.
    bytes, err := json.MarshalIndent(bson.M(j), "", "  ") // 2 spaces for indentation
    if err != nil {
        return fmt.Errorf("Invalid data. Cannot marshall: %w", err)
    }

    // Print the formatted JSON string.
    fmt.Println(string(bytes))
    return nil
}

// PrintAsValue iterates over the Json object and prints each key-value pair.
// This method is useful for quickly inspecting the contents of the Json map in a human-readable format.
func (j Json) PrintAsValue() {
    // Iterate through each key-value pair in the Json map (bson.M).
    for key, value := range j {
        fmt.Printf("%s: %v\n", key, value) // Print each key and its corresponding value.
    }
}

