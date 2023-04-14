package rules_api

import (
	"errors"
	"github.com/Jeffail/gabs"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/utils"
	"net/http"
	"net/url"
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
func GetRules(provider, sourceVersion string) (Rulebook, error) {
	// make the request
	u, _ := url.Parse(commons.RulebooksAPI)
	query := u.Query()
	query.Set("provider", provider)
	query.Set("source_version", sourceVersion)

	u.RawQuery = query.Encode()

	// Prepare the method
	req, _ := http.NewRequest("GET", u.String(), nil)

	// perform the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Rulebook{}, err
	}

	// read body
	body := utils.StreamToByte(res.Body)

	rulebook := Rulebook{
		SourceVersion: sourceVersion,
		Bytes:         body,
	}
	rulebook.TargetVersion = rulebook.GetTargetVersion()

	container, _ := gabs.ParseJSON(body)
	if container.ExistsP("error") {
		return rulebook, errors.New(utils.StripProviderPrefix(provider))
	}

	return rulebook, nil
}
