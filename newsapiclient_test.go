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
		// Fake Error Response
		// When the API Key is missing or not provided
		errorResponse := Response{
			Status:  "error",
			Code:    "apiKeyMissing",
			Message: "Your API key is missing. Append this to the URL with the apiKey param, or use the x-api-key HTTP header.",
		}

		// If the Api Key is not present in headers, raise an error
		if r.Header.Get("X-Api-Key") == "" {
			t.Error("API Key not provided")
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(errorResponse)
	}))

	defer sv.Close()
	// Parse an URL from the test server
	rawURL, _ := url.Parse(sv.URL)

	// Create a a client with the test URL as baseURL
	testClient := &http.Client{Timeout: time.Minute}
	c := NewsClient{
		httpClient: testClient,
		baseURL:    rawURL,
	}
	// Query params object creation
	queryParams := EverythingArgs{}

	// Make the http request without api key
	_, err := c.Everything(queryParams)

	// The error shouldn't be nil
	assert.NotNil(t, err)
}
