package common

import (
	cm "github.com/adrianpk/kamien/common"
)

// ProfileRoot - Profile resource root path.
var ProfileRoot = "profiles"

/* Add routes to '{{.Package}}'/common/routes.go'
	 Cut & paste the following into
	 Routes and RoutesExt definitions
  // Profile
	"profilePath":           ProfilePath,
	"profilePathEdit":       ProfilePathEdit,
	"profilePathNew":        ProfilePathNew,
	"profilePathInitDelete": ProfilePathInitDelete,
	"profilePathID":         ProfilePathID,
*/

// ProfilePath - TODO: Complete comment.
func ProfilePath() string {
	return ResPath(ProfileRoot)
}

// ProfilePathEdit - TODO: Complete comment.
func ProfilePathEdit(model cm.Identifiable) string {
	return ResPathEdit(ProfileRoot, model)
}

// ProfilePathNew - TODO: Complete comment.
func ProfilePathNew() string {
	return ResPathNew(ProfileRoot)
}

// ProfilePathInitDelete - TODO: Complete comment.
func ProfilePathInitDelete(model cm.Identifiable) string {
	return ResPathInitDelete(ProfileRoot, model)
}

// ProfilePathID - TODO: Complete comment.
func ProfilePathID(model cm.Identifiable) string {
	return ResPathID(ProfileRoot, model)
}
