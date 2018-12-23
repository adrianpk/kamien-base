package models

import (
	m "github.com/adrianpk/kamien/models"
	"github.com/markbates/pop/nulls"
	uuid "github.com/satori/go.uuid"
)

type (
	// ResourcePermission - ResourcePermission model
	ResourcePermission struct {
		m.Identification
		m.Detail
		OrganizationID nulls.UUID `db:"organization_id" json:"organizationID,omitempty"`
		ResourceID     nulls.UUID `db:"resource_id" json:"resourceID,omitempty"`
		PermissionID   nulls.UUID `db:"permission_id" json:"permissionID,omitempty"`
		m.LogicalStatus
		m.Audit
		m.Validation
		m.BagModel // Fix: Remove if '_method' hidden field does not cause problems to schema form decoder.
	}
)

// MakeResourcePermission - Returns an empty ResourcePermission.
func MakeResourcePermission() *ResourcePermission {
	return &ResourcePermission{
		Identification: *m.MakeIdentification(),
		Detail:         *m.MakeDetail(),
		OrganizationID: m.NullsZeroUUID(),
		ResourceID:     m.NullsZeroUUID(),
		PermissionID:   m.NullsZeroUUID(),
		LogicalStatus:  *m.MakeLogicalStatus(true, false),
		Audit:          *m.MakeAudit(),
	}
}

// MakeResourcePermissionWithID - Returns an empty ResourcePermission.
func MakeResourcePermissionWithID(id uuid.UUID) *ResourcePermission {
	resourcePermission := ResourcePermission{}
	resourcePermission.SetID(id)
	return &resourcePermission
}

// SetCreateValues - Set values for audit fields.
func (resourcePermission *ResourcePermission) SetCreateValues() {
	resourcePermission.Audit.SetCreateValues()
	resourcePermission.LogicalStatus.SetCreateValues()
}

// SetUpdateValues - Updates audit field.
func (resourcePermission *ResourcePermission) SetUpdateValues() {
	resourcePermission.Audit.SetUpdateValues()
}

// Match - Custom ResourcePermission comparator.
func (resourcePermission *ResourcePermission) Match(tc *ResourcePermission) bool {
	r := resourcePermission.Identification.Match(tc.Identification) &&
		// resourcePermission.Detail.Match(tc.Detail) &&
		resourcePermission.OrganizationID == tc.OrganizationID &&
		resourcePermission.ResourceID == tc.ResourceID &&
		resourcePermission.PermissionID == tc.PermissionID &&
		resourcePermission.Name == tc.Name &&
		resourcePermission.Description == tc.Description
		// resourcePermission.Geo.Match(tc.Geo) &&
		// resourcePermission.TimeBounds.Match(tc.Timebounds)
	return r
}
