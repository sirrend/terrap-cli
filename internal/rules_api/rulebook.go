package rules_api

import (
	"github.com/Jeffail/gabs"
)

// Rulebook is holding the rulebook as bytes alongside relevant metadata
type Rulebook struct {
	SourceVersion string
	Bytes         []byte
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
func (r Rulebook) GetRuleSetByResource(resourceName string) (rule *gabs.Container, err error) {
	jsonRuleBook, err := gabs.ParseJSON(r.Bytes)
	if err != nil {
		return &gabs.Container{}, err
	}

	exists, err := r.doesRuleExistForResource(resourceName)
	if exists {
		return jsonRuleBook.Path(resourceName), nil
	}

	return &gabs.Container{}, err
}

// doesRuleExistForResource
/*
@brief:
	doesRuleExistForResource checks if a rule exist for a given resource
@params:
	resourceName - string - the resource name to check existence for
@returns:
	bool - true if exists, false otherwise
	error - if exists, else nil
*/
func (r Rulebook) doesRuleExistForResource(resourceName string) (bool, error) {
	jsonRuleSet, err := gabs.ParseJSON(r.Bytes)
	if err != nil {
		return false, err
	}

	if jsonRuleSet.Path(resourceName) != nil {
		return true, nil
	}

	return false, nil
}
