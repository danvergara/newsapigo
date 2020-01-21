package newsapigo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllSources(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allSourcesResponse := NewsResponse{
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

		if r.URL.Path != "/v2/sources" {
			t.Error("Bad sources path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(allSourcesResponse)
	}))

	defer sv.Close()

	c := NewsClient{
		BaseURL: sv.URL,
		APIKey:  "FAKE_API_KEY",
	}

	params := url.Values{}
	newsapiResponse, err := c.GetSources(params)

	assert.Nil(t, err)
	assert.Greater(t, len(newsapiResponse.Sources), 0)
}

func TestGetEnglishSources(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		englishSourcesResponse := NewsResponse{
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

		if r.URL.Path != "/v2/sources" {
			t.Error("Bad sources path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(englishSourcesResponse)
	}))

	defer sv.Close()

	c := NewsClient{
		BaseURL: sv.URL,
		APIKey:  "FAKE_API_KEY",
	}

	params := url.Values{}
	params.Add("language", "eng")

	newsapiResponse, err := c.GetSources(params)

	assert.Nil(t, err)
	assert.Greater(t, len(newsapiResponse.Sources), 0)
}

func TestGetEnglishSourceInUS(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		USSourcesResponse := NewsResponse{
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
		if r.URL.Path != "/v2/sources" {
			t.Error("Bad sources path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(USSourcesResponse)
	}))

	defer sv.Close()

	c := NewsClient{
		BaseURL: sv.URL,
		APIKey:  "FAKE_API_KEY",
	}

	params := url.Values{}
	params.Add("langauge", "en")
	params.Add("country", "us")

	newsapiResponse, err := c.GetSources(params)

	assert.Nil(t, err)
	assert.Greater(t, len(newsapiResponse.Sources), 0)
}
