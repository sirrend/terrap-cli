package state

import (
	"github.com/sirrend/terrap-cli/internal/workspace"
	"os"
	"testing"

	"github.com/hashicorp/go-version"
)

var testWorkspace = workspace.Workspace{}

// Test if workspace saving is working
func TestSaveFile(t *testing.T) {
	testWorkspace.Location = "."
	testWorkspace.TerraformVersion = "test"
	testWorkspace.ExecPath = "test"
	testWorkspace.Providers = map[string]*version.Version{}

	err := Save(".test", testWorkspace)
	if err != nil {
		t.Error(err)
	}

	t.Log("Successfully Saved File!")
	testWorkspace = workspace.Workspace{} // clean the workspace for the next load
}

// Test if loading workspace is working
func TestLoadFile(t *testing.T) {
	err := Load(".test", &testWorkspace)
	if err != nil {
		t.Error(err)
	}

	if testWorkspace.TerraformVersion == "test" {
		t.Log("Terraform workspace was loaded successfully!")
	}

	os.Remove(".test")
}
