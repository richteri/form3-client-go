package accounts

import (
	"fmt"
	"github.com/richteri/form3-client-go/pkg/common"
)

type AccountClassification string
type Status string

const (
	Type = "accounts"

	AccountClassificationPersonal AccountClassification = "Personal"
	AccountClassificationBusiness AccountClassification = "Business"

	StatusPending   Status = "pending"
	StatusConfirmed Status = "confirmed"
	StatusFailed    Status = "failed"
)

type AccountAttributes struct {
	Country                 string                `json:"country,omitempty"`
	BaseCurrency            string                `json:"base_currency,omitempty"`
	BankID                  string                `json:"bank_id,omitempty"`
	BankIDCode              string                `json:"bank_id_code,omitempty"`
	AccountNumber           string                `json:"account_number,omitempty"`
	BIC                     string                `json:"bic,omitempty"`
	IBAN                    string                `json:"iban,omitempty"`
	CustomerID              string                `json:"customer_id,omitempty"`
	Name                    [4]string             `json:"name"`
	AlternativeNames        [3]string             `json:"alternative_names"`
	AccountClassification   AccountClassification `json:"account_classification,omitempty"`
	JointAccount            bool                  `json:"joint_account,omitempty"`
	AccountMatchingOptOut   bool                  `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification string                `json:"secondary_identification,omitempty"`
	Switched                bool                  `json:"switched,omitempty"`
	Status                  Status                `json:"status,omitempty"`
}

type Account struct {
	common.Data
	Attributes AccountAttributes `json:"attributes"`
}

type AccountPayload struct {
	Data Account `json:"data"`
	common.Links
}

func NewAccountAttributes() *AccountAttributes {
	return &AccountAttributes{}
}

func NewAccount(organistaionID string, id string, attributes AccountAttributes) *Account {
	return &Account{
		Data: common.Data{
			OrganisationID: organistaionID,
			ID:             id,
			Type:           Type,
		},
		Attributes: attributes,
	}
}

func NewAccountPayload(organistaionID string, id string, attributes AccountAttributes) *AccountPayload {
	return &AccountPayload{
		Data: *NewAccount(organistaionID, id, attributes),
	}
}

type AccountAttributesBuilder struct {
	attributes *AccountAttributes
}

func NewAccountAttributesBuilder() *AccountAttributesBuilder {
	attributes := NewAccountAttributes()
	b := &AccountAttributesBuilder{attributes: attributes}
	return b
}

func (b *AccountAttributesBuilder) Country(country string) *AccountAttributesBuilder {
	b.attributes.Country = country
	return b
}

func (b *AccountAttributesBuilder) BaseCurrency(baseCurrency string) *AccountAttributesBuilder {
	b.attributes.BaseCurrency = baseCurrency
	return b
}

func (b *AccountAttributesBuilder) BankID(bankID string) *AccountAttributesBuilder {
	b.attributes.BankID = bankID
	return b
}

func (b *AccountAttributesBuilder) BankIDCode(bankIDCode string) *AccountAttributesBuilder {
	b.attributes.BankIDCode = bankIDCode
	return b
}

func (b *AccountAttributesBuilder) AccountNumber(accountNumber string) *AccountAttributesBuilder {
	b.attributes.AccountNumber = accountNumber
	return b
}

func (b *AccountAttributesBuilder) BIC(bic string) *AccountAttributesBuilder {
	b.attributes.BIC = bic
	return b
}

func (b *AccountAttributesBuilder) IBAN(iban string) *AccountAttributesBuilder {
	b.attributes.IBAN = iban
	return b
}

func (b *AccountAttributesBuilder) CustomerID(customerID string) *AccountAttributesBuilder {
	b.attributes.CustomerID = customerID
	return b
}

func (b *AccountAttributesBuilder) Name(name [4]string) *AccountAttributesBuilder {
	b.attributes.Name = name
	return b
}

func (b *AccountAttributesBuilder) AlternativeNames(alternativeNames [3]string) *AccountAttributesBuilder {
	b.attributes.AlternativeNames = alternativeNames
	return b
}

func (b *AccountAttributesBuilder) AccountClassification(accountClassification AccountClassification) *AccountAttributesBuilder {
	b.attributes.AccountClassification = accountClassification
	return b
}

func (b *AccountAttributesBuilder) JointAccount(jointAccount bool) *AccountAttributesBuilder {
	b.attributes.JointAccount = jointAccount
	return b
}

func (b *AccountAttributesBuilder) AccountMatchingOptOut(accountMatchingOptOut bool) *AccountAttributesBuilder {
	b.attributes.AccountMatchingOptOut = accountMatchingOptOut
	return b
}

func (b *AccountAttributesBuilder) SecondaryIdentification(secondaryIdentification string) *AccountAttributesBuilder {
	b.attributes.SecondaryIdentification = secondaryIdentification
	return b
}

func (b *AccountAttributesBuilder) Switched(switched bool) *AccountAttributesBuilder {
	b.attributes.Switched = switched
	return b
}

func (b *AccountAttributesBuilder) Status(status Status) *AccountAttributesBuilder {
	b.attributes.Status = status
	return b
}

func (b *AccountAttributesBuilder) BuildWithValidation() (*AccountAttributes, error) {
	err := b.attributes.Validate()
	if err != nil {
		return nil, err
	}
	return b.attributes, nil
}

func (b *AccountAttributesBuilder) Build() *AccountAttributes {
	return b.attributes
}

func (attributes *AccountAttributes) Validate() error {
	if len(attributes.Country) > 3 {
		return fmt.Errorf("country code must use ISO 3166-1 format")
	}
	if len(attributes.BaseCurrency) > 3 {
		return fmt.Errorf("base currency code must use ISO 4217 format")
	}
	if len(attributes.BankID) > 11 {
		return fmt.Errorf("bank ID maximum length is 11 characters")
	}
	if len(attributes.BIC) != 8 && len(attributes.BIC) != 11 {
		return fmt.Errorf("BIC length must be 8 or 11 characters")
	}
	for _, name := range attributes.Name {
		if len(name) > 140 {
			return fmt.Errorf("name line maximum length is 140 characters")
		}
	}
	for _, name := range attributes.AlternativeNames {
		if len(name) > 140 {
			return fmt.Errorf("alternative names line maximum length is 140 characters")
		}
	}
	if attributes.AccountClassification != "" &&
		attributes.AccountClassification != AccountClassificationPersonal &&
		attributes.AccountClassification != AccountClassificationBusiness {
		return fmt.Errorf("attributes classification bust be either 'Personal' or 'Business'")
	}
	if len(attributes.SecondaryIdentification) > 140 {
		return fmt.Errorf("secondary identification maximum length is 140 characters")
	}
	if attributes.Status != "" &&
		attributes.Status != StatusPending &&
		attributes.Status != StatusConfirmed &&
		attributes.Status != StatusFailed {
		return fmt.Errorf("status must be 'pending', 'confirmed' or 'failed'")
	}
	return nil
}
