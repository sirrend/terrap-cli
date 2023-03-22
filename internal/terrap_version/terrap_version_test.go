package terrap_version

import (
	"testing"
)

// Test if the folder creation is working
func TestFolderCreation(t *testing.T) {
	terrap := TerrapVersion{}
	terrap.SetVersion()

	if terrap.Version != "" && terrap.GoVersion != "" && terrap.System != "" && terrap.Product != "" {
		t.Logf("Version was set successfully")
	} else {
		t.Errorf("Something went wrong..")
	}

}
