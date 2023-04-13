package files_handler

import (
	"errors"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/sirrend/terrap-cli/internal/utils"
	"os"
	"path/filepath"
)

// ScanFileForResources
/*
@brief:
	ScanFileForResources parses a tf file in order to get the resources and data source defined in it
@params:
	filename string - the file to parse
@returns:
	resources - []Resource - all the resource found in the given file
	error - if exists, else nil
*/
func ScanFileForResources(fileName string) (resources []Resource, err error) {
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

// ScanFileForModules
/*
@brief:
	ScanFileForModules parses a tf file in order to get the modules defined in it
@params:
	filename string - the file to parse
@returns:
	resources - map[string][]Resource - all the resources found in the given module
	error - if exists, else nil
*/
func ScanFileForModules(fileName string) (filesToResourcesMap map[string][]Resource, err error) {
	var resources []Resource
	filesToResourcesMap = make(map[string][]Resource)

	inspect := tfconfig.NewModule("")

	// load hcl file and parse it
	hclObject, err := convertFileToHCLObject(fileName)
	if err != nil {
		return nil, err
	}

	// load modules from file
	if diag := tfconfig.LoadModuleFromFile(hclObject, inspect); diag.HasErrors() {
		return nil, diag.Errs()[0]
	}

	// go over all modules in the folder in context
	for _, module := range inspect.ModuleCalls {
		if tempModuleResources, err := getLocalModuleResources(utils.GetDirName(fileName), *module); err == nil {
			resources = append(resources, tempModuleResources...)
		} else {
			return filesToResourcesMap, err
		}
	}

	// append resources to the right file
	for _, resource := range resources {
		filesToResourcesMap[resource.Pos.Filename] = append(filesToResourcesMap[resource.Pos.Filename], resource)
	}
	return filesToResourcesMap, nil
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
			if tempModuleResources, err := getLocalModuleResources(dir, *module); err == nil {
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

// FindResourcesPerFile
/*
@brief:
	FindResourcesPerFile parses all the .tf files under a given folder tree recursively in order to get all the Resources
	and Data Sources in it
@params:
	dir string - the folder to parse
@returns:
	resources - map[string][]Resource - map of all resources found in the given folder and its attributes'
	error - if exists, else nil
*/
func FindResourcesPerFile(dir string) (files map[string][]Resource, err error) {
	files = make(map[string][]Resource)
	err = filepath.WalkDir(dir, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isTerraformFile(path) && !utils.IsHiddenPath(path) {
			if fileResources, err := ScanFileForResources(path); err == nil {
				files[utils.GetAbsPath(path)] = fileResources
			} else {
				return err
			}

			if filesMapToResources, err := ScanFileForModules(path); err == nil {
				for file, resources := range filesMapToResources {
					files[file] = resources
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
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

		if info.IsDir() && tfconfig.IsModuleDir(path) && !utils.IsHiddenObject(path) {
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
