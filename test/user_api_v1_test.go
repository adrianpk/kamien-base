package test

import (
	"fmt"
	// "log"
	"net/http"
	// "strings"
	"testing"

	"{{.Package}}/common"
	_ "github.com/lib/pq"
)

var (
	apiV1 = "v1"
)

// clearDB() // Clears DB
// resetDB() // Clears DB + Load Fixtures

func TestIndexUserAPIV1(t *testing.T) {
	clearDB()
	user := createSampleUser()
	user2 := createSampleUser2()
	userAPIURL := buildAPIURL(apiV1, common.UserRoot)
	// t.Log(userAPIURL)
	// Request & Response
	req := buildRequest(userAPIURL, GET)
	authReq(req, admin, "admin")
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
	if !matchString(user.Username.String, b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(user2.Username.String, b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestGetUserAPIV1(t *testing.T) {
	clearDB()
	// Create sample user
	user := createSampleUser()
	userAPIURL := buildResAPIURL(apiV1, common.UserRoot, user.ID.UUID)
	// t.Log(userAPIURL)
	// Request & Response
	req := buildRequest(userAPIURL, GET)
	// tbp.AuthorizeRequest(request, user1, user1Username, user1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
}

func TestCreateUserAPIV1(t *testing.T) {
	clearDB()
	userAPIURL := buildAPIURL(apiV1, common.UserRoot)
	username := "username"
	password := "sevenseas"
	email := "arthurcurry@gmail.com"
	givenName := "Arthur"
	// JSON
	userJSONFmt := `
	{
		"data": {
			"username": "%s",
			"password": "%s",
			"email": "%s",
			"givenName": "%s"
		}
	}
	`
	userJSON := fmt.Sprintf(userJSONFmt, username, password, email, givenName)
	// Request & Response
	req := makeJSONPostReq(userAPIURL, userJSON)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 201-StatusCreated", res.StatusCode)
	}
	u := getUserByUsername(username)
	if u == nil {
		t.Errorf("User not created.")
		return
	}
	a := getAccountByOwnerID(u.ID.UUID)
	if a == nil {
		t.Errorf("User associated account not created.")
		return
	}
	if !checkProfileByOwnerID(a.ID.UUID) {
		t.Errorf("%+v", a)
		t.Errorf("User associated profile not created.")
	}
}

func TestUpdateUserAPIV1(t *testing.T) {
	clearDB()
	// Create sample user
	user := createSampleUser()
	userAPIURL := buildResAPIURL(apiV1, common.UserRoot, user.ID.UUID)
	// Updated properties
	username := "usernameupd"
	email := "usernameupd@gmail.com"
	password := "mera"
	givenName := "Givenupd"
	middleNames := "Middlesupd"
	familyName := "Familyupd"
	// JSON
	userJSONFmt := `
	{
		"data": {
			"username": "%s",
			"password": "%s",
			"email": "%s",
			"givenName": "%s",
			"middleNames": "%s",
			"familyName": "%s"
		}
	}
	`
	userJSON := fmt.Sprintf(userJSONFmt, username, password, email, givenName, middleNames, familyName)
	// Request & Response
	req := makeJSONPutReq(userAPIURL, userJSON)
	// tbp.AuthorizeRequest(request, user1, user1Username, user1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	tc := getUser(user.ID.UUID)
	if usersMatch(user, tc) {
		error := fmt.Sprintf("Username: '%s' | Expected: '%s' - ", user.Username.String, username)
		error += fmt.Sprintf("Email: '%s' | Expected: '%s' - ", user.Email.String, email)
		error += fmt.Sprintf("GivenName: '%s' | Expected: '%s'", user.GivenName.String, givenName)
		error += fmt.Sprintf("MiddleNames: '%s' | Expected: '%s'", user.MiddleNames.String, middleNames)
		error += fmt.Sprintf("FamilyName: '%s' | Expected: '%s'", user.FamilyName.String, familyName)
		t.Error(error)
		t.Errorf("User not updated correctly.")
	}
}

func TestDeleteUserAPIV1(t *testing.T) {
	clearDB()
	// Create sample user
	user := createSampleUser()
	userAPIURL := buildResAPIURL(apiV1, common.UserRoot, user.ID.UUID)
	// t.Log(userAPIURL)
	// Request & Response
	req := buildRequest(userAPIURL, "DELETE")
	authReq(req, admin, "admin")
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if checkUser(user.ID.UUID) {
		t.Errorf("User not deleted.")
	}
}

func TestSignUpUserAPIV1(t *testing.T) {
	clearDB()
	// signupAPIURL := buildAPIURL(apiV1, "/signup")
	signupAPIURL := buildAPIURL(apiV1, common.AuthRoot, "signup")
	t.Log(signupAPIURL)
	// Values
	username := "username"
	password := "sevenseas"
	email := "arthurcurry@gmail.com"
	givenName := "Arthur"
	// JSON
	userJSONFmt := `
	{
		"data": {
			"username": "%s",
			"password": "%s",
			"email": "%s",
			"givenName": "%s"
		}
	}
	`
	userJSON := fmt.Sprintf(userJSONFmt, username, password, email, givenName)
	// Request & Response
	req := makeJSONPostReq(signupAPIURL, userJSON)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 201-StatusCreated", res.StatusCode)
	}
	u := getUserByUsername(username)
	if u == nil {
		t.Errorf("User not created.")
		return
	}
	a := getAccountByOwnerID(u.ID.UUID)
	if a == nil {
		t.Errorf("User associated account not created.")
		return
	}
	if !checkProfileByOwnerID(a.ID.UUID) {
		t.Errorf("User associated profile not created.")
	}
}
func TestLogInUserAPIV1(t *testing.T) {
	clearDB()
	// Create sample user
	user := createSampleUser()
	loginAPIURL := buildAPIURL(apiV1, common.AuthRoot, "login")
	// t.Log(signupAPIURL)
	// t.Log(loginAPIURL)
	username := user.Username.String
	password := user.Password
	email := user.Email.String
	// JSON
	logingJSONFmt := `
	{
		"data": {
			"username": "%s",
			"password": "%s",
			"email": "%s"
		}
	}
	`
	userJSON := fmt.Sprintf(logingJSONFmt, username, password, email)
	// Request & Response
	req := makeJSONPostReq(loginAPIURL, userJSON)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 201-StatusOK", res.StatusCode)
	}
}
