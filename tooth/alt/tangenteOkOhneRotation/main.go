package main

import (
    "bytes"
    "fmt"
    "log"
    //"math"
)

// https://math.stackexchange.com/questions/4086824/how-to-find-the-polar-coordinate-angle-of-the-tangent-of-any-point-on-an-ellipse/4086900#4086900
    // https://math.stackexchange.com/questions/3140614/what-is-the-slope-of-tangent-line-to-a-rotated-ellipse-at-a-specific-point
    // https://math.stackexchange.com/questions/171936/how-do-i-get-a-tangent-to-a-rotated-ellipse-in-a-given-point?rq=1

func main() {

    var b bytes.Buffer
    b.WriteString("<svg width=\"700\" height=\"700\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")

    ellipse := Ellipse{
        Point{350.0, 280.0},
        250.0,
        210.0,
    }
    rotate := 30.0

    // TODO: Zaehne in Rotation versetzen
    first := true
    for i := 0.0; i <= 360.0; i += 0.1 { // .001 is a +-0.00218px precision
        
        // tangente works !!!
        // p := ellipse.PointInAngle(i)//, rotate)
        // fmt.Println(rToD(math.Atan2(ellipse.Height * math.Cos(dToR(i)), -ellipse.Width * math.Sin(dToR(i)))))


        p := ellipse.PointInAbsoluteAngleRotated(i, rotate)
        
        //d := math.Atan2(-ellipse.Width * math.Sin(dToR(rotate)), ellipse.Height * math.Cos(dToR(rotate)))
        //fmt.Println(i, rToD(math.Atan2(ellipse.Height * math.Cos(dToR(i)+d), -ellipse.Width * math.Sin(dToR(i)+d))))

        if first {
            first = false
            b.WriteString(fmt.Sprintf("<circle id=\"%f\" cx=\"%f\" cy=\"%f\"  r=\"5\" fill=\"red\" />\n", i, p.X, p.Y))
            continue
        }
        b.WriteString(fmt.Sprintf("<circle id=\"%f\" cx=\"%.14f\" cy=\"%.14f\" r=\"5\" />\n", i, p.X, p.Y))
    }

    printDot(&b, Point{350.0, 280.0})
    
    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

}
