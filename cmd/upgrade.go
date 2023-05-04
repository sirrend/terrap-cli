/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"context"
	"fmt"
	version "github.com/hashicorp/go-version"
	"github.com/ktr0731/go-updater"
	"github.com/ktr0731/go-updater/brew"
	"github.com/spf13/cobra"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade your Terrap version to the latest version available in the Sirrend tap",
	Run: func(cmd *cobra.Command, args []string) {
		currentVersion, _ := version.NewVersion("0.0.3")

		formula := brew.HomebrewMeans("sirrend/sirrend", "terrap")
		if formulaMeans, err := updater.NewMeans(formula); err == nil {
			u := updater.New(currentVersion, formulaMeans)
			b, v, _ := u.Updatable(context.Background())
			fmt.Println(b, v.String())
			err := u.Update(context.Background())
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
		//fmt.Println(updater)
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upgradeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upgradeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
