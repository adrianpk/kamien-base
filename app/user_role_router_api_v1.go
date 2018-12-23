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
    AddUserRoleAPIV1Router(router) // <- Add this line
    return router
  }
*/

// AddUserRoleAPIV1Router - Initialize and append the userRole's API router.
func AddUserRoleAPIV1Router(parent *mux.Router) {
	// Paths
	userRolesPath := fmt.Sprintf("/api/v1/%s", common.UserRoleRoot)

	// Routers
	userRoleRouter := NewRouter()
	userRoleSubrouter := userRoleRouter.PathPrefix(userRolesPath).Subrouter()

	// Routes
	// addAPIRoute(userRoleSubrouter, path, controller.function, method")
	// path is a standard sub-URL generated in common package
	// and can be replace by a hardcoded string.
	addAPIRoute(userRoleSubrouter, common.IndexPath(), api.IndexUserRolesV1, "GET")     // index
	addAPIRoute(userRoleSubrouter, common.ShowPath(), api.GetUserRoleV1, "GET")         // show
	addAPIRoute(userRoleSubrouter, common.CreatePath(), api.CreateUserRoleV1, "POST")   // create
	addAPIRoute(userRoleSubrouter, common.UpdatePath(), api.UpdateUserRoleV1, "PATCH")  // update
	addAPIRoute(userRoleSubrouter, common.UpdatePath(), api.UpdateUserRoleV1, "PUT")    // update
	addAPIRoute(userRoleSubrouter, common.DeletePath(), api.DeleteUserRoleV1, "DELETE") // delete
	// Middleware
	var methodMiddleware http.HandlerFunc = MethodOverride
	router.PathPrefix(userRolesPath).Handler(negroni.New(
		// negroni.HandlerFunc(boot.Authorize),
		negroni.Wrap(methodMiddleware),
		negroni.Wrap(userRoleRouter),
	))
}
