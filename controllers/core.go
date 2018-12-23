package controllers

import (
	"net/http"

	"bytes"
	"fmt"
	"strings"

	"{{.Package}}/boot"
	"{{.Package}}/models"
	"github.com/adrianpk/kamien"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

const (
	// External templates
	useExtTemplates = false
	// Session
	sessionSecret = "{{.AppNameLowercase}}-secret"
	sessionName   = "{{.AppNameLowercase}}-session"
	cookieName    = "{{.AppNameLowercase}}-session"
	// Misc
	formTrue = "true"
)

var (
	log *kamien.Logger
	// Session
	sessionStore *sessions.CookieStore
)

func init() {
	// Logger
	initLogger()
	// Session
	sessionStore = sessions.NewCookieStore([]byte(sessionSecret))
}

func initLogger() {
	logLevel := boot.Configuration.GetLogLevel()
	log = kamien.GetLogger(logLevel, boot.Env)
}

// Session
func setSession(rw http.ResponseWriter, r *http.Request, user models.User, remember bool) error {
	// Gen and configure session
	session, err := sessionStore.Get(r, sessionName)
	if err != nil {
		return err
	}
	configureSession(session, remember)
	// Set values
	session.Values["userID"] = user.ID.UUID.String()
	err = session.Save(r, rw)
	if err != nil {
		return err
	}
	return nil
}

func clearSession(rw http.ResponseWriter, r *http.Request) error {
	session, err := sessionStore.Get(r, sessionName)
	if err != nil {
		return err
	}
	// Invalidate session values
	session.Values["userID"] = nil
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, rw)
	if err != nil {
		return err
	}
	return nil
}

func configureSession(session *sessions.Session, remember bool) {
	mins := 20
	if remember {
		mins = 10080 // 1 Week
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   mins * 60,
		HttpOnly: true,
	}
}

// Form
func getFormDecoder(ignoreUnknownKeys bool) *schema.Decoder {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(ignoreUnknownKeys)
	return decoder
}

func inputIsTrue(r *http.Request, field string) bool {
	log.Debugf("Remember is %s", r.PostFormValue("remember"))
	return r.PostFormValue(field) == formTrue
}

// Debug

func debugRequestHeader(r *http.Request) {
	log.Debugf("Request header:\n%v", r.Header)
}

func debugRequestBody(r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	str := buf.String()
	log.Debugf("Request body\n%s", str)
}

func debugRequest(r *http.Request) {
	log.Debug(formatRequest(r))
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	// If this is a POST, add post data
	if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
