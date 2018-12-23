package app

import (
	"fmt"
	"net/http"

	api "{{.Package}}/api"
	"{{.Package}}/common"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

/* Edit '{{.Package}}'/app/router.go
   func MainRouter() *mux.Router {
		...
    AddAuthRouter(router) // <- Add this line
    return router
  }
*/

// AddAuthAPIV1Router - Initialize and append the auth's API router.
func AddAuthAPIV1Router(parent *mux.Router) {
	// Paths
	authPath := fmt.Sprintf("/api/v1/%s", common.AuthRoot)

	// Routers
	authRouter := NewRouter()
	authSubrouter := authRouter.PathPrefix(authPath).Subrouter()

	// Routes
	// addAPIRoute(authSubrouter, path, controller.function, method")
	// path is a standard sub-URL generated in common package
	// and can be replace by a hardcoded string.
	addAPIRoute(authSubrouter, common.SignupPath(), api.SignUpUserV1, "POST") // create
	addAPIRoute(authSubrouter, common.LoginPath(), api.LogInUserV1, "POST")   // create
	// Middleware
	var methodMiddleware http.HandlerFunc = MethodOverride
	router.PathPrefix(authPath).Handler(negroni.New(
		// negroni.HandlerFunc(boot.Authorize),
		negroni.Wrap(methodMiddleware),
		negroni.Wrap(authRouter),
	))
}
