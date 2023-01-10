package geofox

import (
	"net/http"
	"strings"
)

func (a *API) Init() ([]byte, error) {
	url := "https://gti.geofox.de/gti/public/init"
	payload := strings.NewReader(`{}`)

	return a.makeRequest(http.MethodPost, url, payload)
}
