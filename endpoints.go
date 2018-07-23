package newsapigo

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var provider = &URLProvider{}

func init() {
	provider.TopHeadlines = "https://newsapi.org/v2/top-headlines?"
	provider.Everything = "https://newsapi.org/v2/everything?"
	provider.Sources = "https://newsapi.org/v2/sources?"
}

var instance *NewsAPIKey

// GetAPIKey returns an pointer to an instance of the NewsApiKey struct
// Makes sure that there is an only on instance
func GetAPIKey() *NewsAPIKey {
	if instance == nil {
		instance = new(NewsAPIKey)
	}
	return instance
}

// GetEverything returns a NewsResponse.
// Search through millions of articles from over 30,000 large and small news
// sources and blogs. This includes breaking news as well as lesser articles.
func GetEverything(payload map[string]string) NewsResponse {
	url := concreteURL(provider.Everything, payloadString(payload), GetAPIKey().Key)
	req := makeRquest(url)
	resp := getAPIResponse(req)
	return encodeJSON(resp)
}

// GetSources return a NewsResponse.
// This endpoint returns the subset of news publishers that top headlines
// (/v2/top-headlines) are available from. It's mainly a convenience endpoint
// that you can use to keep track of the publishers available on the API,
// and you can pipe it straight through to your users.
func GetSources(payload map[string]string) NewsResponse {
	url := concreteURL(provider.Sources, payloadString(payload), GetAPIKey().Key)
	req := makeRquest(url)
	resp := getAPIResponse(req)
	return encodeJSON(resp)
}

// GetTopHeadlines returns a NewsResponse with 20 articles by default.
// Provides live top and breaking headlines for a country, specific category in
// a country, single source, or multiple sources.
func GetTopHeadlines(payload map[string]string) NewsResponse {
	url := concreteURL(provider.TopHeadlines, payloadString(payload), GetAPIKey().Key)
	req := makeRquest(url)
	resp := getAPIResponse(req)
	return encodeJSON(resp)
}

// SetKey set the key of the the unique instance of the NewsApiKey struct
func (a *NewsAPIKey) SetKey(keyString string) {
	a.Key = keyString
}

func concreteURL(basicURL, payloadStr, apiKey string) string {
	var buffer bytes.Buffer
	buffer.WriteString(basicURL)
	buffer.WriteString(payloadStr)
	buffer.WriteString("apikey=")
	buffer.WriteString(apiKey)
	return buffer.String()
}

func encodeJSON(resp *http.Response) NewsResponse {
	var response NewsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println(err)
	}
	return response
}

func getAPIResponse(req *http.Request) *http.Response {
	client := &http.Client{Timeout: time.Second * 3}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}
	return resp
}

func makeRquest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Request failed:", err)
	}
	return req
}

func payloadString(payload map[string]string) string {
	var buffer bytes.Buffer
	for k, v := range payload {
		buffer.WriteString(k)
		buffer.WriteString("=")
		buffer.WriteString(v)
		buffer.WriteString("&")
	}
	return buffer.String()
}
