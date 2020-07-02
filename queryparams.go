package newsapigo

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

// EverythingArgs is a convenient way to build all the query params used to tweak the request for the /everything endpoint
type EverythingArgs struct {
	// Keywords or phrases to search for in the article title and body.
	// Advanced search is supported here:
	// *  Surround phrases with quotes (") for exact match.
	// * Prepend words or phrases that must appear with a + symbol. Eg: +bitcoin
	// * Prepend words that must not appear with a - symbol. Eg: -bitcoin
	// * Alternatively you can use the AND / OR / NOT keywords, and optionally group these with parenthesis. Eg: crypto AND (ethereum OR litecoin) NOT bitcoin.
	// The complete value for q must be URL-encoded.
	Q string
	// Keywords or phrases to search for in the article title only.
	// Advanced search is supported here:
	// * Surround phrases with quotes (") for exact match.
	// * Prepend words or phrases that must appear with a + symbol. Eg: +bitcoin
	// * Prepend words that must not appear with a - symbol. Eg: -bitcoin
	// * Alternatively you can use the AND / OR / NOT keywords, and optionally group these with parenthesis. Eg: crypto AND (ethereum OR litecoin) NOT bitcoin.
	// The complete value for qInTitle must be URL-encoded.
	QInTitle string
	// A comma-seperated string of identifiers (maximum 20) for the news sources or blogs you want headlines from.
	// Use the /sources endpoint to locate these programmatically or look at
	Sources []string
	// A comma-seperated string of domains (eg bbc.co.uk, techcrunch.com, engadget.com) to restrict the search to.
	Domains []string
	// A comma-seperated string of domains (eg bbc.co.uk, techcrunch.com, engadget.com) to remove from the results.
	ExcludeDomains []string
	// A date and optional time for the oldest article allowed.
	// This should be in ISO 8601 format (e.g. 2020-07-01 or 2020-07-01T04:35:51)
	// Default: the oldest according to your plan.
	From time.Time
	// A date and optional time for the newest article allowed.
	// This should be in ISO 8601 format (e.g. 2020-07-01 or 2020-07-01T04:35:51)
	// Default: the newest according to your plan.
	To time.Time
	// The 2-letter ISO-639-1 code of the language you want to get headlines for.
	// Possible options: ardeenesfrheitnlnoptruseudzh. Default: all languages returned.
	Language string
	// The order to sort the articles in. Possible options: relevancy, popularity, publishedAt.
	// relevancy = articles more closely related to q come first.
	// popularity = articles from popular sources and publishers come first.
	// publishedAt = newest articles come first.
	// Default: publishedAt
	SortBy string
	// The number of results to return per page. 20 is the default, 100 is the maximum.
	PageSize int
	// Use this to page through the results.
	Page int
}

// QueryParams adds all query params int a url.Values struct given an EverythingArgs struct
func (args EverythingArgs) QueryParams() url.Values {
	// creates an empty instace of url.Values
	q := make(url.Values)

	if args.Q != "" {
		q.Add("q", args.Q)
	}

	if args.QInTitle != "" {
		q.Add("qInTitle", args.QInTitle)
	}

	if len(args.Sources) > 0 {
		q.Add("sources", strings.Join(args.Sources, ","))
	}

	if len(args.Domains) > 0 {
		q.Add("domains", strings.Join(args.Domains, ","))
	}

	if len(args.ExcludeDomains) > 0 {
		q.Add("excludeDomains", strings.Join(args.ExcludeDomains, ","))
	}

	if !args.From.IsZero() {
		q.Add("from", args.From.Format(time.RFC3339))
	}

	if !args.To.IsZero() {
		q.Add("to", args.To.Format(time.RFC3339))
	}

	if args.Language != "" {
		q.Add("language", args.Language)
	}

	if args.SortBy != "" {
		q.Add("sortBy", args.SortBy)
	}

	if args.PageSize != 0 {
		q.Add("pageSize", strconv.Itoa(args.PageSize))
	}

	if args.Page != 0 {
		q.Add("page", strconv.Itoa(args.Page))
	}

	return q
}

// TopHeadlinesArgs is a convenient way to build all the query params used to tweak the request for the /top-headlines endpoint
type TopHeadlinesArgs struct {
	// The 2-letter ISO 3166-1 code of the country you want to get headlines for.
	// Possible options:
	// ae ar at au be bg br ca ch cn co cu cz de eg fr gb gr hk hu id ie il in it jp kr lt
	// lv ma mx my ng nl no nz ph pl pt ro rs ru sa se sg si sk th tr tw ua us ve za .
	// Note: you can't mix this param with the sources param.
	Country string
	// The category you want to get headlines for.
	// Possible options: business entertainment general health science sports technology.
	// Note: you can't mix this param with the sources param.
	Category string
	// A comma-seperated string of identifiers for the news sources or blogs you want headlines from.
	// Use the /sources endpoint to locate these programmatically or look at the sources index.
	// Note: you can't mix this param with the country or category params.
	Sources []string
	// Keywords or a phrase to search for
	Q string
	// The number of results to return per page (request). 20 is the default, 100 is the maximum.
	PageSize int
	// Use this to page through the results if the total results found is greater than the page size.
	Page int
}

// QueryParams adds all query params int a url.Values struct given an EverythingArgs struct
// Returns a instances of url.Values with all passed query params
func (args TopHeadlinesArgs) QueryParams() url.Values {
	// creates an empty instace of url.Values
	q := make(url.Values)

	if args.Country != "" {
		q.Add("country", args.Country)
	}

	if args.Category != "" {
		q.Add("category", args.Category)
	}

	if len(args.Sources) > 0 {
		q.Add("sources", strings.Join(args.Sources, ","))
	}

	if args.Q != "" {
		q.Add("q", args.Q)
	}

	if args.PageSize != 0 {
		q.Add("pageSize", strconv.Itoa(args.PageSize))
	}

	if args.Page != 0 {
		q.Add("page", strconv.Itoa(args.Page))
	}

	return q
}

// SourcesArgs is a convenient way to build all the query params used to tweak the request for the /sources endpoint
type SourcesArgs struct {
	// Find sources that display news of this category.
	// Possible options: business entertainment general health science sports technology.
	// Default: all categories.
	Category string
	// Find sources that display news in a specific language.
	// Possible options: ar de en es fr he it nl no pt ru se ud zh.
	// Default: all languages.
	Language string
	// Find sources that display news in a specific country.
	// Possible options:
	// ae ar at au be bg br ca ch cn co cu cz de
	// eg fr gb gr hk hu id ie il in it jp kr lt
	// lv ma mx my ng nl no nz ph pl pt ro rs ru
	// sa se sg si sk th tr tw ua us ve za.
	// Default: all countries.
	Country string
}

// QueryParams adds all query params int a url.Values struct given an SourcesArgs struct
// Returns a instances of url.Values with all passed query params
func (args SourcesArgs) QueryParams() url.Values {
	q := make(url.Values)

	if args.Category != "" {
		q.Add("category", args.Category)
	}

	if args.Language != "" {
		q.Add("language", args.Language)
	}

	if args.Country != "" {
		q.Add("country", args.Country)
	}
	return q
}
