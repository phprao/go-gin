package main

import (
	"os"

	"votePlatfom/app/commands"
	"votePlatfom/server"
)

func main() {
	var cmd string
	if len(os.Args) > 2 {
		cmd = os.Args[1]
	}

	if cmd == "" || cmd == "serve" {
		server.RunWithGraceful()
	} else {
		commands.Handle(cmd)
	}
}
