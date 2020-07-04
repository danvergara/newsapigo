package newsapigo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAllSources(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allSourcesResponse := Response{
			Status: "ok",
			Sources: []Source{
				{
					ID:          "abc-news-au",
					Name:        "ABC News (AU)",
					Description: "Australia's most trusted source of local, national and world news. Comprehensive, independent, in-depth analysis, the latest business, sport, weather and more.",
					URL:         "http://www.abc.net.au/news",
					Category:    "general",
					Language:    "en",
					Country:     "au",
				},
			},
		}

		if r.URL.Path != "/sources" {
			t.Error("Bad sources path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(allSourcesResponse)
	}))

	defer sv.Close()
	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}
	c := NewsClient{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "FAKE_API_KEY",
	}

	params := SourcesArgs{}
	newsapiResponse, err := c.Sources(params)

	assert.Nil(t, err)
	assert.Greater(t, len(newsapiResponse.Sources), 0)
}

func TestEnglishSources(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		englishSourcesResponse := Response{
			Status: "ok",
			Sources: []Source{
				{
					ID:          "abc-news-au",
					Name:        "ABC News (AU)",
					Description: "Australia's most trusted source of local, national and world news. Comprehensive, independent, in-depth analysis, the latest business, sport, weather and more.",
					URL:         "http://www.abc.net.au/news",
					Category:    "general",
					Language:    "en",
					Country:     "au",
				},
			},
		}

		if r.URL.Path != "/sources" {
			t.Error("Bad sources path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(englishSourcesResponse)
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}
	c := NewsClient{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "FAKE_API_KEY",
	}

	params := SourcesArgs{
		Language: "eng",
	}

	newsapiResponse, err := c.Sources(params)

	assert.Nil(t, err)
	assert.Greater(t, len(newsapiResponse.Sources), 0)
}

func TestEnglishSourceInUS(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		USSourcesResponse := Response{
			Status: "ok",
			Sources: []Source{
				{
					ID:          "abc-news",
					Name:        "ABC News",
					Description: "Your trusted source for breaking news, analysis, exclusive interviews, headlines, and videos at ABCNews.com.",
					URL:         "https://abcnews.go.com",
					Category:    "general",
					Language:    "en",
					Country:     "us",
				},
			},
		}
		if r.URL.Path != "/sources" {
			t.Error("Bad sources path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(USSourcesResponse)
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}
	c := NewsClient{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "FAKE_API_KEY",
	}

	params := SourcesArgs{
		Language: "eng",
		Country:  "us",
	}

	newsapiResponse, err := c.Sources(params)

	assert.Nil(t, err)
	assert.Greater(t, len(newsapiResponse.Sources), 0)
}

func TestSourcesBadRequest(t *testing.T) {
	// Create a fake server in order to return a fake but realistic response
	// Status Code: Bad Request
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a status code equals to 400 - Bad Request
		w.WriteHeader(http.StatusBadRequest)
		// Set the error response
		errorResponse := newError("The request was unacceptable, often due to a missing or misconfigured parameter", "parameterInvalid")
		encoder := json.NewEncoder(w)

		encoder.Encode(errorResponse)
	}))

	defer sv.Close()
	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}
	// Create the client of the package
	c := NewsClient{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "FAKE_API_KEY",
	}

	// Create a SourcesArgs instance that will be used as query params
	queryParams := SourcesArgs{
		Language: "5hiu4g5tiu34b",
		Country:  "4ui6bghiu5h",
	}
	newsErrorResponse, err := c.Sources(queryParams)

	// assertions section
	assert.NotNil(t, err)
	assert.Equal(t, "error", newsErrorResponse.Status)
	assert.Equal(t, "The request was unacceptable, often due to a missing or misconfigured parameter", newsErrorResponse.Message)
	assert.Equal(t, "parameterInvalid", newsErrorResponse.Code)
}

func TestSourcesUnauthorized(t *testing.T) {
	// Create a fake server in order to return a fake but realistic response
	// Status code: Unauthorized
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a status code equals to 400 - Bad Request
		w.WriteHeader(http.StatusUnauthorized)
		errorResponse := newError("Your API key has been disabled", "apiKeyDisabled")
		encoder := json.NewEncoder(w)
		encoder.Encode(errorResponse)
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}
	// Create the client of the package
	c := NewsClient{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "WRONG_API_KEY",
	}

	// Create a SourcesArgs instance that will be used as query params
	queryParams := SourcesArgs{
		Language: "eng",
		Country:  "us",
	}
	newsErrorResponse, err := c.Sources(queryParams)

	// assertions section
	assert.NotNil(t, err)
	assert.Equal(t, "error", newsErrorResponse.Status)
	assert.Equal(t, "Your API key has been disabled", newsErrorResponse.Message)
	assert.Equal(t, "apiKeyDisabled", newsErrorResponse.Code)
}

func TestSourcesServerError(t *testing.T) {
	// Create a fake server in order to return a fake but realistic response
	// Status code: Server Error
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a status code equals to 500 - Server Error
		w.WriteHeader(http.StatusUnauthorized)
		errorResponse := newError("This shouldn't happen, and if it does then it's our fault, not yours. Try the request again shortly", "unexpectedError")
		encoder := json.NewEncoder(w)
		encoder.Encode(errorResponse)
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}
	// Create the client of the package
	c := NewsClient{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "FAKE_API_KEY",
	}

	// Create a SourcesArgs instance that will be used as query params
	queryParams := SourcesArgs{
		Language: "eng",
		Country:  "us",
	}
	newsErrorResponse, err := c.Sources(queryParams)

	// assertions section
	assert.NotNil(t, err)
	assert.Equal(t, "error", newsErrorResponse.Status)
	assert.Equal(t, "This shouldn't happen, and if it does then it's our fault, not yours. Try the request again shortly", newsErrorResponse.Message)
	assert.Equal(t, "unexpectedError", newsErrorResponse.Code)
}
