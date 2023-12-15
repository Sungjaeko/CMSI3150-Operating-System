#!/bin/sh

go env -w GO111MODULE=off
go install --gcflags="-N -l" github.com/Sungjaeko/hello
go install github.com/Sungjaeko/debugger


hello you found me