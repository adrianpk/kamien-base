package models

import (
	m "github.com/adrianpk/kamien/models"
	"github.com/markbates/pop/nulls"
	uuid "github.com/satori/go.uuid"
)

type (
	// User - User model
	User struct {
		m.Identification
		Username nulls.String `db:"username" json:"username" schema:"username"`
		m.Authentication
		Email       nulls.String `db:"email" json:"email" schema:"email"`
		GivenName   nulls.String `db:"given_name" json:"givenName" schema:"given-name"`
		MiddleNames nulls.String `db:"middle_names" json:"middleNames" schema:"middle-names"`
		FamilyName  nulls.String `db:"family_name" json:"familyName" schema:"family-name"`
		ContextID   nulls.UUID   `db:"context_id" json:"contextID"`
		m.Geo
		m.TimeBounds
		m.LogicalStatus
		m.Audit
		m.Validation
		m.BagModel // Fix: Remove if '_method' hidden field does not cause problems to schema form decoder.
	}
)

// MakeUser - Returns a 'zero value' User.
func MakeUser() *User {
	return &User{
		Identification: *m.MakeIdentification(),
		Username:       m.NullsEmptyString(),
		Authentication: *m.MakeAuthentication(),
		Email:          m.NullsEmptyString(),
		GivenName:      m.NullsEmptyString(),
		MiddleNames:    m.NullsEmptyString(),
		FamilyName:     m.NullsEmptyString(),
		ContextID:      m.NullsZeroUUID(),
		Geo:            *m.MakeGeo(0, 0),
		TimeBounds:     *m.MakeTimeBounds(),
		LogicalStatus:  *m.MakeLogicalStatus(true, false),
		Audit:          *m.MakeAudit(),
	}
}

// MakeUserWithID - Returns an initialized User with ID.
func MakeUserWithID(id uuid.UUID) *User {
	user := User{}
	user.SetID(id)
	return &user
}

// MakeUserUPE - Create a User with username, password and email.
func MakeUserUPE(username, password, passwordConfirmation, email string) *User {
	u := &User{
		Username: m.ToNullsString(username),
		Authentication: m.Authentication{
			Password:             password,
			PasswordConfirmation: passwordConfirmation},
		Email:         m.ToNullsString(email),
		Geo:           *m.MakeGeo(0, 0),
		TimeBounds:    *m.MakeTimeBounds(),
		LogicalStatus: *m.MakeLogicalStatus(true, false),
	}
	// u.GenerateID()
	u.PairContextID()
	// u.ClearPassword()
	return u
}

// SetNames - Set User names.
func (user *User) SetNames(given, middles, family string) {
	user.GivenName = m.ToNullsString(given)
	user.MiddleNames = m.ToNullsString(middles)
	user.FamilyName = m.ToNullsString(family)
}

// SetCreateValues - Default values for models after creation.
func (user *User) SetCreateValues() {
	user.Audit.SetCreateValues()
	user.LogicalStatus.SetCreateValues()
}

// SetUpdateValues - Updates audit field.
func (user *User) SetUpdateValues() {
	user.Audit.SetUpdateValues()
}

// PairContextID - Makes Context ID equals to user
func (user *User) PairContextID() {
	user.ContextID = user.ID
}

// Match - Custom model comparator.
func (user *User) Match(tc *User) bool {
	r := user.Identification.Match(tc.Identification) &&
		user.Username == tc.Username &&
		user.Email == tc.Email &&
		user.GivenName == tc.GivenName &&
		user.MiddleNames == tc.MiddleNames &&
		user.FamilyName == tc.FamilyName &&
		user.ContextID == tc.ContextID &&
		user.Geo.Match(tc.Geo) &&
		user.TimeBounds.Match(tc.TimeBounds)
	return r
}
