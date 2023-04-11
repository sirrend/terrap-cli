package community_modules

import (
	"github.com/Jeffail/gabs"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/sirrend/terrap-cli/internal/utils"
	"os"
	"path/filepath"
)

type Module struct {
	Version              string
	DotTerraformLocation string
}

func (m *Module) Init(dir string, module tfconfig.ModuleCall) error {
	m.Version = module.Version
	abs := utils.GetAbsPath(dir)

	if utils.DoesExist(filepath.Join(abs, ".terraform", "modules", "modules.json")) {
		file, _ := os.ReadFile(filepath.Join(dir, ".terraform", "modules", "modules.json")) // read modules.json generated on init

		if modules, err := gabs.ParseJSON(file); err == nil {
			if paths, err := modules.Path("Modules").Children(); err == nil {
				for _, path := range paths {
					if module.Version == "" {
						if path.Path("Key").String() == module.Name {
							m.DotTerraformLocation = filepath.Join(abs, path.Path("Dir").String())
							break
						}

					} else {
						if path.Path("Key").String() == module.Name && path.Path("Version").String() == module.Version {
							m.DotTerraformLocation = filepath.Join(abs, path.Path("Dir").String())
							break
						}
					}
				}

			} else {
				return err
			}

		} else {
			return err
		}
	}

	return nil
}

//func (m *Module) InitCommunityModule() (err error) {
//	var notYetSupportedMessage string
//	tempWorkspace := workspace.Workspace{}
//
//	tempWorkspace.ExecPath, tempWorkspace.TerraformVersion, err = terraform_utils.TerraformInit(m.DotTerraformLocation)
//	if err != nil {
//		return err
//	}
//
//	providers.FindTfProviders(m.DotTerraformLocation, &tempWorkspace)
//
//	resources, err := handle_files.ScanFolderRecursively(".")
//	if err != nil {
//		_, _ = commons.RED.Println(err)
//	}
//
//	for provider, version := range tempWorkspace.Providers {
//		rulebook, _ := rules_api.GetRules(provider, version.String())
//
//		// validate rulebook downloaded
//		if err != nil {
//			if strings.Contains(err.Error(), utils.StripProviderPrefix(provider)) {
//				notYetSupportedMessage = strings.Join([]string{notYetSupportedMessage, err.Error()}, ", ")
//				continue
//			}
//
//			continue
//		}
//
//		for _, resource := range scanning.GetUniqResources(resources) {
//			ruleset, err := resource.GetRuleset(rulebook, nil)
//			if err != nil {
//				_, _ = commons.RED.Println(err)
//				os.Exit(1)
//			}
//		}
//
//	}
//
//	return nil
//}
