package rules_interaction

import (
	"fmt"
	"github.com/sirrend/terrap-cli/internal/commons"
)

// RuleSet is holding the ruleset as a gabs.Container object alongside relevant metadata
type RuleSet struct {
	ResourceName string `json:"ResourceName"`
	Rules        []Rule `json:"Changes"`
}

// PrettyPrint
/*
@brief:
	PrettyPrint prints the RuleSet object
*/
func (r RuleSet) PrettyPrint() {
	if r.Rules != nil {
		_, _ = commons.GREEN.Print("Resource Name: ")
		fmt.Println(r.ResourceName)
		fmt.Println("  Changes:")

		for _, rule := range r.Rules {
			rule.PrettyPrint()
		}
	}
}
