package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, mon := range cfg.pokedex {
		fmt.Println("  - ", mon.Name)
	}
	return nil
}
