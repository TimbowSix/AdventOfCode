package main

import (
    "aoc/utils"
    "fmt"
    "strconv"
    "strings"
)

func main(){
    lines := utils.GetInputLines(11)
    //lines := utils.GetExampleInputLines(11)

    stones := []*stone{}
    for _, num := range strings.Split(lines[0], " "){
        val, err := strconv.Atoi(num)
        utils.ErrorCheck(err)
        stones = append(stones, &stone{value: val, amount: 1})
    }

    for range 25{
        blink(&stones)
    }
    fmt.Println("Part1 Solution:", countStones(&stones))

    for range 50{
        blink(&stones)
    }
    fmt.Println("Part2 Solution:", countStones(&stones))
}

func blink(stones *[]*stone){
    stoneMap := make(map[int]*stone)
    for _, currStone := range *stones{
        if currStone.value == 0{
            currStone.value = 1
        }else{
            str := strconv.Itoa(currStone.value)
            if len(str) % 2 == 0{
                left, err := strconv.Atoi(str[:len(str)/2])
                utils.ErrorCheck(err)
                right, err := strconv.Atoi(str[len(str)/2:])
                utils.ErrorCheck(err)
                currStone.value = left
                newStone := stone{value: right, amount: currStone.amount}
                addStoneMap(&stoneMap, &newStone)
            }else{
                currStone.value *= 2024
            }
        }
        addStoneMap(&stoneMap, currStone)
    }
    newStones := []*stone{}
    for _, s := range stoneMap{
        newStones = append(newStones, s)
    }
    *stones = newStones
}

func addStoneMap(stoneMap *map[int]*stone, s *stone){
    val, ok := (*stoneMap)[s.value]
    if ok{
        val.amount += s.amount
    }else{
        (*stoneMap)[s.value] = s
    }
}

func countStones(stones *[]*stone) int{
    count := 0
    for _, s := range *stones{
        count += s.amount
    }
    return count
}

type stone struct{
    value int
    amount int
}
