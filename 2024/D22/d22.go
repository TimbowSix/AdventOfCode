package main

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strings"
)

func main(){
    //lines := utils.GetInputLines(22)
    lines := utils.GetExampleInputLines(22)

	computers := make(map[string]*computer)
	for _, line := range lines{
		names := strings.Split(line, "-")

		n1, ok := computers[names[0]]
		if !ok{
			new := computer{name: names[0]}
			n1 = &new
			computers[new.name] = &new
		}

		n2, ok := computers[names[1]]
		if !ok{
			new := computer{name: names[1]}
			n2 = &new
			computers[new.name] = &new
		}

		n1.connections = append(n1.connections, n2)
		n2.connections = append(n2.connections, n1)
	}

}

type computer struct{
	name string
	connections []*computer
}
