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
    AddRolePermissionAPIV1Router(router) // <- Add this line
    return router
  }
*/

// AddRolePermissionAPIV1Router - Initialize and append the rolePermission's API router.
func AddRolePermissionAPIV1Router(parent *mux.Router) {
	// Paths
	rolePermissionsPath := fmt.Sprintf("/api/v1/%s", common.RolePermissionRoot)

	// Routers
	rolePermissionRouter := NewRouter()
	rolePermissionSubrouter := rolePermissionRouter.PathPrefix(rolePermissionsPath).Subrouter()

	// Routes
	// addAPIRoute(rolePermissionSubrouter, path, controller.function, method")
	// path is a standard sub-URL generated in common package
	// and can be replace by a hardcoded string.
	addAPIRoute(rolePermissionSubrouter, common.IndexPath(), api.IndexRolePermissionsV1, "GET")     // index
	addAPIRoute(rolePermissionSubrouter, common.ShowPath(), api.GetRolePermissionV1, "GET")         // show
	addAPIRoute(rolePermissionSubrouter, common.CreatePath(), api.CreateRolePermissionV1, "POST")   // create
	addAPIRoute(rolePermissionSubrouter, common.UpdatePath(), api.UpdateRolePermissionV1, "PATCH")  // update
	addAPIRoute(rolePermissionSubrouter, common.UpdatePath(), api.UpdateRolePermissionV1, "PUT")    // update
	addAPIRoute(rolePermissionSubrouter, common.DeletePath(), api.DeleteRolePermissionV1, "DELETE") // delete
	// Middleware
	var methodMiddleware http.HandlerFunc = MethodOverride
	router.PathPrefix(rolePermissionsPath).Handler(negroni.New(
		// negroni.HandlerFunc(boot.Authorize),
		negroni.Wrap(methodMiddleware),
		negroni.Wrap(rolePermissionRouter),
	))
}
