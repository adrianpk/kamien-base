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

func TestIndexRolePermissionAPIV1(t *testing.T) {
	clearDB()
	rolePermission := createSampleRolePermission()
	rolePermission2 := createSampleRolePermission2()
	rolePermissionAPIURL := buildAPIURL(apiv1, common.RolePermissionRoot)
	// t.Log(rolePermissionAPIURL)
	// Request & Response
	req := buildRequest(rolePermissionAPIURL, GET)
	authReq(req, admin, "admin")
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
	if !matchString(rolePermission.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(rolePermission2.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestGetRolePermissionAPIV1(t *testing.T) {
	clearDB()
	// Create sample rolePermission
	rolePermission := createSampleRolePermission()
	rolePermissionAPIURL := buildResAPIURL(apiv1, common.RolePermissionRoot, rolePermission.ID.UUID)
	// t.Log(rolePermissionAPIURL)
	// Request & Response
	req := buildRequest(rolePermissionAPIURL, GET)
	// tbp.AuthorizeRequest(request, rolePermission1, rolePermission1Name, rolePermission1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
}
func TestCreateRolePermissionAPIV1(t *testing.T) {
	clearDB()
	// URL
	rolePermissionAPIURL := buildAPIURL(apiv1, common.RolePermissionRoot)
	// Values
	organizationID := ""
	roleID := ""
	permissionID := ""
	name := "NameBFJa"
	description := "DescflUe"
	// JSON
	rolePermissionJSONFmt := `
	{
		"data": {
			"organization-id": "%s",
			"role-id": "%s",
			"permission-id": "%s",
			"name": "%s",
			"description": "%s"
		}
	}
	`
	rolePermissionJSON := fmt.Sprintf(rolePermissionJSONFmt, organizationID, roleID, permissionID, name, description)
	// Request & Response
	req := makeJSONPostReq(rolePermissionAPIURL, rolePermissionJSON)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 201-StatusCreated", res.StatusCode)
	}
	if !checkRolePermissionByName(name) {
		t.Errorf("RolePermission not created.")
		return
	}
}

func TestUpdateRolePermissionAPIV1(t *testing.T) {
	clearDB()
	// Create sample rolePermission
	rolePermission := createSampleRolePermission()
	rolePermissionAPIURL := buildResAPIURL(apiv1, common.RolePermissionRoot, rolePermission.ID.UUID)
	// Updated properties
	organizationID := ""
	roleID := ""
	permissionID := ""
	name := "NameBFJa"
	description := "DescflUe"
	// JSON
	rolePermissionJSONFmt := `
	{
		"data": {
			"organization-id": "%s",
			"role-id": "%s",
			"permission-id": "%s",
			"name": "%s",
			"description": "%s"
		}
	}
	`
	rolePermissionJSON := fmt.Sprintf(rolePermissionJSONFmt, organizationID, roleID, permissionID, name, description)
	// Request & Response
	req := makeJSONPutReq(rolePermissionAPIURL, rolePermissionJSON)
	// tbp.AuthorizeRequest(request, rolePermission, rolePermission1Name, rolePermission1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	tc := getRolePermission(rolePermission.ID.UUID)
	if rolePermissionsMatch(rolePermission, tc) {
		error := fmt.Sprintf("Name: '%s' | Expected: '%s' - ", rolePermission.Name.String, name)
		error += fmt.Sprintf("OrganizationID: '%s' | Expected: '%s' - ", rolePermission.OrganizationID.UUID, organizationID)
		error += fmt.Sprintf("RoleID: '%s' | Expected: '%s' - ", rolePermission.RoleID.UUID, roleID)
		error += fmt.Sprintf("PermissionID: '%s' | Expected: '%s' - ", rolePermission.PermissionID.UUID, permissionID)
		error += fmt.Sprintf("Name: '%s' | Expected: '%s' - ", rolePermission.Name.String, name)
		error += fmt.Sprintf("Description: '%s' | Expected: '%s' - ", rolePermission.Description.String, description)
		t.Error(error)
		t.Errorf("RolePermission not updated correctly.")
	}
}

func TestDeleteRolePermissionAPIV1(t *testing.T) {
	clearDB()
	// Create sample rolePermission
	rolePermission := createSampleRolePermission()
	rolePermissionAPIURL := buildResAPIURL(apiv1, common.RolePermissionRoot, rolePermission.ID.UUID)
	// t.Log(rolePermissionAPIURL)
	// Request & Response
	req := buildRequest(rolePermissionAPIURL, "DELETE")
	// tbp.AuthorizeRequest(request, rolePermission1, rolePermission1Name, rolePermission1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if checkRolePermission(rolePermission.ID.UUID) {
		t.Errorf("RolePermission not deleted.")
	}
}
