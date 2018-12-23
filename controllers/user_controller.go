package controllers

import (
	"fmt"
	"net/http"

	"{{.Package}}/common"
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"{{.Package}}/services"
	b "github.com/adrianpk/kamien/boot"
	cmn "github.com/adrianpk/kamien/common"
	ctr "github.com/adrianpk/kamien/controllers"
)

const (
	// Resource
	userRes = "user"
)

var (
	okCrtUserMsg   = ctr.MakeOkMessage("created", userRes)
	okUpdUserMsg   = ctr.MakeOkMessage("updated", userRes)
	okDelUserMsg   = ctr.MakeOkMessage("deleted", userRes)
	errCrtUserMsg  = ctr.MakeErrorMessage("create", userRes)
	errShowUserMsg = ctr.MakeErrorMessage("show", userRes)
	errEditUserMsg = ctr.MakeErrorMessage("show", userRes)
	errUpdUserMsg  = ctr.MakeErrorMessage("update", userRes)
	errDelUserMsg  = ctr.MakeErrorMessage("delete", userRes)
)

// IndexUsers - Renders a list containing all users.
// Handler for HTTP Get - "/users"
func IndexUsers(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) ctr.HandlerResult {
	tx := ctr.RequestTx(r)
	// Repo
	userRepo := repo.MakeUserRepoTx(tx)
	users, _ := userRepo.GetAll()
	// Ok
	okRE := ctr.OkRE(userRes, users, ctr.IndexTmpl, nil)
	// Err
	errRE := ctr.ErrorRE(userRes, nil, ctr.IndexTmpl, nil)
	errRE.AddErrorFlash(ctr.ServerErrorMsg)
	// Return
	return ctr.HR(okRE, errRE)
}

// EditUser - Edit a user.
// Handler for HTTP Get - "/users/{id}/edit"
func EditUser(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) ctr.HandlerResult {
	id := ctr.GetID(r)
	// Repo
	tx := ctr.RequestTx(r)
	userRepo := repo.MakeUserRepoTx(tx)
	user, _ := userRepo.Get(id)
	// Ok
	action := userUpdAction(userRes, user)
	okRE := ctr.OkRE(userRes, user, ctr.EditTmpl, &action)
	// Err
	errRE := ctr.ErrorRedirRE(common.UserPath())
	errRE.AddErrorFlash(errEditUserMsg)
	// Return
	return ctr.HR(okRE, errRE)
}

// NewUser - Presents a new user form.
// Handler for HTTP Get - "/users/new"
func NewUser(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) ctr.HandlerResult {
	user := models.MakeUser()
	action := userCrtAction(userRes)
	// Result
	return ctr.OkHR(userRes, user, ctr.NewTmpl, &action)
}

// ShowUser - Shows a user.
// Handler for HTTP Get - "/users/{id}"
func ShowUser(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) ctr.HandlerResult {
	id := ctr.GetID(r)
	// Repo
	tx := ctr.RequestTx(r)
	userRepo := repo.MakeUserRepoTx(tx)
	user, _ := userRepo.Get(id)
	// Ok
	okRE := ctr.OkRE(userRes, user, ctr.ShowTmpl, nil)
	// Err
	errRE := ctr.ErrorRedirRE(common.UserPath())
	errRE.AddErrorFlash(errShowUserMsg)
	// Return
	return ctr.HR(okRE, errRE)
}

// CreateUser - Create new user.
// Handler for HTTP Post - "/users"
func CreateUser(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) ctr.HandlerResult {
	action := userCrtAction(userRes)
	// Parsing
	err := r.ParseForm()
	if err != nil {
		fmt.Println("1 - ", err)
		return userErrHR(nil, ctr.NewTmpl, &action, err)
	}
	// Decoding
	var user models.User
	err = ctr.NewDecoder().Decode(&user, r.Form)
	if err != nil {
		return userErrHR(user, ctr.NewTmpl, &action, err)
	}
	// Services
	tx := ctr.RequestTx(r)
	uc := services.UserService{Tx: tx}
	uc.Create(&user)
	// Ok
	okRE := ctr.OkRedirRE(common.UserPath())
	okRE.AddInfoFlash(okCrtUserMsg)
	// Err
	errRE := ctr.ErrorRedirRE(common.UserPath())
	errRE.AddErrorFlash(errCrtUserMsg)
	// Return
	return ctr.HR(okRE, errRE)
}

// UpdateUser - Update a user.
// Handler for HTTP Patch/Put - "/users/{id}"
func UpdateUser(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) ctr.HandlerResult {
	id := ctr.GetID(r)
	user := models.MakeUserWithID(id)
	action := userUpdAction(userRes, user)
	// Parsing
	err := r.ParseForm()
	if err != nil {
		return userErrHR(nil, ctr.EditTmpl, &action, err)
	}
	// Decoding
	err = ctr.NewDecoder().Decode(user, r.Form)
	if err != nil {
		return userErrHR(user, ctr.EditTmpl, &action, err)
	}
	// Repo
	tx := ctr.RequestTx(r)
	userRepo := repo.MakeUserRepoTx(tx)
	user.SetID(id)
	err = userRepo.Update(user)
	// Ok
	okRE := ctr.OkRedirRE(common.UserPath())
	okRE.AddInfoFlash(okUpdUserMsg)
	// Err
	errRE := ctr.ErrorRedirRE(common.UserPath())
	errRE.AddErrorFlash(errUpdUserMsg)
	// Return
	return ctr.HR(okRE, errRE)
}

// InitDeleteUser - User delete confirmation.
// Handler for HTTP Get - "/users/{id}/init-delete"
func InitDeleteUser(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) ctr.HandlerResult {
	// Repo
	tx := ctr.RequestTx(r)
	userRepo := repo.MakeUserRepoTx(tx)
	id := ctr.GetID(r)
	user, _ := userRepo.Get(id)
	// Ok
	okRE := ctr.OkRE(userRes, user, ctr.DeleteTmpl, nil)
	// Err
	errRE := ctr.ErrorRedirRE(common.UserPath())
	errRE.AddErrorFlash(errDelUserMsg)
	// Return
	return ctr.HR(okRE, errRE)
}

// DeleteUser - Delete user.
// Handler for HTTP Delete - "/users/{id}"
func DeleteUser(ac *b.AppContext, rw http.ResponseWriter, r *http.Request) ctr.HandlerResult {
	id := ctr.GetID(r)
	// Repo
	tx := ctr.RequestTx(r)
	userRepo := repo.MakeUserRepoTx(tx)
	user, _ := userRepo.Get(id)
	userRepo.Delete(user.ID.UUID)
	// Ok
	okRE := ctr.OkRedirRE(common.UserPath())
	okRE.AddInfoFlash(okDelUserMsg)
	// Err
	errRE := ctr.ErrorRE(userRes, user, ctr.DeleteTmpl, nil)
	errRE.AddErrorFlash(errDelUserMsg)
	// Return
	return ctr.HR(okRE, errRE)
}

// userCrtAction - Create Action: resource index path, POST HTTP method.
func userCrtAction(resource string) ctr.Action {
	return ctr.Action{Target: common.UserPath(), Method: "POST"}
}

// userUpdAction - Upadate Action: resource id path, PUT HTTP method.
func userUpdAction(resource string, model cmn.Identifiable) ctr.Action {
	return ctr.Action{Target: common.UserPathID(model), Method: "PUT"}
}

// userErrHR - Default User Error Handler Result for errors.
func userErrHR(model interface{}, templateName string, action *ctr.Action, err error) ctr.HandlerResult {
	return ctr.DefErrorHR(userRes, model, templateName, action, err)
}
