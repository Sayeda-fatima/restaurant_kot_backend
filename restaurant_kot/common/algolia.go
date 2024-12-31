package common

import (
	"os"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/search"
)

func Algolia() (*search.APIClient, error){

	appID := os.Getenv("APP_ID_ALGOLIA")

	apiKey := os.Getenv("API_KEY_ALGOLIA")

	// initialise the client
	client, err := search.NewClient(appID, apiKey)

	if err != nil{
		return nil, err
	}

	return client, nil
}