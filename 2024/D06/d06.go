package main

import (
    "aoc/utils"
    "fmt"
    "slices"
    "strconv"
)

func main(){
    lines := utils.GetInputLines(6)
    //lines := utils.GetExampleInputLines(6)
    grid := [][]string{}

    for _, line := range lines{
        arr := make([]string, len(line))
        for i, r := range line {
            arr[i] = string(r)

        }
        grid = append(grid, arr)
    }

    var currY int
    var currX int
    for y, row := range grid{
        for x, val := range row{
            if val == "^"{
                currX = x
                currY = y
                break
            }
        }
        if currX != 0{
            break
        }
    }
    var initX int = currX
    var initY int = currY

    //0 up, 1 right, 2 down, 3 left
    var currDir int = 0
    var visitedFields int = 0
    for {
        currField := grid[currY][currX]
        if currField != "X"{
            visitedFields++
            grid[currY][currX] = "X"
        }
        var nextX int
        var nextY int
        for {
            switch currDir {
            case 0:
                //up
                nextX = currX
                nextY = currY - 1
            case 1:
                //right
                nextX = currX + 1
                nextY = currY
            case 2:
                //down
                nextX = currX
                nextY = currY + 1
            case 3:
                //left
                nextX = currX - 1
                nextY = currY
            }
            if nextY < len(grid) && nextY >= 0 && nextX < len(grid[0]) && nextX >= 0 &&  grid[nextY][nextX] == "#"{
                currDir = (currDir + 1) % 4
            }else{
                break
            }
        }

        if !(nextY < len(grid) && nextY >= 0 && nextX < len(grid[0]) && nextX >= 0){
            break
        }else{
            currX = nextX
            currY = nextY
        }

    }

    fmt.Println("Part1 Solution:", visitedFields)

    var loopCount int = 0
    for y, row := range grid{
        for x, field := range row{
            if field == "#" || field == "^"{
                continue
            }
            grid[y][x] = "#"
            currX = initX
            currY = initY
            //0 up, 1 right, 2 down, 3 left
            currDir = 0
            visitedObs := make(map[string][]int)
            foundLoop := false
            for {
                var nextX int
                var nextY int
                for {
                    switch currDir {
                    case 0:
                        //up
                        nextX = currX
                        nextY = currY - 1
                    case 1:
                        //right
                        nextX = currX + 1
                        nextY = currY
                    case 2:
                        //down
                        nextX = currX
                        nextY = currY + 1
                    case 3:
                        //left
                        nextX = currX - 1
                        nextY = currY
                    }
                    if nextY < len(grid) && nextY >= 0 && nextX < len(grid[0]) && nextX >= 0 &&  grid[nextY][nextX] == "#"{
                        id := strconv.Itoa(nextX)+"-"+strconv.Itoa(nextY)
                        val, ok := visitedObs[id]
                        if ok {
                            if slices.Contains(val, currDir){
                                loopCount++
                                foundLoop = true
                                break
                            }
                        }
                        visitedObs[id] = append(visitedObs[id], currDir)
                        currDir = (currDir + 1) % 4
                    }else{
                        break
                    }
                }

                if !(nextY < len(grid) && nextY >= 0 && nextX < len(grid[0]) && nextX >= 0) || foundLoop{
                    break
                }else{
                    currX = nextX
                    currY = nextY
                }
            }
            grid[y][x] = "."
        }

    }

    fmt.Println("Part2 Solution:", loopCount)
}
