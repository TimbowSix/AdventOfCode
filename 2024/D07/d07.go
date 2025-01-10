package main

import (
    "aoc/utils"
    "fmt"
    "math"
    "strconv"
    "strings"
)

type calibration struct{
    expectedResult int
    values []int
}

func checkCalibration(cal *calibration, operatorCount int) bool{
    n := len(cal.values)
    for i := range int(math.Pow(float64(operatorCount), float64(n))){
        t := strconv.FormatInt(int64(i), operatorCount)
        bin := fmt.Sprintf("%0*s", n, t)
        tmpRes := cal.values[0]
        for j := 1; j < n; j++{
            if string(bin[j-1]) == "0"{
                tmpRes += cal.values[j]
            }else if string(bin[j-1]) == "1"{
                tmpRes *= cal.values[j]
            }else if string(bin[j-1]) == "2"{
                n1 := strconv.Itoa(tmpRes)
                n2 := strconv.Itoa(cal.values[j])
                newN, _ := strconv.Atoi(n1+n2)
                tmpRes = newN
            }else{
                fmt.Println("Undefined operator", string(bin[j-1]))
                panic("Undefined operator")
            }
        }
        if tmpRes == cal.expectedResult{
            return true
        }
    }
    return false
}

func main(){
    lines := utils.GetInputLines(7)
    //lines := utils.GetExampleInputLines(7)

    calibrations := []calibration{}
    for _, line := range lines{
        t := strings.Split(line, ": ")
        result, err := strconv.Atoi(t[0])
        utils.ErrorCheck(err)
        numbers := []int{}
        for _, n := range strings.Split(t[1], " "){
            num, err := strconv.Atoi(n)
            utils.ErrorCheck(err)
            numbers = append(numbers, num)
        }
        c := calibration{result, numbers}
        calibrations = append(calibrations, c)
    }

    p1Result := 0
    for _, cal := range calibrations{
        if checkCalibration(&cal, 2){
            p1Result += cal.expectedResult
        }
    }
    fmt.Println("Part1 Solution:", p1Result)

    p2Result := 0
    for _, cal := range calibrations{
        if checkCalibration(&cal, 3){
            p2Result += cal.expectedResult
        }
    }
    fmt.Println("Part2 Solution:", p2Result)
}
