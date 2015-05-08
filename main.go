package main

import (
    "bufio"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

func isText(filename string) bool {
    file, err := os.Open(filename)
    if err != nil {
        println("cannot open", filename)
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    b := scanner.Bytes()
    s := http.DetectContentType(b)
    return strings.HasPrefix(s, "text/")
}

func main() {
    for _, v := range os.Args[1:] {
        if isText(v) {
            fmt.Println(v)
        }
    }
}
