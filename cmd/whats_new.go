/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/fatih/color"
	"github.com/sirrend/terrap-cli/internal/cli_utils"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/files_handler"
	"github.com/sirrend/terrap-cli/internal/rules_api"
	"github.com/sirrend/terrap-cli/internal/state"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/sirrend/terrap-cli/internal/workspace"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// whatsNewCmd represents the whatsNew command
var whatsNewCmd = &cobra.Command{
	Use:   "whats-new",
	Short: "Shows what's up next in the following version of every provider in context",
	Run: func(cmd *cobra.Command, args []string) {
		var workspace workspace.Workspace
		asJson := map[string][]string{}

		providersSet := cmd.Flag("fixed-providers").Changed
		if utils.IsInitialized(".") || providersSet {
			if !providersSet {
				err := state.Load("./.terrap.json", &workspace)
				if err != nil {
					_, _ = commons.RED.Println(err)
				}
			} else {
				workspace = cli_utils.GetFixedProvidersFlag(*cmd)
			}

			for provider, version := range workspace.Providers { // go over every provider in user's folder
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

				ruleSets, err := rulebook.GetAllRuleSets()
				if err != nil {
					os.Exit(1)
				}

				flags := cli_utils.ChangedComponentsFlags(*cmd) // get resources filtering
				for resourcesType, resources := range ruleSets {
					if utils.IsItemInSlice(resourcesType, flags) {
						for resourceName, _ := range resources.(map[string]interface{}) { // go over all ruleSets
							resource := files_handler.Resource{Name: resourceName, Type: resourcesType}
							ruleset, err := resource.GetRuleset(rulebook, nil)
							if err != nil {
								os.Exit(1)
							}

							// fill in json object
							if cmd.Flag("json").Changed {
								if ruleset.Rules != nil {
									asJson[resource.Name] = ruleset.GetNewComponents()
								}

								// print human-readable output
							} else {
								ruleset.PrettyPrintWhatsNew()
								if len(ruleset.Rules) != 0 {
									PRINTED = true

								}
							}
						}
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

			// print not supported message
			if !cmd.Flag("no-not-supported-message").Changed && !cmd.Flag("no-messages").Changed {
				if notYetSupportedMessage != "" {
					message := strings.TrimLeft(notYetSupportedMessage, ", ")
					_, _ = commons.SIRREND.Print("The following providers are not yet supported: ")
					fmt.Println(message, emoji.CryingFace.String())
				}
			}

		} else {
			yellow := color.New(color.FgYellow)
			_, _ = yellow.Println("Hmm..seems like you didn't setup this folder yet.\nPlease execute < terrap init >.")
		}
	},
}

func init() {
	rootCmd.AddCommand(whatsNewCmd)

	// utility flags
	whatsNewCmd.Flags().BoolP("json", "j", false, "Print whats-new output as json.")
	whatsNewCmd.Flags().BoolP("provider", "p", false, "Show only provider changes.")
	whatsNewCmd.Flags().BoolP("data-sources", "d", false, "Show only data source changes.")
	whatsNewCmd.Flags().BoolP("resources", "r", false, "Show only resources changes.")
	whatsNewCmd.Flags().StringSlice("fixed-providers", []string{}, "A comma separated list of fixed providers written in the following format: `<provider>:<version>`.If this flag is used, all other in-context providers are ignored.")

	// extra output flags
	whatsNewCmd.Flags().Bool("no-not-supported-message", false, "Don't print if providers are not supported.")
	whatsNewCmd.Flags().BoolP("no-messages", "n", false, "Don't print any message other than pure command output.")
}
