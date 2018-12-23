package models

import (
	m "github.com/adrianpk/kamien/models"
	"github.com/markbates/pop/nulls"
	uuid "github.com/satori/go.uuid"
)

type (
	// Permission - Permission model
	Permission struct {
		m.Identification
		m.Detail
		OrganizationID nulls.UUID `db:"organization_id" json:"organizationID,omitempty"`
		m.LogicalStatus
    m.Audit
		m.Validation
		m.BagModel // Fix: Remove if '_method' hidden field does not cause problems to schema form decoder.
	}
)

// MakePermission - Returns an empty Permission.
func MakePermission() *Permission {
	return &Permission{
		Identification:   *m.MakeIdentification(),
		Detail:   *m.MakeDetail(),
		OrganizationID: m.NullsZeroUUID(),
		LogicalStatus:  *m.MakeLogicalStatus(true, false),
 		Audit: *m.MakeAudit(),
	}
}

// MakePermissionWithID - Returns an empty Permission.
func MakePermissionWithID(id uuid.UUID) *Permission {
	permission := Permission{}
	permission.SetID(id)
	return &permission
}

// SetCreateValues - Set values for audit fields.
func (permission *Permission) SetCreateValues() {
	permission.Audit.SetCreateValues()
	permission.LogicalStatus.SetCreateValues()
}

// SetUpdateValues - Updates audit field.
func (permission *Permission) SetUpdateValues() {
	permission.Audit.SetUpdateValues()
}

// Match - Custom Permission comparator.
func (permission *Permission) Match(tc *Permission) bool {
	r := permission.Identification.Match(tc.Identification) &&
		// permission.Detail.Match(tc.Detail) &&
		permission.OrganizationID == tc.OrganizationID &&
		permission.Name == tc.Name &&
		permission.Description == tc.Description
		// permission.Geo.Match(tc.Geo) &&
		// permission.TimeBounds.Match(tc.Timebounds)
	return r
}

