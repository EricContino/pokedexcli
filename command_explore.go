package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationResp, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationResp.Name)
	fmt.Println("Found Pokemon: ")
	for _, mon := range locationResp.PokemonEncounters {
		fmt.Println(mon.Pokemon.Name)
	}
	return nil
}
