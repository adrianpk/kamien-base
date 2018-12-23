package controllers

import (
	"net/http"

	"{{.Package}}/common"
	"{{.Package}}/models"
	"{{.Package}}/repo"
	b "github.com/adrianpk/kamien/boot"
	cm "github.com/adrianpk/kamien/common"
	c "github.com/adrianpk/kamien/controllers"
	uuid "github.com/satori/go.uuid"
)

const (
	// Resource
	accountRes = "account"
)

// IndexAccounts - Renders a list containing all accounts.
// Handler for HTTP Get - "/accounts"
func IndexAccounts(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) c.HandlerResult {
	tx := c.RequestTx(r)
	// Repo
	accountRepo := repo.MakeAccountRepoTx(tx)
	accounts, _ := accountRepo.GetAll()
	// Result
	okRE := c.OkRE(accountRes, accounts, c.IndexTmpl, nil)
	errRE := c.ErrorRE(accountRes, nil, c.IndexTmpl, nil)
	errRE.AddErrorFlash(c.ServerErrorMsg)
	return c.HR(okRE, errRE)
}

// EditAccount - Edit a account.
// Handler for HTTP Get - "/accounts/{id}/edit"
func EditAccount(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) c.HandlerResult {
	id := c.GetID(r)
	// Repo
	tx := c.RequestTx(r)
	accountRepo := repo.MakeAccountRepoTx(tx)
	account, _ := accountRepo.Get(id)
	// Result
	action := accountUpdAction(accountRes, account)
	okRE := c.OkRE(accountRes, account, c.EditTmpl, &action)
	errRE := c.ErrorRedirRE(common.AccountPath())
	msg := c.MakeErrorMessage("edit", accountRes)
	errRE.AddErrorFlash(msg)
	return c.HR(okRE, errRE)
}

// NewAccount - Presents a new account form.
// Handler for HTTP Get - "/accounts/new"
func NewAccount(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) c.HandlerResult {
	account := models.MakeAccount()
	action := accountCrtAction(accountRes)
	// Result
	return c.OkHR(accountRes, account, c.NewTmpl, &action)
}

// ShowAccount - Shows a account.
// Handler for HTTP Get - "/accounts/{id}"
func ShowAccount(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) c.HandlerResult {
	id := c.GetID(r)
	// Repo
	tx := c.RequestTx(r)
	accountRepo := repo.MakeAccountRepoTx(tx)
	account, _ := accountRepo.Get(id)
	// Result
	okRE := c.OkRE(accountRes, account, c.ShowTmpl, nil)
	errRE := c.ErrorRedirRE(common.AccountPath())
	msg := c.MakeErrorMessage("show", accountRes)
	errRE.AddErrorFlash(msg)
	return c.HR(okRE, errRE)
}

// CreateAccount - Create new account.
// Handler for HTTP Post - "/accounts"
func CreateAccount(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) c.HandlerResult {
	action := accountCrtAction(accountRes)
	// Parsing
	err := r.ParseForm()
	if err != nil {
		return accountErrHR(nil, c.NewTmpl, &action, err)
	}
	// Decoding
	var account models.Account
	err = c.NewDecoder().Decode(&account, r.Form)
	if err != nil {
		return accountErrHR(account, c.NewTmpl, &action, err)
	}
	// Transaction
	tx := c.RequestTx(r)
	// Update properties
	// Owner
	ownerID := c.GetFormIDValue(r, "owner-id")
	userRepo := repo.MakeUserRepoTx(tx)
	if ownerID != uuid.Nil {
		owner, err := userRepo.Get(ownerID)
		if err != nil {
			return accountErrHR(account, c.NewTmpl, &action, err)
		}
		account.OwnerID = owner.ID
	}
	// Parent
	parentID := c.GetFormIDValue(r, "parent-id")
	accountRepo := repo.MakeAccountRepoTx(tx)
	if parentID != uuid.Nil {
		parent, err := accountRepo.Get(parentID)
		if err != nil {
			return accountErrHR(account, c.NewTmpl, &action, err)
		}
		account.ParentID = parent.ID
	}
	// Repo
	accountRepo.Create(&account)
	// Result
	okRE := c.OkRedirRE(common.AccountPath())
	okRE.AddInfoFlash("Account created")
	errRE := c.ErrorRedirRE(common.AccountPath())
	msg := c.MakeErrorMessage("create", accountRes)
	errRE.AddErrorFlash(msg)
	return c.HR(okRE, errRE)
}

