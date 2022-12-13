package main

import (
	"backendserver"
)

func main() {
	b := backendserver.BackendRoute{}
	b.InitRoute()
	b.StartServer()
}
