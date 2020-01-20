package newsapigo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// NewsClient is the Main User Interface of this package
// This struct is responsible to make the requests to News API
type NewsClient struct {
	APIKey     string
	BaseURL    string
	httpClient *http.Client
}

// GetEverything returns a NewsResponse.
// Search through millions of articles from over 30,000 large and small news
// sources and blogs. This includes breaking news as well as lesser articles.
func (c *NewsClient) GetEverything(queryParams url.Values) (NewsResponse, error) {
	u := c.buildURL(everythingPath, queryParams)

	req, err := buildRquest(u)

	if err != nil {
		return errorNewsResponse(err.Error(), ""), err
	}

	if c.APIKey == "" {
		return errorNewsResponse("Your API Key is missing", "apiKeyMissing"),
			fmt.Errorf("API Key is missing")
	}

	c.setHeader(req)

	resp, err := c.doRequest(req)
	if err != nil {
		return errorNewsResponse(err.Error(), ""), err
	}

	response, err := c.decodeResponse(resp)
	return response, nil
}

// GetSources return a NewsResponse.
// This endpoint returns the subset of news publishers that top headlines
// (/v2/top-headlines) are available from. It's mainly a convenience endpoint
// that you can use to keep track of the publishers available on the API,
// and you can pipe it straight through to your users.
func (c *NewsClient) GetSources(queryParams url.Values) (NewsResponse, error) {
	u := c.buildURL(sourcesPath, queryParams)

	req, err := buildRquest(u)

	if err != nil {
		return errorNewsResponse(err.Error(), ""), err
	}

	if c.APIKey == "" {
		return errorNewsResponse("Your API Key is missing", "apiKeyMissing"),
			fmt.Errorf("API Key is missing")
	}

	c.setHeader(req)

	resp, err := c.doRequest(req)
	if err != nil {
		return errorNewsResponse(err.Error(), ""), err
	}

	response, err := c.decodeResponse(resp)
	return response, nil
}

// GetTopHeadlines returns a NewsResponse with 20 articles by default.
// Provides live top and breaking headlines for a country, specific category in
// a country, single source, or multiple sources.
func (c *NewsClient) GetTopHeadlines(queryParams url.Values) (NewsResponse, error) {
	u := c.buildURL(topHeadlinesPath, queryParams)

	req, err := buildRquest(u)

	if err != nil {
		return errorNewsResponse(err.Error(), ""), err
	}

	if c.APIKey == "" {
		return errorNewsResponse("Your API Key is missing", "apiKeyMissing"),
			fmt.Errorf("API Key is missing")
	}

	c.setHeader(req)

	resp, err := c.doRequest(req)
	if err != nil {
		return errorNewsResponse(err.Error(), ""), err
	}

	response, err := c.decodeResponse(resp)
	return response, nil
}

func (c *NewsClient) buildURL(path string, queryParams url.Values) *url.URL {
	u, err := url.Parse(c.urlBase())

	if err != nil {
		log.Fatal(err)
	}

	u.Path = path
	u.RawQuery = queryParams.Encode()
	return u
}

func (c *NewsClient) decodeResponse(resp *http.Response) (NewsResponse, error) {
	var response NewsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response, err
	}

	return response, nil
}

func (c *NewsClient) doRequest(req *http.Request) (*http.Response, error) {
	c.setClient()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *NewsClient) setClient() {
	if c.httpClient == nil {
		c.httpClient = &http.Client{Timeout: time.Second * 3}
	}
}

func (c *NewsClient) setHeader(req *http.Request) {
	req.Header.Set("X-Api-Key", c.APIKey)
	req.Header.Set("Content-Type", "application/json")
}

func (c *NewsClient) urlBase() string {
	if c.BaseURL == "" {
		return defaultBaseURL
	}

	return c.BaseURL
}
