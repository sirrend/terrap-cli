package rules_api

import (
	"bufio"
	"fmt"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/utils"
	"os"
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
	URL           string `json:"Docs"`
}

// DoesRuleApplyInContext
/*
@brief:
	DoesRuleApplyInContext checks if the rule applies to the user's context
@params:
	filepath - string - the path to the file
	resourceName - string - the resource the rule applies to
	resourceType - string - the type of the resource the rule applies to
@returns:
	bool - true if applies. else false
	error - if exists, otherwise false
*/
func (r Rule) DoesRuleApplyInContext(filePath, resourceName, resourceType string) (bool, error) {
	var lines []string
	splintedPath := strings.Split(r.Path, ".")

	// Open the file for reading
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Read the file into a slice of strings
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// check if rule applies to method
	for pos, line := range lines {
		if strings.Contains(line, fmt.Sprintf("%s \"%s\"", resourceType, resourceName)) {
			codeBlock := utils.GetCodeUntilMatchingBrace(strings.Join(lines[pos:], "\n"))

		OuterLoop:
			for _, component := range splintedPath {
				splintedCodeBlock := strings.Split(codeBlock, "\n")
				for functionPos, functionLine := range splintedCodeBlock {
					if strings.Contains(functionLine, component) && !strings.Contains(functionLine, "#") { // validate not a comment
						if component == splintedPath[len(splintedPath)-1] {
							return true, nil

						} else {
							codeBlock = strings.Join(strings.Split(codeBlock, "\n")[functionPos:], "\n") // continue from next codeBlock line after break
							break
						}
					}

					if splintedCodeBlock[len(splintedCodeBlock)-1] == functionLine {
						break OuterLoop
					}
				}
			}
		}
	}

	return false, nil
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

	if r.URL != "" {
		_, _ = commons.YELLOW.Print("    Documentation: ")
		fmt.Println(r.URL, "\n")
	}
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

		if r.URL != "" {
			_, _ = commons.YELLOW.Print("    Documentation: ")
			fmt.Println(r.URL, "\n")
		}
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
