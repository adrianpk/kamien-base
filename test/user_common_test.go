package test

import (
	"fmt"

	"{{.Package}}/models"
	"{{.Package}}/repo"
	uuid "github.com/satori/go.uuid"
)

func createSampleAdmin() *models.User {
	u := createUser("admin", "password", "admin@gmail.com", "")
	return u
}

func createSampleUser() *models.User {
	u := createUser("username", "password", "username@gmail.com", "")
	return u
}

func createSampleUser2() *models.User {
	u := createUser("username2", "password2", "username2@gmail.com", "2")
	return u
}

func createUser(username, password, email, sufix string) *models.User {
	tx := tenv.GetDBxTx()
	userRepo := repo.MakeUserRepoTx(tx)
	user := models.MakeUserUPE(username, password, password, email)
	name := fmt.Sprintf("Name%s", sufix)
	middles := fmt.Sprintf("Middlenames%s", sufix)
	family := fmt.Sprintf("Family%s", sufix)
	user.SetNames(name, middles, family)
	user.PairContextID()
	userRepo.Create(user)
	// user.ClearPassword()
	err := tx.Commit()
	checkErr(err, "Cannot create sample user.")
	return user
}

func getUser(id uuid.UUID) *models.User {
	tx := tenv.GetDBxTx()
	repo := repo.MakeUserRepoTx(tx)
	user, err := repo.Get(id)
	checkErr(err, "Cannot get user.")
	err = tx.Commit()
	checkErr(err, "Cannot get user.")
	return user
}

func getUserByUsername(username string) *models.User {
	tx := tenv.GetDBxTx()
	repo := repo.MakeUserRepoTx(tx)
	user, err := repo.GetByUsername(username)
	checkErr(err, "Cannot get user.")
	err = tx.Commit()
	checkErr(err, "Cannot get user.")
	return user
}

func checkUser(id uuid.UUID) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeUserRepoTx(tx)
	user, err := repo.Get(id)
	logErr(err, "Cannot get user.")
	err = tx.Commit()
	logErr(err, "Cannot get user.")
	return user != nil && err == nil
}

func checkUserByUsername(username string) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeUserRepoTx(tx)
	user, err := repo.GetByUsername(username)
	logErr(err, "Cannot get user.")
	err = tx.Commit()
	logErr(err, "Cannot get user.")
	return user != nil && err == nil
}

func usersMatch(user, tc *models.User) bool {
	return user.Match(tc)
}
