package accounts

import (
	"strings"
	"testing"
)

func TestAccountAttributes_Validate(t *testing.T) {
	validAttributes := AccountAttributes{
		Country:                 "012",
		BaseCurrency:            "012",
		BankID:                  "01234567890",
		BIC:                     "01234567",
		Name:                    [4]string{"01234567", "", "", ""},
		AlternativeNames:        [3]string{"01234567", "", ""},
		AccountClassification:   AccountClassificationBusiness,
		SecondaryIdentification: "01234567",
		Status:                  StatusConfirmed,
	}

	t.Run("success", func(t *testing.T) {
		err := validAttributes.Validate()

		if err != nil {
			t.Error("attributes should be valid")
		}
	})

	t.Run("country", func(t *testing.T) {
		attributes := validAttributes
		attributes.Country = "1234"

		err := attributes.Validate()

		if err == nil || !strings.Contains(err.Error(), "country") {
			t.Error("should be invalid")
		}
	})

	t.Run("base currency", func(t *testing.T) {
		attributes := validAttributes
		attributes.BaseCurrency = "1234"

		err := attributes.Validate()

		if err == nil || !strings.Contains(err.Error(), "base currency") {
			t.Error("should be invalid")
		}
	})

	t.Run("bank ID", func(t *testing.T) {
		attributes := validAttributes
		attributes.BankID = "123456789012"

		err := attributes.Validate()

		if err == nil || !strings.Contains(err.Error(), "bank ID") {
			t.Error("should be invalid")
		}
	})

	t.Run("BIC", func(t *testing.T) {
		attributes := validAttributes
		attributes.BIC = "012345678"

		err := attributes.Validate()

		if err == nil || !strings.Contains(err.Error(), "BIC") {
			t.Error("should be invalid")
		}
	})

	t.Run("name", func(t *testing.T) {
		attributes := validAttributes
		attributes.Name = [4]string{strings.Repeat("x", 141)}

		err := attributes.Validate()

		if err == nil || !strings.Contains(err.Error(), "name line") {
			t.Error("should be invalid")
		}
	})

	t.Run("alternative names", func(t *testing.T) {
		attributes := validAttributes
		attributes.AlternativeNames = [3]string{strings.Repeat("x", 141)}

		err := attributes.Validate()

		if err == nil || !strings.Contains(err.Error(), "alternative names line") {
			t.Error("should be invalid")
		}
	})

	t.Run("attributes classification", func(t *testing.T) {
		attributes := validAttributes
		attributes.AccountClassification = "012345678"

		err := attributes.Validate()

		if err == nil || !strings.Contains(err.Error(), "attributes classification") {
			t.Error("should be invalid")
		}
	})

	t.Run("status", func(t *testing.T) {
		attributes := validAttributes
		attributes.Status = "012345678"

		err := attributes.Validate()

		if err == nil || !strings.Contains(err.Error(), "status") {
			t.Error("should be invalid")
		}
	})
}
