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

type Circle struct {
    Point
    Radius float64
}

func (c Circle) PointInAngle(a float64) Point {
    x := math.Cos(dToR(a))*c.Radius + c.X
    y := math.Sin(dToR(a))*c.Radius + c.Y

    return Point{x, y}
}

func newCircle(x, y, r float64) *Circle {
    return &Circle{Point{x, y}, r}
}

type Ellipse struct {
    Point
    Width  float64
    Height float64
}

func (e Ellipse) PointInAngle(a float64) Point {
    x := e.Width*math.Cos(dToR(a)) + e.X
    y := e.Height*math.Sin(dToR(a)) + e.Y
    return Point{x, y}
}

func main() {

    var b bytes.Buffer
    b.WriteString("<svg width=\"700\" height=\"1000\" xmlns=\"http://www.w3.org/2000/svg\">")

    elipse := Ellipse{
        Point{350.0, 450.0},
        200.0,
        250.0,
    }


    for i := 0.0; i <= 360; i += 1.0 {
        p := elipse.PointInAngle(i)
        b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", p.X, p.Y))
    }
    
    elipse.Width /= 1.04
    elipse.Height /= 1.04
    for i := 0.0; i <= 360; i += 1.0 {
        p := elipse.PointInAngle(i)
        b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", p.X, p.Y))
    }

    elipse.Width /= 1.04
    elipse.Height /= 1.04
    for i := 0.0; i <= 360; i += 1.0 {
        p := elipse.PointInAngle(i)
        b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", p.X, p.Y))
    }
    

    printDot(&b, Point{350.0, 450.0})
    
    b.WriteString("</svg>")


    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

}
func printDot(b *bytes.Buffer, p Point) {
    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"1\" />\n", p.X, p.Y))
}