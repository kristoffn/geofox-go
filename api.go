package geofox

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	AuthTypeHmacSHA1 = "HmacSHA1"
)

type API struct {
	Username string
	Password string
	AuthType string
}

func New(username, password, authType string) *API {
	return &API{
		Username: username,
		Password: password,
		AuthType: authType,
	}
}

func createSHA1HmacSignature(input []byte, key []byte) string {
	h := hmac.New(sha1.New, key)
	h.Write(input)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func getBodyBytes(req *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}

func (a *API) makeRequest(method, url string, payload *strings.Reader) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := getBodyBytes(req)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("geofox-auth-signature", createSHA1HmacSignature(bodyBytes, []byte(a.Password)))
	req.Header.Add("geofox-auth-user", a.Username)
	req.Header.Add("geofox-auth-type", a.AuthType)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Encoding", "gzip, deflate")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
