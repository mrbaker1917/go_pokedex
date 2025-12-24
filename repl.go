package main

import (
	"strings"
)

func cleanInput(text string) []string {
	split1 := strings.Fields(strings.ToLower(text))
	return split1
}
