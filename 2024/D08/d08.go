package main

import (
    "aoc/utils"
    "fmt"
    "math"
)

type node struct{
    x int
    y int
    freq string
    hasAntinode bool
}

func main(){
    lines := utils.GetInputLines(8)
    //lines := utils.GetExampleInputLines(8)

    allNodes := []*node{}
    antennaNodes := []*node{}
    for y, line := range lines{
        for x, chr := range line{
            new := node{x, y, string(chr), false}
            allNodes = append(allNodes, &new)
            if new.freq != "."{
                antennaNodes = append(antennaNodes, &new)
            }
        }
    }

    for _, a1 := range antennaNodes{
        for _, a2 := range antennaNodes{
            if a1 == a2 || a1.freq != a2.freq{
                continue
            }
            for _, n := range allNodes{
                if n == a1 || n == a2{
                    continue
                }
                if (a2.y - a1.y) * (n.x - a1.x) == (a2.x - a1.x) * (n.y - a1.y){
                    dist1 := int(math.Max(
                        math.Abs(float64(n.x-a1.x)),
                        math.Abs(float64(n.y-a1.y)),
                    ))
                    dist2 := int(math.Max(
                        math.Abs(float64(n.x-a2.x)),
                        math.Abs(float64(n.y-a2.y)),
                    ))

                    if dist1 == dist2*2 || dist2 == dist1*2{
                        n.hasAntinode = true
                    }
                }
            }
        }
    }

    p1Result := 0
    for _, n := range allNodes{
        if n.hasAntinode{
            p1Result++
        }
    }
    //printNodes(&allNodes)
    fmt.Println("Part1 Solution:", p1Result)

    for _, a1 := range antennaNodes{
        for _, a2 := range antennaNodes{
            if a1 == a2 || a1.freq != a2.freq{
                continue
            }
            for _, n := range allNodes{
                if (a2.y - a1.y) * (n.x - a1.x) == (a2.x - a1.x) * (n.y - a1.y){
                    n.hasAntinode = true
                }
            }
        }
    }

    p2Result := 0
    for _, n := range allNodes{
        if n.hasAntinode{
            p2Result++
        }
    }

    fmt.Println("Part2 Solution:", p2Result)
}

func printNodes(nodes *[]*node){
    var p *node
    for _, n := range (*nodes){
        if p != nil && p.y < n.y{
            fmt.Print("\n")
        }
        if n.hasAntinode{
            fmt.Print("#", " ")
        }else{
            fmt.Print(n.freq, " ")
        }
        p = n
    }
    fmt.Println("")
}
