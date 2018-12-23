package api

import (
	"encoding/json"
	"net/http"

	api "github.com/adrianpk/kamien/api"
	mdl "github.com/adrianpk/kamien/models"
	"{{.Package}}/models"
	"{{.Package}}/repo"
	"{{.Package}}/services"
)

type (
	// UserResource - User wrapper for marshalling
	UserResource struct {
		Data models.User `json:"data"`
	}

	// UserListResource - User list wrapper for marshalling
	UserListResource struct {
		Data []models.User `json:"data"`
	}

	// LoginResource - Login resource
	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	// LoginModel for authentication
	LoginModel struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// AuthUserResource for authorized user Post - /users/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	// AuthUserModel for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)

const (
	// Resource
	userRes     = "user"
	userResName = "User"
)

var (
	okCreateUserMsg     = api.MakeOkMessage("created", userResName)
	okGetUserUserMsg    = api.MakeOkMessage("retrieved", userResName)
	okUpdateUserUserMsg = api.MakeOkMessage("updated", userResName)
	okDeleteUserUserMsg = api.MakeOkMessage("deleted", userResName)
	okLoginUserUserMsg  = api.MakeOkMessage("logged in", userResName)
	errListUsersMsg     = api.MakeErrorMessage("list", userResName)
	errCreateUserMsg    = api.MakeErrorMessage("create", userResName)
	errGetUserMsg       = api.MakeErrorMessage("get", userResName)
	errUpdateUserMsg    = api.MakeErrorMessage("update", userResName)
	errDeleteUserMsg    = api.MakeErrorMessage("delete", userResName)
	errLoginUserMsg     = api.MakeErrorMessage("login", userResName)
)

// IndexUsersV1 - Renders a list containing all users.
// Handler for HTTP Get - "/users"
func IndexUsersV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Repo
	userRepo := repo.MakeUserRepoTx(tx)
	users, _ := userRepo.GetAll()
	err := tx.Commit()
	// Error
	if err != nil {
		return userErrHR(api.WrapError(err), errListUsersMsg)
	}
	// Ok
	res := UserListResource{Data: users}
	return api.OkHR(createdSt, res, okCreateUserMsg)
}

// GetUserV1 - Shows a user.
// Handler for HTTP Get - "/users/{id}"
func GetUserV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return userErrHR(api.WrapError(err), errGetUserMsg)
	}
	// Repo
	userRepo := repo.MakeUserRepoTx(tx)
	user, _ := userRepo.Get(id)
	err = tx.Commit()
	// Error
	if err != nil {
		return userErrHR(api.WrapError(err), errGetUserMsg)
	}
	// Ok
	res := UserResource{Data: *user}
	return api.OkHR(okSt, res, okGetUserUserMsg)
}

// CreateUserV1 - Create new User API v1.
// Handler for HTTP Post - "/users"
func CreateUserV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Decode
	var userRes UserResource
	err := json.NewDecoder(r.Body).Decode(&userRes)
	if err != nil {
		return userErrHR(api.WrapError(err), errCreateUserMsg)
	}
	user := &userRes.Data
	// Services
	userSvc := services.UserService{Tx: tx}
	userSvc.Create(user)
	tx.Commit()
	// Error
	if err != nil {
		return userErrHR(api.WrapError(err), errCreateUserMsg)
	}
	// Ok
	user.ClearPassword()
	res := UserResource{Data: *user}
	return api.OkHR(createdSt, res, okCreateUserMsg)
}

// UpdateUserV1 - Update a user.
// Handler for HTTP Patch/Put - "/users/{id}"
func UpdateUserV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return userErrHR(api.WrapError(err), errGetUserMsg)
	}
	// Decode
	var userRes UserResource
	err = json.NewDecoder(r.Body).Decode(&userRes)
	if err != nil {
		return userErrHR(api.WrapError(err), errCreateUserMsg)
	}
	user := &userRes.Data
	// Repo
	userRepo := repo.MakeUserRepoTx(tx)
	user.SetID(id)
	_ = userRepo.Update(user)
	err = tx.Commit()
	// Error
	if err != nil {
		return userErrHR(api.WrapError(err), errUpdateUserMsg)
	}
	// Ok
	res := UserResource{Data: *user}
	return api.OkHR(noContentSt, res, okUpdateUserUserMsg)
}

// DeleteUserV1 - Delete user.
// Handler for HTTP Delete - "/users/{id}"
func DeleteUserV1(rw http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	id, err := api.GetID(r)
	if err != nil {
		return userErrHR(api.WrapError(err), errGetUserMsg)
	}
	// Repo
	userRepo := repo.MakeUserRepoTx(tx)
	user, _ := userRepo.Get(id)
	userRepo.Delete(user.ID.UUID)
	err = tx.Commit()
	// Error
	if err != nil {
		return userErrHR(api.WrapError(err), errDeleteUserMsg)
	}
	// Ok
	res := UserResource{Data: *user}
	return api.OkHR(noContentSt, res, okDeleteUserUserMsg)
}

// Non common REST functions

// SignUpUserV1 - SignUp a new User.
// Handler for HTTP Post - "/users/signup"
func SignUpUserV1(w http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Decode
	var userRes UserResource
	err := json.NewDecoder(r.Body).Decode(&userRes)
	if err != nil {
		return userErrHR(api.WrapError(err), errCreateUserMsg)
	}
	user := &userRes.Data
	// Services
	userSvc := services.UserService{Tx: tx}
	userSvc.Create(user)
	tx.Commit()
	// Error
	if err != nil {
		return userErrHR(api.WrapError(err), errCreateUserMsg)
	}
	// Ok
	user.ClearPassword()
	res := UserResource{Data: *user}
	return api.OkHR(createdSt, res, okCreateUserMsg)
}

// LogInUserV1 - Authenticates the HTTP request with username and password.
// Handler for HTTP Post - "/users/login"
func LogInUserV1(w http.ResponseWriter, r *http.Request) api.HandlerResult {
	tx := api.RequestTx(r)
	// Decode
	var loginRes LoginResource
	// debugRequestBody(r)
	err := json.NewDecoder(r.Body).Decode(&loginRes)
	if err != nil {
		return userErrHR(api.WrapError(err), errLoginUserMsg)
	}
	loginModel := loginRes.Data
	loginUser := models.User{
		Username:       mdl.ToNullsString(loginModel.Username),
		Email:          mdl.ToNullsString(loginModel.Email),
		Authentication: mdl.Authentication{Password: loginModel.Password},
	}
	// Services
	userSvc := services.UserService{Tx: tx}
	userSvc.Login(&loginUser)
	tx.Commit()
	// Error
	if err != nil {
		return userErrHR(api.WrapError(err), errLoginUserMsg)
	}
	// Ok
	res := UserResource{Data: loginUser}
	return api.OkHR(okSt, res, okLoginUserUserMsg)
}

// Private
// userErrHR - Default User Error Handler Result for errors.
func userErrHR(err error, msg string) api.HandlerResult {
	return api.DefErrorHR(err, msg)
}
