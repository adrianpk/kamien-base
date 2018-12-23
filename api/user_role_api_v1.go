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
	// UserRoleResource - UserRole wrapper for marshalling
	UserRoleResource struct {
		Data models.UserRole `json:"data"`
	}

	// UserRoleListResource - UserRole list wrapper for marshalling
	UserRoleListResource struct {
		Data []models.UserRole `json:"data"`
	}
)

const (
	// Resource
	userRoleRes     = "user-role"
	userRoleResName = "UserRole"
)

var (
	okCreateUserRoleMsg  = api.MakeOkMessage("created", userRoleResName)
	okGetUserRoleMsg     = api.MakeOkMessage("retrieved", userRoleResName)
	okUpdateUserRoleMsg  = api.MakeOkMessage("updated", userRoleResName)
	okDeleteUserRoleMsg  = api.MakeOkMessage("deleted", userRoleResName)
	errListUserRolesMsg  = api.MakeErrorMessage("list", userRoleResName)
	errCreateUserRoleMsg = api.MakeErrorMessage("create", userRoleResName)
	errGetUserRoleMsg    = api.MakeErrorMessage("get", userRoleResName)
	errUpdateUserRoleMsg = api.MakeErrorMessage("update", userRoleResName)
	errDeleteUserRoleMsg = api.MakeErrorMessage("delete", userRoleResName)
)

// IndexUserRolesV1 - Renders a list containing all user-roles.
// Handler for HTTP Get - "/user-roles"
func IndexUserRolesV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Repo
	userRoleRepo := repo.MakeUserRoleRepoTx(tx)
	userRoles, err := userRoleRepo.GetAll()
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errListUserRolesMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errListUserRolesMsg)
	}
	// Ok
	res := UserRoleListResource{Data: userRoles}
	return api.OkHR(createdSt, res, okCreateUserRoleMsg)
}

// GetUserRoleV1 - Shows a UserRole.
// Handler for HTTP Get - "/user-roles/{id}"
func GetUserRoleV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errGetUserRoleMsg)
	}
	// Repo
	userRoleRepo := repo.MakeUserRoleRepoTx(tx)
	userRole, err := userRoleRepo.Get(id)
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errGetUserRoleMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errGetUserRoleMsg)
	}
	// Ok
	res := UserRoleResource{Data: *userRole}
	return api.OkHR(okSt, res, okGetUserRoleMsg)
}

// CreateUserRoleV1 - Create new UserRole API v1.
// Handler for HTTP Post - "/user-roles"
func CreateUserRoleV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Decode
	var userRoleRes UserRoleResource
	err := json.NewDecoder(r.Body).Decode(&userRoleRes)
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errCreateUserRoleMsg)
	}
	userRole := &userRoleRes.Data
	// Services
	userRoleSvc := services.UserRoleService{Tx: tx}
	userRoleSvc.Create(userRole)
	err = tx.Commit()
	// Error
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errCreateUserRoleMsg)
	}
	// Ok
	res := UserRoleResource{Data: *userRole}
	return api.OkHR(createdSt, res, okCreateUserRoleMsg)
}

// UpdateUserRoleV1 - Update a UserRole.
// Handler for HTTP Patch/Put - "/user-roles/{id}"
func UpdateUserRoleV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errGetUserRoleMsg)
	}
	// Decode
	var userRoleRes UserRoleResource
	err = json.NewDecoder(r.Body).Decode(&userRoleRes)
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errCreateUserRoleMsg)
	}
	userRole := &userRoleRes.Data
	userRole.SetID(id)
	// Services
	userRoleSvc := services.UserRoleService{Tx: tx}
	userRoleSvc.Update(userRole)
	err = tx.Commit()
	// Error
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errUpdateUserRoleMsg)
	}
	// Ok
	res := UserRoleResource{Data: *userRole}
	return api.OkHR(noContentSt, res, okUpdateUserRoleMsg)
}

// DeleteUserRoleV1 - Delete UserRole.
// Handler for HTTP Delete - "/user-roles/{id}"
func DeleteUserRoleV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errGetUserRoleMsg)
	}
	// Repo
	userRoleRepo := repo.MakeUserRoleRepoTx(tx)
	userRole, err := userRoleRepo.Get(id)
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errDeleteUserRoleMsg)
	}
	userRoleRepo.Delete(userRole.ID.UUID)
	err = tx.Commit()
	// Error
	if err != nil {
		return userRoleErrHR(api.WrapError(err), errDeleteUserRoleMsg)
	}
	// Ok
	res := UserRoleResource{Data: *userRole}
	return api.OkHR(noContentSt, res, okDeleteUserRoleMsg)
}

// Private
// userRoleErrHR - Default UserRoleErrorHandlerResult for errors.
func userRoleErrHR(err error, msg string) api.HandlerResult {
	return api.DefErrorHR(err, msg)
}
