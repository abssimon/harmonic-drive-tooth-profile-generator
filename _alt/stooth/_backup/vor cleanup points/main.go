package main

import (
    "bytes"
    "fmt"
    "log"
    "math"
)

type STooth struct {
    C1   *Circle
    C2   *Circle
    C3   *Circle
    C4   *Circle
}

func (t STooth) rotate(a float64) {
    t.C1.Rotate(a)
    t.C2.Rotate(a)
    t.C3.Rotate(a)
    t.C4.Rotate(a)
}

func drawTooth(b *bytes.Buffer, e Ellipse, i float64) []Point {
    
    // definition
    scale := 100.0
    tipCenter := Point{2.333 * scale, 1.176 * scale}
    tipRadius := 0.866 * scale
    tipStop := 0.925 // 53 degree
    bottomCenter := Point{0.24 * scale, 1.453 * scale}
    bottomRadius := 1.67 * scale
    bottomStop := 0.78 // 45 degree

    tooth := STooth{
        &Circle{&Point{-tipCenter.X, tipCenter.Y}, bottomRadius, 0.0, 0.0},
        &Circle{&Point{bottomCenter.X, bottomCenter.Y}, tipRadius, 0.0, 0.0},
        &Circle{&Point{-bottomCenter.X, bottomCenter.Y}, tipRadius, 0.0, 0.0},
        &Circle{&Point{tipCenter.X, tipCenter.Y}, bottomRadius, 0.0, 0.0},
    }

    // symetrical, just need one tan
    _, tan := tooth.C3.InnerTangentWith(tooth.C4)
    tooth.C1.Start = math.Pi*2.0 - bottomStop
    tooth.C1.Stop = math.Pi*2.0 - tan
    tooth.C2.Start = math.Pi - tipStop
    tooth.C2.Stop = math.Pi - tan
    tooth.C3.Start = tan
    tooth.C3.Stop = tipStop
    tooth.C4.Start = math.Pi + tan
    tooth.C4.Stop = math.Pi + bottomStop
    
    // rotate correctly for ellipse
    tan = e.Tangent(i)  // is i correct?
    tooth.rotate(math.Pi + tan)
    
    points := []Point{}
    // append in right order
    c := tooth.C4.Coordinates()
    for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
        c[i], c[j] = c[j], c[i]
    }
    points = append(points, c...)
    
    c = tooth.C3.Coordinates()
    points = append(points, c...)
    
    c = tooth.C2.Coordinates()
    points = append(points, c...)
    
    c = tooth.C1.Coordinates()
    for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
        c[i], c[j] = c[j], c[i]
    }
    points = append(points, c...)
    
    return points
    
    /*
    for i := len(points) - 1; i >= 0; i-- {
        b.WriteString(fmt.Sprintf(format, points[i].X+p.X, points[i].Y+p.Y))  // ellipse, tooth position
        if first {
            b.WriteString("L ")
            first = false
        }
    }
    points = tooth.C3.Coordinates()
    for _, point := range points {
        b.WriteString(fmt.Sprintf(format, point.X+p.X, point.Y+p.Y))
    }
    points = tooth.C2.Coordinates()
    for _, point := range points {
        b.WriteString(fmt.Sprintf(format, point.X+p.X, point.Y+p.Y))
    }
    points = tooth.C1.Coordinates()
    for i := len(points) - 1; i >= 0; i-- {
        b.WriteString(fmt.Sprintf(format, points[i].X+p.X, points[i].Y+p.Y))
    }*/
}

func main() {

    b := Buffer()
    e := Ellipse{Point{600.0, 500.0}, 300.0, 337.0} // px, py, h, w // 391.0

    b.WriteString(fmt.Sprintf("<line x1=\"%.14f\" y1=\"%.14f\" x2=\"%.14f\" y2=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />", e.X-600, e.Y, e.X+600, e.Y))
    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n", e.X, e.Y, e.Width+35))
    printDot(b, e.Point, "black")

    //step := 10.0
    //rotate := 0.0
    gcounter := 0
    i := 0.0

    p := Point{}
    first := true
    p = e.PointAtAngle(i)

    b.WriteString(fmt.Sprintf("<path id=\"gear%d\" stroke=\"#1f1e1e\" d=\"M", gcounter))

    points := drawTooth(b, e, i)
    for _, point := range points {
    
        b.WriteString(fmt.Sprintf("%.8f %.8f ", point.X+p.X+e.X, point.Y+p.Y+e.Y)) // "%.14f %.14f "
        if first {
            b.WriteString("L ")
            first = false
        }
    }

    b.WriteString("\"/>\n")

/*
    printDot(b, Point{p.X + tooth.C1.Point.X + e.X, p.Y + tooth.C1.Point.Y + e.Y}, "black")
    printDot(b, Point{p.X + tooth.C2.Point.X + e.X, p.Y + tooth.C2.Point.Y + e.Y}, "black")
    printDot(b, Point{p.X + tooth.C3.Point.X + e.X, p.Y + tooth.C3.Point.Y + e.Y}, "black")
    printDot(b, Point{p.X + tooth.C4.Point.X + e.X, p.Y + tooth.C4.Point.Y + e.Y}, "black")
*/
    b.WriteString("</g></svg></body></html>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }
}

func Buffer() *bytes.Buffer {
    var b bytes.Buffer

    b.WriteString(`<!DOCTYPE html>
 <html>
 <body>
 <svg width="1200" height="1000" xmlns="http://www.w3.org/2000/svg">
 <g stroke-width="0.5" fill="none">
`)

    return &b
}
