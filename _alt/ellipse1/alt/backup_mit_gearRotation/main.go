package main

import (
    "bytes"
    "fmt"
    "log"
    "math"
)

func drawTooth(b *bytes.Buffer, e Ellipse, i float64, p Point, rotate float64, first bool) {
    tScale := 8.0

    tooth := newSTooth(
        Point{2.333 * tScale, 1.176 * tScale}, // tipCenter
        0.866*tScale,                          // tipRadius
        Point{0.173 * tScale, 0.223 * tScale}, // tipLimit - point in relation to center
        Point{0.24 * tScale, 1.453 * tScale},  // bottomCenter
        1.67*tScale,                           // bottomRadius
        Point{0.566 * tScale, 0.563 * tScale}, // bottomLimit
    )
    cp := p.CopyRotated(-rotate)
    tan := e.TangentByPoint(cp, i)
    tooth.rotate(180)
    tooth.rotate(rToD(tan) + rotate)

    points := tooth.C4.Coordinates()
    for i := len(points) - 1; i >= 0; i-- {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", p.X+points[i].X+e.X, p.Y+points[i].Y+e.Y)) // with L
        if first {
            b.WriteString("L ")
            first = false
        }
    }

    points = tooth.C3.Coordinates()
    for _, point := range points {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", p.X+point.X+e.X, p.Y+point.Y+e.Y))
    }

    points = tooth.C2.Coordinates()
    for _, point := range points {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", p.X+point.X+e.X, p.Y+point.Y+e.Y))
    }

    points = tooth.C1.Coordinates()
    for i := len(points) - 1; i >= 0; i-- {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", p.X+points[i].X+e.X, p.Y+points[i].Y+e.Y))
    }

}

func main() {

    var b bytes.Buffer

    
    b.WriteString("<svg width=\"1200\" height=\"1000\"  xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")

    e := Ellipse{Point{600.0, 500.0}, 400.0, 375.0} // px, py, h, w
    printDot(&b, e.Point, "")
    teeth := 100

    toothWidth := e.Circumference() / float64(teeth)
    gearRotation := 0.0

    // Minimize Track Error if possible

    //for rotate := 0.0; rotate <= 45.0; rotate += 8.0 {
    for rotate := 0.0; rotate <= 360.0; rotate += 90.0 { // warum gibt das

        p, pLast := Point{}, Point{}
        first := true
        total := 0.0
        counter := 0

        start := 360.0 - rotate
        for {
            p := e.PointByAngleRotated(start, rotate)
            if p.Y > 0.0 {
                start -= 0.0001
            } else {
                start += 0.0001
            }
            if math.Abs(p.Y) < 0.001 {
                break
            }
        }
        start += gearRotation

        for i := start; i <= 360.0; i += 0.001 {

            p = e.PointByAngleRotated(i, rotate)
            if first {
                printDot(&b, Point{p.X + e.X, p.Y + e.Y}, "")
                b.WriteString("<path d=\"M")

                drawTooth(&b, e, i, p, rotate, first)
                first = false
                pLast = p
                counter++
                continue
            }

            total += p.DistanceTo(&pLast) / 2
            pLast = p
            if total >= (float64(counter) * toothWidth) { // Trackerror hier wohl
                drawTooth(&b, e, i, p, rotate, first)
                counter++
            }

            if counter == teeth {
                break
            }
        }

        for i := 0.0; i < start; i += 0.001 {

            p = e.PointByAngleRotated(i, rotate)
            if first {
                printDot(&b, Point{p.X + e.X, p.Y + e.Y}, "")
                b.WriteString("<path d=\"M")

                drawTooth(&b, e, i, p, rotate, first)
                first = false
                pLast = p
                counter++
                continue
            }

            total += p.DistanceTo(&pLast) / 2
            pLast = p
            if total >= (float64(counter) * toothWidth) {
                drawTooth(&b, e, i, p, rotate, first)
                counter++
            }

            if counter == teeth {
                break
            }
        }
        b.WriteString("Z\"/>\n")

        // 100 Zaehne, durch 2 Diff = 50
        // 50 Umdrehungen Ellipse = 1 Rotation
        // 50*360  = 360
        // 1 Umdrehung Ellipse = 360/50 = 7,2 Grad
        // 0,5 (180) Umdrehung Ellipse = 3,6 Grad
        // 0,25 (90) Umdrehung Ellipse = 3,6 Grad
        teethDiff := 2.0
        ratio := float64(teeth) / teethDiff
        perTurn := 360.0 / ratio
        perDegree := degreePerEllipseTurn / 360.0

        gearRotation += perDegree * 90.0 // 90 see loop
    }
    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }
}
