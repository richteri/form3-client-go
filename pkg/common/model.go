package common

import (
	"io"
	"time"
)

// The interface that defines entity-specific response deserialization.
type JsonDecoder interface {
	Decode(response *io.ReadCloser) error
}

type Response struct {
	Links Links `json:"links"`
}

type ErrorResponse struct {
	Code    string `json:"error_code"`
	Message string `json:"error_message"`
}

type Data struct {
	ID             string    `json:"id,omitempty"`
	OrganisationID string    `json:"organisation_id,omitempty"`
	CreatedOn      time.Time `json:"created_on,omitempty"`
	ModifiedOn     time.Time `json:"modified_on,omitempty"`
	Type           string    `json:"type,omitempty"`
	Version        int       `json:"version,omitempty"`
}

type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
	Self  string `json:"self"`
}

type Page struct {
	Number int
	Size   int
}
