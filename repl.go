package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	split1 := strings.Fields(strings.ToLower(text))
	return split1
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startREPL() {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Helps user navigate the Pokedex",
			callback:    help,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command := cleanInput(input)
		if _, ok := commands[command[0]]; !ok {
			fmt.Println("unknown command")
		}
		for k, v := range commands {
			if k == command[0] {
				v.callback()
			}
		}
		fmt.Print("Unknown command\n")

	}
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func help() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exits the Pokedex")
	return nil

}
