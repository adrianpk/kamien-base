#!/bin/sh
go_test () {
  killall -9 test.test 2>>/dev/null
  killall -9 runner-build 2>>/dev/null
  printf "\n%s\n**************************************************\n" $1
  go test {{.Package}}/test --run $1 #--v
  printf "\n"
}

# killall -9 runner-build; go test {{.Package}}/test --run TestIndexResource

go_test TestIndexResourceAPIV1
go_test TestGetResourceAPIV1
go_test TestCreateResourceAPIV1
go_test TestUpdateResourceAPIV1
go_test TestDeleteResourceAPIV1
