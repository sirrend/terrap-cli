package commons

import (
	"os"
	"path"
)

var (
	TERRAP_HOME_FOLDER      = path.Join(getUserHomeDir(), ".terrap")
	TERRAP_PROVIDERS_FOLDER = path.Join(getUserHomeDir(), ".terrap", "providers")
	TERRAP_CONFIG_FILE      = path.Join(getUserHomeDir(), ".terrap", "config")
	RESET_COLOR             = "\033[0m"
	RED_COLOR               = "\033[31m"
	GREEN_COLOR             = "\033[32m"
	YELLOW_COLOR            = "\033[33m"

	RULES_BUCKET        = "terrap-rulebooks"
	RULES_BUCKET_REGION = "eu-central-1"
)

func getUserHomeDir() string {
	u, _ := os.UserHomeDir()
	return u
}
