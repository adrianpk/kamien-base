package repo

import (
	"bytes"
	"fmt"

	"{{.Package}}/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Initialize pq.
	uuid "github.com/satori/go.uuid"
)

// ProfileRepository - Profile repository manager.
type ProfileRepository struct {
	DB *sqlx.DB
	Tx *sqlx.Tx
}

// MakeProfileRepoTx - ProfileRepo constructor.
func MakeProfileRepoTx(tx *sqlx.Tx) ProfileRepository {
	return ProfileRepository{Tx: tx}
}

// GetAll - Get all Profiles from repo.
func (repo *ProfileRepository) GetAll() ([]models.Profile, error) {
	profiles := []models.Profile{}
	err := repo.Tx.Select(&profiles, "SELECT * FROM profiles;")
	return profiles, err
}

// Get - Retrive a Profile from repo by its ID.
func (repo *ProfileRepository) Get(id uuid.UUID) (*models.Profile, error) {
	profile := models.Profile{}
	err := repo.Tx.Get(&profile, "SELECT * FROM profiles WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// GetByName - Retrive a Profile from repo by its Name.
func (repo *ProfileRepository) GetByName(name string) (*models.Profile, error) {
	profile := models.Profile{}
	err := repo.Tx.Get(&profile, "SELECT * FROM profiles WHERE name = $1;", name)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// GetByOwnerID - Retrive a Profile from repo by its owner ID.
func (repo *ProfileRepository) GetByOwnerID(id uuid.UUID) (*models.Profile, error) {
	profile := models.Profile{}
	err := repo.Tx.Get(&profile, "SELECT * FROM profiles WHERE owner_id = $1;", id)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// Create - Create a Profile into the repo.
func (repo *ProfileRepository) Create(profile *models.Profile) error {
	profile.GenerateID()
	profile.SetCreateValues()
	profile.CreatedBy = profile.ID
	profile.UpdatedBy = profile.ID
	profileInsertSQL := "INSERT INTO profiles (id, email, location, bio, moto, website, aniversary_date, host, avatar_path, header_path, owner_id, name, description, geolocation, starts_at, ends_at, is_active, is_logical_deleted, created_by_id, updated_by_id, created_at, updated_at) VALUES (:id, :email, :location, :bio, :moto, :website, :aniversary_date, :host, :avatar_path, :header_path, :owner_id, :name, :description, :geolocation, :starts_at, :ends_at, :is_active, :is_logical_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at);"
	_, err := repo.Tx.NamedExec(profileInsertSQL, profile)
	return err
}

// Update - Update a Profile in repo.
func (repo *ProfileRepository) Update(profile *models.Profile) error {
	profile.SetUpdateValues()
	current, err := repo.Get(profile.ID.UUID)
	if err != nil {
		return err
	}
	// Build update statement
	changes, qty := ProfileChanges(profile, current)
	pos := 0
	last := qty < 2
	var query bytes.Buffer
	query.WriteString("UPDATE profiles SET ")
	for field, structField := range changes {
		var partial string
		if last {
			partial = fmt.Sprintf("%v = %v ", field, structField)
		} else {
			partial = fmt.Sprintf("%v = %v, ", field, structField)
		}
		query.WriteString(partial)
		pos = pos + 1
		last = pos == qty-1
	}
	query.WriteString(fmt.Sprintf("WHERE id = '%s';", profile.ID.UUID))
	repo.Tx.NamedExec(query.String(), &profile)
	// log.Debugf("Update profile query:\n%s", query.String())
	return err
}

// Delete - Delete a Profile from repo.
func (repo *ProfileRepository) Delete(id uuid.UUID) {
	profileDeleteSQL := fmt.Sprintf("DELETE FROM profiles WHERE id = '%s'", id)
	_ = repo.Tx.MustExec(profileDeleteSQL)
}

// ProfileChanges - Creates a map ([string]interface{}) including al changing field.
func ProfileChanges(profile, current *models.Profile) (map[string]string, int) {
	changes := make(map[string]string)
	if current.Email.String != profile.Email.String {
		changes["email"] = ":email"
	}
	if current.Location.String != profile.Location.String {
		changes["location"] = ":location"
	}
	if current.Bio.String != profile.Bio.String {
		changes["bio"] = ":bio"
	}
	if current.Moto.String != profile.Moto.String {
		changes["moto"] = ":moto"
	}
	if current.Website.String != profile.Website.String {
		changes["website"] = ":website"
	}
	if !current.AniversaryDate.Time.Equal(profile.AniversaryDate.Time) {
		changes["aniversary_date"] = ":aniversary_date"
	}
	if current.Host.String != profile.Host.String {
		changes["host"] = ":host"
	}
	if current.AvatarPath.String != profile.AvatarPath.String {
		changes["avatar_path"] = ":avatar_path"
	}
	if current.HeaderPath.String != profile.HeaderPath.String {
		changes["header_path"] = ":header_path"
	}
	if current.OwnerID.UUID != profile.OwnerID.UUID {
		changes["owner_id"] = ":owner_id"
	}
	if current.Name.String != profile.Name.String {
		changes["name"] = ":name"
	}
	if current.Description.String != profile.Description.String {
		changes["description"] = ":description"
	}
	if current.Geolocation.Point.String() != profile.Geolocation.Point.String() {
		changes["geolocation"] = ":geolocation"
	}
	if !current.StartsAt.Time.Equal(profile.StartsAt.Time) {
		changes["starts_at"] = ":starts_at"
	}
	if !current.EndsAt.Time.Equal(profile.EndsAt.Time) {
		changes["ends_at"] = ":ends_at"
	}
	if current.IsActive.Bool != profile.IsActive.Bool {
		changes["is_active"] = ":is_active"
	}
	if current.IsLogicalDeleted.Bool != profile.IsLogicalDeleted.Bool {
		changes["is_logical_deleted"] = ":is_logical_deleted"
	}
	if current.CreatedBy.UUID != profile.CreatedBy.UUID {
		changes["created_by_id"] = ":created_by_id"
	}
	if current.UpdatedBy.UUID != profile.UpdatedBy.UUID {
		changes["updated_by_id"] = ":updated_by_id"
	}
	if !current.CreatedAt.Time.Equal(profile.CreatedAt.Time) {
		changes["created_at"] = ":created_at"
	}
	if !current.UpdatedAt.Time.Equal(profile.UpdatedAt.Time) {
		changes["updated_at"] = ":updated_at"
	}
	return changes, len(changes)
}
