package main

import (
    "bytes"
    "log"
    "math"
)

type Circle struct {
    *Point     // Center
    Radius     float64
    StartPoint *Point
    StopPoint  *Point
}

func (c *Circle) Rotate(a float64) {
    c.Point.Rotate(a)
    c.StartPoint.Rotate(a)
    c.StopPoint.Rotate(a)
}

func (c *Circle) PointInAngleSin(a float64) Point {
    rad := dToR(a)
    //x := math.Cos(rad) * c.Radius
    //y := math.Sin(rad) * c.Radius

/*
    float freq = 20;
    float amp = 20;

    dx = x + (radius + sin(angle * freq) * amp) * cos(angle);
    dy = y + (radius + sin(angle * freq) * amp) * sin(angle);
*/
    freq := 20.0
    amp := 20.0

    x := (c.Radius + math.Sin(rad * freq) * amp) * math.Cos(rad);
    y := (c.Radius + math.Sin(rad * freq) * amp) * math.Sin(rad);

    return Point{x + c.X, y + c.Y}
}

func main() {

    var b bytes.Buffer
    b.WriteString("<svg width=\"1200\" height=\"1000\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<g stroke-width=\"0.5\" fill=\"none\">")

    e := newCircle(500.0, 500.0, 200.0)
    printDot(&b, Point{500.0, 500.0}, "black")
        
    for i := 0.0; i < 360.0; i += 2 {
        p := e.PointInAngle(i)
        printDot(&b, p, "black")
    }

   
    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }
}


