package newsapigo

// Article stores the data in the article. It is an item in the Articles slice.
type Article struct {
	Source      *Source `json:"source,omitempty"`
	Author      string  `json:"author"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	URL         string  `json:"url"`
	URLToImage  string  `json:"urlToImage"`
	PublishedAt string  `json:"publishedAt"`
}

// NewsResponse stores all data received from the request.
type NewsResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Message      string    `json:"message,omitempty"`
	Code         string    `json:"code,omitempty"`
	Articles     []Article `json:"articles,omitempty"`
	Sources      []Source  `json:"sources,omitempty"`
}

// Source stores the data of the source that published the article.
type Source struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Category    string `json:"category,omitempty"`
	Language    string `json:"language,omitempty"`
	Country     string `json:"country,omitempty"`
}

// URLProvider stores the urls of the endpoints
type URLProvider struct {
	TopHeadlines string
	Everything   string
	Sources      string
}

// func FetchNews(country, apiKey string) NewsResponse {
// 	client := &http.Client{Timeout: time.Second * 2}
// 	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&apiKey=%s", country, apiKey)
// 	req, err := http.NewRequest("GET", url, nil)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal("Do: ", err)
// 	}
// 	defer resp.Body.Close()
// 	var response NewsResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
// 		log.Println(err)
// 	}
// 	return response
// }
