#!/bin/bash

TARGET_OSES=(darwin dragonfly freebsd freebsd freebsd linux linux linux linux linux linux linux linux linux linux netbsd netbsd netbsd openbsd openbsd openbsd solaris windows windows)

ARCHES=(amd64 amd64 386 amd64 arm 386 amd64 arm arm64 ppc64 ppc64le mips mipsle mips64 mips64le 386 amd64 arm 386 amd64 arm amd64 386 amd64)

pids=( )

cleanup() {
  echo "aborting..."
  for pid in "${pids[@]}"; do
    kill -9 "$pid" && kill "$pid"
  done
  echo "finished."
}

trap cleanup TERM INT

for ((i = 0; i < ${#TARGET_OSES[@]}; ++i )); do
  export GOOS=${TARGET_OSES[i]};
  export GOARCH=${ARCHES[i]};
  NUM=$((i + 1));

  echo "[$NUM/${#TARGET_OSES[@]}] building $GOOS/$GOARCH";
  go build -o build/"$GOOS"-"$GOARCH"-tateru & pids+=("$!") &
done

wait