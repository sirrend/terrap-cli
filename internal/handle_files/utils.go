package handle_files

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"os"
	"strings"
)

// convertFileToHCLObject
/*
@brief:
	convertFileToHCLObject loads a file in a given location into an *hcl.File object
@params:
	fileName string - the file to load
@returns:
	*hcl.File - the loaded file object
	error - if exists, else nil
*/
func convertFileToHCLObject(fileName string) (*hcl.File, error) {
	parser := hclparse.NewParser()
	hclObject, diags := parser.ParseHCLFile(fileName)
	if diags.HasErrors() {
		return nil, diags.Errs()[0]
	}

	return hclObject, nil
}

// getBlocksFromFile
/*
@brief:
	getBlocksFromFile get all blocks (resources) under the tf given file
@params:
	tfFileName - string - the .tf file name to inspect
@returns:
	[]*hclwrite.Block - all the blocks found under the given file
	error - if exists, else nil
*/
func getBlocksFromFile(tfFileName string) ([]*hclwrite.Block, error) {
	fileBytes, err := os.ReadFile(tfFileName)
	if err != nil {
		return nil, err
	}

	file, diag := hclwrite.ParseConfig(fileBytes, tfFileName, hcl.InitialPos)
	if !diag.HasErrors() {
		return file.Body().Blocks(), nil
	}

	return nil, diag.Errs()[0]

}

// analyzeResources
/*
@brief:
	analyzeResources iterates over a given resources map and initializes it under the Resource struct
@params:
	resources - map[string]*tfconfig.Resource - the map to iterate on
@returns:
	[]Resource - the iteration's result
	error - if exists, else nil
*/
func analyzeResources(resources map[string]*tfconfig.Resource) ([]Resource, error) {
	var analyzedResources []Resource

	// extract resource
	for _, resourceData := range resources {
		r := Resource{}
		blocks, err := getBlocksFromFile(resourceData.Pos.Filename)
		if err != nil {
			return nil, err
		}

		for _, block := range blocks {
			if strings.Contains(resourceData.Type+"."+resourceData.Name, block.Labels()[0]+"."+block.Labels()[1]) {
				r.Init(block, resourceData)
				break // stop searching after allocating the right block
			}
		}

		analyzedResources = append(analyzedResources, r)
	}

	return analyzedResources, nil
}