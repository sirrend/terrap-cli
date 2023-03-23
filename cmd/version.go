/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sirrend/terrap-cli/internal/terrap_version"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version description of terrap",
	Run: func(cmd *cobra.Command, args []string) {
		t := terrap_version.TerrapVersion{}
		t.SetVersion() // set the version details

		if cmd.Flag("json").Changed {
			j, _ := utils.Marshal(t)

			_, err := io.Copy(os.Stdout, j)
			if err != nil {
				log.Fatal(err)
			}
		} else if cmd.Flag("tool-only").Changed {
			fmt.Println(t.Version)

		} else {
			fmt.Println("Version:", t.Product+"-"+t.Version, t.System+"-"+t.GoVersion)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().BoolP("json", "j", false, "Output as json")
	versionCmd.Flags().BoolP("tool-only", "t", false, "Print only the tool version")
}
