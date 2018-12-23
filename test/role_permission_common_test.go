package test

import (
	"fmt"

	mdl "github.com/adrianpk/kamien/models"
	"{{.Package}}/models"
	"{{.Package}}/repo"
	uuid "github.com/satori/go.uuid"
)

func createSampleRolePermission() *models.RolePermission {
	u := createRolePermission("name", "")
	return u
}

func createSampleRolePermission2() *models.RolePermission {
	u := createRolePermission("name", "2")
	return u
}

func createRolePermission(name, sufix string) *models.RolePermission {
	tx := tenv.GetDBxTx()
	rolePermissionRepo := repo.MakeRolePermissionRepoTx(tx)
	rolePermission := models.MakeRolePermission()
	ns := fmt.Sprintf("Name%s", sufix)
	rolePermission.Name = mdl.ToNullsString(ns)
	rolePermissionRepo.Create(rolePermission)
	err := tx.Commit()
	checkErr(err, "Cannot create sample rolePermission.")
	return rolePermission
}

func getRolePermission(id uuid.UUID) *models.RolePermission {
	tx := tenv.GetDBxTx()
	repo := repo.MakeRolePermissionRepoTx(tx)
	rolePermission, err := repo.Get(id)
	checkErr(err, "Cannot get rolePermission.")
	err = tx.Commit()
	checkErr(err, "Cannot get rolePermission.")
	return rolePermission
}

func getRolePermissionByName(name string) *models.RolePermission {
	tx := tenv.GetDBxTx()
	repo := repo.MakeRolePermissionRepoTx(tx)
	rolePermission, err := repo.GetByName(name)
	checkErr(err, "Cannot get rolePermission.")
	err = tx.Commit()
	checkErr(err, "Cannot get rolePermission.")
	return rolePermission
}

func checkRolePermission(id uuid.UUID) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeRolePermissionRepoTx(tx)
	rolePermission, err := repo.Get(id)
	logErr(err, "Cannot get rolePermission.")
	err = tx.Commit()
	logErr(err, "Cannot get rolePermission.")
	return rolePermission != nil && err == nil
}

func checkRolePermissionByName(name string) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeRolePermissionRepoTx(tx)
	rolePermission, err := repo.GetByName(name)
	logErr(err, "Cannot get rolePermission.")
	err = tx.Commit()
	logErr(err, "Cannot get rolePermission.")
	return rolePermission != nil && err == nil
}

func rolePermissionsMatch(rolePermission, tc *models.RolePermission) bool {
	return rolePermission.Match(tc)
}
