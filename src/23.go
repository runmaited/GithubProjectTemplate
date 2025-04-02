package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    var n int = 5
    for i := 0; i < n; i++ {
        fmt.Printf("%d", rand.Intn(10) + 1)
    }
}
