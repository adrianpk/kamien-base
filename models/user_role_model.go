package models

import (
	m "github.com/adrianpk/kamien/models"
	"github.com/markbates/pop/nulls"
	uuid "github.com/satori/go.uuid"
)

type (
	// UserRole - UserRole model
	UserRole struct {
		m.Identification
		m.Detail
		OrganizationID nulls.UUID `db:"organization_id" json:"organizationID,omitempty"`
		UserID         nulls.UUID `db:"user_id" json:"userID,omitempty"`
		RoleID         nulls.UUID `db:"role_id" json:"roleID,omitempty"`
		m.LogicalStatus
		m.Audit
		m.Validation
		m.BagModel // Fix: Remove if '_method' hidden field does not cause problems to schema form decoder.
	}
)

// MakeUserRole - Returns an empty UserRole.
func MakeUserRole() *UserRole {
	return &UserRole{
		Identification: *m.MakeIdentification(),
		Detail:         *m.MakeDetail(),
		OrganizationID: m.NullsZeroUUID(),
		UserID:         m.NullsZeroUUID(),
		RoleID:         m.NullsZeroUUID(),
		LogicalStatus:  *m.MakeLogicalStatus(true, false),
		Audit:          *m.MakeAudit(),
	}
}

// MakeUserRoleWithID - Returns an empty UserRole.
func MakeUserRoleWithID(id uuid.UUID) *UserRole {
	userRole := UserRole{}
	userRole.SetID(id)
	return &userRole
}

// SetCreateValues - Set values for audit fields.
func (userRole *UserRole) SetCreateValues() {
	userRole.Audit.SetCreateValues()
	userRole.LogicalStatus.SetCreateValues()
}

// SetUpdateValues - Updates audit field.
func (userRole *UserRole) SetUpdateValues() {
	userRole.Audit.SetUpdateValues()
}

// Match - Custom UserRole comparator.
func (userRole *UserRole) Match(tc *UserRole) bool {
	r := userRole.Identification.Match(tc.Identification) &&
		// userRole.Detail.Match(tc.Detail) &&
		userRole.OrganizationID == tc.OrganizationID &&
		userRole.UserID == tc.UserID &&
		userRole.RoleID == tc.RoleID &&
		userRole.Name == tc.Name &&
		userRole.Description == tc.Description
		// userRole.Geo.Match(tc.Geo) &&
		// userRole.TimeBounds.Match(tc.Timebounds)
	return r
}
