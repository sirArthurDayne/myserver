package myserver

import (
	"encoding/json"
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

func HandlerMethodNotAllowed(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Hola, la ruta es correcta pero la peticion no es compatible</h1>")
}
func HandlerPostRequest(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(writer, "[ERROR]: %v", err)
	}
	response, parseErr := user.ToJson()
	if parseErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}
