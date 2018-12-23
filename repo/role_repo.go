package repo

import (
	"bytes"
	"fmt"

	"{{.Package}}/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Initialize pq.
	uuid "github.com/satori/go.uuid"
)

// RoleRepository - Role repository manager.
type RoleRepository struct {
	DB *sqlx.DB
	Tx *sqlx.Tx
}

// MakeRoleRepoTx - RoleRepo constructor.
func MakeRoleRepoTx(tx *sqlx.Tx) RoleRepository {
	return RoleRepository{Tx: tx}
}

// GetAll - Get all Roles from repo.
func (repo *RoleRepository) GetAll() ([]models.Role, error) {
	Roles := []models.Role{}
	err := repo.Tx.Select(&Roles, "SELECT * FROM roles;")
	return Roles, err
}

// Get - Retrive a Role from repo by its ID.
func (repo *RoleRepository) Get(id uuid.UUID) (*models.Role, error) {
	Role := models.Role{}
	err := repo.Tx.Get(&Role, "SELECT * FROM roles WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	return &Role, nil
}

// GetByName - Retrive a Role from repo by its Name.
func (repo *RoleRepository) GetByName(name string) (*models.Role, error) {
	Role := models.Role{}
	err := repo.Tx.Get(&Role, "SELECT * FROM roles WHERE name = $1;", name)
	if err != nil {
		return nil, err
	}
	return &Role, nil
}

// Create - Create a Role into the repo.
func (repo *RoleRepository) Create(role *models.Role) error {
	role.GenerateID()
	role.SetCreateValues()
	role.CreatedBy = role.ID
	role.UpdatedBy = role.ID
	roleInsertSQL := "INSERT INTO roles (id, organization_id, name, description, is_active, is_logical_deleted, created_by_id, updated_by_id, created_at, updated_at) VALUES (:id, :organization_id, :name, :description, :is_active, :is_logical_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at);"
	_, err := repo.Tx.NamedExec(roleInsertSQL, role)
	return err
}

// Update - Update a Role in repo.
func (repo *RoleRepository) Update(role *models.Role) error {
	role.SetUpdateValues()
	current, err := repo.Get(role.ID.UUID)
	if err != nil {
		return err
	}
	// Build update statement
	changes, qty := RoleChanges(role, current)
	pos := 0
	last := qty < 2
	var query bytes.Buffer
	query.WriteString("UPDATE roles SET ")
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
	query.WriteString(fmt.Sprintf("WHERE id = '%s';", role.ID.UUID))
	repo.Tx.NamedExec(query.String(), &role)
	// log.Debugf("Update role query:\n%s", query.String())
	return err
}

// Delete - Delete a Role from repo.
func (repo *RoleRepository) Delete(id uuid.UUID) {
	RoleDeleteSQL := fmt.Sprintf("DELETE FROM roles WHERE id = '%s'", id)
	_ = repo.Tx.MustExec(RoleDeleteSQL)
}

// RoleChanges - Creates a map ([string]interface{}) including al changing field.
func RoleChanges(role, current *models.Role) (map[string]string, int) {
	changes := make(map[string]string)
	if current.OrganizationID.UUID != role.OrganizationID.UUID {
		changes["organization_id"] = ":organization_id"
	}
	if current.Name.String != role.Name.String {
		changes["name"] = ":name"
	}
	if current.Description.String != role.Description.String {
		changes["description"] = ":description"
	}
	if current.IsActive.Bool != role.IsActive.Bool {
		changes["is_active"] = ":is_active"
	}
	if current.IsLogicalDeleted.Bool != role.IsLogicalDeleted.Bool {
		changes["is_logical_deleted"] = ":is_logical_deleted"
	}
	if current.CreatedBy.UUID != role.CreatedBy.UUID {
		changes["created_by_id"] = ":created_by_id"
	}
	if current.UpdatedBy.UUID != role.UpdatedBy.UUID {
		changes["updated_by_id"] = ":updated_by_id"
	}
	if !current.CreatedAt.Time.Equal(role.CreatedAt.Time) {
		changes["created_at"] = ":created_at"
	}
	if !current.UpdatedAt.Time.Equal(role.UpdatedAt.Time) {
		changes["updated_at"] = ":updated_at"
	}
	return changes, len(changes)
}
