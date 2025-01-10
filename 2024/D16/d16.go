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
)

func main(){
    grid, start, end := getGrid()

    distance, ok := findPath(start, end)
    if !ok{
        fmt.Println("No Path found")
    }else{
        fmt.Println("Part1 Solution:", distance)
    }
    paintVisitedMap(grid, end)

    fmt.Println("Part2 Solution:", findPaths())
}

func findPath(start *node, end *node) (int, bool) {
    currNode := start
    currNode.distance = 0
    currNode.dir = 1
    visible := []*node{}
    for currNode != end{
        dirs := []*node{currNode.up, currNode.right, currNode.down, currNode.left}
        for dir, n := range dirs{
            if n == nil || n.symbol == "#" || n.latest != nil{
                continue
            }
            dist := currNode.distance
            if dir == currNode.dir{
                dist += 1
            }else{
                dist += 1000 + 1
            }
            if dist < n.distance{
                n.distance = dist
                n.dir = dir
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

func getGrid() (*[][]*node, *node, *node){
    lines := utils.GetInputLines(16)
    //lines := utils.GetExampleInputLines(16)

    var start *node
    var end *node
    grid := [][]*node{}
    for y, line := range lines{
        row := []*node{}
        for x, chr := range line{
            symbol := string(chr)
            new := &node{x: x, y: y, symbol: symbol, distance: math.MaxInt}
            row = append(row, new)
            if symbol == "S"{
                start = new
            }
            if symbol == "E"{
                end = new
            }
        }
        grid = append(grid, row)
    }
    //printMap(&grid)

    for y, row := range grid{
        for x, n := range row{
            l := utils.GetGrid(&grid, x-1, y)
            r := utils.GetGrid(&grid, x+1, y)
            u := utils.GetGrid(&grid, x, y-1)
            d := utils.GetGrid(&grid, x, y+1)
            if l == nil {
                n.left = nil
            }else{
                n.left = *l
            }

            if r == nil {
                n.right = nil
            }else{
                n.right = *r
            }

            if u == nil {
                n.up = nil
            }else{
                n.up = *u
            }

            if d == nil {
                n.down = nil
            }else{
                n.down = *d
            }
        }
    }
    return &grid, start, end
}

func findPaths() int{
    grid, s, e := getGrid()
    fields := [][2]int{}
    max_dist, _ := findPath(s, e)
    for y, row := range *grid{
        for x := range row{
            grid, start, end := getGrid()
            if (*grid)[y][x].symbol == "S" || (*grid)[y][x].symbol == "E" || (*grid)[y][x].symbol == "#"{
                continue
            }
            (*grid)[y][x].symbol = "#"
            dist, ok := findPath(start, end)
            if ok && dist == max_dist{
                curr := end
                for curr != nil{
                    coords := [2]int{curr.x, curr.y}
                    if !slices.Contains(fields, coords){
                        fields = append(fields, coords)
                    }
                    curr = curr.latest
                }
            }

        }
        fmt.Println(y)
    }
    return len(fields)
}


func printMap(grid *[][]*node){
    for _, row := range *grid{
        for _, n := range row{
            fmt.Print(n.symbol, " ")
        }
        fmt.Print("\n")
    }
    fmt.Print("\n")
}

func printVisitedMap(grid *[][]*node, end *node){
    visited := []*node{}
    curr := end
    for curr != nil{
        visited = append(visited, curr)
        curr = curr.latest
    }

    for _, row := range *grid{
        for _, n := range row{
            if slices.Contains(visited, n){
                fmt.Print("x", " ")
            }else{
                fmt.Print(n.symbol, " ")
            }
        }
        fmt.Print("\n")
    }
    fmt.Print("\n")
}

func outputVisitedMap(grid *[][]*node, end *node){
    visited := []*node{}
    curr := end
    for curr != nil{
        visited = append(visited, curr)
        curr = curr.latest
    }
    f, err := os.Create("d16__test")
    utils.ErrorCheck(err)
    defer f.Close()

    for _, row := range *grid{
        for _, n := range row{
            if slices.Contains(visited, n){
                f.WriteString("x")
            }else{
                f.WriteString(n.symbol)
            }
        }
        f.WriteString("\n")
    }
    f.WriteString("\n")
}

func paintVisitedMap(grid *[][]*node, end *node){
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

    f, _ := os.Create("D16__test.png")
    png.Encode(f, img)
}

type node struct{
    x int
    y int
    symbol string

    up *node
    down *node
    left *node
    right *node

    distance int
    dir int
    latest *node
    latestCandidate *node
}
