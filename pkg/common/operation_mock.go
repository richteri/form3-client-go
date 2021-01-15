package common

import "context"

type OperationMock struct {
	err error
}

func NewOperationMock(err error) *OperationMock {
	return &OperationMock{
		err,
	}
}

func (m *OperationMock) Create(ctx context.Context, path string, model interface{}, res JsonDecoder) error {
	return m.err
}

func (m *OperationMock) GetOne(ctx context.Context, path string, id string, res JsonDecoder) error {
	return m.err
}

func (m *OperationMock) GetPage(ctx context.Context, path string, page *Page, res JsonDecoder) error {
	return m.err
}

func (m *OperationMock) Patch(ctx context.Context, path string, id string, version int, update interface{}, res JsonDecoder) error {
	return m.err
}

func (m *OperationMock) Delete(ctx context.Context, path string, id string, version int) error {
	return m.err
}

func (m *OperationMock) OrganisationID() string {
	return ""
}
