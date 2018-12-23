package api

import (
	"encoding/json"
	"net/http"

	api "github.com/adrianpk/kamien/api"
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"{{.Package}}/services"
)

type (
	// AccountResource - Account wrapper for marshalling
	AccountResource struct {
		Data models.Account `json:"data"`
	}

	// AccountListResource - Account list wrapper for marshalling
	AccountListResource struct {
		Data []models.Account `json:"data"`
	}
)

const (
	// Resource
	accountRes     = "account"
	accountResName = "Account"
)

var (
	okCreateAccountMsg        = api.MakeOkMessage("created", accountResName)
	okGetAccountAccountMsg    = api.MakeOkMessage("retrieved", accountResName)
	okUpdateAccountAccountMsg = api.MakeOkMessage("updated", accountResName)
	okDeleteAccountAccountMsg = api.MakeOkMessage("deleted", accountResName)
	errListAccountsMsg        = api.MakeErrorMessage("list", accountResName)
	errCreateAccountMsg       = api.MakeErrorMessage("create", accountResName)
	errGetAccountMsg          = api.MakeErrorMessage("get", accountResName)
	errUpdateAccountMsg       = api.MakeErrorMessage("update", accountResName)
	errDeleteAccountMsg       = api.MakeErrorMessage("delete", accountResName)
)

// IndexAccountsV1 - Renders a list containing all accounts.
// Handler for HTTP Get - "/accounts"
func IndexAccountsV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Repo
	accountRepo := repo.MakeAccountRepoTx(tx)
	accounts, _ := accountRepo.GetAll()
	err := tx.Commit()
	// Error
	if err != nil {
		return accountErrHR(api.WrapError(err), errListAccountsMsg)
	}
	// Ok
	res := AccountListResource{Data: accounts}
	return api.OkHR(createdSt, res, okCreateAccountMsg)
}

// GetAccountV1 - Shows a account.
// Handler for HTTP Get - "/accounts/{id}"
func GetAccountV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return accountErrHR(api.WrapError(err), errGetAccountMsg)
	}
	// Repo
	accountRepo := repo.MakeAccountRepoTx(tx)
	account, _ := accountRepo.Get(id)
	err = tx.Commit()
	// Error
	if err != nil {
		return accountErrHR(api.WrapError(err), errGetAccountMsg)
	}
	// Ok
	res := AccountResource{Data: *account}
	return api.OkHR(okSt, res, okGetAccountAccountMsg)
}

// CreateAccountV1 - Create new Account API v1.
// Handler for HTTP Post - "/accounts"
func CreateAccountV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Decode
	var accountRes AccountResource
	err := json.NewDecoder(r.Body).Decode(&accountRes)
	if err != nil {
		return accountErrHR(api.WrapError(err), errCreateAccountMsg)
	}
	account := &accountRes.Data
	// Services
	accountSvc := services.AccountService{Tx: tx}
	accountSvc.Create(account)
	tx.Commit()
	// Error
	if err != nil {
		return accountErrHR(api.WrapError(err), errCreateAccountMsg)
	}
	// Ok
	res := AccountResource{Data: *account}
	return api.OkHR(createdSt, res, okCreateAccountMsg)
}

// UpdateAccountV1 - Update a account.
// Handler for HTTP Patch/Put - "/accounts/{id}"
func UpdateAccountV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return accountErrHR(api.WrapError(err), errGetAccountMsg)
	}
	// Decode
	var accountRes AccountResource
	err = json.NewDecoder(r.Body).Decode(&accountRes)
	if err != nil {
		return accountErrHR(api.WrapError(err), errCreateAccountMsg)
	}
	account := &accountRes.Data
	// Repo
	accountRepo := repo.MakeAccountRepoTx(tx)
	account.SetID(id)
	_ = accountRepo.Update(account)
	err = tx.Commit()
	// Error
	if err != nil {
		return accountErrHR(api.WrapError(err), errUpdateAccountMsg)
	}
	// Ok
	res := AccountResource{Data: *account}
	return api.OkHR(noContentSt, res, okUpdateAccountAccountMsg)
}

// DeleteAccountV1 - Delete account.
// Handler for HTTP Delete - "/accounts/{id}"
func DeleteAccountV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		panic(err)
		return accountErrHR(api.WrapError(err), errGetAccountMsg)
	}
	// Repo
	accountRepo := repo.MakeAccountRepoTx(tx)
	accountRepo.Delete(id)
	err = tx.Commit()
	// Error
	if err != nil {
		panic(err)
		return accountErrHR(api.WrapError(err), errDeleteAccountMsg)
	}
	// Ok
	res := AccountResource{Data: models.Account{}}
	return api.OkHR(noContentSt, res, okDeleteAccountAccountMsg)
}

// Private
// accountErrHR - Default Account Error Handler Result for errors.
func accountErrHR(err error, msg string) api.HandlerResult {
	return api.DefErrorHR(err, msg)
}
