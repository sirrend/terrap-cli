package workspace

import "github.com/hashicorp/go-version"

// Workspace represents the data to be saved in the state file
type Workspace struct {
	Location         string                      `json:"Location"`
	ExecPath         string                      `json:"ExecPath,omitempty"`
	TerraformVersion string                      `json:"TerraformVersion,omitempty"`
	IsTempProvider   bool                        `json:"IsTempProvider,omitempty"`
	Providers        map[string]*version.Version `json:"Providers"`
}

func (ws Workspace) GetProviderNames() []string {
	var keys []string
	for key := range ws.Providers {
		keys = append(keys, key)
	}

	return keys
}
