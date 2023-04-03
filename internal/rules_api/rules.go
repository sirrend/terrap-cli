package rules_api

import (
	"github.com/Jeffail/gabs"
)

type Rules struct {
	SourceVersion string
	Container     *gabs.Container
}

func (r Rules) GetResourceRules(resource string) interface{} {
	resourceRules := r.Container.Path(resource)

	return resourceRules.Data()
}

func (r Rules) GetResourceRulesJson(resource string) map[string]interface{} {
	data := r.GetResourceRules(resource)
	rulesData := map[string]interface{}{}

	// append resource name to the map start if needed
	if data != nil {
		if dataMap, ok := data.(map[string]interface{}); ok {
			if _, existsInMap := dataMap[resource]; existsInMap {
				rulesData = dataMap

			} else {
				dataWithResourceName := map[string]interface{}{resource: data}
				rulesData = dataWithResourceName
			}
		}
	}

	return rulesData
}
