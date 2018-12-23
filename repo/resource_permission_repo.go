package repo

import (
	"bytes"
	"fmt"

	"{{.Package}}/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Initialize pq.
	uuid "github.com/satori/go.uuid"
)

// ResourcePermissionRepository - ResourcePermission repository manager.
type ResourcePermissionRepository struct {
	DB *sqlx.DB
	Tx *sqlx.Tx
}

// MakeResourcePermissionRepoTx - ResourcePermissionRepo constructor.
func MakeResourcePermissionRepoTx(tx *sqlx.Tx) ResourcePermissionRepository {
	return ResourcePermissionRepository{Tx: tx}
}

// GetAll - Get all ResourcePermissions from repo.
func (repo *ResourcePermissionRepository) GetAll() ([]models.ResourcePermission, error) {
	ResourcePermissions := []models.ResourcePermission{}
	err := repo.Tx.Select(&ResourcePermissions, "SELECT * FROM resource_permissions;")
	return ResourcePermissions, err
}

// Get - Retrive a ResourcePermission from repo by its ID.
func (repo *ResourcePermissionRepository) Get(id uuid.UUID) (*models.ResourcePermission, error) {
	ResourcePermission := models.ResourcePermission{}
	err := repo.Tx.Get(&ResourcePermission, "SELECT * FROM resource_permissions WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	return &ResourcePermission, nil
}

// GetByName - Retrive a ResourcePermission from repo by its Name.
func (repo *ResourcePermissionRepository) GetByName(name string) (*models.ResourcePermission, error) {
	ResourcePermission := models.ResourcePermission{}
	err := repo.Tx.Get(&ResourcePermission, "SELECT * FROM resource_permissions WHERE name = $1;", name)
	if err != nil {
		return nil, err
	}
	return &ResourcePermission, nil
}

// Create - Create a ResourcePermission into the repo.
func (repo *ResourcePermissionRepository) Create(resourcePermission *models.ResourcePermission) error {
	resourcePermission.GenerateID()
	resourcePermission.SetCreateValues()
	resourcePermission.CreatedBy = resourcePermission.ID
	resourcePermission.UpdatedBy = resourcePermission.ID
	resourcePermissionInsertSQL := "INSERT INTO resource_permissions (id, organization_id, resource_id, permission_id, name, description, is_active, is_logical_deleted, created_by_id, updated_by_id, created_at, updated_at) VALUES (:id, :organization_id, :resource_id, :permission_id, :name, :description, :is_active, :is_logical_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at);"
	_, err := repo.Tx.NamedExec(resourcePermissionInsertSQL, resourcePermission)
	return err
}

// Update - Update a ResourcePermission in repo.
func (repo *ResourcePermissionRepository) Update(resourcePermission *models.ResourcePermission) error {
	resourcePermission.SetUpdateValues()
	current, err := repo.Get(resourcePermission.ID.UUID)
	if err != nil {
		return err
	}
	// Build update statement
	changes, qty := ResourcePermissionChanges(resourcePermission, current)
	pos := 0
	last := qty < 2
	var query bytes.Buffer
	query.WriteString("UPDATE resource_permissions SET ")
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
	query.WriteString(fmt.Sprintf("WHERE id = '%s';", resourcePermission.ID.UUID))
	repo.Tx.NamedExec(query.String(), &resourcePermission)
	// log.Debugf("Update resourcePermission query:\n%s", query.String())
	return err
}

// Delete - Delete a ResourcePermission from repo.
func (repo *ResourcePermissionRepository) Delete(id uuid.UUID) {
	ResourcePermissionDeleteSQL := fmt.Sprintf("DELETE FROM resource_permissions WHERE id = '%s'", id)
	_ = repo.Tx.MustExec(ResourcePermissionDeleteSQL)
}

// ResourcePermissionChanges - Creates a map ([string]interface{}) including al changing field.
func ResourcePermissionChanges(resourcePermission, current *models.ResourcePermission) (map[string]string, int) {
	changes := make(map[string]string)
	if current.OrganizationID.UUID != resourcePermission.OrganizationID.UUID {
		changes["organization_id"] = ":organization_id"
	}
	if current.ResourceID.UUID != resourcePermission.ResourceID.UUID {
		changes["resource_id"] = ":resource_id"
	}
	if current.PermissionID.UUID != resourcePermission.PermissionID.UUID {
		changes["permission_id"] = ":permission_id"
	}
	if current.Name.String != resourcePermission.Name.String {
		changes["name"] = ":name"
	}
	if current.Description.String != resourcePermission.Description.String {
		changes["description"] = ":description"
	}
	if current.IsActive.Bool != resourcePermission.IsActive.Bool {
		changes["is_active"] = ":is_active"
	}
	if current.IsLogicalDeleted.Bool != resourcePermission.IsLogicalDeleted.Bool {
		changes["is_logical_deleted"] = ":is_logical_deleted"
	}
	if current.CreatedBy.UUID != resourcePermission.CreatedBy.UUID {
		changes["created_by_id"] = ":created_by_id"
	}
	if current.UpdatedBy.UUID != resourcePermission.UpdatedBy.UUID {
		changes["updated_by_id"] = ":updated_by_id"
	}
	if !current.CreatedAt.Time.Equal(resourcePermission.CreatedAt.Time) {
		changes["created_at"] = ":created_at"
	}
	if !current.UpdatedAt.Time.Equal(resourcePermission.UpdatedAt.Time) {
		changes["updated_at"] = ":updated_at"
	}
	return changes, len(changes)
}
