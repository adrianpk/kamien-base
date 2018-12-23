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
    AdduserRouter(router) // <- Add this line
    return router
  }
*/

// AddUserAPIV1Router - Initialize and append the user's API router.
func AddUserAPIV1Router(parent *mux.Router) {
	// Paths
	usersPath := fmt.Sprintf("/api/v1/%s", common.UserRoot)

	// Routers
	userRouter := NewRouter()
	userSubrouter := userRouter.PathPrefix(usersPath).Subrouter()

	// Routes
	// addAPIRoute(userSubrouter, path, controller.function, method")
	// path is a standard sub-URL generated in common package
	// and can be replace by a hardcoded string.
	addAPIRoute(userSubrouter, common.IndexPath(), api.IndexUsersV1, "GET")     // index
	addAPIRoute(userSubrouter, common.ShowPath(), api.GetUserV1, "GET")         // show
	addAPIRoute(userSubrouter, common.CreatePath(), api.CreateUserV1, "POST")   // create
	addAPIRoute(userSubrouter, common.UpdatePath(), api.UpdateUserV1, "PATCH")  // update
	addAPIRoute(userSubrouter, common.UpdatePath(), api.UpdateUserV1, "PUT")    // update
	addAPIRoute(userSubrouter, common.DeletePath(), api.DeleteUserV1, "DELETE") // delete
	// addAPIRoute(userSubrouter, common.SignupPath(), api.SignUpUserV1, "POST")   // create
	// addAPIRoute(userSubrouter, common.LoginPath(), api.LogInUserV1, "POST")     // create
	// Middleware
	var methodMiddleware http.HandlerFunc = MethodOverride
	router.PathPrefix(usersPath).Handler(negroni.New(
		// negroni.HandlerFunc(boot.Authorize),
		negroni.Wrap(methodMiddleware),
		negroni.Wrap(userRouter),
	))
}
