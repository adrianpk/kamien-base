package test

import (
	"fmt"

	mdl "github.com/adrianpk/kamien/models"
	"{{.Package}}/models"
	"{{.Package}}/repo"
	uuid "github.com/satori/go.uuid"
)

func createSampleResourcePermission() *models.ResourcePermission {
	u := createResourcePermission("name", "")
	return u
}

func createSampleResourcePermission2() *models.ResourcePermission {
	u := createResourcePermission("name", "2")
	return u
}

func createResourcePermission(name, sufix string) *models.ResourcePermission {
	tx := tenv.GetDBxTx()
	resourcePermissionRepo := repo.MakeResourcePermissionRepoTx(tx)
	resourcePermission := models.MakeResourcePermission()
	ns := fmt.Sprintf("Name%s", sufix)
	resourcePermission.Name = mdl.ToNullsString(ns)
	resourcePermissionRepo.Create(resourcePermission)
	err := tx.Commit()
	checkErr(err, "Cannot create sample resourcePermission.")
	return resourcePermission
}

func getResourcePermission(id uuid.UUID) *models.ResourcePermission {
	tx := tenv.GetDBxTx()
	repo := repo.MakeResourcePermissionRepoTx(tx)
	resourcePermission, err := repo.Get(id)
	checkErr(err, "Cannot get resourcePermission.")
	err = tx.Commit()
	checkErr(err, "Cannot get resourcePermission.")
	return resourcePermission
}

func getResourcePermissionByName(name string) *models.ResourcePermission {
	tx := tenv.GetDBxTx()
	repo := repo.MakeResourcePermissionRepoTx(tx)
	resourcePermission, err := repo.GetByName(name)
	checkErr(err, "Cannot get resourcePermission.")
	err = tx.Commit()
	checkErr(err, "Cannot get resourcePermission.")
	return resourcePermission
}

func checkResourcePermission(id uuid.UUID) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeResourcePermissionRepoTx(tx)
	resourcePermission, err := repo.Get(id)
	logErr(err, "Cannot get resourcePermission.")
	err = tx.Commit()
	logErr(err, "Cannot get resourcePermission.")
	return resourcePermission != nil && err == nil
}

func checkResourcePermissionByName(name string) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeResourcePermissionRepoTx(tx)
	resourcePermission, err := repo.GetByName(name)
	logErr(err, "Cannot get resourcePermission.")
	err = tx.Commit()
	logErr(err, "Cannot get resourcePermission.")
	return resourcePermission != nil && err == nil
}

func resourcePermissionsMatch(resourcePermission, tc *models.ResourcePermission) bool {
	return resourcePermission.Match(tc)
}
