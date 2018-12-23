package models

import (
	"time"

	m "github.com/adrianpk/kamien/models"
	"github.com/markbates/pop/nulls"
	uuid "github.com/satori/go.uuid"
)

type (
	// Account - Account model
	Account struct {
		m.Identification
		m.Detail
		AccountType nulls.String `db:"account_type" json:"accountType" schema:"account-type"`
		OwnerID     nulls.UUID   `db:"owner_id" json:"ownerID" schema:"owner-id"`
		ParentID    nulls.UUID   `db:"parent_id" json:"parentID" schema:"parent-id"`
		Email       nulls.String `db:"email" json:"email" schema:"email"`
		m.Geo
		m.TimeBounds
		m.LogicalStatus
		m.Audit
		m.Validation
		m.BagModel // Fix: Remove if '_method' hidden field does not cause problems to schema form decoder.
	}
)

// MakeAccount - Returns an empty Account.
func MakeAccount() *Account {
	return &Account{
		Identification: *m.MakeIdentification(),
		Detail:         *m.MakeDetail(),
		Email:          m.NullsEmptyString(),
		Geo:            *m.MakeGeo(0, 0),
		TimeBounds:     *m.MakeTimeBounds(),
		LogicalStatus:  *m.MakeLogicalStatus(true, false),
		Audit:          *m.MakeAudit(),
	}
}

// MakeAccountWithID - Returns an empty Account.
func MakeAccountWithID(id uuid.UUID) *Account {
	account := Account{}
	account.SetID(id)
	return &account
}

// MakeAccountTOPE - Returns a Account with custom values.
func MakeAccountTOPE(accountType string, ownerID, parentID uuid.UUID, email string) *Account {
	account := &Account{
		AccountType: m.ToNullsString(accountType),
		OwnerID:     m.ToNullsUUID(ownerID),
		ParentID:    m.ToNullsUUID(parentID),
		Email:       m.ToNullsString(email),
		// Geo:           *m.MakeGeo(0, 0),
		// TimeBounds:    *m.MakeTimeBounds(),
		LogicalStatus: *m.MakeLogicalStatus(true, false),
	}
	account.GenerateID()
	return account
}

// SetCreateValues - Set values for audit fields.
func (account *Account) SetCreateValues() {
	account.Audit.SetCreateValues()
	account.LogicalStatus.SetCreateValues()
	account.StartsAt = m.ToNullsTime(time.Now())
	account.EndsAt = m.ToNullsTime(time.Now())
}

// SetUpdateValues - Updates audit field.
func (account *Account) SetUpdateValues() {
	account.Audit.SetUpdateValues()
}

// Match - Custom model comparator.
func (account *Account) Match(tc *Account) bool {
	r := account.Identification.Match(tc.Identification) &&
		account.Name == tc.Name &&
		account.OwnerID == tc.OwnerID &&
		account.ParentID == tc.ParentID &&
		account.Email == tc.Email
	return r
}
