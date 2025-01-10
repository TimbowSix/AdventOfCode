package main

import (
    "aoc/utils"
    "fmt"
    "math"
    "sort"
    "strconv"
    "strings"
)

func main(){
    utils.DownloadInput(1)
    lines := utils.ReadLines("D01/input")
    leftValues := []int{}
    rightValues := []int{}
    for _, line := range lines{
        l, r, _ := strings.Cut(line, "   ")
        l_i, _ := strconv.Atoi(l)
        r_i, _ := strconv.Atoi(r)
        leftValues = append(leftValues, l_i)
        rightValues = append(rightValues, r_i)
    }

    sort.Slice(leftValues, func(i, j int) bool {
        return leftValues[i] < leftValues[j]
    })
    sort.Slice(rightValues, func(i, j int) bool {
        return rightValues[i] < rightValues[j]
    })

    diffSum := 0
    for i := 0; i < len(leftValues); i++{
        diffSum += int(math.Abs(float64(leftValues[i] - rightValues[i])))
    }

    fmt.Println("Part1 Result:", diffSum)

    simScore := 0
    for i := 0; i < len(leftValues); i++{
        lVal := leftValues[i]
        for j := 0; j < len(rightValues); j++{
            rVal := rightValues[j]
            if lVal == rVal{
                simScore += rVal
            }
        }
    }

    fmt.Println("Part2 Result:", simScore)

}
