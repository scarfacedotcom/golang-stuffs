package main

import "fmt"

func main() {
	//number := []int{1, 2, 3, 4, 5, 6, 7, 8}

	users := map[string]int{
		"scar":  1,
		"face":  2,
		"peter": 3,
		"Tor":   4,
	}

	// for i := 0; i < len(number); i++ {
	// 	fmt.Println("here we go", i)
	// }

	for k, v := range users {
		fmt.Printf("k %s, v %d\n", k, v)
	}
}
