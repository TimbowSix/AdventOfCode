package main

import (
    "aoc/utils"
    "fmt"
)

func main(){
    lines := utils.GetInputLines(12)
    //lines := utils.GetExampleInputLines(12)
    grid := [][]*plot{}
    for y, line := range lines{
        row := []*plot{}
        for x, chr := range line{
            row = append(row, &plot{x,y, string(chr), nil})
        }
        grid = append(grid, row)
    }

    regions := []*region{}
    for _, row := range grid{
        for _, plt := range row{
            if plt.region != nil{
                continue
            }
            currRegion := region{plots: make([]*plot, 0)}
            currRegion.plots = append(currRegion.plots, plt)
            plt.region = &currRegion
            buff := []*plot{plt}
            for len(buff) > 0{
                currPlot := buff[0]
                buff = buff[1:]
                coords := [][2]int{
                    {-1, 0},
                    {1, 0},
                    {0, -1},
                    {0, 1},
                }
                for _, c := range coords {
                    next := utils.GetGrid(&grid, currPlot.x+c[0], currPlot.y+c[1])
                    if next == nil{
                        currRegion.perimeter++
                        addBorder(&currRegion, currPlot.x+c[0], currPlot.y+c[1], c)
                    }else{
                        if (*next).region == nil && (*next).plant == currPlot.plant{
                            buff = append(buff, *next)
                            (*next).region = &currRegion
                            currRegion.plots = append(currRegion.plots, *next)
                        }else if (*next).plant != currPlot.plant{
                            currRegion.perimeter++
                            addBorder(&currRegion, currPlot.x+c[0], currPlot.y+c[1], c)
                        }
                    }
                }
            }
            if len(currRegion.plots) > 0{
                regions = append(regions, &currRegion)

            }
        }
    }
    //fmt.Println(len(regions), "regions")
    p1Solution := 0
    for _, r := range regions{
        //fmt.Println(r.plots[0].plant, len(r.plots), r.perimeter)
        p1Solution += r.perimeter*len(r.plots)
    }
    fmt.Println("Part1 Solution:", p1Solution)

    p2Solution := 0
    for _, r := range regions{
        l := findStraights(r)
        p2Solution += l * len(r.plots)
    }
    fmt.Println("Part2 Solution:", p2Solution)

}

func addBorder(reg *region, x int, y int, direction [2]int){
    if direction == [2]int{-1, 0}{//left
        reg.leftBorder = append(reg.leftBorder, [2]int{x,y})
    }else if direction == [2]int{1, 0}{//right
        reg.rightBorder = append(reg.rightBorder, [2]int{x,y})
    }else if direction == [2]int{0, -1}{//up
        reg.upBorder = append(reg.upBorder, [2]int{x,y})
    }else if direction == [2]int{0, 1}{//down
        reg.downBorder = append(reg.downBorder, [2]int{x,y})
    }else{
        panic("Invalid dimension")
    }
}

func findStraights(reg *region) int{
    straights := 0
    straights += findHorizontals(&reg.upBorder)
    straights += findHorizontals(&reg.downBorder)
    straights += findVerticals(&reg.leftBorder)
    straights += findVerticals(&reg.rightBorder)
    return straights
}

func findVerticals(coords *[][2]int) int{
    straights := 0
    arr := (*coords)
    for len(arr) > 0{
        currCoord := arr[0]
        arr = arr[1:]
        upmost := currCoord[1]
        downmost := currCoord[1]
        for i := 0; i<len(arr); i++{
            coord := arr[i]
            if coord[0] == currCoord[0]{
                if coord[1] == downmost+1{
                    downmost = coord[1]
                    arr = append(arr[:i], arr[i+1:]...)
                    i = -1
                }else if coord[1] == upmost-1{
                    upmost = coord[1]
                    arr = append(arr[:i], arr[i+1:]...)
                    i = -1
                }
            }
        }
        straights++
        //fmt.Println("v", upmost, downmost)
    }

    return straights

}

func findHorizontals(coords *[][2]int) int{
    straights := 0
    arr := (*coords)
    for len(arr) > 0{
        currCoord := arr[0]
        arr = arr[1:]
        leftmost := currCoord[0]
        rightmost := currCoord[0]
        for i := 0; i<len(arr); i++{
            coord := arr[i]
            if coord[1] == currCoord[1]{
                if coord[0] == rightmost+1{
                    rightmost = coord[0]
                    arr = append(arr[:i], arr[i+1:]...)
                    i = -1
                }else if coord[0] == leftmost-1{
                    leftmost = coord[0]
                    arr = append(arr[:i], arr[i+1:]...)
                    i = -1
                }
            }
        }
        straights++
        //fmt.Println("h", leftmost, rightmost, "y:", currCoord[0])
    }
    return straights
}

type plot struct{
    x int
    y int
    plant string
    region *region
}

type region struct{
    plots []*plot
    perimeter int
    leftBorder [][2]int
    rightBorder [][2]int
    upBorder [][2]int
    downBorder [][2]int
}

