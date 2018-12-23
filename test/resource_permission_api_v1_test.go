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

func TestIndexResourcePermissionAPIV1(t *testing.T) {
	clearDB()
	resourcePermission := createSampleResourcePermission()
	resourcePermission2 := createSampleResourcePermission2()
	resourcePermissionAPIURL := buildAPIURL(apiv1, common.ResourcePermissionRoot)
	// t.Log(resourcePermissionAPIURL)
	// Request & Response
	req := buildRequest(resourcePermissionAPIURL, GET)
	authReq(req, admin, "admin")
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
	if !matchString(resourcePermission.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(resourcePermission2.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestGetResourcePermissionAPIV1(t *testing.T) {
	clearDB()
	// Create sample resourcePermission
	resourcePermission := createSampleResourcePermission()
	resourcePermissionAPIURL := buildResAPIURL(apiv1, common.ResourcePermissionRoot, resourcePermission.ID.UUID)
	// t.Log(resourcePermissionAPIURL)
	// Request & Response
	req := buildRequest(resourcePermissionAPIURL, GET)
	// tbp.AuthorizeRequest(request, resourcePermission1, resourcePermission1Name, resourcePermission1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
}
func TestCreateResourcePermissionAPIV1(t *testing.T) {
	clearDB()
	// URL
	resourcePermissionAPIURL := buildAPIURL(apiv1, common.ResourcePermissionRoot)
	// Values
	organizationID := ""
	resourceID := ""
	permissionID := ""
	name := "NamerBGU"
	description := "Deschyph"
	// JSON
	resourcePermissionJSONFmt := `
	{
		"data": {
			"organization-id": "%s",
			"resource-id": "%s",
			"permission-id": "%s",
			"name": "%s",
			"description": "%s"
		}
	}
	`
	resourcePermissionJSON := fmt.Sprintf(resourcePermissionJSONFmt, organizationID, resourceID, permissionID, name, description)
	// Request & Response
	req := makeJSONPostReq(resourcePermissionAPIURL, resourcePermissionJSON)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 201-StatusCreated", res.StatusCode)
	}
	if !checkResourcePermissionByName(name) {
		t.Errorf("ResourcePermission not created.")
		return
	}
}

func TestUpdateResourcePermissionAPIV1(t *testing.T) {
	clearDB()
	// Create sample resourcePermission
	resourcePermission := createSampleResourcePermission()
	resourcePermissionAPIURL := buildResAPIURL(apiv1, common.ResourcePermissionRoot, resourcePermission.ID.UUID)
	// Updated properties
	organizationID := ""
	resourceID := ""
	permissionID := ""
	name := "NamerBGU"
	description := "Deschyph"
	// JSON
	resourcePermissionJSONFmt := `
	{
		"data": {
			"organization-id": "%s",
			"resource-id": "%s",
			"permission-id": "%s",
			"name": "%s",
			"description": "%s"
		}
	}
	`
	resourcePermissionJSON := fmt.Sprintf(resourcePermissionJSONFmt, organizationID, resourceID, permissionID, name, description)
	// Request & Response
	req := makeJSONPutReq(resourcePermissionAPIURL, resourcePermissionJSON)
	// tbp.AuthorizeRequest(request, resourcePermission, resourcePermission1Name, resourcePermission1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	tc := getResourcePermission(resourcePermission.ID.UUID)
	if resourcePermissionsMatch(resourcePermission, tc) {
		error := fmt.Sprintf("Name: '%s' | Expected: '%s' - ", resourcePermission.Name.String, name)
		error += fmt.Sprintf("OrganizationID: '%s' | Expected: '%s' - ", resourcePermission.OrganizationID.UUID, organizationID)
		error += fmt.Sprintf("ResourceID: '%s' | Expected: '%s' - ", resourcePermission.ResourceID.UUID, resourceID)
		error += fmt.Sprintf("PermissionID: '%s' | Expected: '%s' - ", resourcePermission.PermissionID.UUID, permissionID)
		error += fmt.Sprintf("Name: '%s' | Expected: '%s' - ", resourcePermission.Name.String, name)
		error += fmt.Sprintf("Description: '%s' | Expected: '%s' - ", resourcePermission.Description.String, description)
		t.Error(error)
		t.Errorf("ResourcePermission not updated correctly.")
	}
}

func TestDeleteResourcePermissionAPIV1(t *testing.T) {
	clearDB()
	// Create sample resourcePermission
	resourcePermission := createSampleResourcePermission()
	resourcePermissionAPIURL := buildResAPIURL(apiv1, common.ResourcePermissionRoot, resourcePermission.ID.UUID)
	// t.Log(resourcePermissionAPIURL)
	// Request & Response
	req := buildRequest(resourcePermissionAPIURL, "DELETE")
	// tbp.AuthorizeRequest(request, resourcePermission1, resourcePermission1Name, resourcePermission1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if checkResourcePermission(resourcePermission.ID.UUID) {
		t.Errorf("ResourcePermission not deleted.")
	}
}
