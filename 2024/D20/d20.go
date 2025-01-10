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
    lines := utils.GetInputLines(20)
    //lines := utils.GetExampleInputLines(20)

    grid := getGrid(&lines)
	var start *node
	var end *node
	for _, row := range *grid{
        for _, n := range row{
			if n.symbol == "S"{
				start = n
			}else if n.symbol == "E"{
				end = n
			}
		}
	}
	baseLength, _ := findPath(start, end)
	//fmt.Println(baseLength)
	//paintMap(grid, end)
	resetGrid(grid)

	savings := make(map[int]int)
	for _, row := range *grid{
		for _, n := range row{
			if n.symbol != "#"{
				continue
			}
			n.symbol = "."
			pLen, _ := findPath(start, end)
			diff := baseLength-pLen
			if pLen < baseLength{
				//fmt.Println(x,y, pLen, baseLength-pLen)
				savings[diff]++
			}
			n.symbol = "#"
			resetGrid(grid)
		}
	}
	//fmt.Println(savings)

	p1Result := 0
	for k, v := range savings{
		if k>=100{
			p1Result += v
		}
	}
	fmt.Println("Part1 Solution:", p1Result)

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

func getGrid(lines *[]string) *[][]*node{
    grid := [][]*node{}
	height := len(*lines)
	width := len((*lines)[0])
    for y := range height{
        row := []*node{}
        for x := range width{
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
			n.symbol = string((*lines)[y][x])
        }
    }
    return &grid
}

func resetGrid(grid *[][]*node){
	for _, row := range *grid{
		for _, n := range row{
			n.distance = math.MaxInt
			n.latestCandidate = nil
			n.latest = nil
		}
	}
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

    f, _ := os.Create("D20__test.png")
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
