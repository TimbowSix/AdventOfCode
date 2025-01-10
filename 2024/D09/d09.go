package main

import (
    "aoc/utils"
    "fmt"
    "slices"
    "strconv"
)

func main(){
    lines := utils.GetInputLines(9)
    //lines := utils.GetExampleInputLines(9)

    currId := 0
    storage := []int{}
    for idx, c := range lines[0]{
        n, err := strconv.Atoi(string(c))
        utils.ErrorCheck(err)
        if idx % 2 == 0{
            for range n{
                storage = append(storage, currId)
            }
            currId++
        }else{
            for range n{
                storage = append(storage, -1)
            }
        }
    }

    //fmt.Println(storage)
    part1Solution := 0
    for idx, n := range storage{
        if n == -1{
            for i := len(storage)-1; i > idx; i--{
                if storage[i] != -1{
                    storage[idx] = storage[i]
                    storage[i] = -1
                    break
                }
            }
        }
        //fmt.Println(storage)
        if storage[idx] != -1{
            part1Solution += idx * storage[idx]
        }
    }

    fmt.Println("Part1 Solution:", part1Solution)

    currId = 0
    files := []*file{}
    for idx, c := range lines[0]{
        len, err := strconv.Atoi(string(c))
        utils.ErrorCheck(err)
        file := file{-1, len}
        if idx % 2 == 0{
            file.id = currId
            currId++
        }
        files = append(files, &file)
    }

    for idx, pFile := range files{
        if pFile.id == -1{
            for i := len(files)-1; i > idx; i--{
                if files[i].id != -1{
                    if files[i].length == pFile.length{
                        files[idx], files[i] = files[i], files[idx]
                        break
                    }else if files[i].length < pFile.length{
                        lenDiff := pFile.length - files[i].length
                        new1 := file{-1, files[i].length}
                        new2 := file{-1, lenDiff}
                        files[idx] = files[i]
                        files[i] = &new1
                        files = slices.Insert(files, idx+1, &new2)
                        break
                    }
                }
            }
        }
        //printFiles(&files)
    }
    part2Solution := 0
    idx := 0
    for _, file := range files{
        if file.id == -1{
            // skip empty space
            idx += file.length
            continue
        }
        for range file.length{
            part2Solution += idx * file.id
            idx++
        }
    }

    fmt.Println("Part2 Solution:", part2Solution)

}

type file struct{
    id int
    length int
}

func printFiles(files *[]*file){
    for _, file := range (*files){
        for range file.length{
            if file.id == -1{
                fmt.Print(".")
            }else{
                fmt.Print(file.id)
            }
        }
    }
    fmt.Print("\n")
}
