package geofox

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI_Init(t *testing.T) {
	// Happy path
	setupTestServer()
	testMux.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		response := InitResponse{
			BeginOfService: "2000-01-01",
			EndOfService:   "2000-01-02",
			ID:             "test-id",
			DataID:         "test-data-id",
			BuildDate:      "2000-01-03",
			BuildTime:      "2000-01-04",
			BuildText:      "test-build",
		}
		response.ReturnCode = "test-return-code"
		responseBytes, err := json.Marshal(response)
		assert.Nil(t, err)
		w.Write(responseBytes) //nolint
	})
	resp, err := testClient.Init(context.Background()) //nolint
	assert.Nil(t, err)
	assert.Equal(t, GeofoxReturnCode("test-return-code"), resp.ReturnCode)
	assert.Equal(t, "2000-01-01", resp.BeginOfService)
	assert.Equal(t, "2000-01-02", resp.EndOfService)
	assert.Equal(t, "test-id", resp.ID)
	assert.Equal(t, "test-data-id", resp.DataID)
	assert.Equal(t, "2000-01-03", resp.BuildDate)
	assert.Equal(t, "2000-01-04", resp.BuildTime)
	assert.Equal(t, "test-build", resp.BuildText)
	teardownTestServer()

	setupTestServer()
	testMux.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	resp, err = testClient.Init(context.Background()) //nolint
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "status code: 404")
	teardownTestServer()
}

func TestAPI_ListStations(t *testing.T) {
	testDataReleaseID := "test-data-release-id"
	// Happy path
	setupTestServer()
	testMux.HandleFunc("/listStations", func(w http.ResponseWriter, r *http.Request) {
		response := LSResponse{
			DataReleaseID: &testDataReleaseID,
			Stations: &[]StationListEntry{
				{}, {}, {},
			},
		}
		responseBytes, err := json.Marshal(response)
		assert.Nil(t, err)
		w.Write(responseBytes) //nolint
	})
	resp, err := testClient.ListStations(context.Background(),
		[]ModificationType{ModTypeMain},
		EPSG4326)
	assert.Nil(t, err)
	assert.Equal(t, &testDataReleaseID, resp.DataReleaseID)
	assert.Equal(t, 3, len(*resp.Stations))
	teardownTestServer()

	// Bad path - error 404
	setupTestServer()
	testMux.HandleFunc("/listStations", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	resp, err = testClient.ListStations(context.Background(),
		[]ModificationType{ModTypeMain},
		EPSG4326)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "status code: 404")
	teardownTestServer()

	// Bad path - unmarshalable response
	setupTestServer()
	testMux.HandleFunc("/listStations", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte{'r', 'a', 'n', 'd', 'o', 'm'}) //nolint
	})
	resp, err = testClient.ListStations(context.Background(),
		[]ModificationType{ModTypeMain},
		EPSG4326)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to marshal body bytes")
	teardownTestServer()
}
