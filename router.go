package main

import (
	"net/http"
)

//ROUTER: RUTA + TIPO DE PETICION  --> Handler
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string, method string) (bool, bool, http.HandlerFunc) {
	_, pathExist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return pathExist, methodExist, handler
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	pathExist, methodExist, handler := r.FindHandler(request.URL.Path, request.Method)
	if !pathExist {
		writer.WriteHeader(http.StatusNotFound)
		HandlerNotFound(writer, request)
		return
	}
	if !methodExist {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		HandlerMethodNotAllowed(writer, request)
		return
	}
	writer.WriteHeader(http.StatusOK)
	handler(writer, request)
}
