package newsapigo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFailApiKey(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorResponse := Response{
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

	rawURL, _ := url.Parse(sv.URL)
	testClient := &http.Client{Timeout: time.Minute}
	c := NewsClient{
		httpClient: testClient,
		baseURL:    rawURL,
	}

	queryParams := EverythingArgs{}
	_, err := c.Everything(queryParams)

	assert.NotNil(t, err)
}
