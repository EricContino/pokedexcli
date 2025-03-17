package commands

import "fmt"

func CommandHelp() error {
    cmds := GetCommandRegistry()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:\n")
    for _, cmd := range cmds {
        fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
    }
    return nil
}
