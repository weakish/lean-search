package main

import (
	"fmt"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/weakish/goaround"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		var logger *log.Logger = log.New(os.Stderr, "", 0)
		logger.Println("Please specify search keywords.")
		os.Exit(64)
	} else {
		var term string = strings.Join(os.Args[1:len(os.Args)], " ")
		var results algoliasearch.QueryRes
		results, err := search(term)
		goaround.FatalIf(err)
		displayHits(results)
	}
}

func search(term string) (algoliasearch.QueryRes, error) {
	appID := "BH4D9OD16A"  // case sensitive
	apiKey := "357b777ed18e79673a2c1de3f6c64478"
	client := algoliasearch.NewClient(appID, apiKey)
	index := client.InitIndex("leancloud")

	params := algoliasearch.Map{
		"hitsPerPage": 1000, // max value
	}

	return index.Search(term, params)
}

func displayHits(res algoliasearch.QueryRes) {
	if res.NbHits == 0 {
		fmt.Printf("No results found for query '%s'\n", res.Query)
	} else {
		for _, hit := range res.Hits {
			// We have to use `interface{}` because the value may contain `nil` or `string`.
			hierarchy := hit["hierarchy"].(map[string]interface{})
			for _, lvl := range []string{"lvl0", "lvl1", "lvl2", "lvl3", "lvl4", "lvl5", "lvl6"} {
				if hierarchy[lvl] == nil {
					fmt.Print("> ")
					break
				} else {
					fmt.Printf(" %s >", hierarchy[lvl])
				}
			}
			fmt.Println(hit["url"])
		}
	}
}
