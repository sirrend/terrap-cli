/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"context"
	"github.com/hashicorp/go-version"
	"github.com/ktr0731/go-updater"
	"github.com/ktr0731/go-updater/brew"
	"github.com/sirrend/terrap-cli/internal/commons"
	terrapVersion "github.com/sirrend/terrap-cli/internal/version"
	"github.com/spf13/cobra"
)

func BrewUpdate() {
	currentVersion, _ := version.NewVersion(terrapVersion.Version)

	formula := brew.HomebrewMeans(commons.BrewFormula, commons.BrewProductName)
	if formulaMeans, err := updater.NewMeans(formula); err == nil {
		u := updater.New(currentVersion, formulaMeans)
		u.UpdateIf = updater.FoundPatchUpdate

		if err = u.Update(context.Background()); err == nil {
			commons.GREEN.Println("Upgraded Terrap to the latest version!")
		} else {
			commons.RED.Println(err)
		}

	} else {
		commons.RED.Println(err)
	}
}

func InstallMakefile(repoDir string) bool {

	return true
}

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade your Terrap version to the latest version available in the Sirrend tap",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("brew").Changed {
			BrewUpdate()
			//} else {
			//	var (
			//		location string
			//		err      error
			//	)
			//
			//	if cmd.Flag("tag").Changed {
			//		location, err = cloaner.CloneRepoByTag("https://github.com/sirrend/terrap-cli.git", cmd.Flag("tag").Value.String())
			//		if err != nil {
			//			commons.RED.Println(err)
			//		}
			//
			//	} else {
			//		location, err = cloaner.CloneLatest("https://github.com/sirrend/terrap-cli.git")
			//		if err != nil {
			//			fmt.Println(err)
			//		}
			//	}
			//
		}
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	upgradeCmd.Flags().BoolP("brew", "b", false, "Use Brew package manager to update Terrap.")
}
