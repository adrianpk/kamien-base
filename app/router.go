package app

import (
	"fmt"
	"net/http"

	api "github.com/adrianpk/kamien/api"
	ctr "github.com/adrianpk/kamien/controllers"
	"{{.Package}}/boot"
	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

// MainRouter - Application main router
func MainRouter() *mux.Router {
	router = NewRouter()
	router.HandleFunc("/", itWorks())
	AddAuthAPIV1Router(router)
	AddUserAPIV1Router(router)
	AddAccountAPIV1Router(router)
	AddProfileAPIV1Router(router)
	AddResourceAPIV1Router(router)
	AddPermissionAPIV1Router(router)
	AddRoleAPIV1Router(router)
	AddRolePermissionAPIV1Router(router)
	AddUserRoleAPIV1Router(router)
	return router
}

// NewRouter - Creates a new mux.router.
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.KeepContext = true
	return r
}

func addAPIRoute(router *mux.Router, path string, handlerFunction api.HandlerFunction, method string) *mux.Route {
	handler := api.AppHandler{HandlerFunction: handlerFunction}
	return router.Handle(path, handler).Methods(method)
}

func addRoute(router *mux.Router, path string, handlerFunction ctr.HandlerFunction, method string) *mux.Route {
	handler := ctr.AppHandler{boot.{{.AppNamePascalCase}}Context, handlerFunction}
	return router.Handle(path, handler).Methods(method)
}

func itWorks() func(http.ResponseWriter, *http.Request) {
	// logger.Debug("Home.....")
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "text/html")
		head := "<head><title>{{.AppNamePascalCase}} @ " + r.Host + "</title></head>"
		body := "<body><div>{{.AppNamePascalCase}} @ " + r.Host + " is working!</div></body>"
		rw.Write([]byte(fmt.Sprintf("<html>%s%s</html>", head, body)))
	}
}
