package repo

import (
	"bytes"
	"fmt"

	"{{.Package}}/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Initialize pq.
	uuid "github.com/satori/go.uuid"
)

// ResourceRepository - Resource repository manager.
type ResourceRepository struct {
	DB *sqlx.DB
	Tx *sqlx.Tx
}

// MakeResourceRepoTx - ResourceRepo constructor.
func MakeResourceRepoTx(tx *sqlx.Tx) ResourceRepository {
	return ResourceRepository{Tx: tx}
}

// GetAll - Get all Resources from repo.
func (repo *ResourceRepository) GetAll() ([]models.Resource, error) {
	Resources := []models.Resource{}
	err := repo.Tx.Select(&Resources, "SELECT * FROM resources;")
	return Resources, err
}

// Get - Retrive a Resource from repo by its ID.
func (repo *ResourceRepository) Get(id uuid.UUID) (*models.Resource, error) {
	Resource := models.Resource{}
	err := repo.Tx.Get(&Resource, "SELECT * FROM resources WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	return &Resource, nil
}

// GetByName - Retrive a Resource from repo by its Name.
func (repo *ResourceRepository) GetByName(name string) (*models.Resource, error) {
	Resource := models.Resource{}
	err := repo.Tx.Get(&Resource, "SELECT * FROM resources WHERE name = $1;", name)
	if err != nil {
		return nil, err
	}
	return &Resource, nil
}

// Create - Create a Resource into the repo.
func (repo *ResourceRepository) Create(resource *models.Resource) error {
	resource.GenerateID()
	resource.SetCreateValues()
	resource.CreatedBy = resource.ID
	resource.UpdatedBy = resource.ID
	resourceInsertSQL := "INSERT INTO resources (id, tag, organization_id, name, description, is_active, is_logical_deleted, created_by_id, updated_by_id, created_at, updated_at) VALUES (:id, :tag, :organization_id, :name, :description, :is_active, :is_logical_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at);"
	_, err := repo.Tx.NamedExec(resourceInsertSQL, resource)
	return err
}

// Update - Update a Resource in repo.
func (repo *ResourceRepository) Update(resource *models.Resource) error {
	resource.SetUpdateValues()
	current, err := repo.Get(resource.ID.UUID)
	if err != nil {
		return err
	}
	// Build update statement
	changes, qty := ResourceChanges(resource, current)
	pos := 0
	last := qty < 2
	var query bytes.Buffer
	query.WriteString("UPDATE resources SET ")
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
	query.WriteString(fmt.Sprintf("WHERE id = '%s';", resource.ID.UUID))
	repo.Tx.NamedExec(query.String(), &resource)
	// log.Debugf("Update resource query:\n%s", query.String())
	return err
}

// Delete - Delete a Resource from repo.
func (repo *ResourceRepository) Delete(id uuid.UUID) {
	ResourceDeleteSQL := fmt.Sprintf("DELETE FROM resources WHERE id = '%s'", id)
	_ = repo.Tx.MustExec(ResourceDeleteSQL)
}

// ResourceChanges - Creates a map ([string]interface{}) including al changing field.
func ResourceChanges(resource, current *models.Resource) (map[string]string, int) {
	changes := make(map[string]string)
	if current.Tag.String != resource.Tag.String {
		changes["tag"] = ":tag"
	}
	if current.OrganizationID.UUID != resource.OrganizationID.UUID {
		changes["organization_id"] = ":organization_id"
	}
	if current.Name.String != resource.Name.String {
		changes["name"] = ":name"
	}
	if current.Description.String != resource.Description.String {
		changes["description"] = ":description"
	}
	if current.IsActive.Bool != resource.IsActive.Bool {
		changes["is_active"] = ":is_active"
	}
	if current.IsLogicalDeleted.Bool != resource.IsLogicalDeleted.Bool {
		changes["is_logical_deleted"] = ":is_logical_deleted"
	}
	if current.CreatedBy.UUID != resource.CreatedBy.UUID {
		changes["created_by_id"] = ":created_by_id"
	}
	if current.UpdatedBy.UUID != resource.UpdatedBy.UUID {
		changes["updated_by_id"] = ":updated_by_id"
	}
	if !current.CreatedAt.Time.Equal(resource.CreatedAt.Time) {
		changes["created_at"] = ":created_at"
	}
	if !current.UpdatedAt.Time.Equal(resource.UpdatedAt.Time) {
		changes["updated_at"] = ":updated_at"
	}
	return changes, len(changes)
}
