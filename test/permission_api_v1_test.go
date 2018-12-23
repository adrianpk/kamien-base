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

func TestIndexPermissionAPIV1(t *testing.T) {
	clearDB()
	permission := createSamplePermission()
	permission2 := createSamplePermission2()
	permissionAPIURL := buildAPIURL(apiv1, common.PermissionRoot)
	// t.Log(permissionAPIURL)
	// Request & Response
	req := buildRequest(permissionAPIURL, GET)
	authReq(req, admin, "admin")
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
	if !matchString(permission.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(permission2.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestGetPermissionAPIV1(t *testing.T) {
	clearDB()
	// Create sample permission
	permission := createSamplePermission()
	permissionAPIURL := buildResAPIURL(apiv1, common.PermissionRoot, permission.ID.UUID)
	// t.Log(permissionAPIURL)
	// Request & Response
	req := buildRequest(permissionAPIURL, GET)
	// tbp.AuthorizeRequest(request, permission1, permission1Name, permission1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
}
func TestCreatePermissionAPIV1(t *testing.T) {
	clearDB()
	// URL
	permissionAPIURL := buildAPIURL(apiv1, common.PermissionRoot)
	// Values
	organizationID := ""
	name := "NameIhlH"
	description := "DescgdSb"
	// JSON
	permissionJSONFmt := `
	{
		"data": {
			"organization-id": "%s",
			"name": "%s",
			"description": "%s"
		}
	}
	`
	permissionJSON := fmt.Sprintf(permissionJSONFmt, organizationID, name, description)
	// Request & Response
	req := makeJSONPostReq(permissionAPIURL, permissionJSON)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 201-StatusCreated", res.StatusCode)
	}
	if !checkPermissionByName(name) {
		t.Errorf("Permission not created.")
		return
	}
}

func TestUpdatePermissionAPIV1(t *testing.T) {
	clearDB()
	// Create sample permission
	permission := createSamplePermission()
	permissionAPIURL := buildResAPIURL(apiv1, common.PermissionRoot, permission.ID.UUID)
	// Updated properties
	organizationID := ""
	name := "NameIhlH"
	description := "DescgdSb"
	// JSON
	permissionJSONFmt := `
	{
		"data": {
			"organization-id": "%s",
			"name": "%s",
			"description": "%s"
		}
	}
	`
	permissionJSON := fmt.Sprintf(permissionJSONFmt, organizationID, name, description)
	// Request & Response
	req := makeJSONPutReq(permissionAPIURL, permissionJSON)
	// tbp.AuthorizeRequest(request, permission, permission1Name, permission1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	tc := getPermission(permission.ID.UUID)
	if permissionsMatch(permission, tc) {
		error := fmt.Sprintf("Name: '%s' | Expected: '%s' - ", permission.Name.String, name)
		error += fmt.Sprintf("OrganizationID: '%s' | Expected: '%s' - ", permission.OrganizationID.UUID, organizationID)
		error += fmt.Sprintf("Name: '%s' | Expected: '%s' - ", permission.Name.String, name)
		error += fmt.Sprintf("Description: '%s' | Expected: '%s' - ", permission.Description.String, description)
		t.Error(error)
		t.Errorf("Permission not updated correctly.")
	}
}

func TestDeletePermissionAPIV1(t *testing.T) {
	clearDB()
	// Create sample permission
	permission := createSamplePermission()
	permissionAPIURL := buildResAPIURL(apiv1, common.PermissionRoot, permission.ID.UUID)
	// t.Log(permissionAPIURL)
	// Request & Response
	req := buildRequest(permissionAPIURL, "DELETE")
	// tbp.AuthorizeRequest(request, permission1, permission1Name, permission1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if checkPermission(permission.ID.UUID) {
		t.Errorf("Permission not deleted.")
	}
}
