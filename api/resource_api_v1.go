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
	// ResourceResource - Resource wrapper for marshalling
	ResourceResource struct {
		Data models.Resource `json:"data"`
	}

	// ResourceListResource - Resource list wrapper for marshalling
	ResourceListResource struct {
		Data []models.Resource `json:"data"`
	}
)

const (
	// Resource
	resourceRes     = "resource"
	resourceResName = "Resource"
)

var (
	okCreateResourceMsg  = api.MakeOkMessage("created", resourceResName)
	okGetResourceMsg     = api.MakeOkMessage("retrieved", resourceResName)
	okUpdateResourceMsg  = api.MakeOkMessage("updated", resourceResName)
	okDeleteResourceMsg  = api.MakeOkMessage("deleted", resourceResName)
	errListResourcesMsg  = api.MakeErrorMessage("list", resourceResName)
	errCreateResourceMsg = api.MakeErrorMessage("create", resourceResName)
	errGetResourceMsg    = api.MakeErrorMessage("get", resourceResName)
	errUpdateResourceMsg = api.MakeErrorMessage("update", resourceResName)
	errDeleteResourceMsg = api.MakeErrorMessage("delete", resourceResName)
)

// IndexResourcesV1 - Renders a list containing all resources.
// Handler for HTTP Get - "/resources"
func IndexResourcesV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Repo
	resourceRepo := repo.MakeResourceRepoTx(tx)
	resources, err := resourceRepo.GetAll()
	if err != nil {
		return resourceErrHR(api.WrapError(err), errListResourcesMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		return resourceErrHR(api.WrapError(err), errListResourcesMsg)
	}
	// Ok
	res := ResourceListResource{Data: resources}
	return api.OkHR(createdSt, res, okCreateResourceMsg)
}

// GetResourceV1 - Shows a Resource.
// Handler for HTTP Get - "/resources/{id}"
func GetResourceV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return resourceErrHR(api.WrapError(err), errGetResourceMsg)
	}
	// Repo
	resourceRepo := repo.MakeResourceRepoTx(tx)
	resource, err := resourceRepo.Get(id)
	if err != nil {
		return resourceErrHR(api.WrapError(err), errGetResourceMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		return resourceErrHR(api.WrapError(err), errGetResourceMsg)
	}
	// Ok
	res := ResourceResource{Data: *resource}
	return api.OkHR(okSt, res, okGetResourceMsg)
}

// CreateResourceV1 - Create new Resource API v1.
// Handler for HTTP Post - "/resources"
func CreateResourceV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Decode
	var resourceRes ResourceResource
	err := json.NewDecoder(r.Body).Decode(&resourceRes)
	if err != nil {
		return resourceErrHR(api.WrapError(err), errCreateResourceMsg)
	}
	resource := &resourceRes.Data
	// Services
	resourceSvc := services.ResourceService{Tx: tx}
	resourceSvc.Create(resource)
	err = tx.Commit()
	// Error
	if err != nil {
		return resourceErrHR(api.WrapError(err), errCreateResourceMsg)
	}
	// Ok
	res := ResourceResource{Data: *resource}
	return api.OkHR(createdSt, res, okCreateResourceMsg)
}

// UpdateResourceV1 - Update a Resource.
// Handler for HTTP Patch/Put - "/resources/{id}"
func UpdateResourceV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return resourceErrHR(api.WrapError(err), errGetResourceMsg)
	}
	// Decode
	var resourceRes ResourceResource
	err = json.NewDecoder(r.Body).Decode(&resourceRes)
	if err != nil {
		return resourceErrHR(api.WrapError(err), errCreateResourceMsg)
	}
	resource := &resourceRes.Data
	resource.SetID(id)
	// Services
	resourceSvc := services.ResourceService{Tx: tx}
	resourceSvc.Update(resource)
	err = tx.Commit()
	// Error
	if err != nil {
		return resourceErrHR(api.WrapError(err), errUpdateResourceMsg)
	}
	// Ok
	res := ResourceResource{Data: *resource}
	return api.OkHR(noContentSt, res, okUpdateResourceMsg)
}

// DeleteResourceV1 - Delete Resource.
// Handler for HTTP Delete - "/resources/{id}"
func DeleteResourceV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return resourceErrHR(api.WrapError(err), errGetResourceMsg)
	}
	// Repo
	resourceRepo := repo.MakeResourceRepoTx(tx)
	resource, err := resourceRepo.Get(id)
	if err != nil {
		return resourceErrHR(api.WrapError(err), errDeleteResourceMsg)
	}
	resourceRepo.Delete(resource.ID.UUID)
	err = tx.Commit()
	// Error
	if err != nil {
		return resourceErrHR(api.WrapError(err), errDeleteResourceMsg)
	}
	// Ok
	res := ResourceResource{Data: *resource}
	return api.OkHR(noContentSt, res, okDeleteResourceMsg)
}

// Private
// resourceErrHR - Default ResourceErrorHandlerResult for errors.
func resourceErrHR(err error, msg string) api.HandlerResult {
	return api.DefErrorHR(err, msg)
}
