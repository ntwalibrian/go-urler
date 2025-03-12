package api

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
)

// reads the urls file 
func FileRead() map[string]string {
	filePath := "urls.json"
	file, err := os.ReadFile(filePath)

	if err != nil {
		log.Println("Error reading file:", err)
		return make(map[string]string)
	}

	urlMappings := make(map[string]string)
	err = json.Unmarshal(file, &urlMappings)

	if err != nil {
		log.Println("Error decoding JSON:", err)
		return make(map[string]string)
	}

	return urlMappings
}

func FileWrite(shortKey string , url string) {
	filePath := "urls.json"

	// Read existing data
	urlMappings := make(map[string]string)
	file, err := os.ReadFile(filePath)
	if err == nil {
		json.Unmarshal(file, &urlMappings)
	}

	// Add new short URL mapping
	urlMappings[shortKey] = url

	// Write updated data back to file
	jsonData, err := json.MarshalIndent(urlMappings, "", "  ")
	if err != nil {
		log.Fatal("Error encoding JSON:", err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	fmt.Println("URL saved successfully: ", urlMappings[shortKey] ," ==> " , shortKey)
}