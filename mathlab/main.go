package main

import (
    "math"
    "fmt"
    "strings"
    "os"
)

type Point struct {
    X float64
    Y float64
}

func (p *Point) Rotate(a float64) {
    x, y := p.X, p.Y
    p.X = x*math.Cos(a) - y*math.Sin(a)
    p.Y = x*math.Sin(a) + y*math.Cos(a)
}

func Ellipse(diameterH float64, diameterV float64, nPoints int) []Point {
    r := make([]Point, nPoints)
    p := LinSpace(0.0,2.0*math.Pi, nPoints)
    for i, v := range p {
        r[i] = Point{math.Cos(v) * diameterH / 2.0, math.Sin(v) * diameterV / 2.0}
    }
    return r
}

// https://commons.wikimedia.org/wiki/File:HarmonicDriveAni.gif
func equidistantEllipse(diameterH float64, diameterV float64, nPoints float64, offset float64) []Point {

    // some ellipse sample points for prediction
    pathXY := Ellipse(diameterH, diameterV, 1000) 

    // and the changes (tiny) of the point distance
    sum := make([]float64, len(pathXY))
    for i, v := range Diff(pathXY) {
        sum[i+1] = math.Hypot(v.X, v.Y) 
    }
    cumLen := CumSum(sum)
    
    // teeth locations 0-1
    circumference := cumLen[len(cumLen)-1]
    loc := LinSpace(0.0, 1.0, int(nPoints)+1)
    loc = loc[:len(loc)-1]
    for i, v := range loc {
        loc[i] = math.Mod(v+offset, 1)*circumference  // 0-1 always for curcumference, with offset
    }

    // x y seperatly
    pathX := make([]float64, len(pathXY))
    pathY := make([]float64, len(pathXY))
    for i, v := range pathXY {
        pathX[i] = v.X
        pathY[i] = v.Y
    }
    
    /*
    for _, v := range loc {
        s := fmt.Sprintf("%.4f", v)
        s = strings.Replace(s, ".", ",", -1)
        fmt.Println(s)
    }
    os.Exit(500)
    */
    
    // cumLen: 0.0000, 0.0119, 0.0238, 0.0357, 0.0477, 0.0596, ... 12.5577 // 1000
    // pathX: 2.1000, 2.1000, 2.0998, 2.0996, 2.0993, 2.0990 ... -2.1000, ... 2.1000 // 1000
    // loc: 0.0000, 0.3139, 0.6279, 0.9418, 1.2558, 1.5697, ... 12.2438 // 42

    // predict X for loc
    x, err := interp1(cumLen, pathX, loc)
    if err != nil {
        panic(err)
    }
    y, err := interp1(cumLen, pathY, loc)
    if err != nil {
        panic(err)
    }

    // merge
    r := make([]Point, len(x))
    for i, v := range x {
        r[i] = Point{v, y[i]}
    }

    return r
}

func main() {

    rigidTheets := 102.0 // 52 looks good with 0.9122 ????
    flexTheets := rigidTheets - 2.0

    nFrames := 100;
    frameAngles := LinSpace(0.0, -math.Pi, nFrames+1)
    frameAngles = frameAngles[:len(frameAngles)-1]

    // Job: fuer zahne richtig H V finden, so das Wellen exakt spitz sind
    diameterH := 4.2  
    diameterV := 4.035 // 4.

    total := []Point{}
    for _, angleWaveGen := range frameAngles {

        angleFlexTeeth := angleWaveGen * (flexTheets - rigidTheets) / flexTheets

        offset := (-angleWaveGen + angleFlexTeeth) / 2.0 / math.Pi // on circumference from 0-1

        r := equidistantEllipse(diameterH, diameterV, flexTheets, offset)

        for i, _ := range r {
            r[i].Rotate(angleWaveGen)
        }

        total = append(total, r...)
    }
    
    svg(total)
}
