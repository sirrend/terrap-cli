package commons

import (
	"github.com/fatih/color"
	"os"
	"path"
)

var (
	TerrapProvidersFolder = path.Join(getUserHomeDir(), ".terrap", "providers")
	TerrapConfigFile      = path.Join(getUserHomeDir(), ".terrap", "config")

	YELLOW  = color.New(color.FgYellow)
	GREEN   = color.New(color.FgGreen)
	RED     = color.New(color.FgRed)
	SIRREND = color.New(color.FgMagenta)
)

func getUserHomeDir() string {
	u, _ := os.UserHomeDir()
	return u
}
