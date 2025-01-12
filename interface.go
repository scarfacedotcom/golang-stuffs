package main

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type MongoDBNumberStore struct {
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("store the number into the data base")
	return nil
}

type ApiSever struct {
	numberStore NumberStorer
}

func main4() {
	apiServer := ApiSever{
		numberStore: MongoDBNumberStore{},
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(numbers)
}
