package repo

import (
	"bytes"
	"fmt"

	"{{.Package}}/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Initialize pq.
	uuid "github.com/satori/go.uuid"
)

// UserRoleRepository - UserRole repository manager.
type UserRoleRepository struct {
	DB *sqlx.DB
	Tx *sqlx.Tx
}

// MakeUserRoleRepoTx - UserRoleRepo constructor.
func MakeUserRoleRepoTx(tx *sqlx.Tx) UserRoleRepository {
	return UserRoleRepository{Tx: tx}
}

// GetAll - Get all UserRoles from repo.
func (repo *UserRoleRepository) GetAll() ([]models.UserRole, error) {
	UserRoles := []models.UserRole{}
	err := repo.Tx.Select(&UserRoles, "SELECT * FROM user_roles;")
	return UserRoles, err
}

// Get - Retrive a UserRole from repo by its ID.
func (repo *UserRoleRepository) Get(id uuid.UUID) (*models.UserRole, error) {
	UserRole := models.UserRole{}
	err := repo.Tx.Get(&UserRole, "SELECT * FROM user_roles WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	return &UserRole, nil
}

// GetByName - Retrive a UserRole from repo by its Name.
func (repo *UserRoleRepository) GetByName(name string) (*models.UserRole, error) {
	UserRole := models.UserRole{}
	err := repo.Tx.Get(&UserRole, "SELECT * FROM user_roles WHERE name = $1;", name)
	if err != nil {
		return nil, err
	}
	return &UserRole, nil
}

// Create - Create a UserRole into the repo.
func (repo *UserRoleRepository) Create(userRole *models.UserRole) error {
	userRole.GenerateID()
	userRole.SetCreateValues()
	userRole.CreatedBy = userRole.ID
	userRole.UpdatedBy = userRole.ID
	userRoleInsertSQL := "INSERT INTO user_roles (id, organization_id, user_id, role_id, name, description, is_active, is_logical_deleted, created_by_id, updated_by_id, created_at, updated_at) VALUES (:id, :organization_id, :user_id, :role_id, :name, :description, :is_active, :is_logical_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at);"
	_, err := repo.Tx.NamedExec(userRoleInsertSQL, userRole)
	return err
}

// Update - Update a UserRole in repo.
func (repo *UserRoleRepository) Update(userRole *models.UserRole) error {
	userRole.SetUpdateValues()
	current, err := repo.Get(userRole.ID.UUID)
	if err != nil {
		return err
	}
	// Build update statement
	changes, qty := UserRoleChanges(userRole, current)
	pos := 0
	last := qty < 2
	var query bytes.Buffer
	query.WriteString("UPDATE user_roles SET ")
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
	query.WriteString(fmt.Sprintf("WHERE id = '%s';", userRole.ID.UUID))
	repo.Tx.NamedExec(query.String(), &userRole)
	// log.Debugf("Update userRole query:\n%s", query.String())
	return err
}

// Delete - Delete a UserRole from repo.
func (repo *UserRoleRepository) Delete(id uuid.UUID) {
	UserRoleDeleteSQL := fmt.Sprintf("DELETE FROM user_roles WHERE id = '%s'", id)
	_ = repo.Tx.MustExec(UserRoleDeleteSQL)
}

// UserRoleChanges - Creates a map ([string]interface{}) including al changing field.
func UserRoleChanges(userRole, current *models.UserRole) (map[string]string, int) {
	changes := make(map[string]string)
	if current.OrganizationID.UUID != userRole.OrganizationID.UUID {
		changes["organization_id"] = ":organization_id"
	}
	if current.UserID.UUID != userRole.UserID.UUID {
		changes["user_id"] = ":user_id"
	}
	if current.RoleID.UUID != userRole.RoleID.UUID {
		changes["role_id"] = ":role_id"
	}
	if current.Name.String != userRole.Name.String {
		changes["name"] = ":name"
	}
	if current.Description.String != userRole.Description.String {
		changes["description"] = ":description"
	}
	if current.IsActive.Bool != userRole.IsActive.Bool {
		changes["is_active"] = ":is_active"
	}
	if current.IsLogicalDeleted.Bool != userRole.IsLogicalDeleted.Bool {
		changes["is_logical_deleted"] = ":is_logical_deleted"
	}
	if current.CreatedBy.UUID != userRole.CreatedBy.UUID {
		changes["created_by_id"] = ":created_by_id"
	}
	if current.UpdatedBy.UUID != userRole.UpdatedBy.UUID {
		changes["updated_by_id"] = ":updated_by_id"
	}
	if !current.CreatedAt.Time.Equal(userRole.CreatedAt.Time) {
		changes["created_at"] = ":created_at"
	}
	if !current.UpdatedAt.Time.Equal(userRole.UpdatedAt.Time) {
		changes["updated_at"] = ":updated_at"
	}
	return changes, len(changes)
}
