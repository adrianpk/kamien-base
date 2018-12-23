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
	// RolePermissionResource - RolePermission wrapper for marshalling
	RolePermissionResource struct {
		Data models.RolePermission `json:"data"`
	}

	// RolePermissionListResource - RolePermission list wrapper for marshalling
	RolePermissionListResource struct {
		Data []models.RolePermission `json:"data"`
	}
)

const (
	// Resource
	rolePermissionRes     = "role-permission"
	rolePermissionResName = "RolePermission"
)

var (
	okCreateRolePermissionMsg  = api.MakeOkMessage("created", rolePermissionResName)
	okGetRolePermissionMsg     = api.MakeOkMessage("retrieved", rolePermissionResName)
	okUpdateRolePermissionMsg  = api.MakeOkMessage("updated", rolePermissionResName)
	okDeleteRolePermissionMsg  = api.MakeOkMessage("deleted", rolePermissionResName)
	errListRolePermissionsMsg  = api.MakeErrorMessage("list", rolePermissionResName)
	errCreateRolePermissionMsg = api.MakeErrorMessage("create", rolePermissionResName)
	errGetRolePermissionMsg    = api.MakeErrorMessage("get", rolePermissionResName)
	errUpdateRolePermissionMsg = api.MakeErrorMessage("update", rolePermissionResName)
	errDeleteRolePermissionMsg = api.MakeErrorMessage("delete", rolePermissionResName)
)

// IndexRolePermissionsV1 - Renders a list containing all role-permissions.
// Handler for HTTP Get - "/role-permissions"
func IndexRolePermissionsV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Repo
	rolePermissionRepo := repo.MakeRolePermissionRepoTx(tx)
	rolePermissions, err := rolePermissionRepo.GetAll()
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errListRolePermissionsMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errListRolePermissionsMsg)
	}
	// Ok
	res := RolePermissionListResource{Data: rolePermissions}
	return api.OkHR(createdSt, res, okCreateRolePermissionMsg)
}

// GetRolePermissionV1 - Shows a RolePermission.
// Handler for HTTP Get - "/role-permissions/{id}"
func GetRolePermissionV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errGetRolePermissionMsg)
	}
	// Repo
	rolePermissionRepo := repo.MakeRolePermissionRepoTx(tx)
	rolePermission, err := rolePermissionRepo.Get(id)
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errGetRolePermissionMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errGetRolePermissionMsg)
	}
	// Ok
	res := RolePermissionResource{Data: *rolePermission}
	return api.OkHR(okSt, res, okGetRolePermissionMsg)
}

// CreateRolePermissionV1 - Create new RolePermission API v1.
// Handler for HTTP Post - "/role-permissions"
func CreateRolePermissionV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Decode
	var rolePermissionRes RolePermissionResource
	err := json.NewDecoder(r.Body).Decode(&rolePermissionRes)
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errCreateRolePermissionMsg)
	}
	rolePermission := &rolePermissionRes.Data
	// Services
	rolePermissionSvc := services.RolePermissionService{Tx: tx}
	rolePermissionSvc.Create(rolePermission)
	err = tx.Commit()
	// Error
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errCreateRolePermissionMsg)
	}
	// Ok
	res := RolePermissionResource{Data: *rolePermission}
	return api.OkHR(createdSt, res, okCreateRolePermissionMsg)
}

// UpdateRolePermissionV1 - Update a RolePermission.
// Handler for HTTP Patch/Put - "/role-permissions/{id}"
func UpdateRolePermissionV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errGetRolePermissionMsg)
	}
	// Decode
	var rolePermissionRes RolePermissionResource
	err = json.NewDecoder(r.Body).Decode(&rolePermissionRes)
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errCreateRolePermissionMsg)
	}
	rolePermission := &rolePermissionRes.Data
	rolePermission.SetID(id)
	// Services
	rolePermissionSvc := services.RolePermissionService{Tx: tx}
	rolePermissionSvc.Update(rolePermission)
	err = tx.Commit()
	// Error
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errUpdateRolePermissionMsg)
	}
	// Ok
	res := RolePermissionResource{Data: *rolePermission}
	return api.OkHR(noContentSt, res, okUpdateRolePermissionMsg)
}

// DeleteRolePermissionV1 - Delete RolePermission.
// Handler for HTTP Delete - "/role-permissions/{id}"
func DeleteRolePermissionV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errGetRolePermissionMsg)
	}
	// Repo
	rolePermissionRepo := repo.MakeRolePermissionRepoTx(tx)
	rolePermission, err := rolePermissionRepo.Get(id)
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errDeleteRolePermissionMsg)
	}
	rolePermissionRepo.Delete(rolePermission.ID.UUID)
	err = tx.Commit()
	// Error
	if err != nil {
		return rolePermissionErrHR(api.WrapError(err), errDeleteRolePermissionMsg)
	}
	// Ok
	res := RolePermissionResource{Data: *rolePermission}
	return api.OkHR(noContentSt, res, okDeleteRolePermissionMsg)
}

// Private
// rolePermissionErrHR - Default RolePermissionErrorHandlerResult for errors.
func rolePermissionErrHR(err error, msg string) api.HandlerResult {
	return api.DefErrorHR(err, msg)
}
