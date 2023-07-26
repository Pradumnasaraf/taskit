package test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestDeleteCmd tests the delete command
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
	payload := ReadTasks()

	if len(payload) != 2 || payload[1]["Task"] != "task3" {
		return
	}

}
