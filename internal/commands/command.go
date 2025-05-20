package commands

type command struct {
	name					string
	description 	string
	callback			func() error
}

