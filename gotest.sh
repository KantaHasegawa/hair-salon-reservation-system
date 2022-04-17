#!/bin/bash

go run test/up/up.go
go run test/seed/main.go
go test ./... -count=1
go run test/down/down.go
