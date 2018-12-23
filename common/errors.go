package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	// ErrDbConnection - Data store error.
	ErrDbConnection = errors.New("Data store error")
	// ErrDataStore - Data store error.
	ErrDataStore = errors.New("Data store error")
	// ErrDataAccess - Data access error.
	ErrDataAccess = errors.New("Data access error")
	// ErrTokenExpired - Access Token is expired.
	ErrTokenExpired = errors.New("Access Token is expired")
	// ErrTokenParsing - Error while parsing the Access Token
	ErrTokenParsing = errors.New("Error while parsing the Access Token")
	// ErrTokenInvalid - Invalid Access Token
	ErrTokenInvalid = errors.New("Invalid Access Token")
	// ErrEntityAlreadySignedUp - Email or Username is already signed up.
	ErrEntityAlreadySignedUp = errors.New("Email or Username is already signed up")
	// ErrRequest - Bad request.
	ErrRequest = errors.New("Bad request")
	// ErrRequestParsing - Error parsing request data.
	ErrRequestParsing = errors.New("Error parsing request data")
	// ErrImageDecoding - Error decoding image data.
	ErrImageDecoding = errors.New("Error decoding image data")
	// ErrResponseMarshalling - Error marshalling response data.
	ErrResponseMarshalling = errors.New("Error marshalling response data")
	// ErrRegistration - Registration error.
	ErrRegistration = errors.New("Registration error")
	// ErrLogin - Login error.
	ErrLogin = errors.New("Login error")
	// ErrLoginInvalidData - Invalid login data.
	ErrLoginInvalidData = errors.New("Invalid login data")
	// ErrLoginDenied - Login denied.
	ErrLoginDenied = errors.New("Login denied")
	// ErrLoginTokenCreate - Error while generating the access token
	ErrLoginTokenCreate = errors.New("Error while generating the access token")
	// ErrLoginSessionCreate - Error while generating session.
	ErrLoginSessionCreate = errors.New("Error while generating session")
	// ErrNotLoggedIn - Not logged in.
	ErrNotLoggedIn = errors.New("Not logged in")
	// ErrUnauthorized - Unauthorized
	ErrUnauthorized = errors.New("Unauthorized")
	// ErrOwnerOnlyCanManage - Error while generating the access token
	ErrOwnerOnlyCanManage = errors.New("Only entity owner are allowed to manage entity")
	// ErrEntityInvalidData - Entity invalid data.
	ErrEntityInvalidData = errors.New("Entity invalid data")
	// ErrEntityNotFound - Entity not found.
	ErrEntityNotFound = errors.New("Entity not found")
	// ErrEntitySelect - Cannot select entities.
	ErrEntitySelect = errors.New("Cannot select entities")
	// ErrEntityCreate - Cannot create user.
	ErrEntityCreate = errors.New("Cannot create entity")
	// ErrEntityUpdate - Cannot update model.
	ErrEntityUpdate = errors.New("Cannot update entity")
	// ErrEntityDelete - Cannot delete entity.
	ErrEntityDelete = errors.New("Cannot delete entity")
	// ErrEntitySetProperty - Cannot set property.
	ErrEntitySetProperty = errors.New("Cannot set property")
	// ErrSearch - Search error.
	ErrSearch = errors.New("Search error")
	// ErrRequestProcessing - Cannot process your request.
	ErrRequestProcessing = errors.New("Cannot process your request")
	// ErrImageProcessing - Error processing image.
	ErrImageProcessing = errors.New("Error processing image")
	// ErrPageNotFoud - Error page not found.
	ErrPageNotFoud = errors.New("Page not found")
	// ErrTemplateExecution - Error template execution.
	ErrTemplateExecution = errors.New("Error executing template")
)

const exposeSourceError = true

type (
	appError struct {
		Error      string `json:"error"`
		Cause      string `json:"cause"`
		HTTPStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
)

// ShowError - Shows error, cause and http code.
func ShowError(rw http.ResponseWriter, err error, cause error, code int) {
	if !exposeSourceError {
		cause = errors.New("[hidden]")
	}
	errObj := appError{
		Error:      err.Error(),
		Cause:      cause.Error(),
		HTTPStatus: code,
	}
	//logger.Debugf("AppError]: %s\n", handlerError)
	// logger.Errorf("%s\n", cause)
	// logger.Error(err)
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		rw.Write(j)
	}
}
