package app

import (
	"fmt"
	"net/http"

	api "{{.Package}}/api"
	"{{.Package}}/common"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

/* Edit '{{.Package}}'}}/app/router.go
   func MainRouter() *mux.Router {
		...
    AddaccountRouter(router) // <- Add this line
    return router
  }
*/

// AddProfileAPIV1Router - Initialize and append the account's API router.
func AddProfileAPIV1Router(parent *mux.Router) {
	// Paths
	accountsPath := fmt.Sprintf("/api/v1/%s", common.ProfileRoot)

	// Routers
	accountRouter := NewRouter()
	accountSubrouter := accountRouter.PathPrefix(accountsPath).Subrouter()

	// Routes
	// addAPIRoute(accountSubrouter, path, controller.function, method")
	// path is a standard sub-URL generated in common package
	// and can be replace by a hardcoded string.
	addAPIRoute(accountSubrouter, common.IndexPath(), api.IndexProfilesV1, "GET")     // index
	addAPIRoute(accountSubrouter, common.ShowPath(), api.GetProfileV1, "GET")         // show
	addAPIRoute(accountSubrouter, common.CreatePath(), api.CreateProfileV1, "POST")   // create
	addAPIRoute(accountSubrouter, common.UpdatePath(), api.UpdateProfileV1, "PATCH")  // update
	addAPIRoute(accountSubrouter, common.UpdatePath(), api.UpdateProfileV1, "PUT")    // update
	addAPIRoute(accountSubrouter, common.DeletePath(), api.DeleteProfileV1, "DELETE") // delete
	// Middleware
	var methodMiddleware http.HandlerFunc = MethodOverride
	router.PathPrefix(accountsPath).Handler(negroni.New(
		// negroni.HandlerFunc(boot.Authorize),
		negroni.Wrap(methodMiddleware),
		negroni.Wrap(accountRouter),
	))
}
