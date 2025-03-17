package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func cleanInput(text string) []string {
    return strings.Fields(strings.ToLower(text))
}

func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex> ")
        scanner.Scan()
        input := scanner.Text()
        clean := cleanInput(input)
        if len(clean) == 0 {
            continue
        }
        commandName := clean[0]
        fmt.Printf("Your command was: %s\n", commandName)
    }
}
