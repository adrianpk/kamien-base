package test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"{{.Package}}/boot"
	"{{.Package}}/models"
	uuid "github.com/satori/go.uuid"
	log "github.com/siddontang/go/log"
)

// RequestMethod - Valid request methods
type RequestMethod string

const (
	// GET request method
	GET RequestMethod = "GET"
	// POST request method
	POST RequestMethod = "POST"
	// PUT request method
	PUT RequestMethod = "PUT"
	// PATCH request method
	PATCH RequestMethod = "PATCH"
	// DELETE request method
	DELETE RequestMethod = "DELETE"
)

var (
	currentEnv = "test"
	ctHeader   = "Content-Type"
	admin      *models.User
)

func init() {
	conf := GetConfiguration()
	tenv = MakeEnvironment(conf)
	boot.InitForTest(conf)
}

func TestMain(m *testing.M) {
	tenv.Start(m)
	createAdmin()
	os.Exit(m.Run())
}

func clearDB() {
	tenv.clearDatabase()
}

func resetDB() {
	tenv.resetDatabase()
}

func createAdmin() {
	admin = createSampleAdmin()
}

func authReq(req *http.Request, user *models.User, role string) {
	authorizeRequest(req, user.ID.UUID, user.Username.String, role)
}

func authorizeRequest(req *http.Request, userID uuid.UUID, username, role string) {
	tenv.AuthorizeRequest(req, userID.String(), username, role)
}

func setContentType(req *http.Request, contentType string) {
	req.Header.Set(ctHeader, contentType)
}

func buildAPIURL(apiVersion, path string, segments ...string) string {
	u := fmt.Sprintf("http://%s/api/%s/%s", tenv.GetServerAddressPort(), apiVersion, path)
	for _, s := range segments {
		u = fmt.Sprintf("%s/%s", u, s)
	}
	return u
}

func buildResAPIURL(apiVersion, path string, id uuid.UUID) string {
	return fmt.Sprintf("%s/%s", buildAPIURL(apiVersion, path), id)
}

func buildURL(path string, segments ...string) string {
	u := fmt.Sprintf("http://%s/%s", tenv.GetServerAddressPort(), path)
	for _, s := range segments {
		u = fmt.Sprintf("%s/%s", u, s)
	}
	return u
}

func buildResURL(path string, id uuid.UUID) string {
	return fmt.Sprintf("%s/%s", buildURL(path), id)
}

func buildResEditURL(path string, id uuid.UUID) string {
	return fmt.Sprintf("%s/%s/edit", buildURL(path), id)
}

func buildResInitDeleteURL(path string, id uuid.UUID) string {
	return fmt.Sprintf("%s/%s/init-delete", buildURL(path), id)
}

func post(req *http.Request) (*http.Response, error) {
	hc := http.Client{}
	return hc.Do(req)
}

func responseBody(res *http.Response) string {
	resBuf := new(bytes.Buffer)
	resBuf.ReadFrom(res.Body)
	defer res.Body.Close()
	return resBuf.String()
}

func recoverControl() {
	if r := recover(); r != nil {
		log.Info("Recovered.")
	}
}

func matchString(toMatch, body string) bool {
	//fmt.Println("To match:", toMatch)
	//fmt.Println("Body", body)
	rs := fmt.Sprintf("(%s)", toMatch)
	r, _ := regexp.Compile(rs)
	return r.MatchString(body)
}

func buildRequest(targetURL string, rm RequestMethod) *http.Request {
	req := httptest.NewRequest(string(rm), targetURL, nil)
	req.RequestURI = ""
	var err error
	req.URL, err = url.Parse(targetURL)
	checkErr(err)
	return req
}

// TODO: These four following functions must be generalized.
func makeJSONPostReq(targetURL string, jsonStr string) *http.Request {
	formSubmitCT := "application/json"
	encodedReader := strings.NewReader(jsonStr)
	req := httptest.NewRequest("POST", targetURL, encodedReader)
	req.RequestURI = ""
	req.URL, _ = url.Parse(targetURL)
	setContentType(req, formSubmitCT)
	return req
}

func makeJSONPutReq(targetURL string, jsonStr string) *http.Request {
	formSubmitCT := "application/json"
	encodedReader := strings.NewReader(jsonStr)
	req := httptest.NewRequest("PUT", targetURL, encodedReader)
	req.RequestURI = ""
	req.URL, _ = url.Parse(targetURL)
	setContentType(req, formSubmitCT)
	return req
}

func makeFormPostReq(targetURL string, formData url.Values) *http.Request {
	formSubmitCT := "application/x-www-form-urlencoded"
	encodedReader := strings.NewReader(formData.Encode())
	req := httptest.NewRequest("POST", targetURL, encodedReader)
	req.RequestURI = ""
	req.URL, _ = url.Parse(targetURL)
	setContentType(req, formSubmitCT)
	return req
}

func makeFormPutReq(targetURL string, formData url.Values) *http.Request {
	formSubmitCT := "application/x-www-form-urlencoded"
	encodedReader := strings.NewReader(formData.Encode())
	req := httptest.NewRequest("PUT", targetURL, encodedReader)
	req.RequestURI = ""
	req.URL, _ = url.Parse(targetURL)
	setContentType(req, formSubmitCT)
	return req
}

func executeRequest(r *http.Request) *http.Response {
	//Configure TLS, etc.
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		InsecureSkipVerify: true,
	// 	},
	// }
	hc := http.Client{
		// Transport: tr,
		Timeout: 5 * time.Second,
	}
	// r.Header.Add("User-Agent", "Golang Test")
	res, err := hc.Do(r)
	checkErr(err, "Cannot execute request")
	return res
}

func extractBody(r *http.Response) string {
	resBuf := new(bytes.Buffer)
	_, err := resBuf.ReadFrom(r.Body)
	defer r.Body.Close()
	checkErr(err, "Cannot read response body")
	return resBuf.String()
}
