package providers

import (
	"encoding/json"
	"testing"

	"github.com/Jeffail/gabs"
	"github.com/sirrend/terrap-cli/internal/terraform_utils"
)

// Test if saving schema is working
func TestProvidersGetting(t *testing.T) {
	execPath, _, err := terraform_utils.TerraformInit("../../terraform-test")

	if err != nil {
		t.Errorf("Test failed on initialization with the following error: %e", err)
	}
	tf := terraform_utils.NewTerraformExecuter("../../terraform-test", execPath)
	schemas, err := SaveSchemas(tf)
	if err != nil {
		t.Errorf(err.Error())
	}

	b, err := json.Marshal(schemas)
	parseJSON, err := gabs.ParseJSON(b)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("Parsed Json: %v", parseJSON.String())
}
