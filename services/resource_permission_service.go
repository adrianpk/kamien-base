package services

import (
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"github.com/jmoiron/sqlx"
)

// ResourcePermissionService - ResourcePermission service.
type ResourcePermissionService struct {
	Tx *sqlx.Tx
}

// Create - Create a ResourcePermission.
// Do additional related tasks if needed.
func (resourcePermissionSvc *ResourcePermissionService) Create(resourcePermission *models.ResourcePermission) (*models.ResourcePermission, error) {
	tx := resourcePermissionSvc.Tx
	resourcePermissionRepo := repo.MakeResourcePermissionRepoTx(tx)
	resourcePermissionRepo.Create(resourcePermission)
	// Result
	return resourcePermission, nil
}

// Update - Update a ResourcePermission.
// Do additional related tasks if needed.
func (resourcePermissionSvc *ResourcePermissionService) Update(resourcePermission *models.ResourcePermission) error {
	tx := resourcePermissionSvc.Tx
	resourcePermissionRepo := repo.MakeResourcePermissionRepoTx(tx)
	err := resourcePermissionRepo.Update(resourcePermission)
	// Result
	return err
}
