package newsapigo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFaileApiKey(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorResponse := NewsResponse{
			Status:  "error",
			Code:    "apiKeyMissing",
			Message: "Your API key is missing. Append this to the URL with the apiKey param, or use the x-api-key HTTP header.",
		}

		encoder := json.NewEncoder(w)

		if r.Header.Get("X-Api-Key") == "" {
			t.Error("API Key not provided")
		}

		encoder.Encode(errorResponse)
	}))

	defer sv.Close()

	c := NewsClient{
		BaseURL: sv.URL,
	}

	queryParams := url.Values{}
	_, err := c.GetEverything(queryParams)

	assert.NotNil(t, err)
}
