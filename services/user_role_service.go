package services

import (
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"github.com/jmoiron/sqlx"
)

// UserRoleService - UserRole service.
type UserRoleService struct {
	Tx *sqlx.Tx
}

// Create - Create a UserRole.
// Do additional related tasks if needed.
func (userRoleSvc *UserRoleService) Create(userRole *models.UserRole) (*models.UserRole, error) {
	tx := userRoleSvc.Tx
	userRoleRepo := repo.MakeUserRoleRepoTx(tx)
	userRoleRepo.Create(userRole)
	// Result
	return userRole, nil
}

// Update - Update a UserRole.
// Do additional related tasks if needed.
func (userRoleSvc *UserRoleService) Update(userRole *models.UserRole) error {
	tx := userRoleSvc.Tx
	userRoleRepo := repo.MakeUserRoleRepoTx(tx)
	err := userRoleRepo.Update(userRole)
	// Result
	return err
}
