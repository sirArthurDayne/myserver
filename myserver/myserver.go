package myserver

import (
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

//generar un servidor mediante puerto y router
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

//combinar ruta con Handler
func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, pathExist := s.router.rules[path]
	if !pathExist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

//Agregar middleware para pasarselos al handler correspondiente
func (s *Server) AddMiddleware(f http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	for _, m := range middleware {
		f = m(f)
	}
	return f
}

//escuchar las peticiones
func (s *Server) Listen() error {
	http.Handle("/", s.router)
	error := http.ListenAndServe(s.port, nil)
	if error != nil {
		return error
	}
	return nil
}
