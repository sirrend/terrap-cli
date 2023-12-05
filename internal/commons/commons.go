package commons

import (
	"github.com/fatih/color"
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
