package main

import (
    "fmt"
)

func main() {
    // 0-360
    rotate := 20.0
    
    for i := 360-rotate; i <= 360; i += 1.0 { 
        fmt.Println(i)
    }
    fmt.Println("----")
    for i := 0.0; i < 360-rotate; i += 1.0 { 
        fmt.Println(i)
    }
}