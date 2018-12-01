#!/usr/bin/env bash

go test -v -coverprofile cover.out ./tests
go tool cover -html=cover.out -o cover.html