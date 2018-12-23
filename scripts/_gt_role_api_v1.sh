#!/bin/sh
go_test () {
  killall -9 test.test 2>>/dev/null
  killall -9 runner-build 2>>/dev/null
  printf "\n%s\n**************************************************\n" $1
  go test {{.Package}}/test --run $1 #--v
  printf "\n"
}

# killall -9 runner-build; go test {{.Package}}/test --run TestIndexRole

go_test TestIndexRoleAPIV1
go_test TestGetRoleAPIV1
go_test TestCreateRoleAPIV1
go_test TestUpdateRoleAPIV1
go_test TestDeleteRoleAPIV1
