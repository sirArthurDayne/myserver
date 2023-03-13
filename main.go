package main

func main() {
	server := NewServer(":3000")

	server.Handle(GET, "/", HandlerRoot)
	server.Handle(GET, "/menu", HandlerMainMenu)
	server.Handle(GET, "/api", server.AddMiddleware(HandlerApi, CheckAuth(), Login()))
	server.Listen()
}
