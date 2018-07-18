package main

import (
	"fmt"

	"github.com/dany2691/newsapigo"
)

func main() {
	// First create a map filled with all the parameters needed.
	// At least one is mandatory.
	// For Top Headlines endpoint, you can choose from: country, category, sources, q, pageSize and page.
	// For Everything endpoint, you can choose from: domains, from, to, language, sources, q,  soryBy, pageSize and page.
	// For Sources endpoint, you can choose from: category, language and country.
	params := make(map[string]string)
	params["country"] = "us"
	//Then, declare a variable with your apiKey from https://newsapi.org
	apiKey := "api-key"
	// To get a response from top headlines endpoint, call GetTopHeadlines
	responseToHeadlines := newsapigo.GetTopHeadlines(params, apiKey)

	// You can  display the status of the response:
	fmt.Printf("Status: %s", responseToHeadlines.Status)

	// To display the title of the first article, yo can do the following:
	fmt.Printf("The title of the first articles is %s", responseToHeadlines.Articles[0].Title)
}
