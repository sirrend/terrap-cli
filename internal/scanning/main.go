package scanning

import (
	"github.com/sirrend/terrap-cli/internal/files_handler"
)

func GetUniqueResources(resources []files_handler.Resource) []files_handler.Resource {
	var tempResourcesSlice []files_handler.Resource
	tempResourcesMap := map[string]files_handler.Resource{}

	for _, resource := range resources {
		if _, inSlice := tempResourcesMap[resource.Name]; !inSlice {
			tempResourcesMap[resource.Name] = resource
		}
	}

	for _, resource := range tempResourcesMap {
		tempResourcesSlice = append(tempResourcesSlice, resource)
	}

	return tempResourcesSlice
}
