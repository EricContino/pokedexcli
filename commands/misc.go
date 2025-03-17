package commands

const PokeApi_LocationAreas_BaseURL string = "https://pokeapi.co/api/v2/location-area/"

type CliCommand struct {
    Name string
    Description string
    Callback func(config *Config) error
}

type Config struct {
    Next string
    Previous string
}

type MapResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
        "map": {
            Name: "map",
            Description: "Displays next 20 locations",
            Callback: CommandMap,
        },
        "mapb": {
            Name: "map",
            Description: "Displays previous 20 locations",
            Callback: CommandMapb,
        },
    }
    return cmds
}
