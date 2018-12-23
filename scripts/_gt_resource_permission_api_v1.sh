#!/bin/sh
go_test () {
  killall -9 test.test 2>>/dev/null
  killall -9 runner-build 2>>/dev/null
  printf "\n%s\n**************************************************\n" $1
  go test {{.Package}}/test --run $1 #--v
  printf "\n"
}

# killall -9 runner-build; go test {{.Package}}/test --run TestIndexResourcePermission

go_test TestIndexResourcePermissionAPIV1
go_test TestGetResourcePermissionAPIV1
go_test TestCreateResourcePermissionAPIV1
go_test TestUpdateResourcePermissionAPIV1
go_test TestDeleteResourcePermissionAPIV1
