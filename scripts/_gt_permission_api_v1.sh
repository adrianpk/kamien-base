#!/bin/sh
go_test () {
  killall -9 test.test 2>>/dev/null
  killall -9 runner-build 2>>/dev/null
  printf "\n%s\n**************************************************\n" $1
  go test {{.Package}}/test --run $1 #--v
  printf "\n"
}

# killall -9 runner-build; go test {{.Package}}/test --run TestIndexPermission

go_test TestIndexPermissionAPIV1
go_test TestGetPermissionAPIV1
go_test TestCreatePermissionAPIV1
go_test TestUpdatePermissionAPIV1
go_test TestDeletePermissionAPIV1
