package newsapigo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "newsapi.org",
	Path:   "/v2/",
}

// NewsClient is the Main User Interface of this package
// This struct is responsible to make the requests to News API
type NewsClient struct {
	apiKey     string
	baseURL    *url.URL
	httpClient *http.Client
}

// New returns a new instance of the NewsClient, given an valid apiKey
// The Client has all the necessary methods to interact with the NewsAPI REST API
func New(apiKey string) *NewsClient {
	c := &http.Client{Timeout: time.Minute}

	return &NewsClient{
		httpClient: c,
		apiKey:     apiKey,
		baseURL:    &baseURL,
	}
}

// Everything returns a NewsResponse.
// Search through millions of articles from over 30,000 large and small news
// sources and blogs. This includes breaking news as well as lesser articles.
func (c *NewsClient) Everything(args EverythingArgs) (Response, error) {
	endpt := c.baseURL.ResolveReference(&url.URL{Path: everythingPath})

	req, err := http.NewRequest("GET", endpt.String(), nil)
	if err != nil {
		return newError(err.Error(), "unexpectedError"), err
	}

	if c.apiKey == "" {
		return newError("Your API Key is missing", "apiKeyMissing"), fmt.Errorf("API Key is missing")
	}

	// add URL headers, query params, then send the request
	req.Header.Add("Accept", "application/json")
	req.Header.Set("X-Api-Key", c.apiKey)
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return newError(err.Error(), "unexpectedError"), err
	}
	// Deserialize the response and return the newsapi data
	defer res.Body.Close()

	var response Response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return newError(err.Error(), "unexpectedError"), err
	}

	return response, nil
}

// Sources return a NewsResponse.
// This endpoint returns the subset of news publishers that top headlines
// (/v2/top-headlines) are available from. It's mainly a convenience endpoint
// that you can use to keep track of the publishers available on the API,
// and you can pipe it straight through to your users.
func (c *NewsClient) Sources(args SourcesArgs) (Response, error) {
	endpt := c.baseURL.ResolveReference(&url.URL{Path: sourcesPath})

	req, err := http.NewRequest("GET", endpt.String(), nil)
	if err != nil {
		return newError(err.Error(), "unexpectedError"), err
	}

	if c.apiKey == "" {
		return newError("Your API Key is missing", "apiKeyMissing"), fmt.Errorf("API Key is missing")
	}

	// add URL headers, query params, then send the request
	req.Header.Add("Accept", "application/json")
	req.Header.Set("X-Api-Key", c.apiKey)
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return newError(err.Error(), "unexpectedError"), err
	}
	// Deserialize the response and return the newsapi data
	defer res.Body.Close()

	var response Response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return newError(err.Error(), "unexpectedError"), err
	}

	return response, nil
}

// TopHeadlines returns a NewsResponse with 20 articles by default.
// Provides live top and breaking headlines for a country, specific category in
// a country, single source, or multiple sources.
func (c *NewsClient) TopHeadlines(args TopHeadlinesArgs) (Response, error) {
	endpt := c.baseURL.ResolveReference(&url.URL{Path: topHeadlinesPath})

	req, err := http.NewRequest("GET", endpt.String(), nil)
	if err != nil {
		return newError(err.Error(), "unexpectedError"), err
	}

	if c.apiKey == "" {
		return newError("Your API Key is missing", "apiKeyMissing"), fmt.Errorf("API Key is missing")
	}

	// add URL headers, query params, then send the request
	req.Header.Add("Accept", "application/json")
	req.Header.Set("X-Api-Key", c.apiKey)
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return newError(err.Error(), "unexpectedError"), err
	}
	// Deserialize the response and return the newsapi data
	defer res.Body.Close()

	var response Response
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return newError(err.Error(), "unexpectedError"), err
	}

	return response, nil
}
