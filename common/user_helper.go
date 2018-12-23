package common

import (
	cm "github.com/adrianpk/kamien/common"
)

// UserRoot - User resource root path.
var UserRoot = "users"

// UserPath - TODO: Complete comment.
func UserPath() string {
	return ResPath(UserRoot)
}

// UserPathEdit - TODO: Complete comment.
func UserPathEdit(model cm.Identifiable) string {
	return ResPathEdit(UserRoot, model)
}

// UserPathNew - TODO: Complete comment.
func UserPathNew() string {
	return ResPathNew(UserRoot)
}

// UserPathInitDelete - TODO: Complete comment.
func UserPathInitDelete(model cm.Identifiable) string {
	return ResPathInitDelete(UserRoot, model)
}

// UserPathID - TODO: Complete comment.
func UserPathID(model cm.Identifiable) string {
	return ResPathID(UserRoot, model)
}
