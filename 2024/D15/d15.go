package main

import (
    "aoc/utils"
    "fmt"
)

func main(){
    lines := utils.GetInputLines(15)
    //lines := utils.GetExampleInputLines(15)

    grid := [][]string{}
    robotPos := [2]int{}
    maxY := 0
    grid2 := [][]string{}
    for y, line := range lines{
        if line == ""{
            maxY = y
            break
        }
        row := []string{}
        row2 := []string{}
        for x, obj := range line{
            symbol := string(obj)
            row = append(row, symbol)
            if symbol == "@"{
                robotPos[0] = x
                robotPos[1] = y
                row2 = append(row2, "@")
                row2 = append(row2, ".")
            }else if symbol == "."{
                row2 = append(row2, ".")
                row2 = append(row2, ".")
            }else if symbol == "#"{
                row2 = append(row2, "#")
                row2 = append(row2, "#")
            }else if symbol == "O"{
                row2 = append(row2, "[")
                row2 = append(row2, "]")
            }
        }
        grid = append(grid, row)
        grid2 = append(grid2, row2)
    }
    //printGrid(&grid2)
    robotPos2 := [2]int{}
    for y, row := range grid2{
        for x, symbol := range row{
            if symbol == "@"{
                robotPos2[0] = x
                robotPos2[1] = y
                break
            }
        }
    }

    for i := maxY+1; i < len(lines); i++{
        line := lines[i]
        for _, c := range line{
            instruction := string(c)
            moveRobot(&grid, instruction, &robotPos)
            moveRobot2(&grid2, instruction, &robotPos2)
            //fmt.Println(instruction)
            //printGrid(&grid2)
        }
    }
    fmt.Println("Part1 Solution:", sumCoords(&grid))
    fmt.Println("Part2 Solution:", sumCoords2(&grid2))

}

func sumCoords(grid *[][]string) int{
    sum := 0
    for y, row := range *grid{
        for x, symbol := range row{
            if symbol == "O"{
                sum += y*100+x
            }
        }
    }
    return sum
}


func moveRobot(grid *[][]string, instruction string, robotPos *[2]int){
    if move(grid, *robotPos, "@", instruction){
        (*grid)[robotPos[1]][robotPos[0]] = "."
        next := [2]int{robotPos[0], robotPos[1]}
        getNext(instruction, &next)
        robotPos[0] = next[0]
        robotPos[1] = next[1]
    }

}

func move(grid *[][]string, pos [2]int, symbol string, instruction string) bool{
    getNext(instruction, &pos)
    x := pos[0]
    y := pos[1]
    next := (*grid)[y][x]
    if next == "#"{
        return false
    }else if next == "."{
        (*grid)[y][x] = symbol
        return true
    }else if next == "O"{
        b := move(grid, pos, "O", instruction)
        if b{
            (*grid)[y][x] = symbol
        }
        return b
    }else{
        panic(fmt.Sprintf("Unknown symbol %s", symbol))
    }
}


func getNext(instruction string, pos *[2]int){
    switch instruction{
    case "<":
        pos[0]--
    case ">":
        pos[0]++
    case "^":
        pos[1]--
    case "v":
        pos[1]++
    }
}

func printGrid(grid *[][]string){
    for _, row := range *grid{
        for _, c := range row{
            fmt.Print(c, " ")
        }
        fmt.Print("\n")
    }
    fmt.Print("\n")
}


func moveRobot2(grid *[][]string, instruction string, robotPos *[2]int){
    var b bool
    if instruction == "<" || instruction == ">"{
        b = move2Horizontal(grid, *robotPos, "@", instruction)
    }else{
        if canMove2Vertical(grid, *robotPos, "@", instruction){
            move2Vertical(grid, *robotPos, "@", instruction)
            b = true
        }
    }
    if b{
        (*grid)[robotPos[1]][robotPos[0]] = "."
        next := [2]int{robotPos[0], robotPos[1]}
        getNext(instruction, &next)
        robotPos[0] = next[0]
        robotPos[1] = next[1]
    }

}

func move2Horizontal(grid *[][]string, pos [2]int, symbol string, instruction string) bool{
    getNext(instruction, &pos)
    x := pos[0]
    y := pos[1]
    next := (*grid)[y][x]
    if next == "#"{
        return false
    }else if next == "."{
        (*grid)[y][x] = symbol
        return true
    }else if next == "["{
        b := move2Horizontal(grid, pos, "[", instruction)
        if b{
            (*grid)[y][x] = symbol
        }
        return b
    }else if next == "]"{
        b := move2Horizontal(grid, pos, "]", instruction)
        if b{
            (*grid)[y][x] = symbol
        }
        return b
    }else{
        panic(fmt.Sprintf("Unknown symbol %s", symbol))
    }
}

func canMove2Vertical(grid *[][]string, pos [2]int, symbol string, instruction string) bool{
    getNext(instruction, &pos)
    x := pos[0]
    y := pos[1]
    next := (*grid)[y][x]
    if next == "#"{
        return false
    }else if next == "."{
        return true
    }else if next == "["{
        posB := [2]int{pos[0]+1, pos[1]}
        a := canMove2Vertical(grid, pos, "[", instruction)
        b := canMove2Vertical(grid, posB, "]", instruction)
        return a&&b
    }else if next == "]"{
        posB := [2]int{pos[0]-1, pos[1]}
        a := canMove2Vertical(grid, pos, "]", instruction)
        b := canMove2Vertical(grid, posB, "[", instruction)
        return a&&b
    }else{
        panic(fmt.Sprintf("Unknown symbol %s", symbol))
    }
}

func move2Vertical(grid *[][]string, pos [2]int, symbol string, instruction string){
    (*grid)[pos[1]][pos[0]] = "."
    getNext(instruction, &pos)
    x := pos[0]
    y := pos[1]
    next := (*grid)[y][x]
    if next == "#"{
        return
    }else if next == "."{
        (*grid)[y][x] = symbol
    }else if next == "["{
        posB := [2]int{pos[0]+1, pos[1]}
        move2Vertical(grid, pos, "[", instruction)
        move2Vertical(grid, posB, "]", instruction)
        (*grid)[y][x] = symbol

    }else if next == "]"{
        posB := [2]int{pos[0]-1, pos[1]}
        move2Vertical(grid, pos, "]", instruction)
        move2Vertical(grid, posB, "[", instruction)
        (*grid)[y][x] = symbol
    }else{
        panic(fmt.Sprintf("Unknown symbol %s", symbol))
    }
}

func sumCoords2(grid *[][]string) int{
    sum := 0
    for y, row := range *grid{
        for x, symbol := range row{
            if symbol == "["{
                sum += y*100+x
            }
        }
    }
    return sum
}
