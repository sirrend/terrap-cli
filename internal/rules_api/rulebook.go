package rules_api

import (
	"encoding/json"
	"github.com/Jeffail/gabs"
	"github.com/sirrend/terrap-cli/internal/utils"
	"os"
)

// Rulebook is holding the rulebook as bytes alongside relevant metadata
type Rulebook struct {
	SourceVersion string
	TargetVersion string
	Bytes         []byte
}

func (r Rulebook) GetTargetVersion() string {
	jsonRuleBook, err := gabs.ParseJSON(r.Bytes)
	if err != nil {
		return ""
	}

	return utils.MustUnquote(jsonRuleBook.Path("RuleBookSettings.TargetVersion").String())
}

// GetRuleSetByResource
/*
@brief:
	GetRuleSetByResource retrieves a rule from the rulebook in context by its name
@params:
	resourceName - string - the resource name to fetch the rule for
@returns:
	rule - *gabs.Container - the rule as a gabs json object
	err - error - if exists, else nil
*/
func (r Rulebook) GetRuleSetByResource(resourceName, resourceType string) (rule *gabs.Container, err error) {
	jsonRuleBook, err := gabs.ParseJSON(r.Bytes)
	if err != nil {
		return &gabs.Container{}, err
	}

	if jsonRuleBook.ExistsP(resourceType + "." + resourceName) {
		return jsonRuleBook.Path(resourceType + "." + resourceName), nil
	}

	return &gabs.Container{}, err
}

// GetAllRuleSets
/*
@brief:
	GetAllRuleSets returns all RuleSets in a rulebook
@returns:
	map[string]interface{} - the rules as a map of interfaces
	error - if exists, else nil
*/
func (r Rulebook) GetAllRuleSets() (map[string]interface{}, error) {
	jsonRuleBook, err := gabs.ParseJSON(r.Bytes)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	if err = json.Unmarshal(jsonRuleBook.Bytes(), &data); err != nil {
		os.Exit(1)
	}

	return data, nil
}
