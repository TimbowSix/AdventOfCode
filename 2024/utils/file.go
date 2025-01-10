package utils

import (
    "bufio"
    "errors"
    "fmt"
    "io"
    "net/http"
    "os"
    "github.com/joho/godotenv"
)


func ReadLines(path string) []string {
    file, err := os.Open(path)
    ErrorCheck(err)
    defer file.Close()
    scanner := bufio.NewScanner(file)

    arr := []string{}
    for scanner.Scan() {
        arr = append(arr, scanner.Text())
    }
    return arr
}

func DownloadInput(day int){
    path := fmt.Sprintf("./D%02d/input", day)
    _, err := os.Stat(path)
    if !errors.Is(err, os.ErrNotExist){
        // file exists
        return
    }
    fmt.Println("Downloading input for day", day)
    godotenv.Load()
    session := os.Getenv("SESSION_COOKIE")

    url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)

    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    ErrorCheck(err)
    req.Header.Set("Cookie", "session=" + session)
    resp, err := client.Do(req)
    ErrorCheck(err)
    data, err := io.ReadAll(resp.Body)
    ErrorCheck(err)
    file, err := os.Create(path)
    ErrorCheck(err)

    file.Write(data)
}

func GetInputLines(day int) []string {
    path := fmt.Sprintf("./D%02d/input", day)
    _, err := os.Stat(path)
    if errors.Is(err, os.ErrNotExist){
        DownloadInput(day)
    }
    return ReadLines(path)
}

func GetInput(day int) string{
    path := fmt.Sprintf("./D%02d/input", day)
    _, err := os.Stat(path)
    if errors.Is(err, os.ErrNotExist){
        DownloadInput(day)
    }
    file, _ := os.ReadFile(path)
    return string(file)
}

func GetExampleInput(day int) string{
	path := fmt.Sprintf("./D%02d/input_example", day)
    _, err := os.Stat(path)
    ErrorCheck(err)
    file, _ := os.ReadFile(path)
    return string(file)
}

func GetExampleInputLines(day int) []string {
    path := fmt.Sprintf("./D%02d/input_example", day)
    _, err := os.Stat(path)
    ErrorCheck(err)
    return ReadLines(path)
}
