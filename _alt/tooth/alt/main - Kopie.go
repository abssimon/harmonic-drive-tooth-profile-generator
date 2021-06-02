package main

import (
    "bytes"
    "fmt"
    "log"
    "math"
)

type Point struct {
    X, Y float64
}

func (p Point) DistanceTo(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Line struct {
    p1 Point
    p2 Point
}

type Circle struct {
    Point  // kann man auch als ganzes ansprechen, oder Center
    Radius float64
}

func (c Circle) TangentTo(c2 Circle) (Line, Line) {
    hp := c.Point.DistanceTo(c2.Point)
    v := c2.Radius / c.Radius
    len1 := hp / 100 * (100.0 / (v + 1.0) * v)

    // angles
    a1 := 180.0 - (rToD(math.Asin(c2.Radius/len1)) + 90.0)
    a2 := rToD(math.Asin((c.Y - c2.Y) / hp))
    a3 := 90 - (a2 + a1)
    a4 := a1 - a2

    sa3 := math.Sin(dToR(a3)) * c2.Radius // todo names
    sb3 := math.Cos(dToR(a3)) * c2.Radius
    sa5 := math.Sin(dToR(a4)) * c2.Radius
    sb5 := math.Cos(dToR(a4)) * c2.Radius

    osa5 := math.Sin(dToR(a3)) * c.Radius
    osb5 := math.Cos(dToR(a3)) * c.Radius
    osa7 := math.Sin(dToR(a4)) * c.Radius
    osb7 := math.Cos(dToR(a4)) * c.Radius

    return Line{Point{c2.X - sa3, c2.Y + sb3}, Point{c.X + osa5, c.Y - osb5}}, Line{Point{c2.X - sb5, c2.Y - sa5}, Point{c.X + osb7, c.Y + osa7}}
}

// Note: more than 180 degree not always possible and rotation is left to right
// See:      90
//        180 + 0
//           -90
func (c Circle) CoordinatesBetween(start Point, stop Point) []Point {
    startAngle := rToD(math.Atan2(start.Y-c.Y, start.X-c.X))
    stopAngle := rToD(math.Atan2(stop.Y-c.Y, stop.X-c.X))
    
    fmt.Println("startAngle", startAngle)
    fmt.Println("stopAngle", stopAngle)
    
    points := []Point{}
    print := false;
    out:
    for i := 0; i < 2; i++ {
        for i := 0; i <= 180; i++ {
            if !print && startAngle >= 0.0 && startAngle <= float64(i) {               
                points = append(points, c.PointByAngle(startAngle))               
                print = true
            }
            if print && stopAngle >= 0.0 && stopAngle <= float64(i) {
                points = append(points, c.PointByAngle(stopAngle))
                break out
            }
            if print {
                points = append(points, c.PointByAngle(float64(i))) 
            }
        }
        for i := -179; i < 0; i++ {
            if !print && startAngle < 0.0 && startAngle <= float64(i) {
                points = append(points, c.PointByAngle(startAngle))
                print = true
            }
            if print && stopAngle < 0.0 && stopAngle <= float64(i) {
                points = append(points, c.PointByAngle(stopAngle))
                break out
            }
            if print {
                points = append(points, c.PointByAngle(float64(i))) 
            }
        }
    }
 
    return points
}

func (c Circle) PointByAngle(a float64) Point {
    x := math.Cos(dToR(a)) * c.Radius
    y := math.Sin(dToR(a)) * c.Radius  
    //fmt.Println(a, x + c.X, y + c.Y)
    return Point{x + c.X, y + c.Y}
}

func main() {

    var b bytes.Buffer
    b.WriteString("<svg width=\"700\" height=\"700\" xmlns=\"http://www.w3.org/2000/svg\">\n")
    b.WriteString("<line x1=\"0\" y1=\"300\" x2=\"700\" y2=\"300\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n")
    b.WriteString("<g transform=\"translate(0,700)\">\n<g transform=\"scale(1,-1)\">\n<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">\n")

    circle1 := Circle{
        Point:  Point{160.0, 200.0},
        Radius: 15.0,
    }
    circle2 := Circle{
        Point:  Point{230.0, 140.0},
        Radius: 50.0,
    }
    
    _, l2 := circle1.TangentTo(circle2)
    
    b.WriteString("<path d=\"M")
    points := circle1.CoordinatesBetween(l2.p2, Point{0, 100}) 
    first := true
    for _, point := range points {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", point.X, point.Y))
        if first {
            first = false
            b.WriteString("L")
        }        
    }
    
    points = circle2.CoordinatesBetween(l2.p1, Point{200, 40}) 
    for _, point := range points {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", point.X, point.Y))
    }

    b.WriteString("\"/>\n</g>\n</g>\n</g>\n</svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

}