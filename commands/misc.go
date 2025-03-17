package commands

type CliCommand struct {
    Name string
    Description string
    Callback func() error
}

func GetCommandRegistry() map[string]CliCommand {
    cmds := map[string]CliCommand{
        "help": {
            Name: "help",
            Description: "Displays a help message",
            Callback: CommandHelp,
        },
        "exit": {
            Name: "exit",
            Description: "Exit the Pokedex",
            Callback: CommandExit,
        },
    }
    return cmds
}
