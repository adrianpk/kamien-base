package models

import (
	m "github.com/adrianpk/kamien/models"
	"github.com/markbates/pop/nulls"
	uuid "github.com/satori/go.uuid"
)

type (
	// Profile - Profile model
	Profile struct {
		m.Identification
		m.Detail
		Email          nulls.String `db:"email" json:"email,omitempty"`
		Location       nulls.String `db:"location" json:"location,omitempty"`
		Bio            nulls.String `db:"bio" json:"bio,omitempty"`
		Moto           nulls.String `db:"moto" json:"moto,omitempty"`
		Website        nulls.String `db:"website" json:"website,omitempty"`
		AniversaryDate nulls.Time   `db:"aniversary_date" json:"aniversaryDate,omitempty"`
		Avatar         nulls.String `db:"avatar" json:"avatar,omitempty"`
		Host           nulls.String `db:"host" json:"host,omitempty"`
		AvatarPath     nulls.String `db:"avatar_path" json:"avatarPath,omitempty"`
		HeaderPath     nulls.String `db:"header_path" json:"headerPath,omitempty"`
		OwnerID        nulls.UUID   `db:"owner_id" json:"ownerID,omitempty"`
		m.Geo
		m.TimeBounds
		m.LogicalStatus
		m.Audit
		m.Validation
		m.BagModel // Fix: Remove if '_method' hidden field does not cause problems to schema form decoder.
	}
)

// MakeProfile - Returns an empty Profile.
func MakeProfile() *Profile {
	return &Profile{
		Identification: *m.MakeIdentification(),
		Detail:         *m.MakeDetail(),
		Email:          m.NullsEmptyString(),
		Location:       m.NullsEmptyString(),
		Bio:            m.NullsEmptyString(),
		Moto:           m.NullsEmptyString(),
		Website:        m.NullsEmptyString(),
		AniversaryDate: m.NullsZeroTime(),
		Avatar:         m.NullsEmptyString(),
		Host:           m.NullsEmptyString(),
		AvatarPath:     m.NullsEmptyString(),
		HeaderPath:     m.NullsEmptyString(),
		OwnerID:        m.NullsZeroUUID(),
		Geo:            *m.MakeGeo(0, 0),
		TimeBounds:     *m.MakeTimeBounds(),
		LogicalStatus:  *m.MakeLogicalStatus(true, false),
		Audit:          *m.MakeAudit(),
	}
}

// MakeProfileWithID - Returns an empty Profile.
func MakeProfileWithID(id uuid.UUID) *Profile {
	profile := Profile{}
	profile.SetID(id)
	return &profile
}

// SetCreateValues - Set values for audit fields.
func (profile *Profile) SetCreateValues() {
	profile.Audit.SetCreateValues()
	profile.LogicalStatus.SetCreateValues()
}

// SetUpdateValues - Updates audit field.
func (profile *Profile) SetUpdateValues() {
	profile.Audit.SetUpdateValues()
}

// Match - Custom Profile comparator.
func (profile *Profile) Match(tc *Profile) bool {
	r := profile.Identification.Match(tc.Identification) &&
		// profile.Detail.Match(tc.Detail) &&
		profile.Email == tc.Email &&
		profile.Location == tc.Location &&
		profile.Bio == tc.Bio &&
		profile.Moto == tc.Moto &&
		profile.Website == tc.Website &&
		profile.AniversaryDate == tc.AniversaryDate &&
		profile.Avatar == tc.Avatar &&
		profile.Host == tc.Host &&
		profile.AvatarPath == tc.AvatarPath &&
		profile.HeaderPath == tc.HeaderPath &&
		profile.OwnerID == tc.OwnerID &&
		profile.Name == tc.Name &&
		profile.Description == tc.Description &&
		profile.Geolocation == tc.Geolocation &&
		profile.StartsAt == tc.StartsAt &&
		profile.EndsAt == tc.EndsAt
		// profile.Geo.Match(tc.Geo) &&
		// profile.TimeBounds.Match(tc.Timebounds)
	return r
}
