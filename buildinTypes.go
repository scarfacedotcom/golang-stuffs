package main

import "fmt"

type Player struct {
	name        string
	health      int
	attackPower float64
}

func (player Player) getHealth() int {
	return player.health
}

func main1() {
	// player := Player{
	// 	name:        "scar face",
	// 	health:      101,
	// 	attackPower: 45.7,
	// }

	users := make(map[string]int)

	users["scar"] = 3
	users["face"] = 2

	// age, ok := users["scar"]

	// if !ok {
	// 	fmt.Println("scar does not exist in the map")
	// } else {
	// 	fmt.Println("scar exist in the map", age)
	// }

	//fmt.Printf("health: %d\n", player.getHealth())
	//fmt.Printf("age, %+d\n", age)

	for k, v := range users {
		fmt.Printf("the key %s and the value %d\n", k, v)
	}
}
