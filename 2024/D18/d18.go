package main

import (
    "aoc/utils"
    "fmt"
    "image"
    "image/color"
    "image/png"
    "math"
    "os"
    "slices"
    "sort"
    "strconv"
    "strings"
)

func main(){
    lines := utils.GetInputLines(18)
    //lines := utils.GetExampleInputLines(18)
    bytePos := [][2]int{}
    for _, line := range lines{
        s := strings.Split(line, ",")
        x, _ := strconv.Atoi(s[0])
        y, _ := strconv.Atoi(s[1])
        bytePos = append(bytePos, [2]int{x,y})
    }

    grid := getGrid(71)
    for i := 0; i<1024; i++{
        b := bytePos[i]
        n := utils.GetGridP(grid, b[0], b[1])
        n.symbol = "#"
    }

    s := (*grid)[0][0]
    e := (*grid)[len(*grid)-1][len((*grid)[0])-1]
    length, ok := findPath(s, e)
    if ok{
        fmt.Println("Part1 Solution:", length)
    }
    paintMap(grid, e)

    last := 1024
    for ok{
        last++
        grid := getGrid(71)
        for i := 0; i<=last; i++{
            b := bytePos[i]
            n := utils.GetGridP(grid, b[0], b[1])
            n.symbol = "#"
        }

        s = (*grid)[0][0]
        e = (*grid)[len(*grid)-1][len((*grid)[0])-1]
        length, ok = findPath(s, e)
        //fmt.Println(last, length, ok)
        if ok{
            //paintMap(grid, e)
        }
    }
    fmt.Printf("Part2 Solution: %d,%d\n", bytePos[last][0], bytePos[last][1])

}

func findPath(start *node, end *node) (int, bool) {
    currNode := start
    currNode.distance = 0
    visible := []*node{}
    for currNode != end{
        for _, n := range currNode.dirs{
            if n == nil || n.symbol == "#" || n.latest != nil{
                continue
            }
            dist := currNode.distance+1

            if dist < n.distance{
                n.distance = dist
                n.latestCandidate = currNode
                if !slices.Contains(visible, n){
                    visible = append(visible, n)
                }
            }
        }
        if len(visible) == 0 {
            //no path
            return 0, false
        }
        sort.Slice(visible, func(i, j int) bool {
            return visible[i].distance < visible[j].distance
        })
        next := visible[0]
        visible = visible[1:]
        currNode = next
        currNode.latest = currNode.latestCandidate
    }
    return currNode.distance, true
}

func getGrid(size int) *[][]*node{
    grid := [][]*node{}
    for y := range size{
        row := []*node{}
        for x := range size{
            new := &node{x: x, y: y, symbol: ".", distance: math.MaxInt}
            row = append(row, new)
        }
        grid = append(grid, row)
    }
    for y, row := range grid{
        for x, n := range row{
            l := utils.GetGridP(&grid, x-1, y)
            r := utils.GetGridP(&grid, x+1, y)
            u := utils.GetGridP(&grid, x, y-1)
            d := utils.GetGridP(&grid, x, y+1)
            n.dirs = []*node{u, r, d, l}
        }
    }
    return &grid
}

func paintMap(grid *[][]*node, end *node){
    visited := []*node{}
    curr := end
    for curr != nil{
        visited = append(visited, curr)
        curr = curr.latest
    }
    height := len(*grid)
    width := len((*grid)[0])
    upLeft := image.Point{0, 0}

    lowRight := image.Point{width, height}

    img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            img.Set(x, y, color.RGBA{255, 0, 0, 255})
        }
    }

    for y, row := range *grid{
        for x, n := range row{
            if slices.Contains(visited, n){
                img.Set(x, y, color.White)
            }else if n.symbol == "#"{
                img.Set(x, y, color.Black)
            }
            if n == visited[0] || n == visited[len(visited)-1]{
                img.Set(x, y, color.RGBA{0, 255, 0, 255})
            }
        }
    }

    f, _ := os.Create("D18__test.png")
    png.Encode(f, img)
}


type node struct{
    x int
    y int
    symbol string

    dirs []*node

    distance int
    latest *node
    latestCandidate *node
}
