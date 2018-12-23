package repo

import (
	"bytes"
	"fmt"

	"{{.Package}}/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Initialize pq.
	uuid "github.com/satori/go.uuid"
)

// AccountRepository - Account repository manager.
type AccountRepository struct {
	DB *sqlx.DB
	Tx *sqlx.Tx
}

// MakeAccountRepoTx - AccountRepo constructor.
func MakeAccountRepoTx(tx *sqlx.Tx) AccountRepository {
	return AccountRepository{Tx: tx}
}

// GetAll - Get all accounts from repo.
func (repo *AccountRepository) GetAll() ([]models.Account, error) {
	accounts := []models.Account{}
	err := repo.Tx.Select(&accounts, "SELECT * FROM accounts;")
	return accounts, err
}

// Get - Retrive a Account from repo by its ID.
func (repo *AccountRepository) Get(id uuid.UUID) (*models.Account, error) {
	account := models.Account{}
	err := repo.Tx.Get(&account, "SELECT * FROM accounts WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetByName - Retrive a Account from repo by its name.
func (repo *AccountRepository) GetByName(name string) (*models.Account, error) {
	account := models.Account{}
	err := repo.Tx.Get(&account, "SELECT * FROM accounts WHERE name = $1;", name)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetByOwnerID - Retrive an Account from repo by its owner ID.
func (repo *AccountRepository) GetByOwnerID(id uuid.UUID) (*models.Account, error) {
	account := models.Account{}
	err := repo.Tx.Get(&account, "SELECT * FROM accounts WHERE owner_id = $1;", id)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetByParentID - Retrive a Account from repo by its parent ID.
func (repo *AccountRepository) GetByParentID(id uuid.UUID) (*models.Account, error) {
	account := models.Account{}
	err := repo.Tx.Get(&account, "SELECT * FROM accounts WHERE parent_id = $1;", id)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// Create - Create a account into the repo.
func (repo *AccountRepository) Create(account *models.Account) error {
	account.GenerateID()
	account.SetCreateValues()
	account.CreatedBy = account.ID
	account.UpdatedBy = account.ID
	accountInsertSQL := "INSERT INTO accounts (id, name, email, owner_id, parent_id, geolocation, starts_at, ends_at, is_active, is_logical_deleted, created_by_id, updated_by_id, created_at, updated_at) VALUES (:id, :name, :email, :owner_id, :parent_id, :geolocation, :starts_at, :ends_at, :is_active, :is_logical_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at);"
	_, err := repo.Tx.NamedExec(accountInsertSQL, account)
	return err
}

// Update - Update a account in repo.
func (repo *AccountRepository) Update(account *models.Account) error {
	account.SetUpdateValues()
	current, err := repo.Get(account.ID.UUID)
	if err != nil {
		return err
	}
	// Build update statement
	changes, qty := AccountChanges(account, current)
	pos := 0
	last := qty < 2
	var query bytes.Buffer
	query.WriteString("UPDATE accounts SET ")
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
	query.WriteString(fmt.Sprintf("WHERE id = '%s';", account.ID.UUID))
	repo.Tx.NamedExec(query.String(), &account)
	// log.Debugf("Update account query:\n%s", query.String())
	return err
}

// Delete - Delete a Account from repo.
func (repo *AccountRepository) Delete(id uuid.UUID) {
	accountDeleteSQL := fmt.Sprintf("DELETE FROM accounts WHERE id = '%s'", id)
	_ = repo.Tx.MustExec(accountDeleteSQL)
}

// AccountChanges - Creates a map ([string]interface{}) including al changing field.
func AccountChanges(account, current *models.Account) (map[string]string, int) {
	changes := make(map[string]string)
	if current.Name.String != account.Name.String {
		changes["name"] = ":name"
	}
	if current.AccountType.String != account.AccountType.String {
		changes["account_type"] = ":account_type"
	}
	if current.OwnerID.UUID != account.OwnerID.UUID {
		changes["owner_id"] = ":owner_id"
	}
	if current.ParentID.UUID != account.ParentID.UUID {
		changes["parent_id"] = ":parent_id"
	}
	if current.Email.String != account.Email.String {
		changes["email"] = ":email"
	}
	if current.Geolocation.Point.String() != account.Geolocation.Point.String() {
		changes["geolocation"] = ":geolocation"
	}
	if !current.StartsAt.Time.Equal(account.StartsAt.Time) {
		changes["starts_at"] = ":starts_at"
	}
	if !current.EndsAt.Time.Equal(account.EndsAt.Time) {
		changes["ends_at"] = ":ends_at"
	}
	if current.IsActive.Bool != account.IsActive.Bool {
		changes["is_active"] = ":is_active"
	}
	if current.IsLogicalDeleted.Bool != account.IsLogicalDeleted.Bool {
		changes["is_logical_deleted"] = ":is_logical_deleted"
	}
	if !current.CreatedAt.Time.Equal(account.CreatedAt.Time) {
		changes["created_at"] = ":created_at"
	}
	if !current.UpdatedAt.Time.Equal(account.UpdatedAt.Time) {
		changes["updated_at"] = ":updated_at"
	}
	return changes, len(changes)
}
