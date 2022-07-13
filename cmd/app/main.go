package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name string `json:"name"`
	age  int    `json:"age"`
}

func main() {
	p := Person{Name: "John", age: 16}
	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	log.Print(string(bytes))
}
