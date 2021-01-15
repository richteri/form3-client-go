package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/richteri/form3-client-go/pkg/common"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

type payload struct {
	Property string `json:"property"`
}

func TestHTTPClient_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewTestClient(func(req *http.Request) *http.Response {
			method := "POST"
			if req.Method != method {
				t.Errorf("request method expected: %s, actual: %s", method, req.Method)
			}

			url := "http://localhost:8080/v1/organistaion/accounts"
			if req.URL.String() != url {
				t.Errorf("request url expected %s, actual %s", url, req.URL.String())
			}

			expected := `{"property":"test"}`
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(req.Body)
			actual := buf.String()
			if actual != expected {
				t.Errorf("request body expected %s, actual %s", expected, actual)
			}

			return &http.Response{
				StatusCode: 201,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString("CREATED")),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}
		})

		httpClient := NewHTTPClient(client, BaseURL, "")

		err := httpClient.Create(nil, "organistaion/accounts", &payload{Property: "test"}, nil)

		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("error", func(t *testing.T) {
		client := NewTestClient(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: 400,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString(`{ "error_message": "Bad request", "error_code": "BAD_REQUEST" }`)),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}
		})

		httpClient := NewHTTPClient(client, BaseURL, "")

		err := httpClient.Create(nil, "organistaion/accounts", nil, nil)

		if err == nil {
			t.Errorf("should return error")
			t.Skip()
		}

		wanted := "POST http://localhost:8080/v1/organistaion/accounts HTTP400 - Bad request"
		if err.Error() != wanted {
			t.Errorf("wanted %s, have %s", wanted, err.Error())
		}
	})
}

func TestHTTPClient_GetPage(t *testing.T) {
	t.Run("default page", func(t *testing.T) {
		client := NewTestClient(func(req *http.Request) *http.Response {
			method := "GET"
			if req.Method != method {
				t.Errorf("request method expected: %s, actual: %s", method, req.Method)
			}

			url := "http://localhost:8080/v1/organistaion/accounts?page[number]=0&page[size]=100"
			if req.URL.String() != url {
				t.Errorf("request url expected %s, actual %s", url, req.URL.String())
			}

			return &http.Response{
				StatusCode: 200,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString("OK")),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}
		})

		httpClient := NewHTTPClient(client, BaseURL, "")

		err := httpClient.GetPage(nil, "organistaion/accounts", nil, nil)

		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("custom page", func(t *testing.T) {
		client := NewTestClient(func(req *http.Request) *http.Response {
			method := "GET"
			if req.Method != method {
				t.Errorf("request method expected: %s, actual: %s", method, req.Method)
			}

			url := "http://localhost:8080/v1/organistaion/accounts?page[number]=1&page[size]=2000"
			if req.URL.String() != url {
				t.Errorf("request url expected %s, actual %s", url, req.URL.String())
			}

			return &http.Response{
				StatusCode: 200,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString("OK")),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}
		})

		httpClient := NewHTTPClient(client, BaseURL, "")

		err := httpClient.GetPage(nil, "organistaion/accounts", &common.Page{Size: 2001, Number: 1}, nil)

		if err != nil {
			t.Error(err.Error())
		}
	})
}

func TestHTTPClient_GetOne(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		id := "887a55f2-b5c6-4246-a06b-94e4bdafb3d5"

		client := NewTestClient(func(req *http.Request) *http.Response {
			method := "GET"
			if req.Method != method {
				t.Errorf("request method expected: %s, actual: %s", method, req.Method)
			}

			url := fmt.Sprintf("http://localhost:8080/v1/organistaion/accounts/%s", id)
			if req.URL.String() != url {
				t.Errorf("request url expected %s, actual %s", url, req.URL.String())
			}

			return &http.Response{
				StatusCode: 200,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString("OK")),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}
		})

		httpClient := NewHTTPClient(client, BaseURL, "")

		err := httpClient.GetOne(nil, "organistaion/accounts", id, nil)

		if err != nil {
			t.Error(err.Error())
		}
	})
}

func TestHTTPClient_Patch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		id := "887a55f2-b5c6-4246-a06b-94e4bdafb3d5"
		version := 42

		client := NewTestClient(func(req *http.Request) *http.Response {
			method := "PATCH"
			if req.Method != method {
				t.Errorf("request method expected: %s, actual: %s", method, req.Method)
			}

			url := fmt.Sprintf("http://localhost:8080/v1/organistaion/accounts/%s?version=%d", id, version)
			if req.URL.String() != url {
				t.Errorf("request url expected %s, actual %s", url, req.URL.String())
			}

			expected := `{"property":"test"}`
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(req.Body)
			actual := buf.String()
			if actual != expected {
				t.Errorf("request body expected %s, actual %s", expected, actual)
			}

			return &http.Response{
				StatusCode: 200,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString("OK")),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}
		})

		httpClient := NewHTTPClient(client, BaseURL, "")

		err := httpClient.Patch(nil, "organistaion/accounts", id, version, &payload{Property: "test"}, nil)

		if err != nil {
			t.Error(err.Error())
		}
	})
}

func TestHTTPClient_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		id := "887a55f2-b5c6-4246-a06b-94e4bdafb3d5"
		version := 42

		client := NewTestClient(func(req *http.Request) *http.Response {
			method := "DELETE"
			if req.Method != method {
				t.Errorf("request method expected: %s, actual: %s", method, req.Method)
			}

			url := fmt.Sprintf("http://localhost:8080/v1/organistaion/accounts/%s?version=%d", id, version)
			if req.URL.String() != url {
				t.Errorf("request url expected %s, actual %s", url, req.URL.String())
			}

			return &http.Response{
				StatusCode: 204,
				// Send response to be tested
				Body: ioutil.NopCloser(bytes.NewBufferString("")),
				// Must be set to non-nil value or it panics
				Header: make(http.Header),
			}
		})

		httpClient := NewHTTPClient(client, BaseURL, "")

		err := httpClient.Delete(nil, "organistaion/accounts", id, version)

		if err != nil {
			t.Error(err.Error())
		}
	})
}
