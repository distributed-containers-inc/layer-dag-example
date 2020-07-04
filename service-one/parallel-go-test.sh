#!/usr/bin/env bash

if [ -z "$SPLIT" ]; then
  echo 'missing required environment variable: $SPLIT'
  exit 1
fi

if [ -z "$SPLIT_NUM" ]; then
  echo 'missing required environment variable: $SPLIT_NUM'
  exit 1
fi

tests=($(go test -list . | head -n -1 | sort))
numtests="${#tests[@]}"
chunksize="$(( numtests / $SPLIT_NUM ))"

start="$(( $SPLIT * $chunksize ))"
num="$chunksize"

if [ $SPLIT = $(( $SPLIT_NUM - 1)) ]; then
  num="$(( $numtests - $start ))"
fi

echo "There are numtests for $SPLIT_NUM splits, currently split $SPLIT"
echo "Running tests $(( $start + 1))-$(( $start + $num ))"

tests_running=( "${tests[@]:$start:$num}")

filter=$(IFS='|' ; echo "${tests_running[*]}")
echo "filter is: $filter"

go test -v -run "$filter"