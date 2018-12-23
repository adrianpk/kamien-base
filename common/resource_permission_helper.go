package common

import (
	cm "github.com/adrianpk/kamien/common"
)

// ResourcePermissionRoot - ResourcePermission resource root path.
var ResourcePermissionRoot = "resource-permissions"

/* Add routes to '{{.Package}}/common/routes.go'
	 Cut & paste the following into
	 Routes and RoutesExt definitions
  // ResourcePermission
	"resourcePermissionPath":           ResourcePermissionPath,
	"resourcePermissionPathEdit":       ResourcePermissionPathEdit,
	"resourcePermissionPathNew":        ResourcePermissionPathNew,
	"resourcePermissionPathInitDelete": ResourcePermissionPathInitDelete,
	"resourcePermissionPathID":         ResourcePermissionPathID,
*/

// ResourcePermissionPath - TODO: Complete comment.
func ResourcePermissionPath() string {
	return ResPath(ResourcePermissionRoot)
}

// ResourcePermissionPathEdit - TODO: Complete comment.
func ResourcePermissionPathEdit(model cm.Identifiable) string {
	return ResPathEdit(ResourcePermissionRoot, model)
}

// ResourcePermissionPathNew - TODO: Complete comment.
func ResourcePermissionPathNew() string {
	return ResPathNew(ResourcePermissionRoot)
}

// ResourcePermissionPathInitDelete - TODO: Complete comment.
func ResourcePermissionPathInitDelete(model cm.Identifiable) string {
	return ResPathInitDelete(ResourcePermissionRoot, model)
}

// ResourcePermissionPathID - TODO: Complete comment.
func ResourcePermissionPathID(model cm.Identifiable) string {
	return ResPathID(ResourcePermissionRoot, model)
}
