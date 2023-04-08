package rules_api

import (
	"fmt"
	"github.com/sirrend/terrap-cli/internal/commons"
	"strings"
)

// Rule holds all needed attribute from a parsed rule
type Rule struct {
	Path          string `json:"Path"`
	Operation     string `json:"Operation"`
	ComponentName string `json:"ComponentName"`
	ComponentType string `json:"ComponentType"`
	Required      bool   `json:"Required"`
	Notification  string `json:"Notification"`
	URL           string `json:"URL"`
}

// IsNew
/*
@brief:
	IsNew checks if the Rule in context represents a new resource / component
*/
func (r Rule) IsNew() bool {
	if !r.IsParameterChange() && r.Operation == "added" {
		return true
	}

	return false
}

// PrettyPrint
/*
@brief:
	PrettyPrint prints the Rule object
*/
func (r Rule) PrettyPrint() {
	_, _ = commons.YELLOW.Print("    Change Path: ")
	fmt.Println(r.Path)

	_, _ = commons.YELLOW.Print("    Operation: ")
	fmt.Println(r.Operation)

	_, _ = commons.YELLOW.Print("    Is This Component Required: ")
	fmt.Println(r.Required)

	_, _ = commons.YELLOW.Print("    Change: ")
	fmt.Println(r.Notification)

	_, _ = commons.YELLOW.Print("    Documentation: ")
	fmt.Println(r.URL, "\n")
}

// PrettyPrintWhatsNew
/*
@brief:
	PrettyPrintWhatsNew prints the Rule object if new
*/
func (r Rule) PrettyPrintWhatsNew() {
	if r.IsNew() {
		_, _ = commons.YELLOW.Print("    Addition: ")
		fmt.Println(r.Notification)

		_, _ = commons.YELLOW.Print("    Documentation: ")
		fmt.Println(r.URL, "\n")
	}
}

// IsParameterChange
/*
@brief:
	IsParameterChange validates that a rule does not concern to a new parameter (computed, deprecated etc.)
@returns:
	bool - true if parameter. else false
*/
func (r Rule) IsParameterChange() bool {
	parameters := []string{"Type",
		"Required",
		"Optional",
		"Computed",
		"ForceNew",
		"Default",
		"Description",
		"InputDefault",
		"MaxItems",
		"MinItems",
		"ComputedWhen",
		"ExactlyOneOf",
		"AtLeastOneOf",
		"RequiredWith",
		"Deprecated",
		"Sensitive",
	}

	for _, parameter := range parameters {
		if strings.ToLower(parameter) == strings.ToLower(r.ComponentName) {
			return true
		}
	}

	return false
}
