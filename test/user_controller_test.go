/**
 * Copyright (c) 2018 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package test

import (
	"net/http"
	"net/url"
	"testing"

	_ "github.com/lib/pq"
	"{{.PackageName}}/common"
)

var (
	userURL string
)

func TestIndexUser(t *testing.T) {
	clearDB()
	userURL = buildURL(common.UserRoot)
	// Request & Response
	req := buildRequest(userURL, GET)
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Users Index", b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestEditUser(t *testing.T) {
	clearDB()
	// Create sample user
	user := createSampleUser()
	userURL = buildResEditURL(common.UserRoot, user.ID.UUID)
	// Request & Response
	req := buildRequest(userURL, GET)
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Edit User", b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(user.Username.String, b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestNewUser(t *testing.T) {
	clearDB()
	userURL = buildURL(common.UserRoot, "new")
	// logger.Println(userURL)
	// Request & Response
	req := buildRequest(userURL, GET)
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("New User", b) {
		t.Errorf("Response differs from expected.")
	}
}
func TestShowUser(t *testing.T) {
	clearDB()
	// Create sample user
	user := createSampleUser()
	userURL = buildResURL(common.UserRoot, user.ID.UUID)
	// logger.Println(userURL)
	// Request & Response
	req := buildRequest(userURL, GET)
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Show User", b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(user.Username.String, b) {
		t.Errorf("Response differs from expected.")
	}
}
func TestCreateUser(t *testing.T) {
	clearDB()
	userURL = buildURL(common.UserRoot)
	username := "username"
	// Form
	userForm := url.Values{
		"username": {username},
		"email":    {"username@gmail.com"},
		"password": {"password"}}
	// Request & Response
	req := makeFormPostReq(userURL, userForm)
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Users Index", b) {
		t.Errorf("Response differs from expected.")
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

func TestUpdateUser(t *testing.T) {
	clearDB()
	// Create sample user
	user := createSampleUser()
	// Updated properties
	updUsername := "usernameupd"
	updEmail := "usernameupd@gmail.com"
	updGivenName := "Givenupd"
	updMiddleNames := "Middlesupd"
	updFamilyName := "Familyupd"
	// Copy
	userURL = buildResURL(common.UserRoot, user.ID.UUID)
	// Form
	userForm := url.Values{
		"_method":      {string(PUT)},
		"username":     {updUsername},
		"email":        {updEmail},
		"given-name":   {updGivenName},
		"middle-names": {updMiddleNames},
		"family-name":  {updFamilyName},
	}
	// Request & Response
	req := makeFormPutReq(userURL, userForm)
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Users Index", b) {
		t.Errorf("Response differs from expected.")
	}
	tc := getUser(user.ID.UUID)
	if usersMatch(user, tc) {
		t.Errorf("User not updated.")
	}
}

func TestInitDeleteUser(t *testing.T) {
	clearDB()
	// Create sample user
	user := createSampleUser()
	userURL = buildResInitDeleteURL(common.UserRoot, user.ID.UUID)
	// logger.Println(userURL)
	// Request & Response
	req := buildRequest(userURL, POST)
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Delete User", b) {
		t.Errorf("Response differs from expected.")
	}
	if !checkUser(user.ID.UUID) {
		t.Errorf("User has been deleted.")
	}
}
func TestDeleteUser(t *testing.T) {
	clearDB()
	// Create sample user
	user := createSampleUser()
	userURL = buildResURL(common.UserRoot, user.ID.UUID)
	// logger.Println(userURL)
	// Form
	userForm := url.Values{
		"_method": {string(DELETE)},
		"id":      {user.ID.UUID.String()},
	}
	// Request & Response
	req := makeFormPostReq(userURL, userForm)
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("User Index", b) {
		t.Errorf("Response differs from expected.")
	}
	if checkUser(user.ID.UUID) {
		t.Errorf("User not deleted.")
	}
}
