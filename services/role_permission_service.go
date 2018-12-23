package services

import (
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"github.com/jmoiron/sqlx"
)

// RolePermissionService - RolePermission service.
type RolePermissionService struct {
	Tx *sqlx.Tx
}

// Create - Create a RolePermission.
// Do additional related tasks if needed.
func (rolePermissionSvc *RolePermissionService) Create(rolePermission *models.RolePermission) (*models.RolePermission, error) {
	tx := rolePermissionSvc.Tx
	rolePermissionRepo := repo.MakeRolePermissionRepoTx(tx)
	rolePermissionRepo.Create(rolePermission)
	// Result
	return rolePermission, nil
}

// Update - Update a RolePermission.
// Do additional related tasks if needed.
func (rolePermissionSvc *RolePermissionService) Update(rolePermission *models.RolePermission) error {
	tx := rolePermissionSvc.Tx
	rolePermissionRepo := repo.MakeRolePermissionRepoTx(tx)
	err := rolePermissionRepo.Update(rolePermission)
	// Result
	return err
}
