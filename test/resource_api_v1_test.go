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

func TestIndexResourceAPIV1(t *testing.T) {
	clearDB()
	resource := createSampleResource()
	resource2 := createSampleResource2()
	resourceAPIURL := buildAPIURL(apiv1, common.ResourceRoot)
	// t.Log(resourceAPIURL)
	// Request & Response
	req := buildRequest(resourceAPIURL, GET)
	authReq(req, admin, "admin")
	res := executeRequest(req)
	b := extractBody(res)
	t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
	if !matchString(resource.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(resource2.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestGetResourceAPIV1(t *testing.T) {
	clearDB()
	// Create sample resource
	resource := createSampleResource()
	resourceAPIURL := buildResAPIURL(apiv1, common.ResourceRoot, resource.ID.UUID)
	// t.Log(resourceAPIURL)
	// Request & Response
	req := buildRequest(resourceAPIURL, GET)
	// tbp.AuthorizeRequest(request, resource1, resource1Name, resource1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
}
func TestCreateResourceAPIV1(t *testing.T) {
	clearDB()
	// URL
	resourceAPIURL := buildAPIURL(apiv1, common.ResourceRoot)
	// Values
	tag := "TagKMyAd"
	organizationID := ""
	name := "NamewpmE"
	description := "DescttlX"
	// JSON
	resourceJSONFmt := `
	{
		"data": {
			"tag": "%s",
			"organization-id": "%s",
			"name": "%s",
			"description": "%s"
		}
	}
	`
	resourceJSON := fmt.Sprintf(resourceJSONFmt, tag, organizationID, name, description)
	// Request & Response
	req := makeJSONPostReq(resourceAPIURL, resourceJSON)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 201-StatusCreated", res.StatusCode)
	}
	if !checkResourceByName(name) {
		t.Errorf("Resource not created.")
		return
	}
}

func TestUpdateResourceAPIV1(t *testing.T) {
	clearDB()
	// Create sample resource
	resource := createSampleResource()
	resourceAPIURL := buildResAPIURL(apiv1, common.ResourceRoot, resource.ID.UUID)
	// Updated properties
	tag := "TagKMyAd"
	organizationID := ""
	name := "NamewpmE"
	description := "DescttlX"
	// JSON
	resourceJSONFmt := `
	{
		"data": {
			"tag": "%s",
			"organization-id": "%s",
			"name": "%s",
			"description": "%s"
		}
	}
	`
	resourceJSON := fmt.Sprintf(resourceJSONFmt, tag, organizationID, name, description)
	// Request & Response
	req := makeJSONPutReq(resourceAPIURL, resourceJSON)
	// tbp.AuthorizeRequest(request, resource, resource1Name, resource1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	tc := getResource(resource.ID.UUID)
	if resourcesMatch(resource, tc) {
		error := fmt.Sprintf("Name: '%s' | Expected: '%s' - ", resource.Name.String, name)
		error += fmt.Sprintf("Tag: '%s' | Expected: '%s' - ", resource.Tag.String, tag)
		error += fmt.Sprintf("OrganizationID: '%s' | Expected: '%s' - ", resource.OrganizationID.UUID, organizationID)
		error += fmt.Sprintf("Name: '%s' | Expected: '%s' - ", resource.Name.String, name)
		error += fmt.Sprintf("Description: '%s' | Expected: '%s' - ", resource.Description.String, description)
		t.Error(error)
		t.Errorf("Resource not updated correctly.")
	}
}

func TestDeleteResourceAPIV1(t *testing.T) {
	clearDB()
	// Create sample resource
	resource := createSampleResource()
	resourceAPIURL := buildResAPIURL(apiv1, common.ResourceRoot, resource.ID.UUID)
	// t.Log(resourceAPIURL)
	// Request & Response
	req := buildRequest(resourceAPIURL, "DELETE")
	// tbp.AuthorizeRequest(request, resource1, resource1Name, resource1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if checkResource(resource.ID.UUID) {
		t.Errorf("Resource not deleted.")
	}
}
