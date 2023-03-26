package commons

import (
	"github.com/fatih/color"
	"os"
	"path"
)

var (
	TERRAP_HOME_FOLDER      = path.Join(getUserHomeDir(), ".terrap")
	TERRAP_PROVIDERS_FOLDER = path.Join(getUserHomeDir(), ".terrap", "providers")
	TERRAP_CONFIG_FILE      = path.Join(getUserHomeDir(), ".terrap", "config")

	RULES_BUCKET        = "terrap-rulebooks"
	RULES_BUCKET_REGION = "eu-central-1"

	YELLOW = color.New(color.FgYellow)
	GREEN  = color.New(color.FgGreen)
	RED    = color.New(color.BgRed)
)

func getUserHomeDir() string {
	u, _ := os.UserHomeDir()
	return u
}
