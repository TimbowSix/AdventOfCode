package main

import (
    "aoc/utils"
    "fmt"
    "slices"
    "strconv"
    "strings"
)

func checkValid(update *[]int, rules *[][2]int) bool{
    prevs := []int{}
    for _, n := range (*update){
        for _, rule := range (*rules){
            // first n not yet seen, but in update set; no duplicate numbers
            if n == rule[1] && !slices.Contains(prevs, rule[0]) && slices.Contains((*update), rule[0]){
                return false
            }
        }
        prevs = append(prevs, n)
    }
    return true
}

func main(){
    lines := utils.GetInputLines(5)
    //lines = utils.GetExampleInputLines(5)

    rules := [][2]int{}
    updates := [][]int{}
    for _, line := range lines{
        if strings.Contains(line, "|"){
            s := strings.Split(line, "|")
            n1, err := strconv.Atoi(s[0])
            utils.ErrorCheck(err)
            n2, err := strconv.Atoi(s[1])
            utils.ErrorCheck(err)
            arr := [2]int{n1, n2}
            rules = append(rules, arr)
        }else if line != ""{
            s := strings.Split(line, ",")
            update := []int{}
            for _, n := range s{
                num, err := strconv.Atoi(n)
                utils.ErrorCheck(err)
                update = append(update, num)
            }
            updates = append(updates, update)
        }
    }

    var  validMiddleSum int = 0
    invalids := [][]int{}
    for _, update := range updates{
        if checkValid(&update, &rules) {
            validMiddleSum += update[len(update) / 2]
        }else{
            invalids = append(invalids, update)
        }
    }

    fmt.Println("Part1 Solution:", validMiddleSum)
    //fmt.Println(invalids)
    var sum int = 0
    for _, update := range invalids{
        for !checkValid(&update, &rules){
            prevs := []int{}
            for secIdx, n := range update{
                t := false
                for _, rule := range rules{
                    if n == rule[1] && !slices.Contains(prevs, rule[0]) && slices.Contains(update, rule[0]){
                        fstIdx := slices.Index(update, rule[0])
                        update[secIdx], update[fstIdx] = update[fstIdx], update[secIdx]
                        t = true
                        break
                    }
                }
                if t{
                    break
                }
                prevs = append(prevs, n)
            }
        }
        sum += update[len(update) / 2]
    }
    //fmt.Println(invalids)

    fmt.Println("Part2 Solution:", sum)

}
