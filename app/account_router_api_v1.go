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
    AddaccountRouter(router) // <- Add this line
    return router
  }
*/

// AddAccountAPIV1Router - Initialize and append the account's API router.
func AddAccountAPIV1Router(parent *mux.Router) {
	// Paths
	accountsPath := fmt.Sprintf("/api/v1/%s", common.AccountRoot)

	// Routers
	accountRouter := NewRouter()
	accountSubrouter := accountRouter.PathPrefix(accountsPath).Subrouter()

	// Routes
	// addAPIRoute(accountSubrouter, path, controller.function, method")
	// path is a standard sub-URL generated in common package
	// and can be replace by a hardcoded string.
	addAPIRoute(accountSubrouter, common.IndexPath(), api.IndexAccountsV1, "GET")     // index
	addAPIRoute(accountSubrouter, common.ShowPath(), api.GetAccountV1, "GET")         // show
	addAPIRoute(accountSubrouter, common.CreatePath(), api.CreateAccountV1, "POST")   // create
	addAPIRoute(accountSubrouter, common.UpdatePath(), api.UpdateAccountV1, "PATCH")  // update
	addAPIRoute(accountSubrouter, common.UpdatePath(), api.UpdateAccountV1, "PUT")    // update
	addAPIRoute(accountSubrouter, common.DeletePath(), api.DeleteAccountV1, "DELETE") // delete
	// addAPIRoute(accountSubrouter, common.SignupPath(), api.SignUpAccountV1, "POST")   // create
	// addAPIRoute(accountSubrouter, common.LoginPath(), api.LogInAccountV1, "POST")     // create
	// Middleware
	var methodMiddleware http.HandlerFunc = MethodOverride
	router.PathPrefix(accountsPath).Handler(negroni.New(
		// negroni.HandlerFunc(boot.Authorize),
		negroni.Wrap(methodMiddleware),
		negroni.Wrap(accountRouter),
	))
}
