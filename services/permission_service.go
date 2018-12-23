package services

import (
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"github.com/jmoiron/sqlx"
)

// PermissionService - Permission service.
type PermissionService struct {
	Tx *sqlx.Tx
}

// Create - Create a Permission.
// Do additional related tasks if needed.
func (permissionSvc *PermissionService) Create(permission *models.Permission) (*models.Permission, error) {
	tx := permissionSvc.Tx
	permissionRepo := repo.MakePermissionRepoTx(tx)
	permissionRepo.Create(permission)
	// Result
	return permission, nil
}

// Update - Update a Permission.
// Do additional related tasks if needed.
func (permissionSvc *PermissionService) Update(permission *models.Permission) error {
	tx := permissionSvc.Tx
	permissionRepo := repo.MakePermissionRepoTx(tx)
	err := permissionRepo.Update(permission)
	// Result
	return err
}
