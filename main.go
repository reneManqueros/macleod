package main

import (
	. "macleod/models"
)

func init() {
	Config.Load()
}

func main() {
	server := Server{}
	server.Serve()
}
