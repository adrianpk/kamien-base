#!/bin/sh
run_test () {
  printf "\n%s\n**************************************************\n" $1
  $1
  printf "\n"
}

# run_test ./scripts/_gt_user_api_v1.sh
