package main

import (
    "aoc/utils"
    "fmt"
    "regexp"
    "strconv"
    "math"
)

func main(){
    lines := utils.GetInput(13)
    //lines := utils.GetExampleInput(13)

    exp := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)
    match := exp.FindAllStringSubmatch(lines, -1)
    machines := []*clawMachine{}
    for _, part := range match{
        //fmt.Println(part)
        xa, _ := strconv.Atoi(part[1])
        ya, _ := strconv.Atoi(part[2])
        xb, _ := strconv.Atoi(part[3])
        yb, _ := strconv.Atoi(part[4])
        xp, _ := strconv.Atoi(part[5])
        yp, _ := strconv.Atoi(part[6])
        buttonA := coords{xa, ya}
        buttonB := coords{xb, yb}
        prize := coords{xp, yp}
        machines = append(machines, &clawMachine{&buttonA, &buttonB, &prize})
    }

    p1Solution := 0
    for _, machine := range machines{
        p1Solution += calcTokens(machine)
    }
    fmt.Println("Part1 Solution:", p1Solution)

    //10000000000000
    for _, machine := range machines{
        machine.prize.x += 10000000000000
        machine.prize.y += 10000000000000
    }

    p2Solution := 0
    for _, machine := range machines{
        p2Solution += calcTokens(machine)
    }
    fmt.Println("Part2 Solution:", p2Solution)
}

func calcTokens(machine *clawMachine) int{
    M := [2][2]float64{
        {float64(machine.buttonA.x), float64(machine.buttonB.x)},
        {float64(machine.buttonA.y), float64(machine.buttonB.y)},
    }
    v := [2]float64{float64(machine.prize.x), float64(machine.prize.y)}

    det := M[0][0] * M[1][1] - M[0][1] * M[1][0]

    M1 := [2][2]float64{
        {M[1][1], -M[0][1]},
        {-M[1][0], M[0][0]},
    }

    a := int(math.Round(1/det*(M1[0][0] * v[0] + M1[0][1] * v[1])))
    b := int(math.Round(1/det*(M1[1][0] * v[0] + M1[1][1] * v[1])))
    //fmt.Println(a,b)
    //fmt.Println(a, a*machine.buttonA.x+b*machine.buttonB.x, machine.prize.x)
    //fmt.Println(b, a*machine.buttonA.y+b*machine.buttonB.y, machine.prize.y)
    if a < 0 || b < 0{
        return 0
    }
    if !(a*machine.buttonA.x+b*machine.buttonB.x == machine.prize.x && a*machine.buttonA.y+b*machine.buttonB.y == machine.prize.y){
        return 0
    }
    //fmt.Println(a, b)
    return a*3+b
}

type clawMachine struct {
    buttonA *coords
    buttonB *coords
    prize *coords
}

type coords struct{
    x int
    y int
}

/*func calcTokens(machine *clawMachine){
    // 8400 = a*94 + b*22
    // prizeX = a*Xa + b*Xb
    // (prizeX-Xa*a)/Xb = b
    for a := range 100{
        bx := (machine.prize.x-machine.buttonA.x*a)/machine.buttonB.x
        by := (machine.prize.y-machine.buttonA.y*a)/machine.buttonB.y
        if bx > 100 || by > 100{
            continue
        }

        tx := machine.prize.x == a*machine.buttonA.x + bx*machine.buttonB.x
        ty := machine.prize.y == a*machine.buttonA.y + by*machine.buttonB.y

        if tx && ty && bx == by{
            price := a*3 + bx
            if price < machine.price || machine.price == 0{
                //fmt.Println(machine.prize.x, machine.prize.y, a, bx, price)
                machine.price = price
            }
        }
    }
}*/
