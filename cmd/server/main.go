package main

import "github.com/ikbarfp/go-clean-arch/app/server"

func main() {
	srv := server.New()
	srv.Run()
}
