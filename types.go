package main

import (
	"net/http"
)

const (
	GET  = http.MethodGet
	POST = http.MethodPost
)

type Middleware func(http.HandlerFunc) http.HandlerFunc
