#!/bin/sh
go_test () {
  killall -9 test.test 2>>/dev/null
  killall -9 runner-build 2>>/dev/null
  printf "\n%s\n**************************************************\n" $1
  go test {{.Package}}/test --run $1 #--v
  printf "\n"
}

# killall -9 runner-build; go test {{.Package}}/test --run TestIndexUserRole

go_test TestIndexUserRoleAPIV1
go_test TestGetUserRoleAPIV1
go_test TestCreateUserRoleAPIV1
go_test TestUpdateUserRoleAPIV1
go_test TestDeleteUserRoleAPIV1
