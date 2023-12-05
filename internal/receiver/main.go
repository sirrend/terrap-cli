package receiver

import (
	"net/http"

	"github.com/sirrend/terrap-cli/internal/files_handler"
	"github.com/sirrend/terrap-cli/internal/parser"
)

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
	// 2. insert resources to the UsedResources map
}

func (r *RulesAPIRequest) fetchRules() {
	// 1. perform request using the internal/requests library
}

func loadRulebook(response *http.Response) *parser.Rulebook {
	// 1. Convert the response object body to bytes representation.
	// 2. Marshal it to a parser.Rulebook struct.

	return &parser.Rulebook{}
}
