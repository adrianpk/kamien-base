#!/bin/sh
go_test () {
  killall -9 test.test 2>>/dev/null
  killall -9 runner-build 2>>/dev/null
  printf "\n%s\n**************************************************\n" $1
  go test {{.PackageName}}/test --run $1 #--v
  printf "\n"
}

go_test TestIndexAccount
go_test TestEditAccount
go_test TestNewAccount
go_test TestShowAccount
go_test TestCreateAccount
go_test TestUpdateAccount
go_test TestInitDeleteAccount
go_test TestDeleteAccount
