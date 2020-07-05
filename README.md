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


First, all what you need is to instantiate an newsapigo client, for that purpose, we implement a convenient function called New:

```go

import (
  "net/url"
  "os"

  "github.com/danvergara/newsapigo"
)

// Instantiates the Client provided by this library
// Don't forget provide your API Key
// You can get one from the official site: https://newsapi.org/register
c := newsapigo.NewClient(os.Getenv("API_KEY"))
```

This client only has three methods (one for each endpoint provided by the API):
TopHeadlines, Everything and Sources.

If you want to customize the request, for convenience, use the corresponding structs for each method. Those structs are converted to url.Vlaues under the hood, using the common public API *QueryParams*.

Each method receives a struct object as shown below:

## Top Headlines

```go
// This corresponds to '/v2/top-headlines'

queryParams := TopHeadlinesArgs{
	Sources: []string{"bbc-news"},
}

response, err := c.TopHeadlines(queryParams)
```

## Everything

```go
// This corresponds to '/v2/everything'

queryParams := EverythingArgs{
	Q: "bitcoin",
}

response, err := c.Everything(queryParams)
```

## Sources

```go
// This corresponds to '/v2/sources'

queryParams := SourcesArgs{
	Language: "eng",
	Country:  "us",
}

response, err := c.Sources(queryParams)
```
For more information on what fields the structs contain, check the file queryparams.go at the root of the project.
