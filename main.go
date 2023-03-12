package main

func main() {
    server := NewServer(":3000")

    server.Handle("/", HandlerRoot)
    server.Handle("/menu", HandlerMainMenu)
    server.Handle("/api", HandlerApi)
    server.Listen()
}
