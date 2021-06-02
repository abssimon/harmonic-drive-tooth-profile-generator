package main

import(
    "fmt"
    "math"
)

func main() {
    a := 4.0
    b := 3.0
    
    // Punkt auf der Ellipse
    tx := -1.8
    ty := 3.2
    
    // b2 x2 + a2 y2 = a2b2
    a2 := math.Pow(a, 2)
    b2 := math.Pow(b, 2)
    
    a2b2 := a2 * b2
    
    fmt.Println(tx, ty, a2b2)
    
    fmt.Println(a2b2/(a2*tx))
    
}