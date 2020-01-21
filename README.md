# News API package for Go

> A Golang client for the News API *https://newsapi.org*

This library provides a convenient way to consume the News API service easily.
For more info about the service, you can check the documentation out.
*https://newsapi.org/docs*

##### Provided under MIT License by Daniel Vergara.

## Installation

```shell
go get github.com/danvergara/newsapigo
```

## Usage

After the installation, import the client into your project:

```go
import (
  "github.com/danvergara/newsapigo"
)
```
## Initialization


```go

import (
  "net/url"
  "os"
)

// Instantiate the Client provided by this library
// Don't forget provide you API Key
// You can get one from the official site: https://newsapi.org/register
c := newsapigo.NewsClient{
        APIKey: os.Getenv("API_KEY")
}

```

This client only has three methods (one for each endpoint provided by the API):
GetTopHeadlines, GetEverything and GetSources.

If you want customize the request, for convenience, create a Values object (from the standard library "net/url") and add the necessary query params.

Each method receives a Values object as shown below:

## Top Headlines

```go
  params := url.Values{}
  params.Add("country", "us")
  response, err := c.GetTopHeadlines(params)
```

## Everything

```go
  params := url.Values{}
  params.Add("q", "bitcoin")
  response, err := c.GetEverything(params)
```

## Sources

```go
  params := url.Values{}
  params.Add("category", "general")
  response, err := c.GetSources(params)
```

Each method returns an instance of the NewsResponse struct:

```go
type NewsResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Message      string    `json:"message,omitempty"`
	Code         string    `json:"code,omitempty"`
	Articles     []Article `json:"articles,omitempty"`
	Sources      []Source  `json:"sources,omitempty"`
}
```
These structs should match with the response from News API:

```go
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
```

```go
type Source struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Category    string `json:"category,omitempty"`
	Language    string `json:"language,omitempty"`
	Country     string `json:"country,omitempty"`
}
```
