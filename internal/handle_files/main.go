package handle_files

import (
	"errors"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"os"
	"path/filepath"
)

// ScanFile
/*
@brief:
	ScanFile parses a tf file in order to get the resources and data source defined in it
@params:
	filename string - the file to parse
@returns:
	resources - []Resource - all the resource found in the given file
	error - if exists, else nil
*/
func ScanFile(fileName string) (resources []Resource, err error) {
	module := tfconfig.NewModule("")

	// load hcl file and parse it
	hclObject, err := convertFileToHCLObject(fileName)
	if err != nil {
		return nil, err
	}

	if diag := tfconfig.LoadModuleFromFile(hclObject, module); diag.HasErrors() {
		return nil, diag.Errs()[0]
	}

	// combine both maps into one
	for k, v := range module.DataResources {
		module.ManagedResources[k] = v
	}

	resources, err = analyzeResources(module.ManagedResources)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

// ScanFolder
/*
@brief:
	ScanFolder parses a folder which contains tf files in order to get the resources and data source defined in it
@params:
	dir string - the folder to parse
@returns:
	resources - []*Resource - map of all resources found in the given folder and its attributes'
	error - if exists, else nil
*/
func ScanFolder(dir string) (resources []Resource, err error) {
	var moduleResources []Resource

	if tfconfig.IsModuleDir(dir) {
		inspect, diag := tfconfig.LoadModule(dir)
		if diag.HasErrors() {
			return resources, diag.Err()
		}

		// go over all modules in the folder in context
		for _, module := range inspect.ModuleCalls {
			if tempModuleResources, err := GetLocalModuleResources(dir, *module); err == nil {
				moduleResources = append(moduleResources, tempModuleResources...)
			} else {
				return resources, err
			}
		}

		// combine both maps into one
		for k, v := range inspect.DataResources {
			inspect.ManagedResources[k] = v
		}

		resources, err = analyzeResources(inspect.ManagedResources)
		if err != nil {
			return nil, err
		}

		resources = append(resources, moduleResources...) // combine bote module resources and current directory resources together
		return resources, nil
	}

	err = errors.New("the given dir does not contain terraform code")
	return resources, err
}

// ScanFolderRecursively
/*
@brief:
	ScanFolderRecursively parses all the .tf files under a given folder tree recursively in order to get all the Resources
	and Data Sources in it
@params:
	dir string - the folder to parse
@returns:
	resources - []*Resource - map of all resources found in the given folder and its attributes'
	error - if exists, else nil
*/
func ScanFolderRecursively(dir string) (resources []Resource, err error) {
	err = filepath.WalkDir(dir, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && tfconfig.IsModuleDir(path) {
			if folderResources, err := ScanFolder(path); err == nil {
				resources = append(resources, folderResources...)
			} else {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return resources, nil
}
