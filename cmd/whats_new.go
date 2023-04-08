/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"github.com/enescakir/emoji"
	"github.com/fatih/color"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/handle_files"
	"github.com/sirrend/terrap-cli/internal/rules_api"
	"github.com/sirrend/terrap-cli/internal/state"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/sirrend/terrap-cli/internal/workspace"
	"github.com/spf13/cobra"
	"os"
)

// whatsNewCmd represents the whatsNew command
var whatsNewCmd = &cobra.Command{
	Use:   "whats-new",
	Short: "Shows what's up next in the following version of every provider in context",
	Run: func(cmd *cobra.Command, args []string) {
		var workspace workspace.Workspace
		asJson := map[string][]string{}

		if utils.IsInitialized(".") {
			err := state.Load("./.terrap.json", &workspace)
			if err != nil {
				_, _ = commons.RED.Println(err)
			}

			for provider, version := range workspace.Providers { // go over every provider in user's folder
				rulebook, err := rules_api.GetRules(provider, version.String())
				if err != nil {
					_, _ = commons.RED.Println(emoji.CrossMark, "Terrap failed to receive changes, please make sure you're connected to the internet.")
					os.Exit(1)
				}

				ruleSets, err := rulebook.GetAllRuleSets()
				if err != nil {
					os.Exit(1)
				}

				for resourceName, _ := range ruleSets { // go over all ruleSets
					resource := handle_files.Resource{Name: resourceName}
					ruleset, err := resource.GetRuleset(rulebook)
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

			// print json object
			if cmd.Flag("json").Changed {
				utils.PrettyPrintStruct(asJson)
				if len(asJson) != 0 {
					PRINTED = true
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
	whatsNewCmd.Flags().BoolP("json", "j", false, "Print whats-new output as json.")
}
