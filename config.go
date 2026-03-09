package main

import (
	"fmt"
	"log"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	Graph string
	BarComponent string
}

func main() {
	conf := Config{
		"test",
		"test",
	}
	out, err := yaml.Marshal(&conf)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	fmt.Println(string(out))
}
