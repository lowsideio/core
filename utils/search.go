package utils

import (
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

var client *search.Client

var index *search.Index

func InitSearch() {
	client = search.NewClient(os.Getenv("ALGOLIA_APP_ID"), os.Getenv("ALGOLIA_API_KEY"))
	index = client.InitIndex("motorcycles")
}

func (c *RequestContext) AlgoliaIndex() *search.Index {
	return index
}
