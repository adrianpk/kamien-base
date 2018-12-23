package services

import (
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"github.com/jmoiron/sqlx"
)

// AccountService - Service used to create accounts
// and its associated models.
type AccountService struct {
	Tx *sqlx.Tx
}

// Create - Create a Account and its associated Account and Profile.
func (us *AccountService) Create(account *models.Account) (*models.Account, error) {
	tx := us.Tx
	accountRepo := repo.MakeAccountRepoTx(tx)
	accountRepo.Create(account)
	// Result
	return account, nil
}
