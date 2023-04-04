/*
Copyright © 2023 Sirrend

*/

package cmd

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/providers"
	"github.com/sirrend/terrap-cli/internal/state"
	"github.com/sirrend/terrap-cli/internal/terraform_utils"
	"github.com/sirrend/terrap-cli/internal/utils"
	"os"
	"path"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

func sirrendLogoPrint() {
	myFigure := figure.NewColorFigure("Sirrend", "", "purple", true)
	myFigure.Print()
}

/*
@brief: terraformInit performs the Terraform init command on the given folder
@
@params: dir - the folder to initialize
*/

func terraformInit(dir string) {
	_, err := os.Stat(path.Join(dir, ".terrap.json"))

	if err != nil {
		_, _ = commons.YELLOW.Print(emoji.Rocket, " Initializing directory...")
		mainWorkspace.ExecPath,
			mainWorkspace.TerraformVersion, err = terraform_utils.TerraformInit(dir) // initiate new terraform tool in context
		_, _ = commons.GREEN.Println("Done!")

		if err != nil {
			fmt.Printf("Failed initializing the given directory with the following error: %e", err)
		}

		_, _ = commons.YELLOW.Print(emoji.Toolbox, " Looking for providers...")
		providers.FindTfProviders(dir, &mainWorkspace) //find all providers and assign to mainWorkspace
		_, _ = commons.GREEN.Println("Done!")

		_, _ = commons.YELLOW.Print(emoji.WavingHand, " Saving workspace...")
		saveInitData() //Save to configuration file
		_, _ = commons.GREEN.Println("Done!")

	} else {
		_, _ = commons.YELLOW.Print("Folder already initialized..")
		os.Exit(0)

	}
}

/*
@brief: saveInitData saves the configuration file of the initialized folder
*/
func saveInitData() {
	err := state.Save(path.Join(mainWorkspace.Location, ".terrap.json"), mainWorkspace)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
@brief: deleteInitData deletes the configuration file of the initialized folder
@
@params: dir - the folder to delete the init file from
*/
func deleteInitData(dir string) {
	_, _ = commons.YELLOW.Print(emoji.Broom, " Cleaning up former configuration...")
	err := os.Remove(path.Join(dir, ".terrap.json"))
	if err != nil {
		_, _ = commons.YELLOW.Println("The given directory is not initialized.")
		os.Exit(1)
	}
	_, _ = commons.GREEN.Println("Done!")
}

// the init command declaration
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize directory",
	Long:  `Initialize directory for sirrend to have all needed files`,

	Run: //check which flags are set and run the appropriate init
	func(cmd *cobra.Command, args []string) {
		if cmd.Flag("current-directory").Changed && cmd.Flag("directory").Changed {
			_, _ = commons.YELLOW.Println("You can't set both -c and -d flags")
			os.Exit(0)

		} else if cmd.Flag("destroy").Changed {
			if !cmd.Flag("current-directory").Changed && !cmd.Flag("directory").Changed {
				_, _ = commons.YELLOW.Println("You must set one of the flags -c / -d")
				os.Exit(0)
			}

			if cmd.Flag("current-directory").Changed {
				c, _ := os.Getwd() // get current directory
				deleteInitData(c)
			} else {
				deleteInitData(cmd.Flag("directory").Value.String())
			}

		} else if cmd.Flag("upgrade").Changed {
			if !cmd.Flag("current-directory").Changed && !cmd.Flag("directory").Changed {
				_, _ = commons.YELLOW.Println("You must set one of the flags -c / -d")
				os.Exit(0)
			}

			if cmd.Flag("current-directory").Changed {
				c, _ := os.Getwd() // get current directory
				deleteInitData(c)
				d, _ := filepath.Abs(cmd.Flag("directory").Value.String())
				mainWorkspace.Location = d
				terraformInit(c)

				fmt.Println()
				_, _ = commons.SIRREND.Println(emoji.BeerMug, "Terrap directory upgraded!")
			} else {
				d, _ := filepath.Abs(cmd.Flag("directory").Value.String())
				deleteInitData(d)
				mainWorkspace.Location = d
				terraformInit(d)

				fmt.Println()
				_, _ = commons.SIRREND.Println(emoji.BeerMug, "Terrap directory upgraded!")
			}

		} else if cmd.Flag("directory").Changed {
			sirrendLogoPrint()
			fmt.Println()

			if utils.IsDir(cmd.Flag("directory").Value.String()) {
				d, _ := filepath.Abs(cmd.Flag("directory").Value.String())
				mainWorkspace.Location = d
				terraformInit(d)
				_, _ = commons.SIRREND.Println("\nTerrap Initialized Successfully!")

			} else {
				_, _ = commons.YELLOW.Println("The given path is not a directory")
				os.Exit(0)
			}

		} else if cmd.Flag("current-directory").Changed {
			sirrendLogoPrint()
			fmt.Println()

			mainWorkspace.Location, _ = os.Getwd() // get current directory
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
	initCmd.Flags().BoolP("current-directory", "c", false, "initialize the current directory")
	initCmd.Flags().StringP("directory", "d", "", "initialize the given directory")
	initCmd.Flags().Bool("destroy", false, "Destroy the terrap context in this folder. Needs to be set alongside the -d or -c flag.")
	initCmd.Flags().BoolP("upgrade", "u", false, "Upgrade the given directory init file")
}
