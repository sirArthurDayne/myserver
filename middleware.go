package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			isLogin := true
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

func Login() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(request.URL.Path, time.Since(start))
			}()
			f(writer, request)
		}
	}
}
