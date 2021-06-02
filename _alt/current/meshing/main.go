package main

import (
	"fmt"
	"math"
)

// https://rechneronline.de/trigonometrie/erweiterte-hyperbelfunktionen.php
// gives back numbers between 0 (the ends) and 1 (middle)
func HyperbelScale(a float64, degreesToNextTooth float64) float64 { //
	x := ((a / degreesToNextTooth) - 0.5) * 6.0 // minus middle -0.5, *6 so nice range for cosh
	max := (2.0 * math.Cosh(3.0/2.0)) - 2       // max on one side to scale output, -2 so zero
	return math.Abs((((2.0 * math.Cosh(x/2.0)) - 2) - max) / max)
}

func main() {
	fmt.Println(HyperbelScale(0.0, 5.5))
	fmt.Println(HyperbelScale(1.0, 5.5))
	fmt.Println(HyperbelScale(5.4, 5.5))
	fmt.Println(HyperbelScale(5.5, 5.5))
}
