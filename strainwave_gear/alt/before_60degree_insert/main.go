package main

import (
    "math"
    "fmt"
)

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
    fmt.Println("entspricht kreis durchmesser ", circumference / math.Pi)

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

    rigidTheets := 102.0
    flexTheets := rigidTheets - 2.0

    nFrames := 1 // 100
    frameAngles := LinSpace(0.0, -math.Pi, nFrames+1)
    frameAngles = frameAngles[:len(frameAngles)-1]

    diameterH := 4.2
    diameterV := 4.035 // calculate?

    // sample points for prediction
    e := Ellipse{Point{0, 0}, diameterH, diameterV}
    ellipseCoordinates := e.Coordinates(1000)

    gears := [][]Point{}
    for _, angleWaveGen := range frameAngles {

        // position on ellipse
        angleFlexTeeth := angleWaveGen * (flexTheets - rigidTheets) / flexTheets // * -0,02
        offset := (-angleWaveGen + angleFlexTeeth) / 2.0 / math.Pi               
        r := toothPosition(ellipseCoordinates, flexTheets, offset)

        // add teeth
        gear := []Point{}
        for _, p := range r {
            sp := Point{p.X * 200.0, p.Y * 200.0}
            g, _ := FlexTooth(sp, math.Pi+e.TangentByPoint(sp))
            gear = append(gear, g...)
        }

        for i := range gear {
            gear[i].Rotate(angleWaveGen) 
        }

        gears = append(gears, gear)
    }

//gears = [][]Point{}
gear := []Point{}
points := []Point{}

    // draw rigid gear
    c := &Circle{&Point{0.0, 0.0}, diameterH / 2, 0.0, math.Pi * 2.0}
    angle := LinSpace(0.0, math.Pi*2.0, int(rigidTheets)+1)
    angle = angle[:len(angle)-1]
      
    for _, a := range angle {
        cp := c.PointInAngle(a)
        
        //gear = append(gear, RigidTooth(Point{cp.X * 200.0, cp.Y * 200.0}, -math.Pi/2.0+a)...)
         gearPoints := RigidTooth(Point{cp.X * 200.0, cp.Y * 200.0}, -math.Pi/2.0+a)
         gear = append(gear, gearPoints...)
         points = append(points, gearPoints...)
    }
    gears = append(gears, gear)
/*
    // rotate all to have one flanc vertical (manually - look "winkel_linie")
    for i := range gears {
        for j := range gears[i] {
            gears[i][j].Rotate(-0.017127357071603733) 
        }
    }
    for i := range points {
        points[i].Rotate(-0.017127357071603733) 
    }


    // rotate all to have one flanc vertical (manually - look "winkel_linie")
    for i := range gears {
        for j := range gears[i] {
            gears[i][j].Rotate(-0.02734514180952) 
        }
    }   
    for i := range points {
        points[i].Rotate(-0.02734514180952) 
    }
*/

/*
    // Production
 
    // gears = [][]Point{}
   

    // draw round ellipse 
    re := &Circle{&Point{0.0, 0.0}, 4.117906474475928 / 2, 0.0, math.Pi * 2.0} // diameter calculation see "entspricht kreis durchmesser "
    angle = LinSpace(0.0, math.Pi*2.0, int(flexTheets)+1)
    
    
// for i, a := range angle {
//    fmt.Println(i, rToD(a))
//}
    
    angle = angle[:len(angle)-1]
    
    
    for _, a := range angle {
        cp := re.PointInAngle(a)
        gearPoints, prodPoints := FlexTooth(Point{cp.X * 200.0, cp.Y * 200.0}, -math.Pi/2.0+a)
        gear = append(gear, gearPoints...) 
        points = append(points, prodPoints...)
    }
    gears = append(gears, gear)
        

    // rotate all to have one flanc vertical (manually - look "winkel_linie")
    for i := range gears {
        for j := range gears[i] {
            gears[i][j].Rotate(0.02476054497625) 
        }
    }
    for i := range points {
        points[i].Rotate(0.02476054497625) 
    }
 /*
    // secound flanc
    for i := range gears {
        for j := range gears[i] {
            gears[i][j].Rotate(-0.04952108995251) 
        }
    }
    for i := range points {
        points[i].Rotate(-0.04952108995251) 
    }    
*/
    svg(gears, points)
}
