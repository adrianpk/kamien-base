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
    AddRoleAPIV1Router(router) // <- Add this line
    return router
  }
*/

// AddRoleAPIV1Router - Initialize and append the role's API router.
func AddRoleAPIV1Router(parent *mux.Router) {
	// Paths
	rolesPath := fmt.Sprintf("/api/v1/%s", common.RoleRoot)

	// Routers
	roleRouter := NewRouter()
	roleSubrouter := roleRouter.PathPrefix(rolesPath).Subrouter()

	// Routes
	// addAPIRoute(roleSubrouter, path, controller.function, method")
	// path is a standard sub-URL generated in common package
	// and can be replace by a hardcoded string.
	addAPIRoute(roleSubrouter, common.IndexPath(), api.IndexRolesV1, "GET")     // index
	addAPIRoute(roleSubrouter, common.ShowPath(), api.GetRoleV1, "GET")         // show
	addAPIRoute(roleSubrouter, common.CreatePath(), api.CreateRoleV1, "POST")   // create
	addAPIRoute(roleSubrouter, common.UpdatePath(), api.UpdateRoleV1, "PATCH")  // update
	addAPIRoute(roleSubrouter, common.UpdatePath(), api.UpdateRoleV1, "PUT")    // update
	addAPIRoute(roleSubrouter, common.DeletePath(), api.DeleteRoleV1, "DELETE") // delete
	// Middleware
	var methodMiddleware http.HandlerFunc = MethodOverride
	router.PathPrefix(rolesPath).Handler(negroni.New(
		// negroni.HandlerFunc(boot.Authorize),
		negroni.Wrap(methodMiddleware),
		negroni.Wrap(roleRouter),
	))
}
