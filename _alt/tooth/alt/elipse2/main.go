package main

import (
    "bytes"
    "fmt"
    "log"
    "math"
    "os"
)

// degree to radian
func dToR(deg float64) float64 {
    return deg * (math.Pi / 180.0)
}

// radian to degree
func rToD(rad float64) float64 {
    return rad * (180.0 / math.Pi)
}

func writeSvg(data []byte) error {
    f, err := os.OpenFile("test.svg", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.Write(data)
    if err != nil {
        return err
    }

    return nil
}

type Point struct {
    X, Y float64
}

type Ellipse struct {
    Point
    Width  float64
    Height float64
}

// http://jsfiddle.net/hkejm20t/88/
func (e Ellipse) draw() {
    steps := 32 // 64
    accuracy := 2
    
    // Ramanujan's approximation
    circumference := math.Pi * (3.0 * (e.X + e.Y) - math.Sqrt((3 * e.X + e.Y) * (e.X + 3 * e.Y)))
    arcLength := circumference / float64(steps)
    d := 0.0
    fmt.Println(circumference, arcLength, d)
    
    oldX := e.X 
    oldY := 0.0
    oldDelta := 0.0
    oldTheta := 0.0

    integrationSteps := steps * int(math.Pow(2, float64(accuracy + 1)))
    
    for i := 0; i < integrationSteps; i++ {
        theta := float64(i / integrationSteps) * (math.Pi * 2.0)
        x := math.Cos(theta) * e.X
        y := math.Sin(theta) * e.Y
        d += math.Sqrt(math.Pow(x - oldX, 2) + math.Pow(y - oldY, 2))
        delta := d - arcLength
        
        if (delta >= 0) {
            t := math.Abs(oldDelta) / (delta - oldDelta)
            theta = oldTheta + (theta - oldTheta) * t
            newX := math.Cos(theta) * e.X
            newY := math.Sin(theta) * e.Y
            bearing := rToD(math.Atan2(newY, newX))
            distance := math.Sqrt(math.Pow(newX, 2) + math.Pow(newY, 2))
            
            // 0.5638960818133139, 0.1365585551247011, 0.49534518068724126, 0.2382354122114775, 25.68514780020028, 0.5496571291829556
            // 0.2256758665552539, 0.2755195377652361, 0.4811419944285728, 0.4760821039136698, 44.697137095023116, 0.6768691073387609
            fmt.Println(t, theta, newX, newY, bearing, distance)
        
            d -= arcLength
        }
        oldX = x
        oldY = y
        oldTheta = theta
        oldDelta = delta
    }
    
    
    
}

func main() {

    var b bytes.Buffer
    b.WriteString("<svg width=\"700\" height=\"700\" xmlns=\"http://www.w3.org/2000/svg\">")

    ellipse := Ellipse{
        Point{0.5, 1.75},
        120.0,
        250.0,
    }
    ellipse.draw()

    
    b.WriteString("</svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

}
func printDot(b *bytes.Buffer, p Point) {
    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"1\" />\n", p.X, p.Y))
}