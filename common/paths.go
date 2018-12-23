package common

import (
	"fmt"

	cm "github.com/adrianpk/kamien/common"
)

// IndexPath - Index path under resource root path.
func IndexPath() string {
	return ""
}

// EditPath - Edit path under resource root path.
func EditPath() string {
	return "/{id}/edit"
}

// NewPath - New path under resource root path.
func NewPath() string {
	return "/new"
}

// ShowPath - Show path under resource root path.
func ShowPath() string {
	return "/{id}"
}

// CreatePath - Create path under resource root path.
func CreatePath() string {
	return ""
}

// UpdatePath - Update path under resource root path.
func UpdatePath() string {
	return "/{id}"
}

// InitDeletePath - Init delete path under resource root path.
func InitDeletePath() string {
	return "/{id}/init-delete"
}

// DeletePath - Delete path under resource root path.
func DeletePath() string {
	return "/{id}"
}

// SignupPath - Signup path.
func SignupPath() string {
	return "/signup"
}

// LoginPath - Login path.
func LoginPath() string {
	return "/login"
}

// ---

// ResPath - TODO: Complete comment.
func ResPath(rootPath string) string {
	return "/" + rootPath + IndexPath()
}

// ResPathEdit - TODO: Complete comment.
func ResPathEdit(rootPath string, model cm.Identifiable) string {
	return fmt.Sprintf("/%s/%s/edit", rootPath, model.GetID())
}

// ResPathNew - TODO: Complete comment.
func ResPathNew(rootPath string) string {
	return fmt.Sprintf("/%s/new", rootPath)
}

// ResPathInitDelete - TODO: Complete comment.
func ResPathInitDelete(rootPath string, model cm.Identifiable) string {
	return fmt.Sprintf("/%s/%s/init-delete", rootPath, model.GetID())
}

// ResPathID - TODO: Complete comment.
func ResPathID(rootPath string, model cm.Identifiable) string {
	return fmt.Sprintf("/%s/%s", rootPath, model.GetID())
}
