package repo

import (
	"bytes"
	"fmt"
	"strings"

	r "github.com/adrianpk/kamien/repo"
	"{{.Package}}/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Initialize pq.
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository - User repository manager.
type UserRepository struct {
	DB *sqlx.DB
	Tx *sqlx.Tx
}

// MakeUserRepoTx - UserRepo constructor.
func MakeUserRepoTx(tx *sqlx.Tx) UserRepository {
	return UserRepository{Tx: tx}
}

// MakeUserRepoDB - UserRepo constructor.
func MakeUserRepoDB(db *sqlx.DB) UserRepository {
	return UserRepository{DB: db}
}

// GetAll - Get all users from repo.
func (repo *UserRepository) GetAll() ([]models.User, error) {
	users := []models.User{}
	err := repo.Tx.Select(&users, "SELECT * FROM users;")
	return users, err
}

// Get - Retrive a user from repo by its ID.
func (repo *UserRepository) Get(id uuid.UUID) (*models.User, error) {
	user := models.User{}
	err := repo.Tx.Get(&user, "SELECT * FROM users WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername - Retrive a user from repo by its username.
func (repo *UserRepository) GetByUsername(username string) (*models.User, error) {
	user := models.User{}
	err := repo.Tx.Get(&user, "SELECT * FROM users WHERE username = $1;", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create - Create a user into the repo.
func (repo *UserRepository) Create(user *models.User) error {
	user.GenerateID()
	user.UpdatePasswordHash()
	user.SetCreateValues()
	user.CreatedBy = user.ID
	user.UpdatedBy = user.ID
	userInsertSQL := "INSERT INTO users (id, username, password_hash, email, given_name, middle_names, family_name, context_id, geolocation, starts_at, ends_at, is_active, is_logical_deleted, created_by_id, updated_by_id, created_at, updated_at) VALUES (:id, :username, :password_hash, :email, :given_name, :middle_names, :family_name, :context_id, :geolocation, :starts_at, :ends_at, :is_active, :is_logical_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at);"
	_, err := repo.Tx.NamedExec(userInsertSQL, user)
	return err
}

// Update - Update a user in repo.
func (repo *UserRepository) Update(user *models.User) error {
	user.SetUpdateValues()
	user.UpdatePasswordHash()
	reference, err := repo.Get(user.ID.UUID)
	if err != nil {
		return err
	}
	// Build update statement
	var st bytes.Buffer
	changes := UserChanges(user, reference)
	st.WriteString("UPDATE users SET ")
	st.WriteString(changes)
	st.WriteString(fmt.Sprintf(" WHERE id = '%s';", user.ID.UUID))
	_, err = repo.Tx.NamedExec(st.String(), &user)
	// fmt.Printf("\nUpdate user query:\n%s\n", st.String())
	return err
}

// Delete - Delete a user from repo.
func (repo *UserRepository) Delete(id uuid.UUID) {
	userDeleteSQL := fmt.Sprintf("DELETE FROM users WHERE id = '%s'", id)
	_ = repo.Tx.MustExec(userDeleteSQL)
}

// UserChanges - Creates a map ([string]interface{}) including al changing field.
func UserChanges(user *models.User, current *models.User) string {
	comma := " " // Add comma
	var st bytes.Buffer
	if current.Username.String != user.Username.String {
		st.WriteString(fmt.Sprintf("username = '%s'", user.Username.String))
		comma = ", "
	}
	if strings.Trim(user.Password, "") != "" && current.PasswordHash != user.PasswordHash {
		st.WriteString(fmt.Sprintf("%spassword_hash = '%s'", comma, user.PasswordHash))
		comma = ", "
	}
	if current.Email.String != user.Email.String {
		st.WriteString(fmt.Sprintf("%semail = '%s'", comma, user.Email.String))
		comma = ", "
	}
	if current.GivenName.String != user.GivenName.String {
		st.WriteString(fmt.Sprintf("%sgiven_name = '%s'", comma, user.GivenName.String))
		comma = ", "
	}
	if current.MiddleNames.String != user.MiddleNames.String {
		st.WriteString(fmt.Sprintf("%smiddle_names = '%s'", comma, user.MiddleNames.String))
		comma = ", "
	}
	if current.FamilyName.String != user.FamilyName.String {
		st.WriteString(fmt.Sprintf("%sfamily_name = '%s'", comma, user.FamilyName.String))
		comma = ", "
	}
	if current.ContextID.UUID != user.ContextID.UUID && user.ContextID.UUID != uuid.Nil {
		st.WriteString(fmt.Sprintf("%scontext_id = '%s'", comma, user.ContextID.UUID))
		comma = ", "
	}
	if true || current.Geolocation.Point.String() != user.Geolocation.Point.String() {
		g := fmt.Sprintf("ST_SetSRID(ST_MakePoint(%f, %f), 4326)",
			user.Geolocation.Point.Lat,
			user.Geolocation.Point.Lng)
		st.WriteString(fmt.Sprintf("%sgeolocation = %s", comma, g))
		comma = ", "
	}
	if current.StartsAt.Time != user.StartsAt.Time {
		fd := r.FormatDate(user.StartsAt)
		st.WriteString(fmt.Sprintf("%sstarts_at = '%s'", comma, fd))
		comma = ", "
	}
	if current.EndsAt.Time != user.EndsAt.Time {
		fd := r.FormatDate(user.EndsAt)
		st.WriteString(fmt.Sprintf("%sends_at = '%s'", comma, fd))
		comma = ", "
	}
	if current.IsActive.Bool != user.IsActive.Bool {
		st.WriteString(fmt.Sprintf("%sis_active = %t", comma, user.IsActive.Bool))
		comma = ", "
	}
	if current.IsLogicalDeleted.Bool != user.IsLogicalDeleted.Bool {
		st.WriteString(fmt.Sprintf("%sis_logical_delete = %t", comma, user.IsLogicalDeleted.Bool))
		comma = ", "
	}
	if current.UpdatedBy.UUID != user.UpdatedBy.UUID && user.UpdatedBy.UUID != uuid.Nil {
		st.WriteString(fmt.Sprintf("%supdated_by_id = '%s'", comma, user.UpdatedBy.UUID.String()))
		comma = ", "
	}
	if current.UpdatedAt.Time != user.UpdatedAt.Time {
		fd := r.FormatDate(user.UpdatedAt)
		st.WriteString(fmt.Sprintf("%supdated_at = '%s'", comma, fd))
	}
	return st.String()
}

// Login - Retrive a User if username/email and provided
func (repo *UserRepository) Login(user *models.User) (*models.User, error) {
	u := &models.User{}
	err := repo.Tx.Get(&u, "SELECT * FROM users WHERE username = $1 OR email=$2 LIMIT 1;", user.Username, user.Email)
	if err != nil {
		return user, err
	}
	// Validate password
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(user.Password))
	if err != nil {
		return user, err
	}
	return u, nil
}
