package test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestRootCmd tests the root command (scrapy)
func TestAddCmd(t *testing.T) {

	expectedOutput := "Task added successfully"
	cmd := exec.Command("taskit", "add", "test")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output))
	assert.Equal(t, got, expectedOutput)

	// Validate the json output
	content, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Print("Error:", err)
	}

	var payload []map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	assert.Equal(t, payload[0]["Task"], "test")

}
