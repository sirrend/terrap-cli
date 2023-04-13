package providers_api

import (
	"errors"
	"github.com/Jeffail/gabs"
	"github.com/sirrend/terrap-cli/internal/utils"
	"net/http"
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
	req, _ := http.NewRequest("GET", "https://ty2nr6s4cvivq7zjxujpxi4aq40uqlqk.lambda-url.eu-west-1.on.aws/", nil)

	// perform the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []Provider{}, err
	}

	// read body
	body := utils.StreamToByte(res.Body)

	container, _ := gabs.ParseJSON(body)
	if container.ExistsP("error") {
		return []Provider{}, errors.New("Oops.. something went wrong on out side. But rest asure we are on it!")
	}

	providers, _ := container.Path("providers").Children()
	for _, provider := range providers {
		p := Provider{
			Name:       utils.MustUnquote(provider.Path("provider").String()),
			MinVersion: utils.MustUnquote(provider.Path("min_version").String()),
			MaxVersion: utils.MustUnquote(provider.Path("max_version").String()),
		}
		parsedProviders = append(parsedProviders, p)
	}

	return parsedProviders, nil
}
