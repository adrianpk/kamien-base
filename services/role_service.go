package services

import (
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"github.com/jmoiron/sqlx"
)

// RoleService - Role service.
type RoleService struct {
	Tx *sqlx.Tx
}

// Create - Create a Role.
// Do additional related tasks if needed.
func (roleSvc *RoleService) Create(role *models.Role) (*models.Role, error) {
	tx := roleSvc.Tx
	roleRepo := repo.MakeRoleRepoTx(tx)
	// Set values
	role.GenerateID()
	roleRepo.Create(role)
	// Result
	return role, nil
}

// Update - Update a Role.
// Do additional related tasks if needed.
func (roleSvc *RoleService) Update(role *models.Role) error {
	tx := roleSvc.Tx
	roleRepo := repo.MakeRoleRepoTx(tx)
	err := roleRepo.Update(role)
	// Result
	return err
}
