package workspace

import "github.com/hashicorp/go-version"

// Workspace represents the data to be saved in the state file
type Workspace struct {
	Location         string
	ExecPath         string
	TerraformVersion string
	Providers        map[string]*version.Version
}

func (ws Workspace) GetProviderNames() []string {
	var keys []string
	for key, _ := range ws.Providers {
		keys = append(keys, key)
	}

	return keys
}
