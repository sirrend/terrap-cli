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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
