package main

import (
	"fmt"
	"net/http"
)

func HandlerRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Hola Mundo desde Root</h1>")
}

func HandlerMainMenu(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Hola Mundo desde Menu Principal</h1>")
}

func HandlerApi(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Hola Mundo desde API Endpoint</h1>")
}
func HandlerNotFound(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Hola, recurso No encontrado</h1>")
}
func HandlerNotLogin(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Hola, Para acceder, primero, debes autenticarte y recargar la pagina</h1>")
}
