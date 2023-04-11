package providers_api

import (
	"errors"
	"github.com/Jeffail/gabs"
	"github.com/sirrend/terrap-cli/internal/utils"
	"net/http"
)

// GetSupportedProviders
/*
@brief:
	GetSupportedProviders retrieves all supported providers
@returns:
	[]string - list of supported Terraform providers
	error - if exists, else nil
*/
func GetSupportedProviders() ([]string, error) {
	var parsedProviders []string

	// Prepare the method
	req, _ := http.NewRequest("GET", "https://ty2nr6s4cvivq7zjxujpxi4aq40uqlqk.lambda-url.eu-west-1.on.aws/", nil)

	// perform the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []string{}, err
	}

	// read body
	body := utils.StreamToByte(res.Body)

	container, _ := gabs.ParseJSON(body)
	if container.ExistsP("error") {
		return []string{}, errors.New("Oops.. something went wrong on out side. But rest asure we are on it!")
	}

	providers, _ := container.Path("providers").Children()
	for _, provider := range providers {
		parsedProviders = append(parsedProviders, provider.String())
	}

	return parsedProviders, nil
}
