package common

import (
	cm "github.com/adrianpk/kamien/common"
)

// ResourceRoot - Resource resource root path.
var ResourceRoot = "resources"

/* Add routes to '{{.Package}}'/common/routes.go'
	 Cut & paste the following into
	 Routes and RoutesExt definitions
  // Resource
	"resourcePath":           ResourcePath,
	"resourcePathEdit":       ResourcePathEdit,
	"resourcePathNew":        ResourcePathNew,
	"resourcePathInitDelete": ResourcePathInitDelete,
	"resourcePathID":         ResourcePathID,
*/

// ResourcePath - TODO: Complete comment.
func ResourcePath() string {
	return ResPath(ResourceRoot)
}

// ResourcePathEdit - TODO: Complete comment.
func ResourcePathEdit(model cm.Identifiable) string {
	return ResPathEdit(ResourceRoot, model)
}

// ResourcePathNew - TODO: Complete comment.
func ResourcePathNew() string {
	return ResPathNew(ResourceRoot)
}

// ResourcePathInitDelete - TODO: Complete comment.
func ResourcePathInitDelete(model cm.Identifiable) string {
	return ResPathInitDelete(ResourceRoot, model)
}

// ResourcePathID - TODO: Complete comment.
func ResourcePathID(model cm.Identifiable) string {
	return ResPathID(ResourceRoot, model)
}
