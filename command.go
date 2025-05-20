package main

type command struct {
	name					string
	description 	string
	callback			func() error
}

