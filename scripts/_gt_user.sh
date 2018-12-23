#!/bin/sh
go_test () {
  killall -9 test.test 2>>/dev/null
  killall -9 runner-build 2>>/dev/null
  printf "\n%s\n**************************************************\n" $1
  go test {{.PackageName}}/test --run $1 #--v
  printf "\n"
}

go_test TestIndexUser
go_test TestEditUser
go_test TestNewUser
go_test TestShowUser
go_test TestCreateUser
go_test TestUpdateUser
go_test TestInitDeleteUser
go_test TestDeleteUser
