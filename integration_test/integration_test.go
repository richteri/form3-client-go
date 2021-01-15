package integration

import (
	"github.com/richteri/form3-client-go/pkg/common"
	"os"
	"testing"

	"github.com/richteri/form3-client-go/pkg/accounts"
	"github.com/richteri/form3-client-go/pkg/client"
)

const (
	organisationID = "dabe0971-b3a8-465f-aa6c-30bf8ba98c89"
	account1ID     = "aab1f5b8-5334-47df-b334-9568ec59ec32"
	account2ID     = "e36ceed3-db72-4340-a38f-324de1c38565"
	account3ID     = "748396c9-a253-4b1e-8175-a6e34ed7cbdc"
)

func TestMain(m *testing.M) {
	baseURL := os.Getenv("API_ADDR")
	setup(baseURL)
	code := m.Run()
	shutdown(baseURL)
	os.Exit(code)
}

func TestAccount_Create(t *testing.T) {
	baseURL := os.Getenv("API_ADDR")

	c := client.NewDefaultHTTPClient(baseURL, organisationID)

	attributes := *accounts.NewAccountAttributesBuilder().
		Name([4]string{"Istvan Richter"}).
		AccountNumber("01234567891").
		BankID("400302").
		BankIDCode("GBDSC").
		BaseCurrency("GBP").
		BIC("NWBKGB43").
		Country("GB").
		CustomerID("236").
		IBAN("GB28NWBK40030212764205").
		Build()

	t.Run("success", func(t *testing.T) {
		id := "9a968c54-3aef-4540-a2c0-f76e9477a018"
		res, err := c.Accounts.Create(nil, id, attributes)
		if err != nil {
			t.Error(err)
		} else if res.Data.ID != id {
			// TODO Check if all properties were sent and read correctly.
			t.Errorf("wanted %s, got %s", res.Data.ID, id)
		}
	})

	t.Run("conflict", func(t *testing.T) {
		_, err := c.Accounts.Create(nil, account1ID, attributes)
		if err == nil {
			t.Error("should handle conflict")
			return
		}
	})
}

func TestAccount_GetPage(t *testing.T) {
	baseURL := os.Getenv("API_ADDR")

	c := client.NewDefaultHTTPClient(baseURL, organisationID)

	t.Run("without page", func(t *testing.T) {
		res, err := c.Accounts.GetPage(nil, nil)
		if err != nil {
			t.Error(err)
		} else {
			if len(res.Data) < 3 {
				t.Error("invalid page size")
			}

			if !contains(res.Data, account1ID) ||
				!contains(res.Data, account2ID) ||
				!contains(res.Data, account3ID) {
				t.Error("missing account")
			}

			// TODO Check if structs were mapped correctly.
		}
	})

	t.Run("with page", func(t *testing.T) {
		res, err := c.Accounts.GetPage(nil, &common.Page{Number: 0, Size: 1})
		if err != nil {
			t.Error(err)
		} else {
			if len(res.Data) != 1 {
				t.Error("invalid page size")
			}
		}
	})
}

func TestAccount_GetOne(t *testing.T) {
	baseURL := os.Getenv("API_ADDR")

	c := client.NewDefaultHTTPClient(baseURL, organisationID)

	t.Run("success", func(t *testing.T) {
		res, err := c.Accounts.GetOne(nil, account1ID)
		if err != nil {
			t.Error(err)
		} else {
			if res.Data.ID != account1ID {
				t.Error("invalid result")
			}

			// TODO Check if structs were mapped correctly.
		}
	})

	t.Run("non-existent", func(t *testing.T) {
		_, err := c.Accounts.GetOne(nil, "3d1358d5-d54c-40d9-9cdd-5d8f5d990622")
		if err == nil {
			t.Error("invalid result")
		}
	})
}

func TestAccount_Delete(t *testing.T) {
	baseURL := os.Getenv("API_ADDR")

	c := client.NewDefaultHTTPClient(baseURL, organisationID)

	t.Run("success", func(t *testing.T) {
		err := c.Accounts.Delete(nil, account1ID, 0)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("non-existent", func(t *testing.T) {
		t.Skip("mock api returns 204 for missing IDs")

		err := c.Accounts.Delete(nil, "a5099a61-365b-4bad-9568-4742a172fe10", 0)
		if err == nil {
			t.Error("invalid result")
		}
	})
}

func contains(s []accounts.Account, id string) bool {
	for _, v := range s {
		if v.ID == id {
			return true
		}
	}

	return false
}
