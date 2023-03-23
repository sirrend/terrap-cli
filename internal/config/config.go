package config

import (
	"github.com/fatih/color"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/state"
	"github.com/sirrend/terrap-cli/internal/utils"
	"log"
	"os"
	"path"
)

// struct to hold all the cli tool configuration parameters
type config struct {
	Username string
	Email    string
	Position string
	SendData bool
}

var c config

/*
@brief: PrintNotConfiguredMessage Prints the error message for not configuring terrap
*/

func PrintNotConfiguredMessage() {
	yellow := color.New(color.FgYellow)
	_, _ = yellow.Println("Hmm..seems like you didn't configure Terrap yet\nPlease execute < terrap config >.")
}

/*
@brief: IsConfigured Checks if the tool has been configured
@
@returns: bool - true if configured, false otherwise
*/

func IsConfigured() bool {
	if utils.DoesExist(commons.TERRAP_CONFIG_FILE) {
		return true
	}

	return false
}

/*
@brief: CreateConfigFolder creates the terrap configuration path in the user's home directory
@
@returns: error if failed
*/

func CreateConfigFolder() error {
	u, _ := os.UserHomeDir()
	if utils.DoesExist(path.Join(u, ".terrap/")) {
		log.Println("Configuration folder already exists..proceeding")

	} else {
		err := os.Mkdir(path.Join(u, ".terrap"), 0755)

		if err == nil {
			log.Printf("Created the user configuration folder in %v", u)
		} else {
			return err
		}
	}

	return nil
}

/*
@brief: SaveConfigurationFile saves the cli tool configuration file in the users home folder
@
@params: path string -> the path to the file to save
@
@returns: error if failed
*/

func SaveConfigurationFile(i interface{}) error {
	u, _ := os.UserHomeDir()

	err := state.Save(path.Join(u, ".terrap", "config"), i)

	if err != nil {
		return err
	}

	return err
}

/*
@brief: Configure configures the cli tool configuration file
@
@params: ct bool -> "configure terraform" is a bool set to indicate if to configure terraform cloud or not
@
@returns: error if failed
*/

func Configure(ct bool) error {
	err := CreateConfigFolder()
	if err != nil {
		return err
	} else {
		c.Username = utils.GetInput("Username: ")
		c.Email = utils.GetInput("Email: ")
		c.Position = utils.GetInput("Position: ")
		sendData := utils.GetInput("Help us improve! Y / N: ")

		if sendData == "y" || sendData == "Y" {
			c.SendData = true
		} else {
			c.SendData = false
		}

		if ct {
			utils.PrettyPrintStruct(c)
		}

		err = SaveConfigurationFile(c)
		if err != nil {
			return err
		}
	}

	return nil
}
