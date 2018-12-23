package common

import (
	cm "github.com/adrianpk/kamien/common"
)

// RoleRoot - Role resource root path.
var RoleRoot = "roles"

/* Add routes to '{{.Package}}/common/routes.go'
	 Cut & paste the following into
	 Routes and RoutesExt definitions
  // Role
	"rolePath":           RolePath,
	"rolePathEdit":       RolePathEdit,
	"rolePathNew":        RolePathNew,
	"rolePathInitDelete": RolePathInitDelete,
	"rolePathID":         RolePathID,
*/

// RolePath - TODO: Complete comment.
func RolePath() string {
	return ResPath(RoleRoot)
}

// RolePathEdit - TODO: Complete comment.
func RolePathEdit(model cm.Identifiable) string {
	return ResPathEdit(RoleRoot, model)
}

// RolePathNew - TODO: Complete comment.
func RolePathNew() string {
	return ResPathNew(RoleRoot)
}

// RolePathInitDelete - TODO: Complete comment.
func RolePathInitDelete(model cm.Identifiable) string {
	return ResPathInitDelete(RoleRoot, model)
}

// RolePathID - TODO: Complete comment.
func RolePathID(model cm.Identifiable) string {
	return ResPathID(RoleRoot, model)
}
