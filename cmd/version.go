/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/sirrend/terrap-cli/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version description of terrap",
	Run: func(cmd *cobra.Command, args []string) {
		t := version.TerrapVersion{}
		t.SetVersion() // set the version details

		if cmd.Flag("json").Changed {
			j, _ := utils.Marshal(t)
			_, _ = io.Copy(os.Stdout, j) // Error handling can be modified as per your requirements
			return
		}

		if cmd.Flag("tool-only").Changed {
			fmt.Println(t.Version)
			return
		}

		fmt.Printf("Version: %s-%s %s-%s",
			t.Product, t.Version,
			t.System, t.GoVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().BoolP("json", "j", false, "Output as json")
	versionCmd.Flags().BoolP("tool-only", "t", false, "Print only the tool version")
}
