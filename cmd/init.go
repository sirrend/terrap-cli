/*
Copyright © 2023 Sirrend

*/

package cmd

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/state"
	"github.com/sirrend/terrap-cli/internal/terraform_utils"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
)

// terraformInit
/*
@brief:
	terraformInit performs the Terraform init command on the given folder
@params:
	dir - the folder to initialize
*/
func terraformInit(dir string) {
	_, err := os.Stat(path.Join(dir, ".terrap.json"))

	if err != nil {
		_, _ = commons.YELLOW.Println(emoji.Rocket, "Initializing directory...")
		mainWorkspace.ExecPath, mainWorkspace.IsTempProvider,
			mainWorkspace.TerraformVersion, err = terraform_utils.TerraformInit(dir) // initiate new terraform tool in context

		if err != nil {
			fmt.Println()
			terraform_utils.TerraformErrorPrettyPrint(err)
			os.Exit(1)

		}

		_, _ = commons.YELLOW.Print(emoji.Toolbox, " Looking for providers...")
		terraform_utils.FindTfProviders(dir, &mainWorkspace) //find all providers and assign to mainWorkspace
		_, _ = commons.GREEN.Println(" Done!")

		_, _ = commons.YELLOW.Print(emoji.WavingHand, " Saving workspace...")
		saveInitData() //Save to configuration file
		_, _ = commons.GREEN.Println(" Done!")

	} else {
		_, _ = commons.YELLOW.Println("Folder already initialized..")
		os.Exit(0)

	}
}

/*
@brief: saveInitData saves the configuration file of the initialized folder
*/
func saveInitData() {
	err := state.Save(path.Join(mainWorkspace.Location, ".terrap.json"), mainWorkspace)
	if err != nil {
		_, _ = commons.RED.Println("Terrap failed saving the current workspace.")
		os.Exit(1)
	}
}

// the init command declaration
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize directory",
	Long:  `Initialize directory for sirrend to have all needed files`,

	Run: //check which flags are set and run the appropriate init
	func(cmd *cobra.Command, args []string) {
		if cmd.Flag("upgrade").Changed {
			var directory string

			if cmd.Flag("directory").Changed {
				directory, _ = filepath.Abs(cmd.Flag("directory").Value.String())
			} else {
				directory, _ = os.Getwd()
			}

			deleteInitData(directory)
			mainWorkspace.Location = directory
			terraformInit(directory)

			fmt.Println()
			_, _ = commons.SIRREND.Println(emoji.BeerMug, "Terrap directory upgraded!")

		} else if cmd.Flag("directory").Changed {
			if utils.IsDir(cmd.Flag("directory").Value.String()) {
				directory, _ := filepath.Abs(cmd.Flag("directory").Value.String())
				mainWorkspace.Location = directory
				terraformInit(directory)
				_, _ = commons.SIRREND.Println("\nTerrap Initialized Successfully!")

			} else {
				_, _ = commons.YELLOW.Println("The given path is not a directory")
				os.Exit(0)
			}

		} else {
			location, err := os.Getwd() // get current directory
			if err != nil {
				_, _ = commons.RED.Print(emoji.AngryFace, "Failed with the following error: ")
				fmt.Println(err.Error())
				os.Exit(1)
			}
			mainWorkspace.Location = location
			terraformInit(mainWorkspace.Location)

			fmt.Println()
			_, _ = commons.SIRREND.Println(emoji.BeerMug, "Terrap Initialized Successfully!")
		}
	},
}

/*
@brief: init adds the command to the root command (terrap)
*/
func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("directory", "d", "", "initialize a given directory")
	initCmd.Flags().BoolP("upgrade", "u", false, "Upgrade the given directory workspace")
}
