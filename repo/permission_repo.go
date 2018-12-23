package repo

import (
	"bytes"
	"fmt"

	"{{.Package}}/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Initialize pq.
	uuid "github.com/satori/go.uuid"
)

// PermissionRepository - Permission repository manager.
type PermissionRepository struct {
	DB *sqlx.DB
	Tx *sqlx.Tx
}

// MakePermissionRepoTx - PermissionRepo constructor.
func MakePermissionRepoTx(tx *sqlx.Tx) PermissionRepository {
	return PermissionRepository{Tx: tx}
}

// GetAll - Get all Permissions from repo.
func (repo *PermissionRepository) GetAll() ([]models.Permission, error) {
	Permissions := []models.Permission{}
	err := repo.Tx.Select(&Permissions, "SELECT * FROM permissions;")
	return Permissions, err
}

// Get - Retrive a Permission from repo by its ID.
func (repo *PermissionRepository) Get(id uuid.UUID) (*models.Permission, error) {
	Permission := models.Permission{}
	err := repo.Tx.Get(&Permission, "SELECT * FROM permissions WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	return &Permission, nil
}

// GetByName - Retrive a Permission from repo by its Name.
func (repo *PermissionRepository) GetByName(name string) (*models.Permission, error) {
	Permission := models.Permission{}
	err := repo.Tx.Get(&Permission, "SELECT * FROM permissions WHERE name = $1;", name)
	if err != nil {
		return nil, err
	}
	return &Permission, nil
}

// Create - Create a Permission into the repo.
func (repo *PermissionRepository) Create(permission *models.Permission) error {
	permission.GenerateID()
	permission.SetCreateValues()
	permission.CreatedBy = permission.ID
	permission.UpdatedBy = permission.ID
	permissionInsertSQL := "INSERT INTO permissions (id, organization_id, name, description, is_active, is_logical_deleted, created_by_id, updated_by_id, created_at, updated_at) VALUES (:id, :organization_id, :name, :description, :is_active, :is_logical_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at);"
	_, err := repo.Tx.NamedExec(permissionInsertSQL, permission)
	return err
}

// Update - Update a Permission in repo.
func (repo *PermissionRepository) Update(permission *models.Permission) error {
	permission.SetUpdateValues()
	current, err := repo.Get(permission.ID.UUID)
	if err != nil {
		return err
	}
	// Build update statement
	changes, qty := PermissionChanges(permission, current)
	pos := 0
	last := qty < 2
	var query bytes.Buffer
	query.WriteString("UPDATE permissions SET ")
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
	query.WriteString(fmt.Sprintf("WHERE id = '%s';", permission.ID.UUID))
	repo.Tx.NamedExec(query.String(), &permission)
	// log.Debugf("Update permission query:\n%s", query.String())
	return err
}

// Delete - Delete a Permission from repo.
func (repo *PermissionRepository) Delete(id uuid.UUID) {
	PermissionDeleteSQL := fmt.Sprintf("DELETE FROM permissions WHERE id = '%s'", id)
	_ = repo.Tx.MustExec(PermissionDeleteSQL)
}

// PermissionChanges - Creates a map ([string]interface{}) including al changing field.
func PermissionChanges(permission, current *models.Permission) (map[string]string, int) {
	changes := make(map[string]string)
	if current.OrganizationID.UUID != permission.OrganizationID.UUID {
		changes["organization_id"] = ":organization_id"
	}
	if current.Name.String != permission.Name.String {
		changes["name"] = ":name"
	}
	if current.Description.String != permission.Description.String {
		changes["description"] = ":description"
	}
	if current.IsActive.Bool != permission.IsActive.Bool {
		changes["is_active"] = ":is_active"
	}
	if current.IsLogicalDeleted.Bool != permission.IsLogicalDeleted.Bool {
		changes["is_logical_deleted"] = ":is_logical_deleted"
	}
	if current.CreatedBy.UUID != permission.CreatedBy.UUID {
		changes["created_by_id"] = ":created_by_id"
	}
	if current.UpdatedBy.UUID != permission.UpdatedBy.UUID {
		changes["updated_by_id"] = ":updated_by_id"
	}
	if !current.CreatedAt.Time.Equal(permission.CreatedAt.Time) {
		changes["created_at"] = ":created_at"
	}
	if !current.UpdatedAt.Time.Equal(permission.UpdatedAt.Time) {
		changes["updated_at"] = ":updated_at"
	}
	return changes, len(changes)
}
