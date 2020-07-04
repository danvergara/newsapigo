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

func TestTopHeadlinesByCountry(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		topHeadlinesByContruyResponse := Response{
			Status:       "ok",
			TotalResults: 38,
			Articles: []Article{
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
				{
					Source: &Source{
						ID:   "cnn",
						Name: "CNN",
					},
					Author:      "Devan Cole, CNN",
					Title:       "House impeachment manager to Trump: If Ukraine call was 'perfect' then 'call the witnesses' - CNN",
					Description: "House impeachment manager Jason Crow on Sunday argued that witnesses should be called in President Donald Trump's upcoming Senate trial that kicks off in earnest on Tuesday.",
					URL:         "https://www.cnn.com/2020/01/19/politics/jason-crown-impeachment-house-manager-cnntv/index.html",
					URLToImage:  "https://cdn.cnn.com/cnnnext/dam/assets/200115111410-jason-crow-file-super-tease.jpg",
					PublishedAt: "2020-01-19T17:44:00Z",
					Content:     "Washington (CNN)House impeachment manager Jason Crow...",
				},
			},
		}

		if r.URL.Path != "/top-headlines" {
			t.Error("Bad top-headlines path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(topHeadlinesByContruyResponse)
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}
	c := NewsClient{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "FAKE_API_KEY",
	}

	params := TopHeadlinesArgs{
		Country: "us",
	}

	newsapiResponse, err := c.TopHeadlines(params)

	assert.Nil(t, err)
	assert.Equal(t, 20, len(newsapiResponse.Articles))
	assert.Equal(t, "ok", newsapiResponse.Status)
	assert.Equal(t, 38, newsapiResponse.TotalResults)
}

func TestTopHeadlinesBySource(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		topHeadlinesBySourceResponse := Response{
			Status:       "ok",
			TotalResults: 38,
			Articles: []Article{
				{
					Source: &Source{
						ID:   "bbc-news",
						Name: "BBC News",
					},
					Author:      "BBC News",
					Title:       "Africa's richest woman 'ripped off her country",
					Description: "Leaked documents reveal how Isabel dos Santos made her fortune through exploitation and corruption.",
					URL:         "http://www.bbc.co.uk/news/world-africa-51128950",
					URLToImage:  "https://ichef.bbci.co.uk/news/1024/branded_news/11BB1/production/_110552627_isabel976.jpg",
					PublishedAt: "2020-01-19T18:07:19Z",
					Content:     "Image copyrightGetty Images",
				},
				{
					Source: &Source{
						ID:   "bbc-news",
						Name: "BBC News",
					},
					Author:      "BBC News",
					Title:       "Africa's richest woman 'ripped off her country",
					Description: "Leaked documents reveal how Isabel dos Santos made her fortune through exploitation and corruption.",
					URL:         "http://www.bbc.co.uk/news/world-africa-51128950",
					URLToImage:  "https://ichef.bbci.co.uk/news/1024/branded_news/11BB1/production/_110552627_isabel976.jpg",
					PublishedAt: "2020-01-19T18:07:19Z",
					Content:     "Image copyrightGetty Images",
				},
				{
					Source: &Source{
						ID:   "bbc-news",
						Name: "BBC News",
					},
					Author:      "BBC News",
					Title:       "Africa's richest woman 'ripped off her country",
					Description: "Leaked documents reveal how Isabel dos Santos made her fortune through exploitation and corruption.",
					URL:         "http://www.bbc.co.uk/news/world-africa-51128950",
					URLToImage:  "https://ichef.bbci.co.uk/news/1024/branded_news/11BB1/production/_110552627_isabel976.jpg",
					PublishedAt: "2020-01-19T18:07:19Z",
					Content:     "Image copyrightGetty Images",
				},
				{
					Source: &Source{
						ID:   "bbc-news",
						Name: "BBC News",
					},
					Author:      "BBC News",
					Title:       "Africa's richest woman 'ripped off her country",
					Description: "Leaked documents reveal how Isabel dos Santos made her fortune through exploitation and corruption.",
					URL:         "http://www.bbc.co.uk/news/world-africa-51128950",
					URLToImage:  "https://ichef.bbci.co.uk/news/1024/branded_news/11BB1/production/_110552627_isabel976.jpg",
					PublishedAt: "2020-01-19T18:07:19Z",
					Content:     "Image copyrightGetty Images",
				},
				{
					Source: &Source{
						ID:   "bbc-news",
						Name: "BBC News",
					},
					Author:      "BBC News",
					Title:       "Africa's richest woman 'ripped off her country",
					Description: "Leaked documents reveal how Isabel dos Santos made her fortune through exploitation and corruption.",
					URL:         "http://www.bbc.co.uk/news/world-africa-51128950",
					URLToImage:  "https://ichef.bbci.co.uk/news/1024/branded_news/11BB1/production/_110552627_isabel976.jpg",
					PublishedAt: "2020-01-19T18:07:19Z",
					Content:     "Image copyrightGetty Images",
				},
				{
					Source: &Source{
						ID:   "bbc-news",
						Name: "BBC News",
					},
					Author:      "BBC News",
					Title:       "Africa's richest woman 'ripped off her country",
					Description: "Leaked documents reveal how Isabel dos Santos made her fortune through exploitation and corruption.",
					URL:         "http://www.bbc.co.uk/news/world-africa-51128950",
					URLToImage:  "https://ichef.bbci.co.uk/news/1024/branded_news/11BB1/production/_110552627_isabel976.jpg",
					PublishedAt: "2020-01-19T18:07:19Z",
					Content:     "Image copyrightGetty Images",
				},
				{
					Source: &Source{
						ID:   "bbc-news",
						Name: "BBC News",
					},
					Author:      "BBC News",
					Title:       "Africa's richest woman 'ripped off her country",
					Description: "Leaked documents reveal how Isabel dos Santos made her fortune through exploitation and corruption.",
					URL:         "http://www.bbc.co.uk/news/world-africa-51128950",
					URLToImage:  "https://ichef.bbci.co.uk/news/1024/branded_news/11BB1/production/_110552627_isabel976.jpg",
					PublishedAt: "2020-01-19T18:07:19Z",
					Content:     "Image copyrightGetty Images",
				},
				{
					Source: &Source{
						ID:   "bbc-news",
						Name: "BBC News",
					},
					Author:      "BBC News",
					Title:       "Africa's richest woman 'ripped off her country",
					Description: "Leaked documents reveal how Isabel dos Santos made her fortune through exploitation and corruption.",
					URL:         "http://www.bbc.co.uk/news/world-africa-51128950",
					URLToImage:  "https://ichef.bbci.co.uk/news/1024/branded_news/11BB1/production/_110552627_isabel976.jpg",
					PublishedAt: "2020-01-19T18:07:19Z",
					Content:     "Image copyrightGetty Images",
				},
				{
					Source: &Source{
						ID:   "bbc-news",
						Name: "BBC News",
					},
					Author:      "BBC News",
					Title:       "Africa's richest woman 'ripped off her country",
					Description: "Leaked documents reveal how Isabel dos Santos made her fortune through exploitation and corruption.",
					URL:         "http://www.bbc.co.uk/news/world-africa-51128950",
					URLToImage:  "https://ichef.bbci.co.uk/news/1024/branded_news/11BB1/production/_110552627_isabel976.jpg",
					PublishedAt: "2020-01-19T18:07:19Z",
					Content:     "Image copyrightGetty Images",
				},
				{
					Source: &Source{
						ID:   "bbc-news",
						Name: "BBC News",
					},
					Author:      "BBC News",
					Title:       "Africa's richest woman 'ripped off her country",
					Description: "Leaked documents reveal how Isabel dos Santos made her fortune through exploitation and corruption.",
					URL:         "http://www.bbc.co.uk/news/world-africa-51128950",
					URLToImage:  "https://ichef.bbci.co.uk/news/1024/branded_news/11BB1/production/_110552627_isabel976.jpg",
					PublishedAt: "2020-01-19T18:07:19Z",
					Content:     "Image copyrightGetty Images",
				},
			},
		}

		if r.URL.Path != "/top-headlines" {
			t.Error("Bad top-headlines path")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(topHeadlinesBySourceResponse)
	}))

	defer sv.Close()

	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}
	c := NewsClient{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "FAKE_API_KEY",
	}

	params := TopHeadlinesArgs{
		Sources: []string{"bbc-news"},
	}
	newsapiResponse, err := c.TopHeadlines(params)

	assert.Nil(t, err)
	assert.Equal(t, 10, len(newsapiResponse.Articles))
	assert.Equal(t, "ok", newsapiResponse.Status)
	assert.Equal(t, 38, newsapiResponse.TotalResults)
}

