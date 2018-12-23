package api

import (
	"encoding/json"
	"net/http"

	"{{.Package}}/models"
	"{{.Package}}/repo"
	"{{.Package}}/services"
	api "github.com/adrianpk/kamien/api"
)

type (
	// RoleResource - Role wrapper for marshalling
	RoleResource struct {
		Data models.Role `json:"data"`
	}

	// RoleListResource - Role list wrapper for marshalling
	RoleListResource struct {
		Data []models.Role `json:"data"`
	}
)

const (
	// Resource
	roleRes     = "role"
	roleResName = "Role"
)

var (
	okCreateRoleMsg          = api.MakeOkMessage("created", roleResName)
	okGetRoleMsg    = api.MakeOkMessage("retrieved", roleResName)
	okUpdateRoleMsg = api.MakeOkMessage("updated", roleResName)
	okDeleteRoleMsg = api.MakeOkMessage("deleted", roleResName)
	errListRolesMsg          = api.MakeErrorMessage("list", roleResName)
	errCreateRoleMsg         = api.MakeErrorMessage("create", roleResName)
	errGetRoleMsg            = api.MakeErrorMessage("get", roleResName)
	errUpdateRoleMsg         = api.MakeErrorMessage("update", roleResName)
	errDeleteRoleMsg         = api.MakeErrorMessage("delete", roleResName)
)

// IndexRolesV1 - Renders a list containing all roles.
// Handler for HTTP Get - "/roles"
func IndexRolesV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Repo
	roleRepo := repo.MakeRoleRepoTx(tx)
	roles, err := roleRepo.GetAll()
	if err != nil {
		return roleErrHR(api.WrapError(err), errListRolesMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		return roleErrHR(api.WrapError(err), errListRolesMsg)
	}
	// Ok
	res := RoleListResource{Data: roles}
	return api.OkHR(createdSt, res, okCreateRoleMsg)
}

// GetRoleV1 - Shows a Role.
// Handler for HTTP Get - "/roles/{id}"
func GetRoleV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return roleErrHR(api.WrapError(err), errGetRoleMsg)
	}
	// Repo
	roleRepo := repo.MakeRoleRepoTx(tx)
	role, err := roleRepo.Get(id)
	if err != nil {
		return roleErrHR(api.WrapError(err), errGetRoleMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		return roleErrHR(api.WrapError(err), errGetRoleMsg)
	}
	// Ok
	res := RoleResource{Data: *role}
	return api.OkHR(okSt, res, okGetRoleMsg)
}

// CreateRoleV1 - Create new Role API v1.
// Handler for HTTP Post - "/roles"
func CreateRoleV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Decode
	var roleRes RoleResource
	err := json.NewDecoder(r.Body).Decode(&roleRes)
	if err != nil {
		return roleErrHR(api.WrapError(err), errCreateRoleMsg)
	}
	role := &roleRes.Data
	// Services
	roleSvc := services.RoleService{Tx: tx}
	roleSvc.Create(role)
	err = tx.Commit()
	// Error
	if err != nil {
		return roleErrHR(api.WrapError(err), errCreateRoleMsg)
	}
	// Ok
	res := RoleResource{Data: *role}
	return api.OkHR(createdSt, res, okCreateRoleMsg)
}

// UpdateRoleV1 - Update a Role.
// Handler for HTTP Patch/Put - "/roles/{id}"
func UpdateRoleV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return roleErrHR(api.WrapError(err), errGetRoleMsg)
	}
	// Decode
	var roleRes RoleResource
	err = json.NewDecoder(r.Body).Decode(&roleRes)
	if err != nil {
		return roleErrHR(api.WrapError(err), errCreateRoleMsg)
	}
	role := &roleRes.Data
	role.SetID(id)
	// Services
	roleSvc := services.RoleService{Tx: tx}
	roleSvc.Update(role)
	err = tx.Commit()
	// Error
	if err != nil {
		return roleErrHR(api.WrapError(err), errUpdateRoleMsg)
	}
	// Ok
	res := RoleResource{Data: *role}
	return api.OkHR(noContentSt, res, okUpdateRoleMsg)
}

// DeleteRoleV1 - Delete Role.
// Handler for HTTP Delete - "/roles/{id}"
func DeleteRoleV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return roleErrHR(api.WrapError(err), errGetRoleMsg)
	}
	// Repo
	roleRepo := repo.MakeRoleRepoTx(tx)
	role, err := roleRepo.Get(id)
	if err != nil {
		return roleErrHR(api.WrapError(err), errDeleteRoleMsg)
	}
	roleRepo.Delete(role.ID.UUID)
	err = tx.Commit()
	// Error
	if err != nil {
		return roleErrHR(api.WrapError(err), errDeleteRoleMsg)
	}
	// Ok
	res := RoleResource{Data: *role}
	return api.OkHR(noContentSt, res, okDeleteRoleMsg)
}

// Private
// roleErrHR - Default RoleErrorHandlerResult for errors.
func roleErrHR(err error, msg string) api.HandlerResult {
	return api.DefErrorHR(err, msg)
}
