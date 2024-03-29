package parser

import (
	"fmt"

	"github.com/sirrend/terrap-cli/internal/commons"
)

// RuleSet is holding the ruleset as a gabs.Container object alongside relevant metadata
type RuleSet struct {
	ResourceName string   `json:"ResourceName"`
	Appearances  []string `json:"Appearances"`
	Rules        []Rule   `json:"Changes"`
}

// GetNewComponents
/*
@brief:
	GetNewComponents finds all the new components in the next version
@returns:
	[]string - a slice of notifications of what's new
*/
func (r RuleSet) GetNewComponents() []string {
	var newComponents []string

	for _, rule := range r.Rules {
		if rule.IsNew() {
			newComponents = append(newComponents, rule.Notification)
		}
	}

	return newComponents
}

// PrettyPrint
/*
@brief:
	PrettyPrint prints the RuleSet object
*/
func (r RuleSet) PrettyPrint(rules []Rule) {
	if r.Rules != nil {
		_, _ = commons.GREEN.Print("- Resource Name: ")
		fmt.Println(r.ResourceName)
		fmt.Println("    Changes:")

		for _, rule := range rules {
			rule.PrettyPrint()
		}

		fmt.Println()
	}
}

// PrettyPrintWhatsNew
/*
@brief:
	PrettyPrintWhatsNew prints the RuleSet new objects
*/
func (r RuleSet) PrettyPrintWhatsNew() {
	if len(r.GetNewComponents()) > 0 {
		_, _ = commons.GREEN.Print("Resource Name: ")
		fmt.Println(r.ResourceName)

		for _, rule := range r.Rules {
			rule.PrettyPrintWhatsNew()
		}
	}
}
