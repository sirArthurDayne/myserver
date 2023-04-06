package main

import (
    "fmt"
    "github.com/sirArthurDayne/webserver/myserver"
)

func main() {
	server := myserver.NewServer(":3000")

    fmt.Println("loading server...")
	server.Handle(myserver.GET, "/", myserver.HandlerRoot)
	server.Handle(myserver.GET, "/menu", myserver.HandlerMainMenu)
	server.Handle(myserver.POST, "/create", myserver.HandlerPostRequest)
	server.Handle(myserver.GET, "/api", server.AddMiddleware(myserver.HandlerApi, myserver.CheckAuth(), myserver.Login()))
	server.Listen()
}
