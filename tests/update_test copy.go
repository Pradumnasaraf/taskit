package test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestUpdateCmd tests the update command
func TestUpdateCmd(t *testing.T) {

	expectedOutput := "Task updated successfully"
	cmd := exec.Command("taskit", "update", "1", "update-task")

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

	assert.Equal(t, payload[0]["Task"], "update-task")

}
