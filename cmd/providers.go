/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/providers_api"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

// providersCmd represents the providers command
var providersCmd = &cobra.Command{
	Use:   "providers",
	Short: "The providers command enables you to see which providers are in context and what providers are supported.",
	Run: func(cmd *cobra.Command, args []string) {
		providers, _ := providers_api.GetSupportedProviders()

		_, _ = commons.SIRREND.Println("The following providers are currently supported by Terrap: ")
		for index, provider := range providers {
			_, _ = commons.SIRREND.Print("  " + cast.ToString(index+1))
			fmt.Println(". " + utils.MustUnquote(provider))
		}

		_, _ = commons.HighMagenta.Println("\n", emoji.Man, emoji.Woman, "Our Sirrend RockStars are hard at work, expanding our engine's capability to connect with more providers.")
		_, _ = commons.HighMagenta.Println(" Exciting things are coming! Stay tuned: https://www.sirrend.com/")
	},
}

func init() {
	rootCmd.AddCommand(providersCmd)
}
