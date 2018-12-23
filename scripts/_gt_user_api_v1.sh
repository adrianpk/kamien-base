#!/bin/sh
go_test () {
  killall -9 test.test 2>>/dev/null
  killall -9 runner-build 2>>/dev/null
  printf "\n%s\n**************************************************\n" $1
  go test {{.Package}}/test --run $1 #--v
  printf "\n"
}

# killall -9 runner-build; go test {{.Package}}/test --run TestIndexUser

go_test TestIndexUserAPIV1
go_test TestGetUserAPIV1
go_test TestCreateUserAPIV1
go_test TestUpdateUserAPIV1
go_test TestDeleteUserAPIV1
go_test TestSignUpUserAPIV1
go_test TestLogInUserAPIV1