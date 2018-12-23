package services

import (
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"github.com/jmoiron/sqlx"
)

// UserService - Service used to create users
// and its associated models.
type UserService struct {
	Tx *sqlx.Tx
}

// Create - Create a User and its associated Account and Profile.
func (us *UserService) Create(user *models.User) (*models.User, error) {
	tx := us.Tx
	userRepo := repo.MakeUserRepoTx(tx)
	userRepo.Create(user)
	user.ClearPassword()
	user.GenerateID()
	// Associated Account
	accountRepo := repo.MakeAccountRepoTx(tx)
	var account = models.Account{}
	account.OwnerID = user.ID
	account.Name = user.Username
	account.Email = user.Email
	accountRepo.Create(&account)
	// Associated Account Profile
	profileRepo := repo.MakeProfileRepoTx(tx)
	var profile = models.Profile{}
	profile.Name = user.Username
	profile.Email = user.Email
	profile.OwnerID = account.ID
	profileRepo.Create(&profile)
	// Result
	return user, nil
}

// Login - Create a User and its associated Account and Profile.
func (us *UserService) Login(user *models.User) (*models.User, error) {
	tx := us.Tx
	userRepo := repo.MakeUserRepoTx(tx)
	return userRepo.Login(user)
}
