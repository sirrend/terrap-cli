package rules_interaction

import (
	"fmt"
	"github.com/sirrend/terrap-cli/internal/commons"
)

// Rule holds all needed attribute from a parsed rule
type Rule struct {
	Path          string `json:"Path"`
	ComponentName string `json:"ComponentName"`
	ComponentType string `json:"ComponentType"`
	Required      bool   `json:"Required"`
	Notification  string `json:"Notification"`
}

// PrettyPrint
/*
@brief:
	PrettyPrint prints the Rule object
*/
func (r Rule) PrettyPrint() {
	_, _ = commons.YELLOW.Print("    Change Path: ")
	fmt.Println(r.Path)

	_, _ = commons.YELLOW.Print("    Is This Component Required: ")
	fmt.Println(r.Required)

	_, _ = commons.YELLOW.Print("    Change: ")
	fmt.Println(r.Notification, "\n")
}
