package common

import (
	cm "github.com/adrianpk/kamien/common"
)

// AccountRoot - Account resource root path.
var AccountRoot = "accounts"

// AccountPath - TODO: Complete comment.
func AccountPath() string {
	return ResPath(AccountRoot)
}

// AccountPathEdit - TODO: Complete comment.
func AccountPathEdit(model cm.Identifiable) string {
	return ResPathEdit(AccountRoot, model)
}

// AccountPathNew - TODO: Complete comment.
func AccountPathNew() string {
	return ResPathNew(AccountRoot)
}

// AccountPathInitDelete - TODO: Complete comment.
func AccountPathInitDelete(model cm.Identifiable) string {
	return ResPathInitDelete(AccountRoot, model)
}

// AccountPathID - TODO: Complete comment.
func AccountPathID(model cm.Identifiable) string {
	return ResPathID(AccountRoot, model)
}
