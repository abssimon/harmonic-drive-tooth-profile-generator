package main

import(
    //"fmt"
    "math"
)

// https://rechneronline.de/trigonometrie/erweiterte-hyperbelfunktionen.php
// gives back numbers between 0 (the ends) and 1 (middle)
func HyperbelScale(a float64, degreesToNextTooth float64) (float64) { // 
    x := ((a / degreesToNextTooth) - 0.5) * 8.0 // minus middle -0.5, *8 so nice range for cosh
    max := (2.0 * math.Cosh(4.0 / 2.0)) - 2
    return math.Abs((((2.0 * math.Cosh(x / 2.0)) - 2) - max) / max)
}

// wie damit auf winkel kommen?
// dir richtige summe 
// schritte bis zum naechsten Punkt vorab aufaddieren und dann prozent von
/*
func main() {
    fmt.Println(HyperbelScale(0.0, 5.5))
    fmt.Println(HyperbelScale(1.5, 5.5))
    fmt.Println(HyperbelScale(5.4, 5.5))
    fmt.Println(HyperbelScale(5.5, 5.5))
}*/