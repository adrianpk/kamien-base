package app

import (
	"fmt"
	"net/http"

	"{{.Package}}/boot"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

var (
	server  http.Server
	domains = make(Domains)
)

// InitServer - Todo: complete comments.
func InitServer() chan error {
	MakeDomains()
	errs := make(chan error)
	// HTTP server
	go func() {
		p := fmt.Sprintf(":%s", boot.Configuration.GetServerPort())
		log.Infof("HTTP Server initialiting on port: %s...\n", p)
		if err := http.ListenAndServe(p, http.HandlerFunc(redirectToHTTPS)); err != nil {
			log.Infof("HTTPS Server error %+v", err)
			errs <- err
		}
	}()
	// HTTPS server
	go func() {
		p := fmt.Sprintf(":%s", boot.Configuration.GetServerSSLPort())
		log.Infof("HTTPS Server initialiting on port: %s...\n", p)
		// Fix: Check this before release
		// if err := http.ListenAndServeTLS(":8080", "./resources/certificates/cert.pem", "./resources/certificates/key.pem", domains); err != nil {
		if err := http.ListenAndServe(p, domains); err != nil {
			log.Infof("HTTPS Server error %+v", err)
			errs <- err
		}
	}()
	return errs
}

// MakeDomains - Todo: Complete comments
func MakeDomains() {
	p := boot.Configuration.GetServerSSLPort()
	for _, d := range boot.Configuration.GetDomains() {
		m := http.NewServeMux()
		m.Handle("/", MainHandler())
		key := fmt.Sprintf("%s:%s", d, p)
		log.Infof("Domain '%s' configured.", key)
		domains[key] = m
	}
}

// MainHandler - App handler.
func MainHandler() http.Handler {
	n := negroni.Classic()
	// Fix: Check this before release
	// n.Use(CORSProtection())
	// n.UseHandler(CSRFProtection(MainRouter())) // Todo: Must be enabled in production (Env)
	n.UseHandler(MainRouter())
	return n
}

// CSRFProtection - Cross-site request forgery protecction to be used by middleware.
func CSRFProtection(h http.Handler) http.Handler {
	return csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))(h)
}

// CORSProtection - Cross-Origin Resource Sharing protecction to be used by middleware.
func CORSProtection() *cors.Cors {
	var origins []string
	origins = make([]string, 8, 8)
	origins = append(origins, "https://127.0.0.1:8080")
	origins = append(origins, "https://localhost:8080")
	origins = append(origins, "https://www.{{.AppNameLowercase}}100.com:8080")
	origins = append(origins, "https://www.{{.AppNameLowercase}}101.com:8080")
	// origins = append(origins, "http://127.0.0.1:8081")
	// origins = append(origins, "http://localhost:8081")
	// origins = append(origins, "http://www.{{.AppNameLowercase}}100.com:8081")
	// origins = append(origins, "http://www.{{.AppNameLowercase}}101.com:8081")
	return cors.New(cors.Options{
		AllowedOrigins: origins,
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Content-Type", "X-Auth-Token", "Access-Control-Allow-Origin", "X-Requested-With", "Access-Control-Expose-Headers", "Authorization"}})
}

// Router - Application main router
func Router() *mux.Router {
	router := NewRouter()
	router.HandleFunc("/", home())
	return router
}

func home() func(http.ResponseWriter, *http.Request) {
	// log.Info("Home.....")
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "text/html")
		head := "<head><title>{{.AppNamePascalcase}}</title></head>"
		body := "<body><div>{{.AppNamePascalCase}} is working!</div></body>"
		rw.Write([]byte(fmt.Sprintf("<html>%s%s</html>", head, body)))
	}
}

// Domains - Todo: complete comment.
type Domains map[string]http.Handler

func (domains Domains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domain := r.Host
	log.Infof("Request to host %s", domain)
	if mux := domains[domain]; mux != nil {
		mux.ServeHTTP(w, r)
	} else {
		// Handle 404
		http.Error(w, "Not found", 404)
	}
}

func redirectToHTTPS(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Infof("redirect to: %s", target)
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

// InitTestServer - Todo: complete comments.
func InitTestServer() error {
	MakeDomains()
	p := fmt.Sprintf(":%s", boot.Configuration.GetServerPort())
	return http.ListenAndServe(p, domains)
}
