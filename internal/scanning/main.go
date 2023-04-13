package scanning

import (
	"github.com/sirrend/terrap-cli/internal/files_handler"
	"github.com/sirrend/terrap-cli/internal/utils"
)

func WhereDoesResourceAppear(resources []files_handler.Resource) map[string][]string {
	appearances := make(map[string][]string)

	for _, resource := range resources {
		if !utils.IsItemInSlice(utils.GetAbsPath(resource.Pos.Filename), appearances[resource.Name]) {
			appearances[resource.Name] = append(appearances[resource.Name], utils.GetAbsPath(resource.Pos.Filename))
		}
	}

	return appearances
}

func GetUniqResources(resources []files_handler.Resource) []files_handler.Resource {
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
