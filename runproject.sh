#!/bin/sh

go fmt myserver/*.go
go fmt main/*.go
go run main/main.go
