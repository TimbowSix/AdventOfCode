package main

import (
    "aoc/utils"
    "fmt"
    "regexp"
    "strconv"
)

func main(){
    file := utils.GetInput(3)

    re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
    match := re.FindAllStringSubmatch(file, -1)
    result := 0
    for _, part := range match{
        n1, _ := strconv.Atoi(part[1])
        n2, _ := strconv.Atoi(part[2])
        result += n1 * n2
    }

    fmt.Println("Part1 Result:", result)

    re = regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\))|(?:don't\(\))|(?:do\(\))`)
    active := true
    result = 0
    for _, part := range re.FindAllStringSubmatch(file, -1){
        if part[0] == "do()"{
            active = true
        }else if part[0] == "don't()"{
            active = false
        }else if active{
            n1, _ := strconv.Atoi(part[1])
            n2, _ := strconv.Atoi(part[2])
            result += n1 * n2
        }
    }

    fmt.Println("Part2 Result:", result)
}
