package test

import (
	mdl "github.com/adrianpk/kamien/models"
	"{{.Package}}/models"
	"{{.Package}}/repo"
	uuid "github.com/satori/go.uuid"
)

func createSampleAccount(accountType string, user *models.User, account *models.Account) *models.Account {
	u := createAccount(accountType, user, account)
	return u
}

func createAccount(accountType string, user *models.User, account *models.Account) *models.Account {
	tx := tenv.GetDBxTx()
	accountRepo := repo.MakeAccountRepoTx(tx)
	a := models.MakeAccount()
	a.AccountType = mdl.ToNullsString(accountType)
	a.OwnerID = user.ID
	if account != nil {
		a.ParentID = account.ID
	}
	accountRepo.Create(a)
	err := tx.Commit()
	checkErr(err, "Cannot create sample account.")
	return a
}

func getAccount(id uuid.UUID) *models.Account {
	tx := tenv.GetDBxTx()
	repo := repo.MakeAccountRepoTx(tx)
	account, err := repo.Get(id)
	checkErr(err, "Cannot get account.")
	err = tx.Commit()
	checkErr(err, "Cannot get account.")
	return account
}

func getAccountByName(name string) *models.Account {
	tx := tenv.GetDBxTx()
	repo := repo.MakeAccountRepoTx(tx)
	account, err := repo.GetByName(name)
	checkErr(err, "Cannot get account.")
	err = tx.Commit()
	checkErr(err, "Cannot get account.")
	return account
}

func getAccountByOwnerID(ownerID uuid.UUID) *models.Account {
	tx := tenv.GetDBxTx()
	repo := repo.MakeAccountRepoTx(tx)
	account, err := repo.GetByOwnerID(ownerID)
	checkErr(err, "Cannot get account.")
	err = tx.Commit()
	checkErr(err, "Cannot get account.")
	return account
}

func getAccountByParentID(parentID uuid.UUID) *models.Account {
	tx := tenv.GetDBxTx()
	repo := repo.MakeAccountRepoTx(tx)
	account, err := repo.GetByParentID(parentID)
	checkErr(err, "Cannot get account.")
	err = tx.Commit()
	checkErr(err, "Cannot get account.")
	return account
}

func checkAccount(id uuid.UUID) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeAccountRepoTx(tx)
	account, err := repo.Get(id)
	logErr(err, "Cannot get account.")
	err = tx.Commit()
	logErr(err, "Cannot get account.")
	return account != nil && err == nil
}

func checkAccountByName(name string) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeAccountRepoTx(tx)
	account, err := repo.GetByName(name)
	logErr(err, "Cannot get account.")
	err = tx.Commit()
	logErr(err, "Cannot get account.")
	return account != nil && err == nil
}

func checkAccountByOwnerID(ownerID uuid.UUID) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeAccountRepoTx(tx)
	account, err := repo.GetByOwnerID(ownerID)
	logErr(err, "Cannot get account.")
	err = tx.Commit()
	logErr(err, "Cannot get account.")
	return account != nil && err == nil
}

func checkAccountByParentID(parentID uuid.UUID) bool {
	tx := tenv.GetDBxTx()
	repo := repo.MakeAccountRepoTx(tx)
	account, err := repo.GetByParentID(parentID)
	logErr(err, "Cannot get account.")
	err = tx.Commit()
	logErr(err, "Cannot get account.")
	return account != nil && err == nil
}

func accountsMatch(account, tc *models.Account) bool {
	return account.Match(tc)
}
