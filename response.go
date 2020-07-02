package newsapigo

import "fmt"

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

// Response stores all data received from the request.
type Response struct {
	Status       string    `json:"status"`
	Code         string    `json:"code"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles,omitempty"`
	Sources      []Source  `json:"sources,omitempty"`
	Message      string    `json:"message"`
	StatusCode   int       `json:"status_code"`
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

func newError(msg, code string) Response {
	return Response{
		Status:  "error",
		Message: msg,
		Code:    code,
	}
}

// ErrorResponse implements the Error interfaces so that it can be returned as an error
type ErrorResponse struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (err *ErrorResponse) Error() string {
	return fmt.Sprintf("%s (%s) API error: %s", err.Status, err.Code, err.Message)
}
