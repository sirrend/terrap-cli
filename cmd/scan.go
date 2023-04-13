/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/fatih/color"
	"github.com/sirrend/terrap-cli/internal/annotate"
	"github.com/sirrend/terrap-cli/internal/cli_commons"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/handle_files"
	"github.com/sirrend/terrap-cli/internal/rules_api"
	"github.com/sirrend/terrap-cli/internal/scanning"
	"github.com/sirrend/terrap-cli/internal/state"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/sirrend/terrap-cli/internal/workspace"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	PRINTED                = false
	upgradeMessage         = ""
	notYetSupportedMessage = ""
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan your IaC code to find provider changes",

	Run: func(cmd *cobra.Command, args []string) {
		var workspace workspace.Workspace
		asJson := map[string]map[string]interface{}{}

		providersSet := cmd.Flag("fixed-providers").Changed
		if utils.IsInitialized(".") || providersSet {
			resources, err := handle_files.ScanFolderRecursively(".")
			if err != nil {
				_, _ = commons.RED.Println(err)
			}

			// find resource appearances
			resourceAppearances := scanning.WhereDoesResourceAppear(resources)

			if !providersSet {
				err = state.Load("./.terrap.json", &workspace)
				if err != nil {
					_, _ = commons.RED.Println(err)
				}
			} else {
				workspace = cli_commons.GetFixedProvidersFlag(*cmd)
			}

			// go over every provider in user's folder
			for provider, version := range workspace.Providers {
				rulebook, err := rules_api.GetRules(provider, version.String())

				// validate rulebook downloaded
				if err != nil {
					if strings.Contains(err.Error(), utils.StripProviderPrefix(provider)) {
						notYetSupportedMessage = strings.Join([]string{notYetSupportedMessage, err.Error()}, ", ")
						continue
					}

					// TODO: add error here
					continue
				}

				flags := cli_commons.ChangedComponentsFlags(*cmd)
				if !cmd.Flag("annotate").Changed {
					for _, resource := range scanning.GetUniqResources(resources) {
						if utils.IsItemInSlice(resource.Type, flags) {

							ruleset, err := resource.GetRuleset(rulebook, resourceAppearances)
							if err != nil {
								_, _ = commons.RED.Println(err)
								os.Exit(1)
							}

							// fill in json object
							if cmd.Flag("json").Changed {
								if ruleset.Rules != nil {
									asJson[ruleset.ResourceName] = map[string]interface{}{
										"Rules":       ruleset.Rules,
										"Appearances": resourceAppearances[resource.Name],
									}
								}

								// print human-readable output
							} else {
								ruleset.PrettyPrint()
								if len(ruleset.Rules) != 0 {
									PRINTED = true
								}
							}
						} else {
							PRINTED = true // to avoid wrong possible upgrade message
						}
					}

					// print json object
					if cmd.Flag("json").Changed {
						utils.PrettyPrintStruct(asJson)
						if len(asJson) != 0 {
							PRINTED = true
						}
					}

					// check if upgrade is possible for the provider in context
					if !PRINTED {
						if rulebook.TargetVersion != "" {
							upgradeMessage += fmt.Sprintf("  %s %s: ",
								emoji.UpArrow, utils.StripProviderPrefix(provider))
							upgradeMessage += commons.GREEN.Sprintf("%s -> %s\n", version, rulebook.TargetVersion)

							PRINTED = false // for next provider
						}
					}

				} else {
					for _, resource := range resources {
						ruleset, err := resource.GetRuleset(rulebook, resourceAppearances)
						if err != nil {
							_, _ = commons.RED.Println(err)
							os.Exit(1)
						}

						annotate.AddAnnotationByRuleSet(resource, ruleset)

					}
				}
			}

			// print safe upgrade message
			if !cmd.Flag("no-safe-upgrade-message").Changed && !cmd.Flag("no-messages").Changed {
				if len(upgradeMessage) != 0 {
					_, _ = commons.SIRREND.Println("The following providers are safe to upgrade: ")
					fmt.Println(upgradeMessage)
				}
			}

			// print not supported message
			if !cmd.Flag("no-not-supported-message").Changed && !cmd.Flag("no-messages").Changed {
				if notYetSupportedMessage != "" {
					message := strings.TrimLeft(notYetSupportedMessage, ", ")
					_, _ = commons.SIRREND.Println("========== This message can be suppressed using --no-not-supported-message ==========")
					_, _ = commons.SIRREND.Print("The following providers are not yet supported: ")
					fmt.Println(message, emoji.CryingFace.String())
					_, _ = commons.HighMagenta.Print("Check again soon! ")
					fmt.Println("we're actively working on increasing our Providers support " + emoji.BuildingConstruction.String())
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

	// utility flags
	scanCmd.Flags().BoolP("json", "j", false, "Print scan output as json.")
	scanCmd.Flags().BoolP("annotate", "a", false, "Annotate the code itself.")
	scanCmd.Flags().BoolP("provider", "p", false, "Show only provider changes.")
	scanCmd.Flags().BoolP("data-sources", "d", false, "Show only data source changes.")
	scanCmd.Flags().BoolP("resources", "r", false, "Show only resources changes.")
	scanCmd.Flags().StringSlice("fixed-providers", []string{}, "Set fixed provider version in the following format: `provider:version`.\nIf this flag is used, all other in-context providers are ignored.")

	// extra output flags
	scanCmd.Flags().Bool("no-safe-upgrade-message", false, "Don't print which providers are safe to upgrade.")
	scanCmd.Flags().Bool("no-not-supported-message", false, "Don't print if providers are not supported.")
	scanCmd.Flags().BoolP("no-messages", "n", false, "Don't print any message other than pure command output.")
}
