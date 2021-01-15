package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/richteri/form3-client-go/pkg/accounts"
	"github.com/richteri/form3-client-go/pkg/common"
)

const (
	BaseURL    = "http://localhost:8080"
	APIVersion = "v1"
	MimeType   = "application/vnd.api+json; charset=utf-8"
)

// An HTTP-based client that uses composition for operating on different entity types and endpoints.
type HTTPClient struct {
	BaseURL    string
	APIVersion string
	MimeType   string
	Client     *http.Client

	organisationID string

	Accounts *accounts.AccountOperation
}

func (c *HTTPClient) OrganisationID() string {
	return c.organisationID
}

// Builds a client to be used with custom `http.Client` instances.
func NewHTTPClient(httpClient *http.Client, baseURL string, organisationID string) *HTTPClient {
	if baseURL == "" {
		baseURL = BaseURL
	}

	c := &HTTPClient{
		BaseURL:        baseURL,
		APIVersion:     APIVersion,
		MimeType:       MimeType,
		Client:         httpClient,
		organisationID: organisationID,
	}

	c.Accounts = accounts.NewAccountOperation(c)

	return c
}

// Builds a client using the default `http.Client` instance.
func NewDefaultHTTPClient(baseURL string, organisationID string) *HTTPClient {
	return NewHTTPClient(&http.Client{
		Timeout: time.Minute,
	}, baseURL, organisationID)
}

func (c *HTTPClient) Create(ctx context.Context, path string, payload interface{}, jd common.JsonDecoder) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, c.APIVersion, path)

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	if err := c.sendRequest(ctx, req, jd); err != nil {
		return err
	}

	return nil
}

func (c *HTTPClient) GetPage(ctx context.Context, path string, page *common.Page, jd common.JsonDecoder) error {
	number := 0
	// avoid overloading API
	size := 100
	if page != nil {
		number = page.Number
		if page.Size > 2000 {
			size = 2000
		} else {
			size = page.Size
		}
	}

	url := fmt.Sprintf("%s/%s/%s?page[number]=%d&page[size]=%d", c.BaseURL, c.APIVersion, path, number, size)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if err := c.sendRequest(ctx, req, jd); err != nil {
		return err
	}

	return nil
}

func (c *HTTPClient) GetOne(ctx context.Context, path string, id string, jd common.JsonDecoder) error {
	url := fmt.Sprintf("%s/%s/%s/%s", c.BaseURL, c.APIVersion, path, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if err := c.sendRequest(ctx, req, jd); err != nil {
		return err
	}

	return nil
}

func (c *HTTPClient) Patch(ctx context.Context, path string, id string, version int, update interface{}, jd common.JsonDecoder) error {
	body, err := json.Marshal(update)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/%s/%s/%s?version=%d", c.BaseURL, c.APIVersion, path, id, version)

	req, err := http.NewRequest("PATCH", url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	if err := c.sendRequest(ctx, req, jd); err != nil {
		return err
	}

	return nil
}

func (c *HTTPClient) Delete(ctx context.Context, path string, id string, version int) error {
	url := fmt.Sprintf("%s/%s/%s/%s?version=%d", c.BaseURL, c.APIVersion, path, id, version)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	if err := c.sendRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// Implements the generic part of the RESTful request-response communication.
func (c *HTTPClient) sendRequest(ctx context.Context, req *http.Request, jd common.JsonDecoder) error {
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Set("Content-Type", c.MimeType)
	req.Header.Set("Accept", c.MimeType)

	res, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK &&
		res.StatusCode != http.StatusNoContent &&
		res.StatusCode != http.StatusCreated {

		errRes := common.ErrorResponse{}
		err := json.NewDecoder(res.Body).Decode(&errRes)

		prefix := fmt.Sprintf("%s %s HTTP%d -", req.Method, req.URL.String(), res.StatusCode)

		if err != nil {
			return fmt.Errorf("%s unknown error", prefix)
		}

		return fmt.Errorf("%s %s", prefix, errRes.Message)
	}

	if jd != nil {
		err := jd.Decode(&res.Body)
		if err != nil {
			return err
		}
	}

	return nil
}
