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
	// ResourcePermissionResource - ResourcePermission wrapper for marshalling
	ResourcePermissionResource struct {
		Data models.ResourcePermission `json:"data"`
	}

	// ResourcePermissionListResource - ResourcePermission list wrapper for marshalling
	ResourcePermissionListResource struct {
		Data []models.ResourcePermission `json:"data"`
	}
)

const (
	// Resource
	resourcePermissionRes     = "resource-permission"
	resourcePermissionResName = "ResourcePermission"
)

var (
	okCreateResourcePermissionMsg  = api.MakeOkMessage("created", resourcePermissionResName)
	okGetResourcePermissionMsg     = api.MakeOkMessage("retrieved", resourcePermissionResName)
	okUpdateResourcePermissionMsg  = api.MakeOkMessage("updated", resourcePermissionResName)
	okDeleteResourcePermissionMsg  = api.MakeOkMessage("deleted", resourcePermissionResName)
	errListResourcePermissionsMsg  = api.MakeErrorMessage("list", resourcePermissionResName)
	errCreateResourcePermissionMsg = api.MakeErrorMessage("create", resourcePermissionResName)
	errGetResourcePermissionMsg    = api.MakeErrorMessage("get", resourcePermissionResName)
	errUpdateResourcePermissionMsg = api.MakeErrorMessage("update", resourcePermissionResName)
	errDeleteResourcePermissionMsg = api.MakeErrorMessage("delete", resourcePermissionResName)
)

// IndexResourcePermissionsV1 - Renders a list containing all resource-permissions.
// Handler for HTTP Get - "/resource-permissions"
func IndexResourcePermissionsV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Repo
	resourcePermissionRepo := repo.MakeResourcePermissionRepoTx(tx)
	resourcePermissions, err := resourcePermissionRepo.GetAll()
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errListResourcePermissionsMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errListResourcePermissionsMsg)
	}
	// Ok
	res := ResourcePermissionListResource{Data: resourcePermissions}
	return api.OkHR(createdSt, res, okCreateResourcePermissionMsg)
}

// GetResourcePermissionV1 - Shows a ResourcePermission.
// Handler for HTTP Get - "/resource-permissions/{id}"
func GetResourcePermissionV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errGetResourcePermissionMsg)
	}
	// Repo
	resourcePermissionRepo := repo.MakeResourcePermissionRepoTx(tx)
	resourcePermission, err := resourcePermissionRepo.Get(id)
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errGetResourcePermissionMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errGetResourcePermissionMsg)
	}
	// Ok
	res := ResourcePermissionResource{Data: *resourcePermission}
	return api.OkHR(okSt, res, okGetResourcePermissionMsg)
}

// CreateResourcePermissionV1 - Create new ResourcePermission API v1.
// Handler for HTTP Post - "/resource-permissions"
func CreateResourcePermissionV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Decode
	var resourcePermissionRes ResourcePermissionResource
	err := json.NewDecoder(r.Body).Decode(&resourcePermissionRes)
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errCreateResourcePermissionMsg)
	}
	resourcePermission := &resourcePermissionRes.Data
	// Services
	resourcePermissionSvc := services.ResourcePermissionService{Tx: tx}
	resourcePermissionSvc.Create(resourcePermission)
	err = tx.Commit()
	// Error
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errCreateResourcePermissionMsg)
	}
	// Ok
	res := ResourcePermissionResource{Data: *resourcePermission}
	return api.OkHR(createdSt, res, okCreateResourcePermissionMsg)
}

// UpdateResourcePermissionV1 - Update a ResourcePermission.
// Handler for HTTP Patch/Put - "/resource-permissions/{id}"
func UpdateResourcePermissionV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errGetResourcePermissionMsg)
	}
	// Decode
	var resourcePermissionRes ResourcePermissionResource
	err = json.NewDecoder(r.Body).Decode(&resourcePermissionRes)
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errCreateResourcePermissionMsg)
	}
	resourcePermission := &resourcePermissionRes.Data
	resourcePermission.SetID(id)
	// Services
	resourcePermissionSvc := services.ResourcePermissionService{Tx: tx}
	resourcePermissionSvc.Update(resourcePermission)
	err = tx.Commit()
	// Error
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errUpdateResourcePermissionMsg)
	}
	// Ok
	res := ResourcePermissionResource{Data: *resourcePermission}
	return api.OkHR(noContentSt, res, okUpdateResourcePermissionMsg)
}

// DeleteResourcePermissionV1 - Delete ResourcePermission.
// Handler for HTTP Delete - "/resource-permissions/{id}"
func DeleteResourcePermissionV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errGetResourcePermissionMsg)
	}
	// Repo
	resourcePermissionRepo := repo.MakeResourcePermissionRepoTx(tx)
	resourcePermission, err := resourcePermissionRepo.Get(id)
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errDeleteResourcePermissionMsg)
	}
	resourcePermissionRepo.Delete(resourcePermission.ID.UUID)
	err = tx.Commit()
	// Error
	if err != nil {
		return resourcePermissionErrHR(api.WrapError(err), errDeleteResourcePermissionMsg)
	}
	// Ok
	res := ResourcePermissionResource{Data: *resourcePermission}
	return api.OkHR(noContentSt, res, okDeleteResourcePermissionMsg)
}

// Private
// resourcePermissionErrHR - Default ResourcePermissionErrorHandlerResult for errors.
func resourcePermissionErrHR(err error, msg string) api.HandlerResult {
	return api.DefErrorHR(err, msg)
}
