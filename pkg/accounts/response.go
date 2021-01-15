package accounts

import (
	"encoding/json"
	"io"

	"github.com/richteri/form3-client-go/pkg/common"
)

type AccountResponse struct {
	common.Response
	Data Account `json:"data"`
}

func (response *AccountResponse) Decode(body *io.ReadCloser) error {
	err := json.NewDecoder(*body).Decode(&response)
	return err
}

type AccountListResponse struct {
	common.Response
	Data []Account `json:"data"`
}

func (response *AccountListResponse) Decode(body *io.ReadCloser) error {
	err := json.NewDecoder(*body).Decode(&response)
	return err
}
