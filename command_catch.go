package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)
	res := rand.Intn(pokemonResp.BaseExperience)
	if res > 40 {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		cfg.pokedex[pokemonResp.Name] = pokemonResp
	} else {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	}
	return nil
}
