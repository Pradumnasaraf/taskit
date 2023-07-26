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
func TestDeleteCmd(t *testing.T) {

	expectedOutput := "Task deleted successfully"
	cmd := exec.Command("taskit", "delete", "2")

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

	if len(payload) != 2 || payload[1]["Task"] != "task3" {
		return
	}

}
