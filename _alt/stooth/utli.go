package main

import (
    "math"
)

const intScaleFactor = 10000000000000

// float64 to int for clipper
func scaleUp(value float64) int {
    return int(value * intScaleFactor)
}

// degree to radian
func dToR(deg float64) float64 {
    return deg * (math.Pi / 180.0)
}

// radian to degree
func rToD(rad float64) float64 {
    return rad * (180.0 / math.Pi)
}
