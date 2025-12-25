package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	config := NewConfig()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config)
			if err == nil {
				config.TrackCommand(commandName)
				continue
			}
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type Config struct {
	CommandCalls map[string]int
}

func NewConfig() *Config {
	return &Config{
		CommandCalls: make(map[string]int),
	}
}

func (c *Config) TrackCommand(commandName string) {
	c.CommandCalls[commandName]++
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "returns 20 locations",
			callback:    commandMap,
		},
	}
}
