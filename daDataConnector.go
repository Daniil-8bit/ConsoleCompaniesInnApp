package consolecompaniesinnapp

import (
	"context"
	"errors"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
)

var creds = client.Credentials{
	ApiKeyValue:    "",
	SecretKeyValue: "",
}

var api = dadata.NewSuggestApi(client.WithCredentialProvider(&creds))

func getInfo(name string) (string, string, error) {
	params := suggest.RequestParams{Query: name}
	result, err := api.Party(context.Background(), &params)

	if err != nil {
		return name, name, err
	}

	if len(result) < 1 {
		return "no inn", "", errors.New("no suggestions are found")
	}

	return result[0].Data.Inn, result[0].Data.Name.FullWithOpf, nil
}
