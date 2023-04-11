package scanning

import (
	"github.com/sirrend/terrap-cli/internal/handle_files"
	"github.com/sirrend/terrap-cli/internal/utils"
)

func WhereDoesResourceAppear(resources []handle_files.Resource) map[string][]string {
	appearances := make(map[string][]string)

	for _, resource := range resources {
		if !utils.IsItemInSlice(utils.GetAbsPath(resource.Pos.Filename), appearances[resource.Name]) {
			appearances[resource.Name] = append(appearances[resource.Name], utils.GetAbsPath(resource.Pos.Filename))
		}
	}

	return appearances
}

func GetUniqResources(resources []handle_files.Resource) []handle_files.Resource {
	var tempResourcesSlice []handle_files.Resource
	tempResourcesMap := map[string]handle_files.Resource{}

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
