package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Person interface {
	Inactive()
}

type Client struct {
	Name    string
	Age     int
	Active  bool
	Address Address
}

type Company struct {
	Name string
}

func (c *Company) Inactive() {
}

func (c *Client) Inactive() {
	c.Active = false
}

func Desactivate(p Person) {
	p.Inactive()
}

func main() {
	ranni := Client{
		Name:   "Ranni",
		Age:    25,
		Active: true,
	}
	myCompany := Company{
		Name: "My Company",
	}

	Desactivate(&ranni)
	Desactivate(&myCompany)

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", ranni.Name, ranni.Age, ranni.Active)
}
