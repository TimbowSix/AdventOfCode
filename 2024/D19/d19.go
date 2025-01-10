package main

import (
    "aoc/utils"
    "fmt"
    "regexp"
    "strings"
)

func main(){
    lines := utils.GetInputLines(19)
    //lines := utils.GetExampleInputLines(19)
    patternParts := []string{}
    patterns := []string{}
    b := false
    rpat := "^(?:"
    for _, line := range lines{
        if line == ""{
            b = true
        }else if b{
            patterns = append(patterns, line)
        }else{
            s := strings.Split(line, ", ")
            patternParts = append(patternParts, s...)
            for _, s := range s{
                rpat += "(?:"
                rpat += s
                rpat += ")|"
            }
        }
    }
    rpat += ")+$"
    //fmt.Println(patternParts)
    //fmt.Println(patterns)
    //fmt.Println(rpat)
    exp := regexp.MustCompile(rpat)
    p1Solution := 0
    for _, pat := range patterns{
        if exp.MatchString(pat){
            p1Solution++
        }
    }
    fmt.Println("Part1 Solution:", p1Solution)

    p2Solution := 0
    cache := make(map[string]int)
    for _, pat := range patterns{
        //fmt.Println(pat)
        p2Solution += patternPossible(pat, &patternParts, &cache)
    }
    fmt.Println("Part2 Solution:", p2Solution)

}

func patternPossible(pattern string, parts *[]string, cache *map[string]int) int{
    if pattern == ""{
        return 1
    }
    val, ok := (*cache)[pattern]
    if ok{
        return val
    }
    c := 0
    for _, part := range *parts{
        if strings.HasPrefix(pattern, part){
            new := pattern[len(part):]
            n := patternPossible(new, parts, cache)
            (*cache)[new] = n
            c+=n
        }
    }
    return c
}

