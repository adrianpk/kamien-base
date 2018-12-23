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
    AddResourcePermissionAPIV1Router(router) // <- Add this line
    return router
  }
*/

// AddResourcePermissionAPIV1Router - Initialize and append the resourcePermission's API router.
func AddResourcePermissionAPIV1Router(parent *mux.Router) {
	// Paths
	resourcePermissionsPath := fmt.Sprintf("/api/v1/%s", common.ResourcePermissionRoot)

	// Routers
	resourcePermissionRouter := NewRouter()
	resourcePermissionSubrouter := resourcePermissionRouter.PathPrefix(resourcePermissionsPath).Subrouter()

	// Routes
	// addAPIRoute(resourcePermissionSubrouter, path, controller.function, method")
	// path is a standard sub-URL generated in common package
	// and can be replace by a hardcoded string.
	addAPIRoute(resourcePermissionSubrouter, common.IndexPath(), api.IndexResourcePermissionsV1, "GET")     // index
	addAPIRoute(resourcePermissionSubrouter, common.ShowPath(), api.GetResourcePermissionV1, "GET")         // show
	addAPIRoute(resourcePermissionSubrouter, common.CreatePath(), api.CreateResourcePermissionV1, "POST")   // create
	addAPIRoute(resourcePermissionSubrouter, common.UpdatePath(), api.UpdateResourcePermissionV1, "PATCH")  // update
	addAPIRoute(resourcePermissionSubrouter, common.UpdatePath(), api.UpdateResourcePermissionV1, "PUT")    // update
	addAPIRoute(resourcePermissionSubrouter, common.DeletePath(), api.DeleteResourcePermissionV1, "DELETE") // delete
	// Middleware
	var methodMiddleware http.HandlerFunc = MethodOverride
	router.PathPrefix(resourcePermissionsPath).Handler(negroni.New(
		// negroni.HandlerFunc(boot.Authorize),
		negroni.Wrap(methodMiddleware),
		negroni.Wrap(resourcePermissionRouter),
	))
}
