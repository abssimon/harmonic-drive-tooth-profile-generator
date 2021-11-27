package main

import (
    "math"
)

func main() {
    
    e := Ellipse{Point{600.0, 500.0}, 327.0, 347.0} // px, py, h, w 
    teeth := 100 
    modulo := e.Circumference() / float64(teeth)

    gear := []Point{}
    pLast := Point{}
    tcount := 0
    sumDistance := 0.0
    
    // build gear
    start := 0.0;
    first := true
    for i := start; i <= math.Pi*2; i += 0.000017453 {  // ~ 0.001 degree

        p := e.PointAtAngle(i)
        
        // first teeth
        if first {
            gear = append(gear, EllipseTooth(e, i, p)...)
            tcount++
            first = false
            pLast = p
            continue
        }

        // set next teeth by distance
        sumDistance += p.DistanceTo(&pLast) / 2.0  // TODO, mal checken ob das in summe circumferecne gibt !!!!!!!!!!!!!!!!!!!!!!!!!
        pLast = p

        if sumDistance >= (float64(tcount) * modulo) {
            gear = append(gear, EllipseTooth(e, i, p)...)
            tcount++
        }

        if tcount == teeth {
            break
        }
    }
    
    // rotate
    for i, _ := range gear {
        gear[i].Rotate(math.Pi/4)
    }
    
    svg(e, gear)
}