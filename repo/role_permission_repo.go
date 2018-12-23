package repo

import (
	"bytes"
	"fmt"

	"{{.Package}}/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Initialize pq.
	uuid "github.com/satori/go.uuid"
)

// RolePermissionRepository - RolePermission repository manager.
type RolePermissionRepository struct {
	DB *sqlx.DB
	Tx *sqlx.Tx
}

// MakeRolePermissionRepoTx - RolePermissionRepo constructor.
func MakeRolePermissionRepoTx(tx *sqlx.Tx) RolePermissionRepository {
	return RolePermissionRepository{Tx: tx}
}

// GetAll - Get all RolePermissions from repo.
func (repo *RolePermissionRepository) GetAll() ([]models.RolePermission, error) {
	RolePermissions := []models.RolePermission{}
	err := repo.Tx.Select(&RolePermissions, "SELECT * FROM role_permissions;")
	return RolePermissions, err
}

// Get - Retrive a RolePermission from repo by its ID.
func (repo *RolePermissionRepository) Get(id uuid.UUID) (*models.RolePermission, error) {
	RolePermission := models.RolePermission{}
	err := repo.Tx.Get(&RolePermission, "SELECT * FROM role_permissions WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	return &RolePermission, nil
}

// GetByName - Retrive a RolePermission from repo by its Name.
func (repo *RolePermissionRepository) GetByName(name string) (*models.RolePermission, error) {
	RolePermission := models.RolePermission{}
	err := repo.Tx.Get(&RolePermission, "SELECT * FROM role_permissions WHERE name = $1;", name)
	if err != nil {
		return nil, err
	}
	return &RolePermission, nil
}

// Create - Create a RolePermission into the repo.
func (repo *RolePermissionRepository) Create(rolePermission *models.RolePermission) error {
	rolePermission.GenerateID()
	rolePermission.SetCreateValues()
	rolePermission.CreatedBy = rolePermission.ID
	rolePermission.UpdatedBy = rolePermission.ID
	rolePermissionInsertSQL := "INSERT INTO role_permissions (id, organization_id, role_id, permission_id, name, description, is_active, is_logical_deleted, created_by_id, updated_by_id, created_at, updated_at) VALUES (:id, :organization_id, :role_id, :permission_id, :name, :description, :is_active, :is_logical_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at);"
	_, err := repo.Tx.NamedExec(rolePermissionInsertSQL, rolePermission)
	return err
}

// Update - Update a RolePermission in repo.
func (repo *RolePermissionRepository) Update(rolePermission *models.RolePermission) error {
	rolePermission.SetUpdateValues()
	current, err := repo.Get(rolePermission.ID.UUID)
	if err != nil {
		return err
	}
	// Build update statement
	changes, qty := RolePermissionChanges(rolePermission, current)
	pos := 0
	last := qty < 2
	var query bytes.Buffer
	query.WriteString("UPDATE role_permissions SET ")
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
	query.WriteString(fmt.Sprintf("WHERE id = '%s';", rolePermission.ID.UUID))
	repo.Tx.NamedExec(query.String(), &rolePermission)
	// log.Debugf("Update rolePermission query:\n%s", query.String())
	return err
}

// Delete - Delete a RolePermission from repo.
func (repo *RolePermissionRepository) Delete(id uuid.UUID) {
	RolePermissionDeleteSQL := fmt.Sprintf("DELETE FROM role_permissions WHERE id = '%s'", id)
	_ = repo.Tx.MustExec(RolePermissionDeleteSQL)
}

// RolePermissionChanges - Creates a map ([string]interface{}) including al changing field.
func RolePermissionChanges(rolePermission, current *models.RolePermission) (map[string]string, int) {
	changes := make(map[string]string)
	if current.OrganizationID.UUID != rolePermission.OrganizationID.UUID {
		changes["organization_id"] = ":organization_id"
	}
	if current.RoleID.UUID != rolePermission.RoleID.UUID {
		changes["role_id"] = ":role_id"
	}
	if current.PermissionID.UUID != rolePermission.PermissionID.UUID {
		changes["permission_id"] = ":permission_id"
	}
	if current.Name.String != rolePermission.Name.String {
		changes["name"] = ":name"
	}
	if current.Description.String != rolePermission.Description.String {
		changes["description"] = ":description"
	}
	if current.IsActive.Bool != rolePermission.IsActive.Bool {
		changes["is_active"] = ":is_active"
	}
	if current.IsLogicalDeleted.Bool != rolePermission.IsLogicalDeleted.Bool {
		changes["is_logical_deleted"] = ":is_logical_deleted"
	}
	if current.CreatedBy.UUID != rolePermission.CreatedBy.UUID {
		changes["created_by_id"] = ":created_by_id"
	}
	if current.UpdatedBy.UUID != rolePermission.UpdatedBy.UUID {
		changes["updated_by_id"] = ":updated_by_id"
	}
	if !current.CreatedAt.Time.Equal(rolePermission.CreatedAt.Time) {
		changes["created_at"] = ":created_at"
	}
	if !current.UpdatedAt.Time.Equal(rolePermission.UpdatedAt.Time) {
		changes["updated_at"] = ":updated_at"
	}
	return changes, len(changes)
}
