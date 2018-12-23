package common

import (
	cm "github.com/adrianpk/kamien/common"
)

// UserRoleRoot - UserRole resource root path.
var UserRoleRoot = "user-roles"

/* Add routes to '{{.Package}}/common/routes.go'
	 Cut & paste the following into
	 Routes and RoutesExt definitions
  // UserRole
	"userRolePath":           UserRolePath,
	"userRolePathEdit":       UserRolePathEdit,
	"userRolePathNew":        UserRolePathNew,
	"userRolePathInitDelete": UserRolePathInitDelete,
	"userRolePathID":         UserRolePathID,
*/

// UserRolePath - TODO: Complete comment.
func UserRolePath() string {
	return ResPath(UserRoleRoot)
}

// UserRolePathEdit - TODO: Complete comment.
func UserRolePathEdit(model cm.Identifiable) string {
	return ResPathEdit(UserRoleRoot, model)
}

// UserRolePathNew - TODO: Complete comment.
func UserRolePathNew() string {
	return ResPathNew(UserRoleRoot)
}

// UserRolePathInitDelete - TODO: Complete comment.
func UserRolePathInitDelete(model cm.Identifiable) string {
	return ResPathInitDelete(UserRoleRoot, model)
}

// UserRolePathID - TODO: Complete comment.
func UserRolePathID(model cm.Identifiable) string {
	return ResPathID(UserRoleRoot, model)
}
