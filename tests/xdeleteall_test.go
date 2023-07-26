package test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestRootCmd tests the root command (scrapy)
func TestDeleteAllCmd(t *testing.T) {

	expectedOutput := "All tasks deleted successfully"
	cmd := exec.Command("taskit", "deleteall")

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

	assert.Equal(t, len(payload), 0, "Expected 0 tasks")

}
