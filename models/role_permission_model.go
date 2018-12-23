package models

import (
	m "github.com/adrianpk/kamien/models"
	"github.com/markbates/pop/nulls"
	uuid "github.com/satori/go.uuid"
)

type (
	// RolePermission - RolePermission model
	RolePermission struct {
		m.Identification
		m.Detail
		OrganizationID nulls.UUID `db:"organization_id" json:"organizationID,omitempty"`
		RoleID         nulls.UUID `db:"role_id" json:"roleID,omitempty"`
		PermissionID   nulls.UUID `db:"permission_id" json:"permissionID,omitempty"`
		m.LogicalStatus
		m.Audit
		m.Validation
		m.BagModel // Fix: Remove if '_method' hidden field does not cause problems to schema form decoder.
	}
)

// MakeRolePermission - Returns an empty RolePermission.
func MakeRolePermission() *RolePermission {
	return &RolePermission{
		Identification: *m.MakeIdentification(),
		Detail:         *m.MakeDetail(),
		OrganizationID: m.NullsZeroUUID(),
		RoleID:         m.NullsZeroUUID(),
		PermissionID:   m.NullsZeroUUID(),
		LogicalStatus:  *m.MakeLogicalStatus(true, false),
		Audit:          *m.MakeAudit(),
	}
}

// MakeRolePermissionWithID - Returns an empty RolePermission.
func MakeRolePermissionWithID(id uuid.UUID) *RolePermission {
	rolePermission := RolePermission{}
	rolePermission.SetID(id)
	return &rolePermission
}

// SetCreateValues - Set values for audit fields.
func (rolePermission *RolePermission) SetCreateValues() {
	rolePermission.Audit.SetCreateValues()
	rolePermission.LogicalStatus.SetCreateValues()
}

// SetUpdateValues - Updates audit field.
func (rolePermission *RolePermission) SetUpdateValues() {
	rolePermission.Audit.SetUpdateValues()
}

// Match - Custom RolePermission comparator.
func (rolePermission *RolePermission) Match(tc *RolePermission) bool {
	r := rolePermission.Identification.Match(tc.Identification) &&
		// rolePermission.Detail.Match(tc.Detail) &&
		rolePermission.OrganizationID == tc.OrganizationID &&
		rolePermission.RoleID == tc.RoleID &&
		rolePermission.PermissionID == tc.PermissionID &&
		rolePermission.Name == tc.Name &&
		rolePermission.Description == tc.Description
		// rolePermission.Geo.Match(tc.Geo) &&
		// rolePermission.TimeBounds.Match(tc.Timebounds)
	return r
}
