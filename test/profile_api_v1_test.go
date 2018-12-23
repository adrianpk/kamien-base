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
	apiv1 = "v1"
)

// clearDB() // Clears DB
// resetDB() // Clears DB + Load Fixtures

func TestIndexProfileAPIV1(t *testing.T) {
	clearDB()
	profile := createSampleProfile()
	profile2 := createSampleProfile2()
	profileAPIURL := buildAPIURL(apiv1, common.ProfileRoot)
	// t.Log(profileAPIURL)
	// Request & Response
	req := buildRequest(profileAPIURL, GET)
	authReq(req, admin, "admin")
	res := executeRequest(req)
	b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
	if !matchString(profile.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
	if !matchString(profile2.Name.String, b) {
		t.Errorf("Response differs from expected.")
	}
}

func TestGetProfileAPIV1(t *testing.T) {
	clearDB()
	// Create sample profile
	profile := createSampleProfile()
	profileAPIURL := buildResAPIURL(apiv1, common.ProfileRoot, profile.ID.UUID)
	// t.Log(profileAPIURL)
	// Request & Response
	req := buildRequest(profileAPIURL, GET)
	// tbp.AuthorizeRequest(request, profile1, profile1Name, profile1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
}
func TestCreateProfileAPIV1(t *testing.T) {
	clearDB()
	// URL
	profileAPIURL := buildAPIURL(apiv1, common.ProfileRoot)
	// Values
	email := "EmaiOMxX"
	bio := "BioeEqKP"
	moto := "Motopobi"
	website := "WebsROGg"
	aniversaryDate := "2018-11-30T01:18:27+01:00"
	host := "HostWgtM"
	avatarPath := "AvatlINN"
	headerPath := "HeadGMnc"
	ownerID := ""
	name := "NameSChN"
	description := "DesckEDD"
	startsAt := "2018-11-30T01:18:27+01:00"
	endsAt := "2018-11-30T01:18:27+01:00"
	// JSON
	profileJSONFmt := `
	{
		"data": {
			"email": "%s",
			"bio": "%s",
			"moto": "%s",
			"website": "%s",
			"aniversary-date": "%s",
			"host": "%s",
			"avatar-path": "%s",
			"header-path": "%s",
			"owner-id": "%s",
			"name": "%s",
			"description": "%s",
			"starts-at": "%s",
			"ends-at": "%s"
		}
	}
	`
	profileJSON := fmt.Sprintf(profileJSONFmt, email, bio, moto, website, aniversaryDate, host, avatarPath, headerPath, ownerID, name, description, startsAt, endsAt)
	// Request & Response
	req := makeJSONPostReq(profileAPIURL, profileJSON)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 201-StatusCreated", res.StatusCode)
	}
	if !checkProfileByName(name) {
		t.Errorf("Profile not created.")
		return
	}
}

func TestUpdateProfileAPIV1(t *testing.T) {
	clearDB()
	// Create sample profile
	profile := createSampleProfile()
	profileAPIURL := buildResAPIURL(apiv1, common.ProfileRoot, profile.ID.UUID)
	// Updated properties
	email := "EmaiOMxX"
	bio := "BioeEqKP"
	moto := "Motopobi"
	website := "WebsROGg"
	aniversaryDate := "2018-11-30T01:18:27+01:00"
	host := "HostWgtM"
	avatarPath := "AvatlINN"
	headerPath := "HeadGMnc"
	ownerID := ""
	name := "NameSChN"
	description := "DesckEDD"
	startsAt := "2018-11-30T01:18:27+01:00"
	endsAt := "2018-11-30T01:18:27+01:00"
	// JSON
	profileJSONFmt := `
	{
		"data": {
			"email": "%s",
			"bio": "%s",
			"moto": "%s",
			"website": "%s",
			"aniversary-date": "%s",
			"host": "%s",
			"avatar-path": "%s",
			"header-path": "%s",
			"owner-id": "%s",
			"name": "%s",
			"description": "%s",
			"starts-at": "%s",
			"ends-at": "%s"
		}
	}
	`
	profileJSON := fmt.Sprintf(profileJSONFmt, email, bio, moto, website, aniversaryDate, host, avatarPath, headerPath, ownerID, name, description, startsAt, endsAt)
	// Request & Response
	req := makeJSONPutReq(profileAPIURL, profileJSON)
	// tbp.AuthorizeRequest(request, profile, profile1Name, profile1Role)
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	tc := getProfile(profile.ID.UUID)
	if profilesMatch(profile, tc) {
		error := fmt.Sprintf("Name: '%s' | Expected: '%s' - ", profile.Name.String, name)
		error += fmt.Sprintf("Email: '%s' | Expected: '%s' - ", profile.Email.String, email)
		error += fmt.Sprintf("Bio: '%s' | Expected: '%s' - ", profile.Bio.String, bio)
		error += fmt.Sprintf("Moto: '%s' | Expected: '%s' - ", profile.Moto.String, moto)
		error += fmt.Sprintf("Website: '%s' | Expected: '%s' - ", profile.Website.String, website)
		error += fmt.Sprintf("AniversaryDate: '%s' | Expected: '%s' - ", profile.AniversaryDate.Time, aniversaryDate)
		error += fmt.Sprintf("Host: '%s' | Expected: '%s' - ", profile.Host.String, host)
		error += fmt.Sprintf("AvatarPath: '%s' | Expected: '%s' - ", profile.AvatarPath.String, avatarPath)
		error += fmt.Sprintf("HeaderPath: '%s' | Expected: '%s' - ", profile.HeaderPath.String, headerPath)
		error += fmt.Sprintf("OwnerID: '%s' | Expected: '%s' - ", profile.OwnerID.UUID, ownerID)
		error += fmt.Sprintf("Name: '%s' | Expected: '%s' - ", profile.Name.String, name)
		error += fmt.Sprintf("Description: '%s' | Expected: '%s' - ", profile.Description.String, description)
		error += fmt.Sprintf("StartsAt: '%s' | Expected: '%s' - ", profile.StartsAt.Time, startsAt)
		error += fmt.Sprintf("EndsAt: '%s' | Expected: '%s' - ", profile.EndsAt.Time, endsAt)
		t.Error(error)
		t.Errorf("Profile not updated correctly.")
	}
}

func TestDeleteProfileAPIV1(t *testing.T) {
	clearDB()
	// Create sample profile
	profile := createSampleProfile()
	profileAPIURL := buildResAPIURL(apiv1, common.ProfileRoot, profile.ID.UUID)
	// t.Log(profileAPIURL)
	// Request & Response
	req := buildRequest(profileAPIURL, "DELETE")
	authReq(req, admin, "admin")
	res := executeRequest(req)
	// b := extractBody(res)
	// t.Log(b)
	// Assertions
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusOK", res.StatusCode)
	}
	if checkProfile(profile.ID.UUID) {
		t.Errorf("Profile not deleted.")
	}
}
