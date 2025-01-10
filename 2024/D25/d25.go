package main

import (
    "aoc/utils"
    "fmt"
    "strings"
)

func main(){
    lines := utils.GetInput(25)
    //lines := utils.GetExampleInput(25)

    locks := [][5]int{}
    keys := [][5]int{}

    rawSchemas := strings.Split(lines, "\n\n")
    for _, raw := range rawSchemas{
        schematic := strings.Split(raw, "\n")
        t := &locks
        if schematic[0][0] == '.'{//key
            t = &keys
        }
        heights := [5]int{}
        for i := 1; i<6; i++{
            row := schematic[i]
            for idx, chr := range row{
                if chr == '#'{
                    heights[idx]++
                }
            }
        }
        *t = append(*t, heights)
    }

    //fmt.Println(locks)
    //fmt.Println(keys)

    p1Solution := 0
    for _, key := range keys{
        for _, lock := range locks{
            valid := true
            for i := 0; i < 5; i++{
                if key[i]+lock[i] > 5{
                    valid = false
                    break
                }
            }
            if valid{
                p1Solution++
            }
        }
    }

    fmt.Println("Part1 Solution:", p1Solution)

}
