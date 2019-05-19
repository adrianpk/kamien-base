/**
 * Copyright (c) 2018 Adrian K <adrian.git@kuguar.dev>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package test

import (
	"net/http"
	"net/url"
	"testing"

	_ "github.com/adrianpk/kamien"
	_ "github.com/lib/pq"
	_ "{{.PackageName}}/app"
	_ "{{.PackageName}}/boot"
	"{{.PackageName}}/common"
)

var (
	accountURL string
)

func TestIndexAccount(t *testing.T) {
	clearDB()
	owner := createSampleUser()
	account := createSampleAccount("user", owner, nil)
	accountURL = buildURL(common.AccountRoot)
	// Request & Response
	req := buildRequest(accountURL, GET)
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Accounts Index", b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(account.ID.UUID.String(), b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestEditAccount(t *testing.T) {
	clearDB()
	// Create sample account
	owner := createSampleUser()
	account := createSampleAccount("user", owner, nil)
	accountURL = buildResEditURL(common.AccountRoot, account.ID.UUID)
	// t.Log(accountURL)
	// Request & Response
	req := buildRequest(accountURL, GET)
	res := executeRequest(req)
	b := extractBody(res)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Edit Account", b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(account.ID.UUID.String(), b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestNewAccount(t *testing.T) {
	clearDB()
	accountURL = buildURL(common.AccountRoot, "new")
	// t.Log(accountURL)
	// Request & Response
	req := buildRequest(accountURL, GET)
	res := executeRequest(req)
	b := extractBody(res)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("New Account", b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestShowAccount(t *testing.T) {
	clearDB()
	// Create sample account
	owner := createSampleUser()
	account := createSampleAccount("user", owner, nil)
	accountURL = buildResURL(common.AccountRoot, account.ID.UUID)
	// t.Log(accountURL)
	// Request & Response
	req := buildRequest(accountURL, GET)
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Show Account", b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(account.ID.UUID.String(), b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestCreateAccount(t *testing.T) {
	clearDB()
	// Create a sample account owner
	user := createSampleUser()
	// Create a sample parent account
	account := createSampleAccount("user", user, nil)
	accountURL = buildURL(common.AccountRoot)
	ownerID := user.ID.UUID
	parentID := account.ID.UUID
	// Form
	accountForm := url.Values{
		"owner-id":  {ownerID.String()},
		"parent-id": {parentID.String()},
	}
	// Request & Response
	req := makeFormPostReq(accountURL, accountForm)
	res := executeRequest(req)
	b := extractBody(res)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Accounts Index", b) {
		t.Errorf("Response differs from expected.")
	}
	if !checkAccountByOwnerID(ownerID) {
		t.Errorf("Cannot find account by its owner")
		return
	}
	if !checkAccountByParentID(parentID) {
		t.Errorf("Cannot find account by its parent account")
		return
	}
}

func TestUpdateAccount(t *testing.T) {
	resetDB()
	// Create a sample account owner
	user := createSampleUser()
	// Create a sample parent account
	account := createSampleAccount("user", user, nil)
	// Create a sample alternative account owner
	user2 := createSampleUser2()
	// Create a sample alternative parenta account
	account2 := createSampleAccount("user", user2, nil)
	// Updated properties
	ownerID := user2.ID.UUID.String()
	parentID := account2.ID.UUID.String()
	// Request & Response
	accountURL = buildResURL(common.AccountRoot, account.ID.UUID)
	// Form
	accountForm := url.Values{
		"_method":      {string(PUT)},
		"account_type": {"organization"},
		"owner-id":     {ownerID},
		"parent-id":    {parentID},
		"email":        {user2.Email.String},
	}
	// Request & Response
	req := makeFormPutReq(accountURL, accountForm)
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Accounts Index", b) {
		t.Errorf("Response differs from expected.")
	}
	tc := getAccount(account.ID.UUID)
	if accountsMatch(account, tc) {
		t.Errorf("Account not updated.")
	}
}

func TestInitDeleteAccount(t *testing.T) {
	clearDB()
	// Create sample account
	owner := createSampleUser()
	account := createSampleAccount("user", owner, nil)
	accountURL = buildResInitDeleteURL(common.AccountRoot, account.ID.UUID)
	// logger.Println(accountURL)
	// Request & Response
	req := buildRequest(accountURL, POST)
	res := executeRequest(req)
	b := extractBody(res)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Delete Account", b) {
		t.Errorf("Response differs from expected.")
	}
	if !checkAccount(account.ID.UUID) {
		t.Errorf("Account has been deleted.")
	}
}

func TestDeleteAccount(t *testing.T) {
	clearDB()
	// Create sample account
	owner := createSampleUser()
	account := createSampleAccount("user", owner, nil)
	accountURL = buildResURL(common.AccountRoot, account.ID.UUID)
	// logger.Println(accountURL)
	// Form
	accountForm := url.Values{
		"_method": {string(DELETE)},
		"id":      {account.ID.UUID.String()},
	}
	// Request & Response
	req := makeFormPostReq(accountURL, accountForm)
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if !matchString("Account Index", b) {
		t.Errorf("Response differs from expected.")
	}
	if checkAccount(account.ID.UUID) {
		t.Errorf("Account not deleted.")
	}
}
