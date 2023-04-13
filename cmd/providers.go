/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"github.com/enescakir/emoji"
	"github.com/sirrend/terrap-cli/internal/cli_utils"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/providers_api"
	"github.com/sirrend/terrap-cli/internal/state"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/sirrend/terrap-cli/internal/workspace"
	"github.com/spf13/cobra"
	"strings"
)

// getContextCmd shows the providers which are currently in context
var getContextCmd = &cobra.Command{
	Use:   "get-context",
	Short: "Shows what providers are currently in the Terrap context",
	Run: func(cmd *cobra.Command, args []string) {
		var workspace workspace.Workspace
		var tableData [][]string
		table := cli_utils.GetTable([]string{"Provider", "Version"}) // initialize new table

		if utils.IsInitialized(".") {
			err := state.Load("./.terrap.json", &workspace)
			if err != nil {
				_, _ = commons.RED.Println(err)
			}

			for provider, version := range workspace.Providers { // collect providers
				provider = strings.ReplaceAll(provider, "registry.terraform.io/", "")

				if cmd.Flag("filter").Changed {
					if strings.Contains(provider, cmd.Flag("filter").Value.String()) {
						tableData = append(tableData, []string{provider, version.String()})
					}
				} else {
					tableData = append(tableData, []string{provider, version.String()})
				}
			}

			if len(tableData) > 0 {
				_, _ = commons.SIRREND.Println("Currently working with the following providers:")
				table.AppendBulk(tableData)
				table.Render()
			} else {
				if cmd.Flag("filter").Changed {
					_, _ = commons.YELLOW.Println("No providers matched your filter..", emoji.FaceWithoutMouth)
				} else {
					_, _ = commons.YELLOW.Println("No providers found in context", emoji.FaceWithoutMouth)
				}
			}

		} else {
			_, _ = commons.YELLOW.Println("Oops.. the current folder is not initialized, please execute <terrap init>.")
		}
	},
}

// providersCmd represents the providers command
var getSupportedProvidersCmd = &cobra.Command{
	Use:   "get-supported",
	Short: "Outputs which providers are supported by terrap.",
	Run: func(cmd *cobra.Command, args []string) {
		var tableData [][]string
		table := cli_utils.GetTable([]string{"Provider", "Min Version", "Max Version"}) // initialize new table
		providers, _ := providers_api.GetSupportedProviders()

		// go over providers retrieved from API
		for _, provider := range providers {
			if cmd.Flag("filter").Changed {
				if strings.Contains(provider.Name, cmd.Flag("filter").Value.String()) {
					tableData = append(tableData, []string{provider.Name, provider.MinVersion, provider.MaxVersion})
				}
			} else {
				tableData = append(tableData, []string{provider.Name, provider.MinVersion, provider.MaxVersion})
			}
		}

		// print results
		if len(tableData) > 0 {
			_, _ = commons.SIRREND.Println("The following providers are currently supported by Terrap: ")
			table.AppendBulk(tableData)
			table.Render()
		} else {
			if cmd.Flag("filter").Changed {
				_, _ = commons.YELLOW.Println("No providers matched your filter..", emoji.FaceWithoutMouth)
			}
		}

		_, _ = commons.HighMagenta.Println("\n", emoji.Man, emoji.Woman, "Our Sirrend RockStars are hard at work, expanding our engine's capability to connect with more providers.")
		_, _ = commons.HighMagenta.Println(" Exciting things are coming! Stay tuned: https://www.sirrend.com/")
	},
}

// providersCmd represents the providers command
var providersCmd = &cobra.Command{
	Use:   "providers",
	Short: "The providers command enables you to see which providers are in context and what providers are supported.",
}

func init() {
	rootCmd.AddCommand(providersCmd)

	providersCmd.AddCommand(getSupportedProvidersCmd)
	getSupportedProvidersCmd.Flags().StringP("filter", "f", "", "Show only the providers which match the filter applied.")

	providersCmd.AddCommand(getContextCmd)
	getContextCmd.Flags().StringP("filter", "f", "", "Show only the providers which match the filter applied.")
}
