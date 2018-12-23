package test

import (
	"fmt"

	mdl "github.com/adrianpk/kamien/models"
	"{{.Package}}/models"
	"{{.Package}}/repo"
	uuid "github.com/satori/go.uuid"
)

func createSampleProfile() *models.Profile {
	u := createProfile("name", "")
	return u
}

func createSampleProfile2() *models.Profile {
	u := createProfile("name", "2")
	return u
}

func createProfile(name, sufix string) *models.Profile {
	tx := tenv.GetDBxTx()
	profileRepo := repo.MakeProfileRepoTx(tx)
	profile := models.MakeProfile()
	ns := fmt.Sprintf("Name%s", sufix)
	profile.Name = mdl.ToNullsString(ns)
	profileRepo.Create(profile)
	err := tx.Commit()
	checkErr(err, "Cannot create sample profile.")
	return profile
}

func getProfile(id uuid.UUID) *models.Profile {
	tx := tenv.GetDBxTx()
	repo := repo.MakeProfileRepoTx(tx)
	profile, err := repo.Get(id)
	checkErr(err, "Cannot get profile.")
	err = tx.Commit()
	checkErr(err, "Cannot get profile.")
	return profile
}

func getProfileByName(name string) *models.Profile {
	tx := tenv.GetDBxTx()
	repo := repo.MakeProfileRepoTx(tx)
	profile, err := repo.GetByName(name)
	checkErr(err, "Cannot get profile.")
	err = tx.Commit()
	checkErr(err, "Cannot get profile.")
	return profile
}

func getProfileByOwnerID(ownerID uuid.UUID) *models.Profile {
	tx := tenv.GetDBxTx()
	repo := repo.MakeProfileRepoTx(tx)
	account, err := repo.GetByOwnerID(ownerID)
	checkErr(err, "Cannot get profile.")
	err = tx.Commit()
	checkErr(err, "Cannot get profile.")
	return account
}

func checkProfile(id uuid.UUID) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeProfileRepoTx(tx)
	profile, err := repo.Get(id)
	logErr(err, "Cannot get profile.")
	err = tx.Commit()
	logErr(err, "Cannot get profile.")
	return profile != nil && err == nil
}

func checkProfileByName(name string) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeProfileRepoTx(tx)
	profile, err := repo.GetByName(name)
	logErr(err, "Cannot get profile.")
	err = tx.Commit()
	logErr(err, "Cannot get profile.")
	return profile != nil && err == nil
}

func checkProfileByOwnerID(ownerID uuid.UUID) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeProfileRepoTx(tx)
	profile, err := repo.GetByOwnerID(ownerID)
	logErr(err, "Cannot get profile.")
	err = tx.Commit()
	logErr(err, "Cannot get profile.")
	return profile != nil && err == nil
}

func profilesMatch(profile, tc *models.Profile) bool {
	return profile.Match(tc)
}
