package commons

import (
	"github.com/fatih/color"
	"os"
)

var (
	ProviderAPI  = "https://api.sirrend.com/supported-providers"
	RulebooksAPI = "https://api.sirrend.com/rulebook"
	GitHubOwner  = "sirrend"
	GitHubRepo   = "terrap-cli"

	YELLOW      = color.New(color.FgYellow)
	GREEN       = color.New(color.FgGreen)
	RED         = color.New(color.FgRed)
	SIRREND     = color.New(color.FgHiMagenta)
	HighMagenta = color.New(color.FgMagenta)
)

func getUserHomeDir() string {
	u, _ := os.UserHomeDir()
	return u
}
