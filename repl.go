package main

import (
    "bufio"
    "github.com/EricContino/pokedexcli/commands"
    "fmt"
    "os"
    "strings"
)

func cleanInput(text string) []string {
    return strings.Fields(strings.ToLower(text))
}

func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)
    cmds := commands.GetCommandRegistry()
    config := commands.Config{
        Next: commands.PokeApi_LocationAreas_BaseURL,
        Previous: commands.PokeApi_LocationAreas_BaseURL,
    }
    for {
        fmt.Print("Pokedex> ")
        scanner.Scan()
        input := scanner.Text()
        clean := cleanInput(input)
        if len(clean) == 0 {
            continue
        }
        commandName := clean[0]
        if cmd, exists := cmds[commandName]; exists {
            cmd.Callback(&config)
        } else {
            fmt.Println("Unknown command")
        }
    }
}
