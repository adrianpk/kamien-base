package services

import (
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"github.com/jmoiron/sqlx"
)

// ResourceService - Resource service.
type ResourceService struct {
	Tx *sqlx.Tx
}

// Create - Create a Resource.
// Do additional related tasks if needed.
func (resourceSvc *ResourceService) Create(resource *models.Resource) (*models.Resource, error) {
	tx := resourceSvc.Tx
	resourceRepo := repo.MakeResourceRepoTx(tx)
	// Set values
	resource.GenerateID()
	resource.GenerateTag()
	resourceRepo.Create(resource)
	// Result
	return resource, nil
}

// Update - Update a Resource.
// Do additional related tasks if needed.
func (resourceSvc *ResourceService) Update(resource *models.Resource) error {
	tx := resourceSvc.Tx
	resourceRepo := repo.MakeResourceRepoTx(tx)
	err := resourceRepo.Update(resource)
	// Result
	return err
}
