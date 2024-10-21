package main

import (
	"fmt"
	"os"
	"strings"

	rcon "github.com/gtaylor/factorio-rcon"
)

func main() {
	r, err := rcon.Dial("192.168.0.25:27015")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	rconPassword := os.Args[1]
	commands := os.Args[2:]

	err = r.Authenticate(rconPassword)
	if err != nil {
		panic(err)
	}

	commandsStr := ""
	for _, command := range commands {
		commandsStr += command + " "
	}
	commandsStr = strings.TrimSpace(commandsStr)

	response, err := r.Execute(commandsStr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %+v\n", response)
}
