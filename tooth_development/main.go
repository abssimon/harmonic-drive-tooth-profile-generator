package main

import(
    "math"
)


func main() {
        
    e := Ellipse{Point{600.0, 500.0}, 300.0, 337.0} // px, py, h, w 

    

    tooth := EllipseTooth(e, 0, Point{0,0}, 125.0)
    
    // rotate    
    tooth.C1.Rotate(-math.Pi/2)
    tooth.C2.Rotate(-math.Pi/2)
    tooth.C3.Rotate(-math.Pi/2)
    tooth.C4.Rotate(-math.Pi/2)
    
    
    svg(e, tooth)
}