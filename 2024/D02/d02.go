package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main(){
    lines := utils.GetInputLines(2)
    reports := make([][]int, len(lines))
    for lineiIdx, line := range lines{
        vals := strings.Split(line, " ")
        reports[lineiIdx] = make([]int, len(vals))
        for reportIdx, n := range vals{
            num, _ := strconv.Atoi(n)
            reports[lineiIdx][reportIdx] = num
        }
    }

    safeReports := 0
    for _, report := range reports{
        increasing := report[0] < report[1]
        valid := true
        for i := 1; i < len(report); i++{
            diff := int(math.Abs(float64(report[i]-report[i-1])))
            if diff < 1 || diff > 3{
                valid = false
                break
            }
            if increasing && report[i] < report[i-1]{
                valid = false
                break
            }
            if !increasing && report[i] > report[i-1]{
                valid = false
                break
            }

        }
        if valid{
            safeReports++
        }

    }

    fmt.Println("Part1 Result:", safeReports)


    dampenerSafe := 0
    for _, report := range reports{
        for ignoreIdx := -1; ignoreIdx < len(report); ignoreIdx++{
            prevVal := -1
            valid := true
            increasing := 0
            decreasing := 0
            for idx, val := range report{
                if idx == ignoreIdx{
                    continue
                }
                if prevVal > -1{
                    diff := int(math.Abs(float64(val - prevVal)))
                    if diff < 1 || diff > 3{
                        valid = false
                        break
                    }
                    if val > prevVal{
                        increasing++
                    }
                    if val < prevVal{
                        decreasing++
                    }
                }
                prevVal = val
            }
            if valid && (increasing == 0 || decreasing == 0){
                dampenerSafe++
                break
            }
        }
    }

    fmt.Println("Part2 Result:", dampenerSafe)
}
