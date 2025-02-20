package geofox

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testMux    *http.ServeMux
	testServer *httptest.Server
	testClient *API
)

func setupTestServer() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	testClient, _ = New("testuser", "testpassword")
	testClient.BaseURL = testServer.URL
}

func teardownTestServer() {
	testServer.Close()
}

func TestAPI_DefaultHeaders(t *testing.T) {
	// Default headers
	setupTestServer()
	testMux.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, http.MethodPost, r.Method,
			"Expected method 'POST', got %s",
			r.Method)

		assert.Equal(t, "application/json", r.Header.Get("Accept"),
			"Expected 'application/json' in geofox-auth-user header, got %s",
			r.Header.Get("Accept"))

		assert.Equal(t, "testuser", r.Header.Get("geofox-auth-user"),
			"Expected 'testuser' in geofox-auth-user header, got %s",
			r.Header.Get("geofox-auth-user"))

		assert.Equal(t, "xuEmPipFxJbEEnmHlGOAqaquhoE=", r.Header.Get("geofox-auth-signature"),
			"Expected 'signature' in geofox-auth-signature header, got %s",
			r.Header.Get("geofox-auth-user"))

		assert.Equal(t, AuthTypeHmacSHA1, r.Header.Get("geofox-auth-type"),
			"Expected '%s' in geofox-auth-type header, got %s", AuthTypeHmacSHA1,
			r.Header.Get("geofox-auth-type"))

		assert.Equal(t, "application/json", r.Header.Get("Content-Type"),
			"Expected 'application/json' in Content-Type header, got %s",
			r.Header.Get("Content-Type"))

		assert.Equal(t, "gzip, deflate", r.Header.Get("Accept-Encoding"),
			"Expected 'gzip, deflate' in Accept-Encoding header, got %s",
			r.Header.Get("Accept-Encoding"))

		assert.Equal(t, "web", r.Header.Get("X-Platform"),
			"Expected 'web' in X-Platform header, got %s",
			r.Header.Get("X-Platform"))

	})
	testClient.Init(context.Background()) //nolint
	teardownTestServer()
}

func TestNew(t *testing.T) {
	api, err := New("testuser", "testpassword")
	assert.Nil(t, err)
	assert.NotNil(t, api)
	assert.Equal(t, "testuser", api.Username)
	assert.Equal(t, "testpassword", api.Password)
	assert.Equal(t, AuthTypeHmacSHA1, api.AuthType)
	assert.Equal(t, false, api.debug)
}

func TestCreateSignature(t *testing.T) {
	cases := []struct {
		name              string
		inputMessage      []byte
		inputKey          []byte
		expectedSignature string
	}{
		{
			name:              "empty message - empty key",
			inputMessage:      []byte{},
			inputKey:          []byte{},
			expectedSignature: "+9sdGxiqbAgyS31ktx+3Y3BpDh0=",
		},
		{
			name:              "nil message - empty key",
			inputMessage:      nil,
			inputKey:          []byte{},
			expectedSignature: "+9sdGxiqbAgyS31ktx+3Y3BpDh0=",
		},
		{
			name:              "empty message - nil key",
			inputMessage:      []byte{},
			inputKey:          nil,
			expectedSignature: "+9sdGxiqbAgyS31ktx+3Y3BpDh0=",
		},
		{
			name:              "apple message - empty key",
			inputMessage:      []byte("apple"),
			inputKey:          []byte{},
			expectedSignature: "owHgjIirOQbvpNrkZ/PMomgjgno=",
		},
		{
			name:              "apple message - nil key",
			inputMessage:      []byte("apple"),
			inputKey:          nil,
			expectedSignature: "owHgjIirOQbvpNrkZ/PMomgjgno=",
		},
		{
			name:              "empty message - apple key",
			inputMessage:      []byte{},
			inputKey:          []byte("apple"),
			expectedSignature: "OlIRfD63z/DxE8x/5DT5e41ksJs=",
		},
		{
			name:              "nil message - apple key",
			inputMessage:      nil,
			inputKey:          []byte("apple"),
			expectedSignature: "OlIRfD63z/DxE8x/5DT5e41ksJs=",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			res := createSHA1HmacSignature(tt.inputMessage, tt.inputKey)
			assert.Equal(t, tt.expectedSignature, res)
		})
	}
}
