package common

import (
	"github.com/alecthomas/template"
	htmlTemplate "github.com/arschles/go-bindata-html-template"
)

// Routes - Map of all resources functions used for path generations.
var Routes = htmlTemplate.FuncMap{
	// User
	"userPath":           UserPath,
	"userPathEdit":       UserPathEdit,
	"userPathNew":        UserPathNew,
	"userPathInitDelete": UserPathInitDelete,
	"userPathID":         UserPathID,
	// Account
	"accountPath":           AccountPath,
	"accountPathEdit":       AccountPathEdit,
	"accountPathNew":        AccountPathNew,
	"accountPathInitDelete": AccountPathInitDelete,
	"accountPathID":         AccountPathID,
	// Profile
	"profilePath":           ProfilePath,
	"profilePathEdit":       ProfilePathEdit,
	"profilePathNew":        ProfilePathNew,
	"profilePathInitDelete": ProfilePathInitDelete,
	"profilePathID":         ProfilePathID,
	// Resource
	"resourcePath":           ResourcePath,
	"resourcePathEdit":       ResourcePathEdit,
	"resourcePathNew":        ResourcePathNew,
	"resourcePathInitDelete": ResourcePathInitDelete,
	"resourcePathID":         ResourcePathID,
	// Permission
	"permissionPath":           PermissionPath,
	"permissionPathEdit":       PermissionPathEdit,
	"permissionPathNew":        PermissionPathNew,
	"permissionPathInitDelete": PermissionPathInitDelete,
	"permissionPathID":         PermissionPathID,
	// Role
	"rolePath":           RolePath,
	"rolePathEdit":       RolePathEdit,
	"rolePathNew":        RolePathNew,
	"rolePathInitDelete": RolePathInitDelete,
	"rolePathID":         RolePathID,
	// RolePermission
	"rolePermissionPath":           RolePermissionPath,
	"rolePermissionPathEdit":       RolePermissionPathEdit,
	"rolePermissionPathNew":        RolePermissionPathNew,
	"rolePermissionPathInitDelete": RolePermissionPathInitDelete,
	"rolePermissionPathID":         RolePermissionPathID,
}

// RoutesExt - Map of all resources functions used for path generations.
var RoutesExt = template.FuncMap{
	// User
	"userPath":           UserPath,
	"userPathEdit":       UserPathEdit,
	"userPathNew":        UserPathNew,
	"userPathInitDelete": UserPathInitDelete,
	"userPathID":         UserPathID,
	// Account
	"accountPath":           AccountPath,
	"accountPathEdit":       AccountPathEdit,
	"accountPathNew":        AccountPathNew,
	"accountPathInitDelete": AccountPathInitDelete,
	"accountPathID":         AccountPathID,
	// Profile
	"profilePath":           ProfilePath,
	"profilePathEdit":       ProfilePathEdit,
	"profilePathNew":        ProfilePathNew,
	"profilePathInitDelete": ProfilePathInitDelete,
	"profilePathID":         ProfilePathID,
	// Resource
	"resourcePath":           ResourcePath,
	"resourcePathEdit":       ResourcePathEdit,
	"resourcePathNew":        ResourcePathNew,
	"resourcePathInitDelete": ResourcePathInitDelete,
	"resourcePathID":         ResourcePathID,
	// Permission
	"permissionPath":           PermissionPath,
	"permissionPathEdit":       PermissionPathEdit,
	"permissionPathNew":        PermissionPathNew,
	"permissionPathInitDelete": PermissionPathInitDelete,
	"permissionPathID":         PermissionPathID,
	// Role
	"rolePath":           RolePath,
	"rolePathEdit":       RolePathEdit,
	"rolePathNew":        RolePathNew,
	"rolePathInitDelete": RolePathInitDelete,
	"rolePathID":         RolePathID,
	// RolePermission
	"rolePermissionPath":           RolePermissionPath,
	"rolePermissionPathEdit":       RolePermissionPathEdit,
	"rolePermissionPathNew":        RolePermissionPathNew,
	"rolePermissionPathInitDelete": RolePermissionPathInitDelete,
	"rolePermissionPathID":         RolePermissionPathID,
}
