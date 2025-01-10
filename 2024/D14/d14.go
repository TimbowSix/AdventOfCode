package main

import (
    "aoc/utils"
    "fmt"
    "image"
    "image/color"
    "image/png"
    "os"
    "regexp"
    "strconv"
)

const height int = 103//7
const width int = 101//11

func main(){
    lines := utils.GetInput(14)
    //lines := utils.GetExampleInput(14)

    exp := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
    match := exp.FindAllStringSubmatch(lines, -1)
    robots := []*Robot{}
    for _, part := range match{
        x, _ := strconv.Atoi(part[1])
        y, _ := strconv.Atoi(part[2])
        vx, _ := strconv.Atoi(part[3])
        vy, _ := strconv.Atoi(part[4])
        robots = append(robots, &Robot{x,y,vx,vy})
    }
    //fmt.Println(len(robots), "robots")


    for i := range 100{
        for _, robot := range robots{
            updateRobot(robot)
        }
        printRobots(&robots, i+1)
    }

    tl := 0
    tr := 0
    bl := 0
    br := 0
    for _, robot := range robots{
        //fmt.Println(robot.x, robot.y)
        if robot.x < (width-1)/2{
            if robot.y < (height-1)/2{
                tl++
            }else if robot.y > height/2{
                bl++
            }
        }else if robot.x > width/2{
            if robot.y < (height-1)/2{
                tr++
            }else if robot.y > height/2{
                br++
            }
        }
    }
    p1Solution := tl*tr*bl*br
    fmt.Println("Part1 Solution:", p1Solution)

    for i := range 9900{
        for _, robot := range robots{
            updateRobot(robot)
        }
        printRobots(&robots, i+100+1)
    }

}

func updateRobot(robot *Robot){
    robot.x = robot.x + robot.vx
    if robot.x < 0{
        robot.x = width+robot.x
    }else if robot.x >= width{
        robot.x = robot.x%width
    }
    robot.y = robot.y + robot.vy
    if robot.y < 0{
        robot.y = height+robot.y
    }else if robot.y >= height{
        robot.y = robot.y%height
    }
}

func printRobots(arr *[]*Robot, i int){
    upLeft := image.Point{0, 0}
    lowRight := image.Point{width, height}

    img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            img.Set(x, y, color.Black)
        }
    }

    for _, robot := range *arr{
        x := robot.x
        y := robot.y
        img.Set(x, y, color.White)
    }

    n := fmt.Sprintf("D14/__test/image_%d.png", i)
    f, _ := os.Create(n)
    png.Encode(f, img)
}

type Robot struct{
    x int
    y int
    vx int
    vy int
}
