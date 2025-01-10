package main

import (
    "aoc/utils"
    "fmt"
)

func lookup(lines *[]string, x int, y int) bool{
    buff := []rune{}
    if y-3 >= 0{
        for i := range 4{
            char := rune((*lines)[y-i][x])
            buff = append(buff, char)
        }
        if string(buff) == "XMAS"{
            return true
        }
    }
    return false
}
func lookdown(lines *[]string, x int, y int) bool{
    buff := []rune{}
    if y+3 < len((*lines)){
        for i := range 4{
            char := rune((*lines)[y+i][x])
            buff = append(buff, char)
        }
        if string(buff) == "XMAS"{
            return true
        }
    }
    return false
}
func lookleft(lines *[]string, x int, y int) bool{
    line := (*lines)[y]
    end := x-3
    if end >= 0{
        sub := line[end:x+1]
        sub = utils.ReverseString(sub)
        if sub == "XMAS"{
            return true
        }

    }
    return false
}
func lookright(lines *[]string, x int, y int) bool{
    line := (*lines)[y]
    end := x+4
    if end <= len((*lines)){
        sub := line[x: end]
        if sub == "XMAS"{
            return true
        }
    }
    return false
}
func lookdownleft(lines *[]string, x int, y int) bool{
    if y + 3 < len((*lines)) && x - 3 >= 0{
        l := (*lines)
        buff := []rune{
            rune(l[y][x]),
            rune(l[y+1][x-1]),
            rune(l[y+2][x-2]),
            rune(l[y+3][x-3]),
        }
        if string(buff) == "XMAS"{
            return true
        }
    }
    return false
}
func lookdownright(lines *[]string, x int, y int) bool{
    l := (*lines)
    if y + 3 < len(l) && x + 3 < len(l[0]){
        buff := []rune{
            rune(l[y][x]),
            rune(l[y+1][x+1]),
            rune(l[y+2][x+2]),
            rune(l[y+3][x+3]),
        }
        if string(buff) == "XMAS"{
            return true
        }
    }
    return false
}
func lookupleft(lines *[]string, x int, y int) bool{
    l := (*lines)
    if y - 3 >= 0 && x - 3 >= 0{
        buff := []rune{
            rune(l[y][x]),
            rune(l[y-1][x-1]),
            rune(l[y-2][x-2]),
            rune(l[y-3][x-3]),
        }
        if string(buff) == "XMAS"{
            return true
        }
    }
    return false
}
func lookupright(lines *[]string, x int, y int) bool{
    l := (*lines)
    if y - 3 >= 0 && x + 3 < len(l[0]){
        buff := []rune{
            rune(l[y][x]),
            rune(l[y-1][x+1]),
            rune(l[y-2][x+2]),
            rune(l[y-3][x+3]),
        }
        if string(buff) == "XMAS"{
            return true
        }
    }
    return false
}

func main(){
    lines := utils.GetInputLines(4)

    xmasCount := 0

    for lIdx, line := range lines{
        for cIdx, char := range line{
            if string(char) == "X"{
                if lookdown(&lines, cIdx, lIdx){
                    xmasCount++
                    //fmt.Println("found down at", cIdx, lIdx)
                }
                if lookup(&lines, cIdx, lIdx){
                    xmasCount++
                    //fmt.Println("found up at", cIdx, lIdx)
                }
                if lookleft(&lines, cIdx, lIdx){
                    xmasCount++
                    //fmt.Println("found left at", cIdx, lIdx)
                }
                if lookright(&lines, cIdx, lIdx){
                    xmasCount++
                    //fmt.Println("found right at", cIdx, lIdx)
                }
                if lookupleft(&lines, cIdx, lIdx){
                    xmasCount++
                    //fmt.Println("found upleft at", cIdx, lIdx)
                }
                if lookupright(&lines, cIdx, lIdx){
                    xmasCount++
                    //fmt.Println("found upright at", cIdx, lIdx)
                }
                if lookdownleft(&lines, cIdx, lIdx){
                    xmasCount++
                    //fmt.Println("found downleft at", cIdx, lIdx)
                }
                if lookdownright(&lines, cIdx, lIdx){
                    xmasCount++
                    //fmt.Println("found downright at", cIdx, lIdx)
                }
            }
        }
    }

    fmt.Println("Solution Part1:", xmasCount)

    x_masCount := 0
    for lIdx, line := range lines{
        for cIdx, char := range line{
            if string(char) == "A"{
                if lIdx - 1 >= 0 && cIdx -1 >= 0 && lIdx + 1 < len(lines) && cIdx + 1 < len(lines[0]){
                    s1 := []byte{
                        lines[lIdx-1][cIdx-1],
                        lines[lIdx][cIdx],
                        lines[lIdx+1][cIdx+1],
                    }
                    s2 := []byte{
                        lines[lIdx-1][cIdx+1],
                        lines[lIdx][cIdx],
                        lines[lIdx+1][cIdx-1],
                    }
                    //fmt.Println(string(s1), string(s2))
                    if (string(s1) == "MAS" || utils.ReverseString(string(s1)) == "MAS") && (string(s2) == "MAS" || utils.ReverseString(string(s2)) == "MAS"){
                        x_masCount++
                    }
                }
            }
        }
    }

    fmt.Println("Solution Part2:", x_masCount)
}
