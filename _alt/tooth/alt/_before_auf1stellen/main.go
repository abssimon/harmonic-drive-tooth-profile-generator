package main

import (
    "bytes"
    "fmt"
    "log"
)

const scaleFactor = 10000000000000 // 1e9

func scaleUp(value float64) int {
    return int(value * scaleFactor)
}

func main() {

    // Modulo! - Abstand zwischen den Zaehnen gross und klein muss gleich bleiben!!!
    // Finde Umfang ... damit starten
    //  Dreisatz, Punkte finden die im gewissen Masse den gleichen Abstand auf der Ellipse haben... (ellipse eckig machen)

    var b bytes.Buffer
    b.WriteString("<svg width=\"1500\" height=\"2000\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<line x1=\"0\" y1=\"47\" x2=\"1500\" y2=\"47\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")

    // Circles
    elipse := newEllipse(750.0, 900.0, 470.0*1.2, 490.0*1.2)
    tooths := 120
    scale := 30.0

    b.WriteString("<path d=\"M")
    first := true
    for i := 360.0; i >= 0.0; i -= (360.0 / float64(tooths)) { // ERROR, nicht die Gradzahl, die variiert bei einer Elypes, Abstaende unterschiedlich lang

        tooth := newSTooth(
            Point{700 / scale, 353 / scale}, // tipCenter
            260.0/scale,                     // tipRadius
            Point{52 / scale, 67 / scale},   // tipLimit - point in relation to center
            Point{72 / scale, 436 / scale},  // bottomCenter
            501.0/scale,                     // bottomRadius
            Point{170 / scale, 169 / scale}, // bottomLimit
        )
        tooth.rotate(i - 90)  // ERROR, der Zahn muss auch auf der Oberflaeche gewinkelt sein

        p := elipse.PointInAngle(i) // fangt rechts an...
        moveX := p.X
        moveY := p.Y

        // tooth.MoveX = 780.0
        // tooth.MoveY = 147.0
        // tooth.Coordinates()

        points := tooth.C1.Coordinates()
        for _, point := range points {
            b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX+point.X, moveY+point.Y))
            if first {
                b.WriteString("L ")
                first = false
            }
        }

        points = tooth.C2.Coordinates()
        for i := len(points) - 1; i >= 0; i-- {
            b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX+points[i].X, moveY+points[i].Y))
        }

        points = tooth.C3.Coordinates()
        for i := len(points) - 1; i >= 0; i-- {
            b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX+points[i].X, moveY+points[i].Y))
        }

        points = tooth.C4.Coordinates()
        for _, point := range points {
            b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX+point.X, moveY+point.Y))
        }
    }
    b.WriteString("\"/>\n")

    circle := newCircle(750.0, 900.0, 490.0*1.2)

    for i := 0.0; i <= 360.0; i += (360.0 / 122) {
        p := circle.PointInAngle(i)
        b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", p.X, p.Y))
    }

    // Debug
    /*
       printDot(&b, tooth.Tan2.p1, moveX, moveY)
       printDot(&b, tooth.Tan2.p2, moveX, moveY)

       printDot(&b, tooth.C1.Point, moveX, moveY)
       printDot(&b, tooth.C2.Point, moveX, moveY)
       printDot(&b, tooth.C3.Point, moveX, moveY)
       printDot(&b, tooth.C4.Point, moveX, moveY)
    */
    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

}
