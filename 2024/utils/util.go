package utils

import (
    "fmt"
    "math"
)

func ErrorCheck(e error) {
    if e != nil {
        fmt.Println(e)
        panic(e)
    }
}

func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func GetIfInRange[T any](arr *[]T, idx int) *T{
    if idx >= 0 && idx < len(*arr){
        return &(*arr)[idx]
    }
    return nil
}

func GetGrid[T any](grid *[][]T, x int, y int) *T{
    row := GetIfInRange(grid, y)
    if row == nil{
        return nil
    }
    return GetIfInRange(row, x)
}

func GetGridP[T any](grid *[][]*T, x int, y int) *T{
    row := GetIfInRange(grid, y)
    if row == nil{
        return nil
    }
    val := GetIfInRange(row, x)
    if val == nil{
        return nil
    }
    return *val
}

func IntPow(x int, y int) int{
    return int(math.Pow(float64(x), float64(y)))
}
