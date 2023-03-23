/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"github.com/sirrend/terrap-cli/internal/handle_user_files"
	"github.com/sirrend/terrap-cli/internal/utils"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan your IaC code to find provider changes",
	Run: func(cmd *cobra.Command, args []string) {
		resources, _ := handle_user_files.ScanFolderRecursively("./terraform-test")
		utils.PrettyPrintStruct(resources)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// scan command flags
	scanCmd.Flags().BoolP("table", "t", false, "Print scan output as a table.")
	scanCmd.Flags().BoolP("json", "j", false, "Print scan output as json.")
	scanCmd.Flags().BoolP("annotate", "a", false, "Annotate the code itself.")
}
