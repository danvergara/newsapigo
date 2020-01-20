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
	Content     string  `json:"contentt"`
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
