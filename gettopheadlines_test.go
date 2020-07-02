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
