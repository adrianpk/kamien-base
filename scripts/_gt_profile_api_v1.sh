#!/bin/sh
go_test () {
  killall -9 test.test 2>>/dev/null
  killall -9 runner-build 2>>/dev/null
  printf "\n%s\n**************************************************\n" $1
  go test {{.Package}}/test --run $1 #--v
  printf "\n"
}

# killall -9 runner-build; go test {{.Package}}/test --run TestIndexProfile

go_test TestIndexProfileAPIV1
go_test TestGetProfileAPIV1
go_test TestCreateProfileAPIV1
go_test TestUpdateProfileAPIV1
go_test TestDeleteProfileAPIV1
