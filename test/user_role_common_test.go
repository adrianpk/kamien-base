package test

import (
	"fmt"

	mdl "github.com/adrianpk/kamien/models"
	"{{.Package}}/models"
	"{{.Package}}/repo"
	uuid "github.com/satori/go.uuid"
)

func createSampleUserRole() *models.UserRole {
	u := createUserRole("name", "")
	return u
}

func createSampleUserRole2() *models.UserRole {
	u := createUserRole("name", "2")
	return u
}

func createUserRole(name, sufix string) *models.UserRole {
	tx := tenv.GetDBxTx()
	userRoleRepo := repo.MakeUserRoleRepoTx(tx)
	userRole := models.MakeUserRole()
	ns := fmt.Sprintf("Name%s", sufix)
	userRole.Name = mdl.ToNullsString(ns)
	userRoleRepo.Create(userRole)
	err := tx.Commit()
	checkErr(err, "Cannot create sample userRole.")
	return userRole
}

func getUserRole(id uuid.UUID) *models.UserRole {
	tx := tenv.GetDBxTx()
	repo := repo.MakeUserRoleRepoTx(tx)
	userRole, err := repo.Get(id)
	checkErr(err, "Cannot get userRole.")
	err = tx.Commit()
	checkErr(err, "Cannot get userRole.")
	return userRole
}

func getUserRoleByName(name string) *models.UserRole {
	tx := tenv.GetDBxTx()
	repo := repo.MakeUserRoleRepoTx(tx)
	userRole, err := repo.GetByName(name)
	checkErr(err, "Cannot get userRole.")
	err = tx.Commit()
	checkErr(err, "Cannot get userRole.")
	return userRole
}

func checkUserRole(id uuid.UUID) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeUserRoleRepoTx(tx)
	userRole, err := repo.Get(id)
	logErr(err, "Cannot get userRole.")
	err = tx.Commit()
	logErr(err, "Cannot get userRole.")
	return userRole != nil && err == nil
}

func checkUserRoleByName(name string) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeUserRoleRepoTx(tx)
	userRole, err := repo.GetByName(name)
	logErr(err, "Cannot get userRole.")
	err = tx.Commit()
	logErr(err, "Cannot get userRole.")
	return userRole != nil && err == nil
}

func userRolesMatch(userRole, tc *models.UserRole) bool {
	return userRole.Match(tc)
}
