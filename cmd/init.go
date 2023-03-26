/*
Copyright © 2023 Sirrend

*/

package cmd

import (
	"fmt"
	"github.com/sirrend/terrap-cli/internal/providers"
	"github.com/sirrend/terrap-cli/internal/state"
	"github.com/sirrend/terrap-cli/internal/terraform_utils"
	"github.com/sirrend/terrap-cli/internal/utils"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

/*
@brief: terraformInit performs the Terraform init command on the given folder
@
@params: dir - the folder to initialize
*/

func terraformInit(dir string) {
	_, err := os.Stat(path.Join(dir, ".terrap.json"))

	if err != nil {
		mainWorkspace.ExecPath,
			mainWorkspace.TerraformVersion, err = terraform_utils.TerraformInit(dir) // initiate new terraform tool in context

		if err != nil {
			fmt.Printf("Failed initializing the given directory with the following error: %e", err)
		}

		providers.FindTfProviders(dir, &mainWorkspace) //find all providers and assign to mainWorkspace
		saveInitData()                                 //Save to configuration file

	} else {
		log.Println("Folder already initialized..")
		err := state.Load(path.Join(dir, ".terrap.json"), &mainWorkspace)
		if err != nil {
			log.Fatal(err)
		}
	}
}

/*
@brief: saveInitData saves the configuration file of the initialized folder
*/
func saveInitData() {
	err := state.Save(path.Join(mainWorkspace.Location, ".terrap.json"), mainWorkspace)
	if err != nil {
		log.Fatal(err)
	}
}

/*
@brief: deleteInitData deletes the configuration file of the initialized folder
@
@params: dir - the folder to delete the init file from
*/
func deleteInitData(dir string) {
	err := os.Remove(path.Join(dir, ".terrap.json"))
	if err != nil {
		log.Fatal("The given directory is not initialized.")
	}
}

// the init command declaration
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize directory",
	Long:  `Initialize directory for sirrend to have all needed files`,

	Run: //check which flags are set and run the appropriate init
	func(cmd *cobra.Command, args []string) {
		if cmd.Flag("current-directory").Changed && cmd.Flag("directory").Changed {
			log.Fatal("You can't set both -c and -d flags")

		} else if cmd.Flag("destroy").Changed {
			if !cmd.Flag("current-directory").Changed && !cmd.Flag("directory").Changed {
				log.Fatal("You must set one of the flags -c / -d")
			}

			if cmd.Flag("current-directory").Changed {
				c, _ := os.Getwd() // get current directory
				deleteInitData(c)
			} else {
				deleteInitData(cmd.Flag("directory").Value.String())
			}

		} else if cmd.Flag("upgrade").Changed {
			if !cmd.Flag("current-directory").Changed && !cmd.Flag("directory").Changed {
				log.Fatal("You must set one of the flags -c / -d")
			}

			if cmd.Flag("current-directory").Changed {
				c, _ := os.Getwd() // get current directory
				deleteInitData(c)
				d, _ := filepath.Abs(cmd.Flag("directory").Value.String())
				mainWorkspace.Location = d
				terraformInit(c)

				fmt.Println("Terrap directory upgraded!")
			} else {
				d, _ := filepath.Abs(cmd.Flag("directory").Value.String())
				deleteInitData(d)
				mainWorkspace.Location = d
				terraformInit(d)
				fmt.Println("Terrap directory upgraded!")
			}

		} else if cmd.Flag("directory").Changed {
			if utils.IsDir(cmd.Flag("directory").Value.String()) {
				d, _ := filepath.Abs(cmd.Flag("directory").Value.String())
				mainWorkspace.Location = d
				terraformInit(d)

			} else {
				log.Fatal("Not a directory")
			}

		} else if cmd.Flag("current-directory").Changed {
			c, _ := os.Getwd() // get current directory
			mainWorkspace.Location = c
			terraformInit(c)
		}
	},
}

/*
@brief: init adds the command to the root command (terrap)
*/
func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("current-directory", "c", false, "initialize the current directory")
	initCmd.Flags().StringP("directory", "d", "", "initialize the given directory")
	initCmd.Flags().Bool("destroy", false, "Destroy the terrap context in this folder. Needs to be set alongside the -d or -c flag.")
	initCmd.Flags().BoolP("upgrade", "u", false, "Upgrade the given directory init file")
}
