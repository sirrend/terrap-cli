package commons

import (
	"github.com/fatih/color"
	"os"
)

var (
	API         = "https://api.sirrend.com"
	GitHubOwner = "sirrend"
	GitHubRepo  = "terrap-cli"

	YELLOW      = color.New(color.FgYellow)
	GREEN       = color.New(color.FgGreen)
	RED         = color.New(color.FgHiRed)
	SIRREND     = color.New(color.FgHiMagenta)
	HighMagenta = color.New(color.FgMagenta)
)

func getUserHomeDir() string {
	u, _ := os.UserHomeDir()
	return u
}
