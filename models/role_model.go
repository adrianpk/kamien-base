package models

import (
	m "github.com/adrianpk/kamien/models"
	"github.com/markbates/pop/nulls"
	uuid "github.com/satori/go.uuid"
)

type (
	// Role - Role model
	Role struct {
		m.Identification
		m.Detail
		OrganizationID nulls.UUID `db:"organization_id" json:"organizationID,omitempty"`
		m.LogicalStatus
		m.Audit
		m.Validation
		m.BagModel // Fix: Remove if '_method' hidden field does not cause problems to schema form decoder.
	}
)

// MakeRole - Returns an empty Role.
func MakeRole() *Role {
	return &Role{
		Identification: *m.MakeIdentification(),
		Detail:         *m.MakeDetail(),
		OrganizationID: m.NullsZeroUUID(),
		LogicalStatus:  *m.MakeLogicalStatus(true, false),
		Audit:          *m.MakeAudit(),
	}
}

// MakeRoleWithID - Returns an empty Role.
func MakeRoleWithID(id uuid.UUID) *Role {
	role := Role{}
	role.SetID(id)
	return &role
}

// SetCreateValues - Set values for audit fields.
func (role *Role) SetCreateValues() {
	role.Audit.SetCreateValues()
	role.LogicalStatus.SetCreateValues()
}

// SetUpdateValues - Updates audit field.
func (role *Role) SetUpdateValues() {
	role.Audit.SetUpdateValues()
}

// Match - Custom Role comparator.
func (role *Role) Match(tc *Role) bool {
	r := role.Identification.Match(tc.Identification) &&
		// role.Detail.Match(tc.Detail) &&
		role.OrganizationID == tc.OrganizationID &&
		role.Name == tc.Name &&
		role.Description == tc.Description
		// role.Geo.Match(tc.Geo) &&
		// role.TimeBounds.Match(tc.Timebounds)
	return r
}
