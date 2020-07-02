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

func TestGetAllSources(t *testing.T) {
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

func TestGetEnglishSources(t *testing.T) {
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

func TestGetEnglishSourceInUS(t *testing.T) {
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