// UpdateAccount - Update a account.
// Handler for HTTP Patch/Put - "/accounts/{id}"
func UpdateAccount(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) c.HandlerResult {
	id := c.GetID(r)
	account := models.MakeAccountWithID(id)
	action := accountUpdAction(accountRes, account)
	// Parsing
	err := r.ParseForm()
	if err != nil {
		return accountErrHR(nil, c.EditTmpl, &action, err)
	}
	// Decoding
	err = c.NewDecoder().Decode(account, r.Form)
	if err != nil {
		return accountErrHR(account, c.EditTmpl, &action, err)
	}
	// Repo
	tx := c.RequestTx(r)
	account.SetID(id)
	// Update properties
	// Owner
	ownerID := c.GetFormIDValue(r, "owner-id")
	userRepo := repo.MakeUserRepoTx(tx)
	if ownerID != uuid.Nil {
		owner, err := userRepo.Get(ownerID)
		if err != nil {
			log.Errorf("OID: %+v", err)
			return accountErrHR(account, c.NewTmpl, &action, err)
		}
		account.OwnerID = owner.ID
	}
	// Parent
	parentID := c.GetFormIDValue(r, "parent-id")
	accountRepo := repo.MakeAccountRepoTx(tx)
	if parentID != uuid.Nil {
		parent, err := accountRepo.Get(parentID)
		if err != nil {
			log.Errorf("PID: %+v", err)
			return accountErrHR(account, c.NewTmpl, &action, err)
		}
		account.ParentID = parent.ID
	}
	err = accountRepo.Update(account)
	// Result
	okRE := c.OkRedirRE(common.AccountPath())
	okRE.AddInfoFlash("Account updated")
	errRE := c.ErrorRedirRE(common.AccountPath())
	msg := c.MakeErrorMessage("update", accountRes)
	errRE.AddErrorFlash(msg)
	return c.HR(okRE, errRE)
}

// InitDeleteAccount - Account delete confirmation.
// Handler for HTTP Get - "/accounts/{id}/init-delete"
func InitDeleteAccount(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) c.HandlerResult {
	// Repo
	tx := c.RequestTx(r)
	accountRepo := repo.MakeAccountRepoTx(tx)
	id := c.GetID(r)
	account, _ := accountRepo.Get(id)
	// Result
	okRE := c.OkRE(accountRes, account, c.DeleteTmpl, nil)
	errRE := c.ErrorRedirRE(common.AccountPath())
	msg := c.MakeErrorMessage("delete", accountRes)
	errRE.AddErrorFlash(msg)
	return c.HR(okRE, errRE)
}

// DeleteAccount - Delete account.
// Handler for HTTP Delete - "/accounts/{id}"
func DeleteAccount(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) c.HandlerResult {
	id := c.GetID(r)
	// Repo
	tx := c.RequestTx(r)
	accountRepo := repo.MakeAccountRepoTx(tx)
	account, _ := accountRepo.Get(id)
	accountRepo.Delete(account.ID.UUID)
	// Result
	okRE := c.OkRedirRE(common.AccountPath())
	okRE.AddInfoFlash("Account created")
	errRE := c.OkRE(accountRes, account, c.DeleteTmpl, nil)
	msg := c.MakeErrorMessage("delete", accountRes)
	errRE.AddErrorFlash(msg)
	return c.HR(okRE, errRE)
}

// Private

// accountCrtAction - Create Action: resource index path, POST HTTP method.
func accountCrtAction(resource string) c.Action {
	return c.Action{Target: common.AccountPath(), Method: "POST"}
}

// accountUpdAction - Upadate Action: resource id path, PUT HTTP method.
func accountUpdAction(resource string, model cm.Identifiable) c.Action {
	return c.Action{Target: common.AccountPathID(model), Method: "PUT"}
}

// accountErrHR - Default User Error Handler Result for errors.
func accountErrHR(model interface{}, templateName string, action *c.Action, err error) c.HandlerResult {
	return c.DefErrorHR(accountRes, model, templateName, action, err)
}
