package main

import "t-bonatti/gopj/server"

func main() {
	server := server.New()
	server.Run(":8013")
}
