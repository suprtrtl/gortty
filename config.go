package main

import (
	"errors"
	"log"
	"os"

	"go.yaml.in/yaml/v4"
)

var defaultData = []int{432, 17, 298, 401, 56, 73, 489, 120, 345, 210, 67, 154, 399, 278, 44, 311, 92, 407, 188, 265, 134, 358, 21, 476, 303, 84, 250, 168, 392, 59, 147, 326, 415, 203, 12, 437, 290, 361, 75, 224}

type Config struct { // `config.yaml`
	Graph        string // graph: bar
	BarComponent string // barcomponent: ▊
	Delay        int    // delay: 50
	Spacing      int    // spacing: 2
}

func ConfigFromYaml(path string) Config {
	c := Config{}
	data, err := os.ReadFile(path)
	// Crappy solution fix later
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return c
}

func (c *Config) ToModel() (m model, err error) {

	dims := Dimension{0,0,c.Spacing}

	m = model{
		delay: c.Delay,
		dims: dims,
	}

	switch c.Graph {
	case "bar":
		m.graph = BarGraph{
			component: c.BarComponent,
		}
	default:
		return model{}, errors.New("Invalid Config @ \"graph\"");
	}

	return m, nil
}
