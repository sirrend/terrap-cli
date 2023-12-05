package receiver

import (
	"errors"

	"github.com/Jeffail/gabs"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/parser"
	"github.com/sirrend/terrap-cli/internal/requests"
	"github.com/sirrend/terrap-cli/internal/utils"
)

// GetRules
/*
@brief:

	GetRules retrieves a rulebook from the API

@params:

	provider - string - the provider name
	sourceVersion - string - the source version

@returns:

	Rulebook - the initializes Rulebook
	error - if exists, else nil
*/
func GetRules(provider, sourceVersion string) (parser.Rulebook, error) {
	// make the request
	res, err := requests.PerformRequestWithParams(commons.RulebooksAPI, map[string]string{"provider": provider, "source_version": sourceVersion})
	if err != nil {
		return parser.Rulebook{}, err
	}

	body := utils.StreamToByte(res.Body)

	rulebook := parser.Rulebook{
		SourceVersion: sourceVersion,
		Bytes:         body,
	}
	rulebook.TargetVersion = rulebook.GetTargetVersion()

	container, _ := gabs.ParseJSON(body)
	if container.ExistsP("error") {
		return rulebook, errors.New(utils.StripProviderPrefix(provider) + ":" + rulebook.SourceVersion)
	}

	return rulebook, nil
}
