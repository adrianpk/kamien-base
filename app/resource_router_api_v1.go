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
    AddResourceAPIV1Router(router) // <- Add this line
    return router
  }
*/

// AddResourceAPIV1Router - Initialize and append the resource's API router.
func AddResourceAPIV1Router(parent *mux.Router) {
	// Paths
	resourcesPath := fmt.Sprintf("/api/v1/%s", common.ResourceRoot)

	// Routers
	resourceRouter := NewRouter()
	resourceSubrouter := resourceRouter.PathPrefix(resourcesPath).Subrouter()

	// Routes
	// addAPIRoute(resourceSubrouter, path, controller.function, method")
	// path is a standard sub-URL generated in common package
	// and can be replace by a hardcoded string.
	addAPIRoute(resourceSubrouter, common.IndexPath(), api.IndexResourcesV1, "GET")     // index
	addAPIRoute(resourceSubrouter, common.ShowPath(), api.GetResourceV1, "GET")         // show
	addAPIRoute(resourceSubrouter, common.CreatePath(), api.CreateResourceV1, "POST")   // create
	addAPIRoute(resourceSubrouter, common.UpdatePath(), api.UpdateResourceV1, "PATCH")  // update
	addAPIRoute(resourceSubrouter, common.UpdatePath(), api.UpdateResourceV1, "PUT")    // update
	addAPIRoute(resourceSubrouter, common.DeletePath(), api.DeleteResourceV1, "DELETE") // delete
	// Middleware
	var methodMiddleware http.HandlerFunc = MethodOverride
	router.PathPrefix(resourcesPath).Handler(negroni.New(
		// negroni.HandlerFunc(boot.Authorize),
		negroni.Wrap(methodMiddleware),
		negroni.Wrap(resourceRouter),
	))
}
