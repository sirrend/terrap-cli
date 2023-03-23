package rules_interaction

import (
	"github.com/Jeffail/gabs"
	"github.com/sirrend/terrap-cli/internal/utils"
)

// RuleSet is holding the ruleset as a gabs.Container object alongside relevant metadata
type RuleSet struct {
	ResourceName string
	Rules        *gabs.Container
}

func (r RuleSet) Execute(action string) {
	var rules []string

	if rulesSlice, err := r.Rules.Path("Rules").Children(); err == nil {
		for _, rule := range rulesSlice {
			rules = append(rules, rule.Path("Notify").String())
		}

		if action == "json" {
			tempJson, err := gabs.New().Set(rules, r.ResourceName)
			if err != nil {
				return
			}

			utils.PrettyPrintStruct(tempJson)
		}
	}
}
