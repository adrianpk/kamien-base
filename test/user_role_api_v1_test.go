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

func TestIndexUserRoleAPIV1(t *testing.T) {
	clearDB()
	userRole := createSampleUserRole()
	userRole2 := createSampleUserRole2()
	userRoleAPIURL := buildAPIURL("v1", common.UserRoleRoot)
	// t.Log(userRoleAPIURL)
	// Request & Response
	req := buildRequest(userRoleAPIURL, GET)
	authReq(req, admin, "admin")
	res := executeRequest(req)
	b := extractBody(res)
	t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
	if !matchString(userRole.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(userRole2.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestGetUserRoleAPIV1(t *testing.T) {
	clearDB()
	// Create sample userRole
	userRole := createSampleUserRole()
	userRoleAPIURL := buildResAPIURL("v1", common.UserRoleRoot, userRole.ID.UUID)
	// t.Log(userRoleAPIURL)
	// Request & Response
	req := buildRequest(userRoleAPIURL, GET)
	// tbp.AuthorizeRequest(request, userRole1, userRole1Name, userRole1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
}
func TestCreateUserRoleAPIV1(t *testing.T) {
	clearDB()
	// URL
	userRoleAPIURL := buildAPIURL(apiv1, common.UserRoleRoot)
	// Values
	organizationID := ""
	userID := ""
	roleID := ""
	name := "NamefbVZ"
	description := "DescdjLR"
	// JSON
	userRoleJSONFmt := `
	{
		"data": {
			"organization-id": "%s",
			"user-id": "%s",
			"role-id": "%s",
			"name": "%s",
			"description": "%s"
		}
	}
	`
	userRoleJSON := fmt.Sprintf(userRoleJSONFmt, organizationID, userID, roleID, name, description)
	// Request & Response
	req := makeJSONPostReq(userRoleAPIURL, userRoleJSON)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 201-StatusCreated", res.StatusCode)
	}
	if !checkUserRoleByName(name) {
		t.Errorf("UserRole not created.")
		return
	}
}

func TestUpdateUserRoleAPIV1(t *testing.T) {
	clearDB()
	// Create sample userRole
	userRole := createSampleUserRole()
	userRoleAPIURL := buildResAPIURL(apiv1, common.UserRoleRoot, userRole.ID.UUID)
	// Updated properties
	organizationID := ""
	userID := ""
	roleID := ""
	name := "NamefbVZ"
	description := "DescdjLR"
	// JSON
	userRoleJSONFmt := `
	{
		"data": {
			"organization-id": "%s",
			"user-id": "%s",
			"role-id": "%s",
			"name": "%s",
			"description": "%s"
		}
	}
	`
	userRoleJSON := fmt.Sprintf(userRoleJSONFmt, organizationID, userID, roleID, name, description)
	// Request & Response
	req := makeJSONPutReq(userRoleAPIURL, userRoleJSON)
	// tbp.AuthorizeRequest(request, userRole, userRole1Name, userRole1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	tc := getUserRole(userRole.ID.UUID)
	if userRolesMatch(userRole, tc) {
		error := fmt.Sprintf("Name: '%s' | Expected: '%s' - ", userRole.Name.String, name)
		error += fmt.Sprintf("OrganizationID: '%s' | Expected: '%s' - ", userRole.OrganizationID.UUID, organizationID)
		error += fmt.Sprintf("UserID: '%s' | Expected: '%s' - ", userRole.UserID.UUID, userID)
		error += fmt.Sprintf("RoleID: '%s' | Expected: '%s' - ", userRole.RoleID.UUID, roleID)
		error += fmt.Sprintf("Name: '%s' | Expected: '%s' - ", userRole.Name.String, name)
		error += fmt.Sprintf("Description: '%s' | Expected: '%s' - ", userRole.Description.String, description)
		t.Error(error)
		t.Errorf("UserRole not updated correctly.")
	}
}

func TestDeleteUserRoleAPIV1(t *testing.T) {
	clearDB()
	// Create sample userRole
	userRole := createSampleUserRole()
	userRoleAPIURL := buildResAPIURL("v1", common.UserRoleRoot, userRole.ID.UUID)
	// t.Log(userRoleAPIURL)
	// Request & Response
	req := buildRequest(userRoleAPIURL, "DELETE")
	// tbp.AuthorizeRequest(request, userRole1, userRole1Name, userRole1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if checkUserRole(userRole.ID.UUID) {
		t.Errorf("UserRole not deleted.")
	}
}
