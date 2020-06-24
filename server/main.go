package main

import (
	"fmt"
	// "log"
	"net/http"

	"./servertest"
	"./rpc"
)

func main() {
	server := &servertest.Server{} // implements Haberdasher interface
	twirpHandler := rpc.NewTestAPIServer(server, nil)
	fmt.Println("server started")
	http.ListenAndServe(":8080", twirpHandler)
  }
