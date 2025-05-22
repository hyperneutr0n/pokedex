package commands

import "github.com/hyperneutr0n/pokedex/internal/pokecache"

type command struct {
	name					string
	description 	string
	callback			func([]string, *Config) error
}

type Config struct {
	Next			string
	Previous	string
	Cache			*pokecache.Cache
}