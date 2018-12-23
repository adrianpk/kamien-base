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
	// PermissionResource - Permission wrapper for marshalling
	PermissionResource struct {
		Data models.Permission `json:"data"`
	}

	// PermissionListResource - Permission list wrapper for marshalling
	PermissionListResource struct {
		Data []models.Permission `json:"data"`
	}
)

const (
	// Resource
	permissionRes     = "permission"
	permissionResName = "Permission"
)

var (
	okCreatePermissionMsg  = api.MakeOkMessage("created", permissionResName)
	okGetPermissionMsg     = api.MakeOkMessage("retrieved", permissionResName)
	okUpdatePermissionMsg  = api.MakeOkMessage("updated", permissionResName)
	okDeletePermissionMsg  = api.MakeOkMessage("deleted", permissionResName)
	errListPermissionsMsg  = api.MakeErrorMessage("list", permissionResName)
	errCreatePermissionMsg = api.MakeErrorMessage("create", permissionResName)
	errGetPermissionMsg    = api.MakeErrorMessage("get", permissionResName)
	errUpdatePermissionMsg = api.MakeErrorMessage("update", permissionResName)
	errDeletePermissionMsg = api.MakeErrorMessage("delete", permissionResName)
)

// IndexPermissionsV1 - Renders a list containing all permissions.
// Handler for HTTP Get - "/permissions"
func IndexPermissionsV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Repo
	permissionRepo := repo.MakePermissionRepoTx(tx)
	permissions, err := permissionRepo.GetAll()
	if err != nil {
		return permissionErrHR(api.WrapError(err), errListPermissionsMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		return permissionErrHR(api.WrapError(err), errListPermissionsMsg)
	}
	// Ok
	res := PermissionListResource{Data: permissions}
	return api.OkHR(createdSt, res, okCreatePermissionMsg)
}

// GetPermissionV1 - Shows a Permission.
// Handler for HTTP Get - "/permissions/{id}"
func GetPermissionV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return permissionErrHR(api.WrapError(err), errGetPermissionMsg)
	}
	// Repo
	permissionRepo := repo.MakePermissionRepoTx(tx)
	permission, err := permissionRepo.Get(id)
	if err != nil {
		return permissionErrHR(api.WrapError(err), errGetPermissionMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		return permissionErrHR(api.WrapError(err), errGetPermissionMsg)
	}
	// Ok
	res := PermissionResource{Data: *permission}
	return api.OkHR(okSt, res, okGetPermissionMsg)
}

// CreatePermissionV1 - Create new Permission API v1.
// Handler for HTTP Post - "/permissions"
func CreatePermissionV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Decode
	var permissionRes PermissionResource
	err := json.NewDecoder(r.Body).Decode(&permissionRes)
	if err != nil {
		return permissionErrHR(api.WrapError(err), errCreatePermissionMsg)
	}
	permission := &permissionRes.Data
	// Services
	permissionSvc := services.PermissionService{Tx: tx}
	permissionSvc.Create(permission)
	err = tx.Commit()
	// Error
	if err != nil {
		return permissionErrHR(api.WrapError(err), errCreatePermissionMsg)
	}
	// Ok
	res := PermissionResource{Data: *permission}
	return api.OkHR(createdSt, res, okCreatePermissionMsg)
}

// UpdatePermissionV1 - Update a Permission.
// Handler for HTTP Patch/Put - "/permissions/{id}"
func UpdatePermissionV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return permissionErrHR(api.WrapError(err), errGetPermissionMsg)
	}
	// Decode
	var permissionRes PermissionResource
	err = json.NewDecoder(r.Body).Decode(&permissionRes)
	if err != nil {
		return permissionErrHR(api.WrapError(err), errCreatePermissionMsg)
	}
	permission := &permissionRes.Data
	permission.SetID(id)
	// Services
	permissionSvc := services.PermissionService{Tx: tx}
	permissionSvc.Update(permission)
	err = tx.Commit()
	// Error
	if err != nil {
		return permissionErrHR(api.WrapError(err), errUpdatePermissionMsg)
	}
	// Ok
	res := PermissionResource{Data: *permission}
	return api.OkHR(noContentSt, res, okUpdatePermissionMsg)
}

// DeletePermissionV1 - Delete Permission.
// Handler for HTTP Delete - "/permissions/{id}"
func DeletePermissionV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return permissionErrHR(api.WrapError(err), errGetPermissionMsg)
	}
	// Repo
	permissionRepo := repo.MakePermissionRepoTx(tx)
	permission, err := permissionRepo.Get(id)
	if err != nil {
		return permissionErrHR(api.WrapError(err), errDeletePermissionMsg)
	}
	permissionRepo.Delete(permission.ID.UUID)
	err = tx.Commit()
	// Error
	if err != nil {
		return permissionErrHR(api.WrapError(err), errDeletePermissionMsg)
	}
	// Ok
	res := PermissionResource{Data: *permission}
	return api.OkHR(noContentSt, res, okDeletePermissionMsg)
}

// Private
// permissionErrHR - Default PermissionErrorHandlerResult for errors.
func permissionErrHR(err error, msg string) api.HandlerResult {
	return api.DefErrorHR(err, msg)
}
