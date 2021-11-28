package main

import (
    "math"
)
// hello

// Teeth mashing is pretty hard. The only working solution is found here
// https://commons.wikimedia.org/wiki/File:HarmonicDriveAni.gif
func toothPosition(ellipseCoordinates []Point, nPoints float64, offset float64) []Point {

    // changes (tiny) of coordinates
    sum := make([]float64, len(ellipseCoordinates))
    for i, v := range Diff(ellipseCoordinates) {
        sum[i+1] = math.Hypot(v.X, v.Y)
    }
    cumLen := CumSum(sum)
    circumference := cumLen[len(cumLen)-1]
    //fmt.Println("diamter for ellipse ", circumference / math.Pi)

    // location on curcumference, with offset
    pos := LinSpace(0.0, 1.0, int(nPoints)+1)
    pos = pos[:len(pos)-1]
    for i, v := range pos {
        pos[i] = math.Mod(v+offset, 1) * circumference
    }

    coordX := make([]float64, len(ellipseCoordinates))
    coordY := make([]float64, len(ellipseCoordinates))
    for i, v := range ellipseCoordinates {
        coordX[i] = v.X
        coordY[i] = v.Y
    }

    // predict X for pos
    x, err := interp1(cumLen, coordX, pos)
    if err != nil {
        panic(err)
    }
    y, err := interp1(cumLen, coordY, pos)
    if err != nil {
        panic(err)
    }

    r := make([]Point, len(x))
    for i, v := range x {
        r[i] = Point{v, y[i]}
    }

    return r
}

func main() {

    rigidTheets := 102.0 // 1:30: 62.0
    flexTheets := rigidTheets - 2.0

    nFrames :=  1 // for animation 100
    frameAngles := LinSpace(0.0, -math.Pi, nFrames+1)
    frameAngles = frameAngles[:len(frameAngles)-1]
    
    // draw flex gear
    diameterH := 4.2
    diameterV := 4.035 // 1:30: 3.94  

    // sample points for prediction
    e := Ellipse{Point{0, 0}, diameterH, diameterV}
    ellipseCoordinates := e.Coordinates(1000)

    gears := []Gear{}
    gear := Gear{}
    for _, angleWaveGen := range frameAngles {
    
        // position prediction on ellipse
        angleFlexTeeth := angleWaveGen * (flexTheets - rigidTheets) / flexTheets // * -0,02
        offset := (-angleWaveGen + angleFlexTeeth) / 2.0 / math.Pi               
        pos := toothPosition(ellipseCoordinates, flexTheets, offset)

        gear = Gear{}
        for _, p := range pos {
            gear.Tooths = append(gear.Tooths, NewTooth(math.Pi+e.TangentByPoint(p), 1.56, p))
        }
        gear.Angle = angleWaveGen
        gears = append(gears, gear)
    }
    
    // draw rigid gear
    c := &Circle{&Point{0.0, 0.0}, diameterH / 2, 0.0, math.Pi * 2.0}
    angle := LinSpace(0.0, math.Pi*2.0, int(rigidTheets)+1)
    angle = angle[:len(angle)-1]
    
    gear = Gear{}
    for _, a := range angle {
        gear.Tooths = append(gear.Tooths, NewTooth(-math.Pi/2.0+a, 0.95, c.PointInAngle(a)))
    }
    gears = append(gears, gear)
   

    // draw undeformed flexgear for production
    gears = []Gear{}
    gear = Gear{}
   
    c2 := &Circle{
        &Point{0.0, 0.0}, 
        4.117906474475928 / 2, // diameter for "round ellipse" see toothPosition()->circumference
        0.0, 
        math.Pi * 2.0,
    } 
    angle = LinSpace(0.0, math.Pi*2.0, int(flexTheets)+1)
    angle = angle[:len(angle)-1]
        
    for _, a := range angle {
        gear.Tooths = append(gear.Tooths, NewTooth(-math.Pi/2.0+a, 1.56, c2.PointInAngle(a)))
    }
    gears = append(gears, gear)
    
        
    svg(gears)
}
