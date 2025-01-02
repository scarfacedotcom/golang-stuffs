package main

import "fmt"

func getDamage(weaponType string) int {
	switch weaponType {
	case "axe":
		return 30
	case "sword":
		return 20
	case "knife":
		return 10
	case "wood":
		return 5
	case "stick":
		return 1
	default:
		panic("weapon does not exist")
	}
}

func main() {
	fmt.Println("the damage caused is:", getDamage("axe"))
}
