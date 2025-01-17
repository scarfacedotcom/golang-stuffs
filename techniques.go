package main

import "fmt"

type Entity struct {
	name    string
	id      string
	version string
	posx    int
	posy    int
}

type SpecialEntity struct {
	Entity
	specialField float64
}

func main() {
	e := Entity{
		name:    "scar face",
		id:      "id 1",
		version: "version 1.1",
		posx:    200,
		posy:    400,
	}

	fmt.Printf("%+v\n", e)
}
