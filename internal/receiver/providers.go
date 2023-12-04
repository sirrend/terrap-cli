package receiver

import (
	"errors"

	"github.com/Jeffail/gabs"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/requests"
	"github.com/sirrend/terrap-cli/internal/utils"
)

// Provider represents a provider fetched from the API and all it's attributes
type Provider struct {
	Name       string `json:"Provider"`
	MinVersion string `json:"MinVersion"`
	MaxVersion string `json:"MaxVersion"`
}

// GetSupportedProviders
/*
@brief:
	GetSupportedProviders retrieves all supported providers
@returns:
	[]string - list of supported Terraform providers
	error - if exists, else nil
*/
func GetSupportedProviders() ([]Provider, error) {
	var parsedProviders []Provider

	// Prepare the method
	res, err := requests.PerformRequest("GET", commons.ProviderAPI, nil)
	if err != nil {
		return []Provider{}, err
	}

	// read body
	body := utils.StreamToByte(res.Body)

	container, _ := gabs.ParseJSON(body)
	if container.ExistsP("error") {
		return []Provider{}, errors.New("oops.. something went wrong on our side. But rest asure we are on it")
	}

	providers, _ := container.Path("providers").Children()
	for _, provider := range providers {
		unquote := func(param string) string {
			return utils.MustUnquote(provider.Path(param).String())
		}

		parsedProviders = append(parsedProviders, Provider{
			Name:       unquote("provider"),
			MinVersion: unquote("min_version"),
			MaxVersion: unquote("max_version"),
		})
	}

	return parsedProviders, nil
}
