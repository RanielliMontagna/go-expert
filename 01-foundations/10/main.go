package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Client struct {
	Name    string
	Age     int
	Active  bool
	Address Address
}

func (c *Client) InactiveClient() {
	c.Active = false
}

func main() {
	ranni := Client{
		Name:   "Ranni",
		Age:    25,
		Active: true,
	}

	ranni.InactiveClient()

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", ranni.Name, ranni.Age, ranni.Active)
}
