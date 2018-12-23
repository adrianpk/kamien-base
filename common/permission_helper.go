package common

import (
	cm "github.com/adrianpk/kamien/common"
)

// PermissionRoot - Permission resource root path.
var PermissionRoot = "permissions"

/* Add routes to '{{.Package}}/common/routes.go'
	 Cut & paste the following into
	 Routes and RoutesExt definitions
  // Permission
	"permissionPath":           PermissionPath,
	"permissionPathEdit":       PermissionPathEdit,
	"permissionPathNew":        PermissionPathNew,
	"permissionPathInitDelete": PermissionPathInitDelete,
	"permissionPathID":         PermissionPathID,
*/

// PermissionPath - TODO: Complete comment.
func PermissionPath() string {
	return ResPath(PermissionRoot)
}

// PermissionPathEdit - TODO: Complete comment.
func PermissionPathEdit(model cm.Identifiable) string {
	return ResPathEdit(PermissionRoot, model)
}

// PermissionPathNew - TODO: Complete comment.
func PermissionPathNew() string {
	return ResPathNew(PermissionRoot)
}

// PermissionPathInitDelete - TODO: Complete comment.
func PermissionPathInitDelete(model cm.Identifiable) string {
	return ResPathInitDelete(PermissionRoot, model)
}

// PermissionPathID - TODO: Complete comment.
func PermissionPathID(model cm.Identifiable) string {
	return ResPathID(PermissionRoot, model)
}
