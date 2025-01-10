package main

import (
    "aoc/utils"
    "fmt"
    "os"
    "slices"
    "strconv"
    "strings"
)

func main(){
    lines := utils.GetInputLines(17)
    //lines := utils.GetExampleInputLines(17)

    regA, _ := strconv.Atoi(lines[0][12:])
    regB, _ := strconv.Atoi(lines[1][12:])
    regC, _ := strconv.Atoi(lines[2][12:])
    programm := []int{}
    codes := strings.Split(lines[4][9:], ",")
    for i := 0; i < len(codes); i+=2{
        opcode, _ := strconv.Atoi(codes[i])
        operand, _ := strconv.Atoi(codes[i+1])
        programm = append(programm, opcode, operand)
    }

    result := solve(&programm, regA, regB, regC)
    stringSlice := make([]string, len(result))
    for i, num := range result {
        stringSlice[i] = strconv.Itoa(num)
    }
    fmt.Println("Part1 Solution:", strings.Join(stringSlice, ","))

    //part2test()
    fmt.Println("Part2 Solution:", part2(&programm, regA, regB, regC))
}

func solve(program *[]int, regA int, regB int, regC int) []int{
    getComboOperand := func (operand int) int{
        switch operand{
        case 0,1,2,3:
            return operand
        case 4:
            return regA
        case 5:
            return regB
        case 6:
            return regC
        default:
            panic("Illegal combo operand")
        }
    }

    insPointer := 0
    outputBuff := []int{}
    for insPointer < len(*program){
        //fmt.Println(regA)
        opcode := (*program)[insPointer]
        operand := (*program)[insPointer+1]
        switch opcode{
        case 0:
            //adv
            //fmt.Println("adv", regA, regA / utils.IntPow(2, getComboOperand(operand)))
            regA = regA / utils.IntPow(2, getComboOperand(operand))
            insPointer+=2
        case 1:
            //bxl
            //fmt.Println("bxl", regB, regB^operand)
            regB = regB^operand
            insPointer+=2
        case 2:
            //bst
            //fmt.Println("bst", regB, getComboOperand(operand) % 8)
            regB = getComboOperand(operand) % 8
            insPointer+=2
        case 3:
            //jnz
            //fmt.Println("jnz", insPointer, operand / 2)
            if regA != 0{
                insPointer = operand / 2
            }else{
                insPointer+=2
            }
        case 4:
            //bxc
            //fmt.Println("bxc", regB, regB^regC)
            regB = regB^regC
            insPointer+=2
        case 5:
            //out
            //fmt.Println("out", getComboOperand(operand), getComboOperand(operand) % 8)
            s :=  getComboOperand(operand) % 8
            outputBuff = append(outputBuff, s)
            insPointer+=2
        case 6:
            //bdv
            //fmt.Println("bdv", regB, regA / utils.IntPow(2, getComboOperand(operand)))
            regB = regA / utils.IntPow(2, getComboOperand(operand))
            insPointer+=2
        case 7:
            //cdv
            //fmt.Println("cdv", regC, regA / utils.IntPow(2, getComboOperand(operand)))
            regC = regA / utils.IntPow(2, getComboOperand(operand))
            insPointer+=2
        }
    }
    return outputBuff
}


func part2(program *[]int, regA int, regB int, regC int) int{
    candidates := []int{0}
    c := 1
    lowest := 999999999999999999
    for i := len(*program)-1; i >= 0; i--{
        //fmt.Println(candidates)
        val := (*program)[i]
        //fmt.Println("val", val)
        buff := append([]int{}, candidates...)
        candidates = []int{}
        for _, n := range buff{
            //fmt.Println("n", n)
            for j := range 8{
                //fmt.Println("j", j)
                nVal := n << 3
                nVal = nVal|j
                //fmt.Println("nval", nVal)
                tmp := solve(program, nVal, regB, regC)
                //fmt.Println("tmp", tmp)
                if tmp[len(tmp)-c] == val{
                    candidates = append(candidates, nVal)
                }
            }
        }
        c++
        for _, v := range candidates{
            tmp := solve(program, v, regB, regC)
            if slices.Equal(*program, tmp){
                //fmt.Println(tmp, v)
                if v < lowest{
                    lowest = v
                }
            }
        }
    }
    //fmt.Println(lowest)
    //fmt.Println(*program)
    //fmt.Println(solve(program, lowest, regB, regC))

    return lowest
}

func part2test(){
    lines := utils.GetInputLines(17)
    //lines := utils.GetExampleInputLines(17)

    //1, 8, 64, 512, 4096
    //1 0 1 3 5 4 7 6 3 1 1 7 5 4 7 6 5 2 1 3 5 4 6 6 7 3 1 7 5 4 6 6 1 4 1 3 5 4 5 7 3 5 1 7 5 4 5 7 5 6 1 3 5 4 4 7 7 7 1 7 5 4 4 7

    f, err := os.Create("d17__test6")
    utils.ErrorCheck(err)
    defer f.Close()

    program := []int{}
    codes := strings.Split(lines[4][9:], ",")
    for i := 0; i < len(codes); i+=2{
        opcode, _ := strconv.Atoi(codes[i])
        operand, _ := strconv.Atoi(codes[i+1])
        program = append(program, opcode, operand)
    }

    b, _ := strconv.Atoi(lines[1][12:])
    c, _ := strconv.Atoi(lines[2][12:])

    for i := 911111111111111; i < 999999999999999*2; i++{
        //fmt.Println(i)
        result := solve(&program, i, b, c)

        //fmt.Print(result[0], " ")
        stringSlice := make([]string, len(result))
        for i, num := range result {
            stringSlice[i] = strconv.Itoa(num)
        }
        f.WriteString(strconv.Itoa(i))
        f.WriteString(" ")
        f.WriteString(strings.Join(stringSlice, ","))
        f.WriteString("\n")

        if slices.Equal(program, result){
            //break
        }
    }
}
