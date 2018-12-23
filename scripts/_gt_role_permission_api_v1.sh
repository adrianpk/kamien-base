#!/bin/sh
go_test () {
  killall -9 test.test 2>>/dev/null
  killall -9 runner-build 2>>/dev/null
  printf "\n%s\n**************************************************\n" $1
  go test {{.Package}}/test --run $1 #--v
  printf "\n"
}

# killall -9 runner-build; go test {{.Package}}/test --run TestIndexRolePermission

go_test TestIndexRolePermissionAPIV1
go_test TestGetRolePermissionAPIV1
go_test TestCreateRolePermissionAPIV1
go_test TestUpdateRolePermissionAPIV1
go_test TestDeleteRolePermissionAPIV1
