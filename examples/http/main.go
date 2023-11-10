package main

import (
	media "github.com/colber/go-sdk/client"
)

func main() {

	client, err := media.NewClient()

	filter :=map[string][]string{}
	client.Find(filter)
}