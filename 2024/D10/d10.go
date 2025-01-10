package main

import (
    "aoc/utils"
    "fmt"
)

func main(){
    lines := utils.GetInputLines(10)
    //lines := utils.GetExampleInputLines(10)

    grid := [][]*position{}
    trailheads := []*position{}
    for y, line := range lines{
        row := []*position{}
        for x, chr := range line{
            height := int(chr - '0')
            pos := position{x, y, height, false}
            row = append(row, &pos)
            if height == 0{
                trailheads = append(trailheads, &pos)
            }
        }
        grid = append(grid, row)
    }

    part1Solution := 0
    for _, head := range trailheads{
        buff := []*position{head}
        for len(buff) != 0{
            pos := buff[0]
            buff = buff[1:]
            pos.visited = true

            //up
            checkNextPos(&grid, &buff, pos.x, pos.y-1, pos.height)
            //down
            checkNextPos(&grid, &buff, pos.x, pos.y+1, pos.height)
            //left
            checkNextPos(&grid, &buff, pos.x-1, pos.y, pos.height)
            //right
            checkNextPos(&grid, &buff, pos.x+1, pos.y, pos.height)

        }
        //printGrid(&grid)
        part1Solution += resetVisited(&grid)
    }
    fmt.Println("Part1 Solution:", part1Solution)


    part2Solution := 0

    for _, head := range trailheads{
        part2Solution += findPath(&grid, head)
    }

    fmt.Println("Part2 Solution:", part2Solution)
}

func findPath(grid *[][]*position, currPos *position) int{
    if currPos.height == 9{
        return 1
    }
    val := 0
    var up **position = utils.GetGrid(grid, currPos.x, currPos.y-1)
    var down **position = utils.GetGrid(grid, currPos.x, currPos.y+1)
    var left **position = utils.GetGrid(grid, currPos.x-1, currPos.y)
    var right **position = utils.GetGrid(grid, currPos.x+1, currPos.y)

    tmp := []**position{up, down, left, right}
    for _, next := range tmp{
        if next != nil && (*next).height == currPos.height+1{
            val += findPath(grid, *next)
        }
    }
    return val
}


func checkNextPos(grid *[][]*position, buff *[]*position, x int, y int, currHeight int){
    var next **position = utils.GetGrid(grid, x, y)
    if next != nil{
        nextPos := *next
        if !nextPos.visited && nextPos.height == currHeight+1{
            *buff = append(*buff, nextPos)
        }
    }
}

func resetVisited(grid *[][]*position) int{
    score := 0
    for _, row := range *grid{
        for _, pos := range row{
            if pos.visited && pos.height == 9{
                score++
            }
            pos.visited = false
        }
    }
    return score
}

func printGrid(grid *[][]*position){
    for _, row := range *grid{
        for _, pos := range row{
            if pos.visited{
                fmt.Print("x", " ")
            }else{
                fmt.Print(pos.height, " ")
            }
        }
        fmt.Print("\n")
    }
    fmt.Print("\n")
}

type position struct{
    x int
    y int
    height int
    visited bool
}
