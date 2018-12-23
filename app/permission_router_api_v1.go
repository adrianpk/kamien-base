package app

import (
	"fmt"
	"net/http"

	api "{{.Package}}/api"
	"{{.Package}}/common"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

/* Edit '{{.Package}}/app/router.go'
   func MainRouter() *mux.Router {
		...
    AddPermissionAPIV1Router(router) // <- Add this line
    return router
  }
*/

// AddPermissionAPIV1Router - Initialize and append the permission's API router.
func AddPermissionAPIV1Router(parent *mux.Router) {
	// Paths
	permissionsPath := fmt.Sprintf("/api/v1/%s", common.PermissionRoot)

	// Routers
	permissionRouter := NewRouter()
	permissionSubrouter := permissionRouter.PathPrefix(permissionsPath).Subrouter()

	// Routes
	// addAPIRoute(permissionSubrouter, path, controller.function, method")
	// path is a standard sub-URL generated in common package
	// and can be replace by a hardcoded string.
	addAPIRoute(permissionSubrouter, common.IndexPath(), api.IndexPermissionsV1, "GET")     // index
	addAPIRoute(permissionSubrouter, common.ShowPath(), api.GetPermissionV1, "GET")         // show
	addAPIRoute(permissionSubrouter, common.CreatePath(), api.CreatePermissionV1, "POST")   // create
	addAPIRoute(permissionSubrouter, common.UpdatePath(), api.UpdatePermissionV1, "PATCH")  // update
	addAPIRoute(permissionSubrouter, common.UpdatePath(), api.UpdatePermissionV1, "PUT")    // update
	addAPIRoute(permissionSubrouter, common.DeletePath(), api.DeletePermissionV1, "DELETE") // delete
	// Middleware
	var methodMiddleware http.HandlerFunc = MethodOverride
	router.PathPrefix(permissionsPath).Handler(negroni.New(
		// negroni.HandlerFunc(boot.Authorize),
		negroni.Wrap(methodMiddleware),
		negroni.Wrap(permissionRouter),
	))
}
