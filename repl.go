package main

import "strings"

func cleanInput(text string) []string {
	result := strings.Fields(text)
	return result
}