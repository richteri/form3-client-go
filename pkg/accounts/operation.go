package accounts

import (
	"context"

	"github.com/richteri/form3-client-go/pkg/common"
)

const (
	defaultPath = "organisation/accounts"
)

type AccountOperation struct {
	path   string
	client common.Operation
}

// Builds a CRUD operator that targets the account endpoint.
func NewAccountOperation(c common.Operation) *AccountOperation {
	aop := &AccountOperation{
		path:   defaultPath,
		client: c,
	}

	return aop
}

func (op *AccountOperation) Create(ctx context.Context, id string, model AccountAttributes) (*AccountResponse, error) {
	res := &AccountResponse{}

	payload := NewAccountPayload(op.client.OrganisationID(), id, model)

	err := op.client.Create(ctx, op.path, payload, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (op *AccountOperation) GetPage(ctx context.Context, page *common.Page) (*AccountListResponse, error) {
	res := &AccountListResponse{}

	err := op.client.GetPage(ctx, op.path, page, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (op *AccountOperation) GetOne(ctx context.Context, id string) (*AccountResponse, error) {
	res := &AccountResponse{}

	err := op.client.GetOne(ctx, op.path, id, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (op *AccountOperation) Patch(ctx context.Context, id string, version int, update AccountAttributes) (*AccountResponse, error) {
	res := &AccountResponse{}

	payload := NewAccountPayload(op.client.OrganisationID(), id, update)

	err := op.client.Patch(ctx, op.path, id, version, payload, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (op *AccountOperation) Delete(ctx context.Context, id string, version int) error {
	err := op.client.Delete(ctx, op.path, id, version)

	return err
}
