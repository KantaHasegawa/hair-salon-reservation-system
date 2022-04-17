#!/bin/bash

go run test/up/up.go
go run test/seed/main.go
go test ./...
go run test/down/down.go
