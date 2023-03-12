package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			isLogin := false
			fmt.Println("Checking Authentication...")
			if !isLogin {
				fmt.Println("Authentication failed!")
				HandlerNotLogin(writer, request)
				return
			}
			f(writer, request)
		}
	}
}
