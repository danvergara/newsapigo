package newsapigo

import "fmt"

// Article stores the data in the article. It is an item in the Articles slice.
type Article struct {
	Source *Source `json:"source,omitempty"`
	// The author of the article
	Author string `json:"author"`
	// The headline or title of the article
	Title string `json:"title"`
	// A description or snippet from the article.
	Description string `json:"description"`
	// The direct URL to the article
	URL string `json:"url"`
	// The URL to a relevant image for the article
	URLToImage string `json:"urlToImage"`
	// The date and time that the article was published, in UTC (+000)
	PublishedAt string `json:"publishedAt"`
	// The unformatted content of the article, where available. This is truncated to 200 chars
	Content string `json:"contentt"`
}

// Response stores all data received from the request.
type Response struct {
	// If the request was successful or not. Options: ok, error
	// In the case of error a code and message property will be populated
	Status string `json:"status"`
	// A short code identifying the type of error returned, if any
	Code string `json:"code"`
	// The total number of results available for your request
	TotalResults int `json:"totalResults"`
	// The results of the request
	Articles []Article `json:"articles,omitempty"`
	// The results of the request
	Sources []Source `json:"sources,omitempty"`
	// The error message if any
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// Source stores the data of the source that published the article.
type Source struct {
	// The identifier of the news source. You can use this with our other endpoints
	ID string `json:"id,omitempty"`
	// The name of the news source
	Name string `json:"name"`
	// A description of the news source
	Description string `json:"description,omitempty"`
	// The URL of the homepage
	URL string `json:"url,omitempty"`
	// The type of news to expect from this news source
	Category string `json:"category,omitempty"`
	// The language that this news source writes in
	Language string `json:"language,omitempty"`
	// The country this news source is based in (and primarily writes about)
	Country string `json:"country,omitempty"`
}

func newError(msg, code string) Response {
	return Response{
		Status:  "error",
		Message: msg,
		Code:    code,
	}
}

// ErrorResponse implements the Error interfaces so that it can be returned as an error
type ErrorResponse struct {
	// A short code identifying the type of error returned
	Code string `json:"code"`
	// If the request was successful or not. Options: ok, error.
	// In the case of ok, the below two properties will not be present
	Status string `json:"status"`
	// A fuller description of the error, usually including how to fix it
	Message string `json:"message"`
}

func (err *ErrorResponse) Error() string {
	return fmt.Sprintf("%s (%s) API error: %s", err.Status, err.Code, err.Message)
}
