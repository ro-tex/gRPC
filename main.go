package main

import (
	"gitlab.com/ro-tex/grpc/client"
	"gitlab.com/ro-tex/grpc/server"
)

func main() {
	go server.Run()
	client.Run()
}
