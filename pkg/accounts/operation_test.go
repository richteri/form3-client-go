package accounts

import (
	"fmt"
	"github.com/richteri/form3-client-go/pkg/common"
	"testing"
)

func TestAccountOperation_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewAccountOperation(common.NewOperationMock(nil))

		_, err := client.Create(nil, "", AccountAttributes{})
		if err != nil {
			t.Error("should not return error")
		}
	})

	t.Run("error", func(t *testing.T) {
		client := NewAccountOperation(common.NewOperationMock(fmt.Errorf("error")))

		_, err := client.Create(nil, "", AccountAttributes{})
		if err == nil {
			t.Error("should return the error")
		}
	})
}

func TestAccountOperation_GetPage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewAccountOperation(common.NewOperationMock(nil))

		_, err := client.GetPage(nil, nil)
		if err != nil {
			t.Error("should not return error")
		}
	})

	t.Run("error", func(t *testing.T) {
		client := NewAccountOperation(common.NewOperationMock(fmt.Errorf("error")))

		_, err := client.GetPage(nil, nil)
		if err == nil {
			t.Error("should return the error")
		}
	})
}

func TestAccountOperation_GetOne(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewAccountOperation(common.NewOperationMock(nil))

		_, err := client.GetOne(nil, "")
		if err != nil {
			t.Error("should not return error")
		}
	})

	t.Run("error", func(t *testing.T) {
		client := NewAccountOperation(common.NewOperationMock(fmt.Errorf("error")))

		_, err := client.GetOne(nil, "")
		if err == nil {
			t.Error("should return the error")
		}
	})
}

func TestAccountOperation_Patch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewAccountOperation(common.NewOperationMock(nil))

		_, err := client.Patch(nil, "", 0, AccountAttributes{})
		if err != nil {
			t.Error("should not return error")
		}
	})

	t.Run("error", func(t *testing.T) {
		client := NewAccountOperation(common.NewOperationMock(fmt.Errorf("error")))

		_, err := client.Patch(nil, "", 0, AccountAttributes{})
		if err == nil {
			t.Error("should return the error")
		}
	})
}

func TestAccountOperation_Delete(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		client := NewAccountOperation(common.NewOperationMock(fmt.Errorf("error")))

		err := client.Delete(nil, "", 0)
		if err == nil {
			t.Error("should return the error")
		}
	})
}
