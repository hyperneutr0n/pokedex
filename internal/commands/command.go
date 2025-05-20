package commands

type command struct {
	name					string
	description 	string
	callback			func(*Config) error
}

type Config struct {
	Next			string
	Previous	string
}