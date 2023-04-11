package commons

import (
	"github.com/fatih/color"
	"os"
	"path"
)

var (
	TerrapProvidersFolder = path.Join(getUserHomeDir(), ".terrap", "providers")
	TerrapConfigFile      = path.Join(getUserHomeDir(), ".terrap", "config")

	YELLOW      = color.New(color.FgYellow)
	GREEN       = color.New(color.FgGreen)
	RED         = color.New(color.FgHiRed)
	SIRREND     = color.New(color.FgHiMagenta)
	HighMagenta = color.New(color.FgMagenta)
	GitHubOwner = "sirrend"
	GitHubRepo  = "terrap-cli"
)

func getUserHomeDir() string {
	u, _ := os.UserHomeDir()
	return u
}
