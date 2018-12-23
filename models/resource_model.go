package models

import (
	m "github.com/adrianpk/kamien/models"
	"github.com/markbates/pop/nulls"
	uuid "github.com/satori/go.uuid"
)

type (
	// Resource - Resource model
	Resource struct {
		m.Identification
		m.Detail
		Tag            nulls.String `db:"tag" json:"tag,omitempty"`
		OrganizationID nulls.UUID   `db:"organization_id" json:"organizationID,omitempty"`
		m.LogicalStatus
		m.Audit
		m.Validation
		m.BagModel // Fix: Remove if '_method' hidden field does not cause problems to schema form decoder.
	}
)

// MakeResource - Returns an empty Resource.
func MakeResource() *Resource {
	return &Resource{
		Identification: *m.MakeIdentification(),
		Detail:         *m.MakeDetail(),
		Tag:            m.NullsEmptyString(),
		OrganizationID: m.NullsZeroUUID(),
		LogicalStatus:  *m.MakeLogicalStatus(true, false),
		Audit:          *m.MakeAudit(),
	}
}

// GenerateTag - Generates Resource's tag based on last8 digits of its ID.
func (resource *Resource) GenerateTag() {
	if !resource.IsNew() && len(resource.ID.UUID.String()) == 36 {
		resource.Tag = nulls.NewString(resource.ID.UUID.String()[28:36])
	}
}

// MakeResourceWithID - Returns an empty Resource.
func MakeResourceWithID(id uuid.UUID) *Resource {
	resource := Resource{}
	resource.SetID(id)
	return &resource
}

// SetCreateValues - Set values for audit fields.
func (resource *Resource) SetCreateValues() {
	resource.Audit.SetCreateValues()
	resource.LogicalStatus.SetCreateValues()
}

// SetUpdateValues - Updates audit field.
func (resource *Resource) SetUpdateValues() {
	resource.Audit.SetUpdateValues()
}

// Match - Custom Resource comparator.
func (resource *Resource) Match(tc *Resource) bool {
	r := resource.Identification.Match(tc.Identification) &&
		// resource.Detail.Match(tc.Detail) &&
		resource.Tag == tc.Tag &&
		resource.OrganizationID == tc.OrganizationID &&
		resource.Name == tc.Name &&
		resource.Description == tc.Description
		// resource.Geo.Match(tc.Geo) &&
		// resource.TimeBounds.Match(tc.Timebounds)
	return r
}
