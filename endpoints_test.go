package newsapigo

import "testing"

var payload map[string]string

// Your api key goes here
var apiKey = "XXXXXXXXXXXXXXXXXXXXXXX"
var keyTest = GetAPIKey()

func TestGetApiKey(t *testing.T) {
	key1 := GetAPIKey()

	if key1 == nil {
		t.Errorf("Expected pointer to ApiKey after calling GetApiKey, not nil")
	}
	key1.SetKey("Hey")

	key2 := GetAPIKey()

	if key1.Key != key2.Key {
		t.Errorf("The instances are different")
	}
}

func TestGetTopHeadlines(t *testing.T) {
	keyTest.SetKey(apiKey)
	payload = make(map[string]string)
	payload["country"] = "mx"
	response := GetTopHeadlines(payload)
	if response.TotalResults != 20 {
		t.Errorf("Numbers of articles in response, got %d, expected: %d", response.TotalResults, 20)
	}
}

func TestGetSources(t *testing.T) {
	keyTest.SetKey(apiKey)
	payload = make(map[string]string)
	payload["country"] = "us"
	payload["category"] = "general"
	response := GetSources(payload)
	if len(response.Sources) == 0 {
		t.Errorf("Length of the sources, got: %d", len(response.Sources))
	}
}

func TestGetEverything(t *testing.T) {
	keyTest.SetKey(apiKey)
	payload = make(map[string]string)
	payload["q"] = "bitcoin"
	payload["pageSize"] = "5"
	response := GetEverything(payload)
	if len(response.Articles) != 5 {
		t.Errorf("Length of the articles slice, got: %d, expected: %d, message: %s", len(response.Articles), 5, response.Message)
	}
}
