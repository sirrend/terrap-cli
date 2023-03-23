/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sirrend/terrap-cli/internal/handle_files"
	"github.com/sirrend/terrap-cli/internal/rules_interaction"
	"github.com/sirrend/terrap-cli/internal/state"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/sirrend/terrap-cli/internal/workspace"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan your IaC code to find provider changes",

	Run: func(cmd *cobra.Command, args []string) {
		var workspace workspace.Workspace

		if utils.IsInitialized(".") {
			resources, _ := handle_files.ScanFolderRecursively(".")

			err := state.Load("./.terrap.json", &workspace)
			if err != nil {
				fmt.Println(err)
			}

			for provider, version := range workspace.Providers {
				rulebookName := rules_interaction.FindRulebookFile(provider, version.String())

				rulebook, _ := rules_interaction.GetRulebook(rulebookName, version.String(), "")

				for _, resource := range resources {
					ruleset, err := resource.GetRuleset(rulebook)
					if err != nil {
						fmt.Println(err)
					}

					if cmd.Flag("json").Changed {
						ruleset.Execute("json")
					}
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
	scanCmd.Flags().BoolP("table", "t", false, "Print scan output as a table.")
	scanCmd.Flags().BoolP("json", "j", false, "Print scan output as json.")
	scanCmd.Flags().BoolP("annotate", "a", false, "Annotate the code itself.")
}
