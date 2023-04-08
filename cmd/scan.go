/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/fatih/color"
	"github.com/sirrend/terrap-cli/internal/annotate"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/handle_files"
	"github.com/sirrend/terrap-cli/internal/rules_api"
	"github.com/sirrend/terrap-cli/internal/state"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/sirrend/terrap-cli/internal/workspace"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	PRINTED        = false
	upgradeMessage = ""
)

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

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan your IaC code to find provider changes",

	Run: func(cmd *cobra.Command, args []string) {
		var workspace workspace.Workspace
		asJson := map[string][]rules_api.Rule{}

		if utils.IsInitialized(".") {
			resources, err := handle_files.ScanFolderRecursively(".")
			if err != nil {
				_, _ = commons.RED.Println(err)
			}

			err = state.Load("./.terrap.json", &workspace)
			if err != nil {
				_, _ = commons.RED.Println(err)
			}

			// go over every provider in user's folder
			for provider, version := range workspace.Providers {
				rulebook, _ := rules_api.GetRules(provider, version.String())

				if !cmd.Flag("annotate").Changed {
					for _, resource := range GetUniqResources(resources) {
						ruleset, err := resource.GetRuleset(rulebook)
						if err != nil {
							_, _ = commons.RED.Println(err)
							os.Exit(1)
						}

						// fill in json object
						if cmd.Flag("json").Changed {
							if ruleset.Rules != nil {
								asJson[ruleset.ResourceName] = ruleset.Rules
							}

							// print human-readable output
						} else {
							ruleset.PrettyPrint()
							if len(ruleset.Rules) != 0 {
								PRINTED = true
							}
						}
					}

					// print json object
					if cmd.Flag("json").Changed {
						utils.PrettyPrintStruct(asJson)
						if len(asJson) != 0 {
							PRINTED = true
						}
					}

					if !PRINTED {
						upgradeMessage += fmt.Sprintf("  %s  %s: ",
							emoji.UpArrow, strings.ReplaceAll(provider, "registry.terraform.io/", ""))
						upgradeMessage += commons.GREEN.Sprintf("%s -> %s\n", version, rulebook.TargetVersion)

						PRINTED = false // for next provider
					}
				} else {
					for _, resource := range resources {
						ruleset, err := resource.GetRuleset(rulebook)
						if err != nil {
							_, _ = commons.RED.Println(err)
							os.Exit(1)
						}

						annotate.AddAnnotationByRuleSet(resource, ruleset)

					}
				}
			}

			if !cmd.Flag("no-safe-upgrade-message").Changed {
				if len(upgradeMessage) != 0 {
					_, _ = commons.SIRREND.Println("You are safe to upgrade the following providers: ")
					fmt.Print(upgradeMessage)
				}
			}

		} else {
			yellow := color.New(color.FgYellow)
			_, _ = yellow.Println("Hmm..seems like you didn't setup this folder yet.\nPlease execute < terrap init >.")
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// scan command flags
	scanCmd.Flags().BoolP("json", "j", false, "Print scan output as json.")
	scanCmd.Flags().BoolP("annotate", "a", false, "Annotate the code itself.")
	scanCmd.Flags().BoolP("no-safe-upgrade-message", "n", false, "Don't print which providers are safe to upgrade")
}
