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

// clearDB() // Clears DB
// resetDB() // Clears DB + Load Fixtures

func TestIndexAccountAPIV1(t *testing.T) {
	clearDB()
	user := createSampleUser()
	account := createSampleAccount("user", user, nil)
	user2 := createSampleUser2()
	account2 := createSampleAccount("account", user2, nil)
	// URL
	accountAPIURL := buildAPIURL(apiV1, common.AccountRoot)
	t.Log(accountAPIURL)
	// Request & Response
	req := buildRequest(accountAPIURL, GET)
	authReq(req, admin, "admin")
	res := executeRequest(req)
	b := extractBody(res)
	//t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
	if !matchString(account.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(account2.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestGetAccountAPIV1(t *testing.T) {
	clearDB()
	// Create a sample account owner
	user := createSampleUser()
	// Create a sample parent account
	account := createSampleAccount("user", user, nil)
	// URL
	accountAPIURL := buildResAPIURL(apiV1, common.AccountRoot, account.ID.UUID)
	// t.Log(accountAPIURL)
	// Request & Response
	req := buildRequest(accountAPIURL, GET)
	authReq(req, admin, "admin")
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
	if !matchString(account.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestCreateAccountAPIV1(t *testing.T) {
	clearDB()
	// Create a sample account owner
	user := createSampleUser()
	// Create a sample parent account
	account := createSampleAccount("user", user, nil)
	// URL
	accountAPIURL := buildAPIURL(apiV1, common.AccountRoot)
	ownerID := user.ID.UUID
	parentID := account.ID.UUID
	// JSON
	accountJSONFmt := `
	{
		"data": {
			"ownerID": "%s",
			"parentID": "%s"
		}
	}
	`
	accountJSON := fmt.Sprintf(accountJSONFmt, ownerID.String(), parentID.String())
	// Request & Response
	req := makeJSONPostReq(accountAPIURL, accountJSON)
	//authReq(req, admin, "admin")
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 201-StatusCreated", res.StatusCode)
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

func TestUpdateAccountAPIV1(t *testing.T) {
	clearDB()
	// Create a sample account owner
	user := createSampleUser()
	// Create a sample parent account
	account := createSampleAccount("user", user, nil)
	// Create a sample alternative account owner
	user2 := createSampleUser2()
	// Create a sample alternative parenta account
	account2 := createSampleAccount("user", user2, nil)
	// URL
	accountAPIURL := buildResAPIURL(apiV1, common.AccountRoot, account.ID.UUID)
	// Updated properties
	ownerID := user2.ID.UUID.String()
	parentID := account2.ID.UUID.String()
	t.Logf("New owner should be %s", ownerID)
	t.Logf("New parent should be %s", parentID)
	// JSON
	accountJSONFmt := `
	{
		"data": {
			"ownerID": "%s",
			"parentID": "%s"
		}
	}
	`
	accountJSON := fmt.Sprintf(accountJSONFmt, ownerID, parentID)
	// Request & Response
	req := makeJSONPutReq(accountAPIURL, accountJSON)
	authReq(req, admin, "admin")
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusNoContent", res.StatusCode)
	}
	tc := getAccount(account.ID.UUID)
	if accountsMatch(account, tc) {
		error := fmt.Sprintf("OwnerID: '%s' | Expected: '%s', ", tc.OwnerID.UUID.String(), ownerID)
		error += fmt.Sprintf("ParentID: '%s' | Expected: '%s' ", tc.ParentID.UUID.String(), parentID)
		t.Error(error)
		t.Errorf("Account not updated correctly.")
	}
}

func TestDeleteAccountAPIV1(t *testing.T) {
	clearDB()
	// Create a sample account owner
	user := createSampleUser()
	// Create a sample parent account
	account := createSampleAccount("user", user, nil)
	// URL
	accountAPIURL := buildResAPIURL(apiV1, common.AccountRoot, account.ID.UUID)
	// t.Log(accountAPIURL)
	// Request & Response
	req := buildRequest(accountAPIURL, "DELETE")
	authReq(req, admin, "admin")
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if checkAccount(account.ID.UUID) {
		t.Errorf("Account not deleted.")
	}
}
