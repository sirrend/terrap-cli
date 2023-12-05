/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"os"
	"path"
	"path/filepath"

	"github.com/enescakir/emoji"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/state"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/sirrend/terrap-cli/internal/utils/terraform"
	"github.com/sirrend/terrap-cli/internal/workspace"

	"github.com/spf13/cobra"
)

/*
@brief: deleteInitData deletes the configuration file of the initialized folder
@
@params: dir - the folder to delete the init file from
*/
func deleteInitData(dir string) {
	var ws workspace.Workspace
	if utils.IsInitialized(dir) {
		_, _ = commons.YELLOW.Println(emoji.Broom, "Cleaning up former configuration...")

		err := state.Load(filepath.Join(dir, ".terrap.json"), &ws)
		if err != nil {
			_, _ = commons.RED.Println(err)
			os.Exit(1)
		}

		if ws.IsTempProvider {
			err = terraform.RemoveTempTerraformExecutor(ws.ExecPath)
			if err != nil {
				_, _ = commons.RED.Println(err)
				os.Exit(1)
			}
			_, _ = commons.YELLOW.Println(emoji.CheckMark, " Temporary Terraform executor removed")
		}

		err = os.Remove(path.Join(dir, ".terrap.json"))
		if err != nil {
			_, _ = commons.RED.Println(emoji.CrossMark, " Failed to clean up the current workspace.")
			os.Exit(1)
		}
		_, _ = commons.YELLOW.Println(emoji.CheckMark, " Configuration file removed")

		// delete Terraform files if exist
		_ = os.Remove(path.Join(dir, ".terraform.lock.hcl"))
		_ = os.RemoveAll(path.Join(dir, ".terraform"))

		_, _ = commons.GREEN.Println("\nWorkspace removed.")

	} else {
		_, _ = commons.YELLOW.Println("The given directory is not initialized.")
		os.Exit(0)
	}
}

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroys the given Terrap workspace",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("directory").Changed {
			deleteInitData(cmd.Flag("directory").Value.String())
		} else {
			currentPath, _ := os.Getwd()
			deleteInitData(currentPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)

	destroyCmd.Flags().StringP("directory", "d", "", "Supply a directory to destroy")
}
