package test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestAddCmd tests the add command
func TestAddCmd(t *testing.T) {

	expectedOutput := "Task added successfully"
	cmd := exec.Command("taskit", "add", "task3")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output))
	assert.Equal(t, got, expectedOutput)

	// Validate the json output
	payload := ReadTasks()

	// Validate the json output 3rd task
	assert.Equal(t, payload[2]["Task"], "task3")

}
