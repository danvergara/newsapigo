package newsapigo

import "testing"

var payload map[string]string
var apiKey = "49767a30a6d342f184b535120b6472e0"

func TestGetTopHeadlines(t *testing.T) {
	payload = make(map[string]string)
	payload["country"] = "mx"
	response := GetTopHeadlines(payload, apiKey)
	if response.TotalResults != 20 {
		t.Errorf("Numbers of articles in response, got %d, expected: %d", response.TotalResults, 20)
	}
}

func TestGetSources(t *testing.T) {
	payload = make(map[string]string)
	payload["country"] = "us"
	payload["category"] = "general"
	response := GetSources(payload, apiKey)
	if len(response.Sources) == 0 {
		t.Errorf("Length of the sources, got: %d", len(response.Sources))
	}
}

func TestGetEverything(t *testing.T) {
	payload = make(map[string]string)
	payload["q"] = "bitcoin"
	payload["pageSize"] = "5"
	response := GetEverything(payload, apiKey)
	if len(response.Articles) != 5 {
		t.Errorf("Length of the articles slice, got: %d, expected: %d, message: %s", len(response.Articles), 5, response.Message)
	}
}
