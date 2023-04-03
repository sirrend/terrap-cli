package rules_api

import (
	"github.com/sirrend/terrap-cli/internal/utils"
	"net/http"
	"net/url"
)

// GetRules
/*
@brief:
	GetRules retrieves a rulebook from the API
@params:
	rulebookName - string - the rulebook name
	sourceVersion - string - the source version
@returns:
	Rules - the downloaded rules
	error - if exists, else nil
*/
func GetRules(provider, sourceVersion string) (Rulebook, error) {
	// make the request
	u, _ := url.Parse("https://oumoaz1pu2.execute-api.eu-west-1.amazonaws.com/prod")
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

	return Rulebook{
		SourceVersion: sourceVersion,
		Bytes:         body,
	}, nil
}