func TestTopHeadlinesBadRequest(t *testing.T) {
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
	params := TopHeadlinesArgs{
		Country: "3g56hj34",
	}

	newsErrorResponse, err := c.TopHeadlines(params)

	// assertions section
	assert.NotNil(t, err)
	assert.Equal(t, "error", newsErrorResponse.Status)
	assert.Equal(t, "The request was unacceptable, often due to a missing or misconfigured parameter", newsErrorResponse.Message)
	assert.Equal(t, "parameterInvalid", newsErrorResponse.Code)
}

func TestTopHeadlinesUnauthorized(t *testing.T) {
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

	// Create a TopHeadlinesArgs instance that will be used as query params
	params := TopHeadlinesArgs{
		Sources: []string{"bbc-news"},
	}
	newsErrorResponse, err := c.TopHeadlines(params)
	// assertions section
	assert.NotNil(t, err)
	assert.Equal(t, "error", newsErrorResponse.Status)
	assert.Equal(t, "Your API key has been disabled", newsErrorResponse.Message)
	assert.Equal(t, "apiKeyDisabled", newsErrorResponse.Code)
}

func TestTopHeadlinesServerError(t *testing.T) {
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

	// Create a TopHeadlinesArgs instance that will be used as query params
	params := TopHeadlinesArgs{
		Country: "us",
	}

	newsErrorResponse, err := c.TopHeadlines(params)
	// assertions section
	assert.NotNil(t, err)
	assert.Equal(t, "error", newsErrorResponse.Status)
	assert.Equal(t, "This shouldn't happen, and if it does then it's our fault, not yours. Try the request again shortly", newsErrorResponse.Message)
	assert.Equal(t, "unexpectedError", newsErrorResponse.Code)
}
