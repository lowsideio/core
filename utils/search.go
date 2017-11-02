package utils

import (
  "os"

  "github.com/labstack/echo"
  "github.com/algolia/algoliasearch-client-go/algoliasearch"
)

var client algoliasearch.Client;
var index algoliasearch.Index;

func InitSearch() {
  client = algoliasearch.NewClient(os.Getenv("ALGOLIA_CLIENT_ID"), os.Getenv("ALGOLIA_SECRET_API_KEY"))
  index = client.InitIndex("motorcycles")
}


type AlgoliaContext struct {
  echo.Context
}
func (c *AlgoliaContext) AlgoliaIndex() algoliasearch.Index {
  return index
}
