package accounts

import (
	"encoding/json"
	"io"

	"github.com/richteri/form3-client-go/pkg/common"
)

// Wraps a single account into a response structure.
type AccountResponse struct {
	common.Response
	Data Account `json:"data"`
}

// Implements a single account response specific deserialization.
func (response *AccountResponse) Decode(body *io.ReadCloser) error {
	err := json.NewDecoder(*body).Decode(&response)
	return err
}

// Wraps a list of accounts into a response structure.
type AccountListResponse struct {
	common.Response
	Data []Account `json:"data"`
}

// Implements deserialization for a response containing a list of accounts.
func (response *AccountListResponse) Decode(body *io.ReadCloser) error {
	err := json.NewDecoder(*body).Decode(&response)
	return err
}
