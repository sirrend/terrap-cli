package receiver

import "github.com/sirrend/terrap-cli/internal/files_handler"

type RulesAPIRequest struct {
	Provider       string              `json:"provider"`
	CurrentVersion string              `json:"current_version"`
	DesiredVersion string              `json:"desired_version"`
	UsedResources  map[string][]string `json:"used_resources"`
}

func CreateRulesRequest() *RulesAPIRequest {
	return &RulesAPIRequest{}
}

func (r *RulesAPIRequest) fillUsedResources(files map[string][]files_handler.Resource) {
	// 1. extract all used components by iterating over the files argument
	// 2. insert resources to the UsedResource map
}

func (r *RulesAPIRequest) fetchRules() {
	// 1. perform request using the internal/requests library
}
