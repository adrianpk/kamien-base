#!/bin/sh
go_test () {
  killall -9 test.test 2>>/dev/null
  killall -9 runner-build 2>>/dev/null
  printf "\n%s\n**************************************************\n" $1
  go test {{.Package}}/test --run $1 #--v
  printf "\n"
}

# killall -9 runner-build; go test {{.Package}}/test --run TestIndexAccount

go_test TestIndexAccountAPIV1
go_test TestGetAccountAPIV1
go_test TestCreateAccountAPIV1
go_test TestUpdateAccountAPIV1
go_test TestDeleteAccountAPIV1

