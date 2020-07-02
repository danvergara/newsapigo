package newsapigo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetEverythingBitcoin(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		exampleResponse := Response{
			Status:       "ok",
			TotalResults: 3421,
			Articles: []Article{
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
				{
					Source: &Source{
						ID:   "techcrunch",
						Name: "TechCrunch",
					},
					Author:      "Jonathan Shieber",
					Title:       "All tulips must wilt",
					Description: "It’s a bad day in the world of cryptocurrency.",
					URL:         "http://techcrunch.com/2019/12/18/all-tulips-must-wilt/",
					URLToImage:  "https://techcrunch.com/wp-content/uploads/2019/12/Screen-Shot-2019-12-18-at-9.39.21-AM.png?w=669",
					PublishedAt: "2019-12-18T15:17:17Z",
					Content:     "It’s a bad day in the world of cryptocurrency.",
				},
			},
		}

		jsonResponse, err := json.Marshal(exampleResponse)

		if err != nil {
			panic("Oops")
		}

		if r.URL.Path != "/everything" {
			t.Error("Bad everything path")
		}

		fmt.Fprint(w, string(jsonResponse))
	}))

	defer sv.Close()
	rawURL, _ := url.Parse(sv.URL)

	queryParams := EverythingArgs{
		Q: "bitcoin",
	}
	testClient := &http.Client{Timeout: time.Minute}
	c := NewsClient{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "FAKE_API_KEY",
	}

	newsapiResponse, err := c.Everything(queryParams)

	assert.Equal(t, "ok", newsapiResponse.Status)
	assert.Equal(t, 3421, newsapiResponse.TotalResults)
	assert.Equal(t, 20, len(newsapiResponse.Articles))
	assert.Nil(t, err)
}

func TestGetEverythingApple(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		exampleResponse := Response{
			Status:       "ok",
			TotalResults: 2639,
			Articles: []Article{
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
				{
					Source: &Source{
						ID:   "the-verge",
						Name: "The Verge",
					},
					Author:      "Chaim Gartenberg",
					Title:       "Apple TV Plus’ The Banker will hit theaters in March after being delayed following controversy",
					Description: "Apple TV Plus’ biggest movie yet, The Banker, will hit theaters on March 6th after a delay while Apple investigated",
					URL:         "https://www.theverge.com/2020/1/16/21069352/apple-tv-plus-the-banker-theaters-march-6-delay-controversy-accusations",
					URLToImage:  "https://cdn.vox-cdn.com/thumbor/sosnVoYPNN2NksbsdS5yeJM9hWk=/0x25:1280x695/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/19395920/maxresdefault.jpg",
					PublishedAt: "2020-01-16T21:38:23Z",
					Content:     " theatrical release on March 6th, followed by a streaming debut on March 20th...",
				},
			},
		}

		encoder := json.NewEncoder(w)

		if r.URL.Path != "/everything" {
			t.Error("Bad everything path")
		}

		encoder.Encode(exampleResponse)
	}))

	defer sv.Close()
	rawURL, _ := url.Parse(sv.URL)

	testClient := &http.Client{Timeout: time.Minute}
	c := NewsClient{
		baseURL:    rawURL,
		httpClient: testClient,
		apiKey:     "FAKE_API_KEY",
	}

	queryParams := EverythingArgs{
		Q:      "apple",
		From:   time.Now(),
		To:     time.Now().Add(24 * time.Hour),
		SortBy: "popularity",
	}

	newsapiResponse, err := c.Everything(queryParams)

	assert.Nil(t, err)
	assert.Equal(t, 20, len(newsapiResponse.Articles))
	assert.Equal(t, "ok", newsapiResponse.Status)
	assert.Equal(t, 2639, newsapiResponse.TotalResults)
}
