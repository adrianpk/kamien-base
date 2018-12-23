package common

import (
	cm "github.com/adrianpk/kamien/common"
)

// RolePermissionRoot - RolePermission resource root path.
var RolePermissionRoot = "role-permissions"

/* Add routes to '{{.Package}}/common/routes.go'
	 Cut & paste the following into
	 Routes and RoutesExt definitions
  // RolePermission
	"rolePermissionPath":           RolePermissionPath,
	"rolePermissionPathEdit":       RolePermissionPathEdit,
	"rolePermissionPathNew":        RolePermissionPathNew,
	"rolePermissionPathInitDelete": RolePermissionPathInitDelete,
	"rolePermissionPathID":         RolePermissionPathID,
*/

// RolePermissionPath - TODO: Complete comment.
func RolePermissionPath() string {
	return ResPath(RolePermissionRoot)
}

// RolePermissionPathEdit - TODO: Complete comment.
func RolePermissionPathEdit(model cm.Identifiable) string {
	return ResPathEdit(RolePermissionRoot, model)
}

// RolePermissionPathNew - TODO: Complete comment.
func RolePermissionPathNew() string {
	return ResPathNew(RolePermissionRoot)
}

// RolePermissionPathInitDelete - TODO: Complete comment.
func RolePermissionPathInitDelete(model cm.Identifiable) string {
	return ResPathInitDelete(RolePermissionRoot, model)
}

// RolePermissionPathID - TODO: Complete comment.
func RolePermissionPathID(model cm.Identifiable) string {
	return ResPathID(RolePermissionRoot, model)
}
