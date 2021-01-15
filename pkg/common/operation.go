package common

import "context"

type Operation interface {
	Create(ctx context.Context, path string, payload interface{}, res JsonDecoder) error
	GetOne(ctx context.Context, path string, id string, res JsonDecoder) error
	GetPage(ctx context.Context, path string, page *Page, res JsonDecoder) error
	Patch(ctx context.Context, path string, id string, version int, payload interface{}, res JsonDecoder) error
	Delete(ctx context.Context, path string, id string, version int) error
	OrganisationID() string
}
