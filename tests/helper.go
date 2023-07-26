package test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Read tasks.json and return the tasks
func ReadTasks() []map[string]interface{} {
	content, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Print("Error:", err)
	}

	var payload []map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}
