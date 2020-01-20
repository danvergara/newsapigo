package newsapigo

import (
	"net/http"
	"net/url"
)

func buildRquest(url *url.URL) (*http.Request, error) {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return req, err
	}
	return req, nil
}

func errorNewsResponse(msg, code string) NewsResponse {
	return NewsResponse{
		Status:  "error",
		Message: msg,
		Code:    code,
	}
}
