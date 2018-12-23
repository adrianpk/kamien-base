package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	api "github.com/adrianpk/kamien/api"
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"{{.Package}}/services"
)

type (
	// ProfileResource - Profile wrapper for marshalling
	ProfileResource struct {
		Data models.Profile `json:"data"`
	}

	// ProfileListResource - Profile list wrapper for marshalling
	ProfileListResource struct {
		Data []models.Profile `json:"data"`
	}
)

const (
	// Resource
	profileRes     = "profile"
	profileResName = "Profile"
)

var (
	okCreateProfileMsg  = api.MakeOkMessage("created", profileResName)
	okGetProfileMsg     = api.MakeOkMessage("retrieved", profileResName)
	okUpdateProfileMsg  = api.MakeOkMessage("updated", profileResName)
	okDeleteProfileMsg  = api.MakeOkMessage("deleted", profileResName)
	errListProfilesMsg  = api.MakeErrorMessage("list", profileResName)
	errCreateProfileMsg = api.MakeErrorMessage("create", profileResName)
	errGetProfileMsg    = api.MakeErrorMessage("get", profileResName)
	errUpdateProfileMsg = api.MakeErrorMessage("update", profileResName)
	errDeleteProfileMsg = api.MakeErrorMessage("delete", profileResName)
)

// IndexProfilesV1 - Renders a list containing all profiles.
// Handler for HTTP Get - "/profiles"
func IndexProfilesV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Repo
	profileRepo := repo.MakeProfileRepoTx(tx)
	profiles, err := profileRepo.GetAll()
	if err != nil {
		return profileErrHR(api.WrapError(err), errListProfilesMsg)
	}
	err = tx.Commit()
	// Error
	if err != nil {
		fmt.Println(err)
		return profileErrHR(api.WrapError(err), errListProfilesMsg)
	}
	// Ok
	res := ProfileListResource{Data: profiles}
	return api.OkHR(createdSt, res, okCreateProfileMsg)
}

// GetProfileV1 - Shows a Profile.
// Handler for HTTP Get - "/profiles/{id}"
func GetProfileV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return profileErrHR(api.WrapError(err), errGetProfileMsg)
	}
	// Repo
	profileRepo := repo.MakeProfileRepoTx(tx)
	profile, _ := profileRepo.Get(id)
	err = tx.Commit()
	// Error
	if err != nil {
		return profileErrHR(api.WrapError(err), errGetProfileMsg)
	}
	// Ok
	res := ProfileResource{Data: *profile}
	return api.OkHR(okSt, res, okGetProfileMsg)
}

// CreateProfileV1 - Create new Profile API v1.
// Handler for HTTP Post - "/profiles"
func CreateProfileV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Decode
	var profileRes ProfileResource
	err := json.NewDecoder(r.Body).Decode(&profileRes)
	if err != nil {
		return profileErrHR(api.WrapError(err), errCreateProfileMsg)
	}
	profile := &profileRes.Data
	// Services
	profileSvc := services.ProfileService{Tx: tx}
	profileSvc.Create(profile)
	err = tx.Commit()
	// Error
	if err != nil {
		return profileErrHR(api.WrapError(err), errCreateProfileMsg)
	}
	// Ok
	res := ProfileResource{Data: *profile}
	return api.OkHR(createdSt, res, okCreateProfileMsg)
}

// UpdateProfileV1 - Update a Profile.
// Handler for HTTP Patch/Put - "/profiles/{id}"
func UpdateProfileV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return profileErrHR(api.WrapError(err), errGetProfileMsg)
	}
	// Decode
	var profileRes ProfileResource
	err = json.NewDecoder(r.Body).Decode(&profileRes)
	if err != nil {
		return profileErrHR(api.WrapError(err), errCreateProfileMsg)
	}
	profile := &profileRes.Data
	profile.SetID(id)
	// Services
	profileSvc := services.ProfileService{Tx: tx}
	profileSvc.Update(profile)
	err = tx.Commit()
	// Error
	if err != nil {
		return profileErrHR(api.WrapError(err), errUpdateProfileMsg)
	}
	// Ok
	res := ProfileResource{Data: *profile}
	return api.OkHR(noContentSt, res, okUpdateProfileMsg)
}

// DeleteProfileV1 - Delete Profile.
// Handler for HTTP Delete - "/profiles/{id}"
func DeleteProfileV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return profileErrHR(api.WrapError(err), errGetProfileMsg)
	}
	// Repo
	profileRepo := repo.MakeProfileRepoTx(tx)
	profile, _ := profileRepo.Get(id)
	profileRepo.Delete(profile.ID.UUID)
	err = tx.Commit()
	// Error
	if err != nil {
		return profileErrHR(api.WrapError(err), errDeleteProfileMsg)
	}
	// Ok
	res := ProfileResource{Data: *profile}
	return api.OkHR(noContentSt, res, okDeleteProfileMsg)
}

// Private
// profileErrHR - Default ProfileErrorHandlerResult for errors.
func profileErrHR(err error, msg string) api.HandlerResult {
	return api.DefErrorHR(err, msg)
}
