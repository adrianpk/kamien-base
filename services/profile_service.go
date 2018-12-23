package services

import (
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"github.com/jmoiron/sqlx"
)

// ProfileService - Profile service.
type ProfileService struct {
	Tx *sqlx.Tx
}

// Create - Create a Profile.
// Do additional related tasks if needed.
func (profileSvc *ProfileService) Create(profile *models.Profile) (*models.Profile, error) {
	tx := profileSvc.Tx
	profileRepo := repo.MakeProfileRepoTx(tx)
	profileRepo.Create(profile)
	// Result
	return profile, nil
}

// Update - Update a Profile.
// Do additional related tasks if needed.
func (profileSvc *ProfileService) Update(profile *models.Profile) error {
	tx := profileSvc.Tx
	profileRepo := repo.MakeProfileRepoTx(tx)
	err := profileRepo.Update(profile)
	// Result
	return err
}
