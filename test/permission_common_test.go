package test

import (
	"fmt"

	mdl "github.com/adrianpk/kamien/models"
	"{{.Package}}/models"
	"{{.Package}}/repo"
	uuid "github.com/satori/go.uuid"
)

func createSamplePermission() *models.Permission {
	u := createPermission("name", "")
	return u
}

func createSamplePermission2() *models.Permission {
	u := createPermission("name", "2")
	return u
}

func createPermission(name, sufix string) *models.Permission {
	tx := tenv.GetDBxTx()
	permissionRepo := repo.MakePermissionRepoTx(tx)
	permission := models.MakePermission()
	ns := fmt.Sprintf("Name%s", sufix)
	permission.Name = mdl.ToNullsString(ns)
	permissionRepo.Create(permission)
	err := tx.Commit()
	checkErr(err, "Cannot create sample permission.")
	return permission
}

func getPermission(id uuid.UUID) *models.Permission {
	tx := tenv.GetDBxTx()
	repo := repo.MakePermissionRepoTx(tx)
	permission, err := repo.Get(id)
	checkErr(err, "Cannot get permission.")
	err = tx.Commit()
	checkErr(err, "Cannot get permission.")
	return permission
}

func getPermissionByName(name string) *models.Permission {
	tx := tenv.GetDBxTx()
	repo := repo.MakePermissionRepoTx(tx)
	permission, err := repo.GetByName(name)
	checkErr(err, "Cannot get permission.")
	err = tx.Commit()
	checkErr(err, "Cannot get permission.")
	return permission
}

func checkPermission(id uuid.UUID) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakePermissionRepoTx(tx)
	permission, err := repo.Get(id)
	logErr(err, "Cannot get permission.")
	err = tx.Commit()
	logErr(err, "Cannot get permission.")
	return permission != nil && err == nil
}

func checkPermissionByName(name string) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakePermissionRepoTx(tx)
	permission, err := repo.GetByName(name)
	logErr(err, "Cannot get permission.")
	err = tx.Commit()
	logErr(err, "Cannot get permission.")
	return permission != nil && err == nil
}

func permissionsMatch(permission, tc *models.Permission) bool {
	return permission.Match(tc)
}
