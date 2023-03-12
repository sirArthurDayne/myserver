package main

func main() {
    server := NewServer(":3000")
    server.Handle("/", HandleRoot)
    server.Handle("/menu", HandleMainMenu)
    server.Handle("/api", HandleApi)
    server.Listen()
}
