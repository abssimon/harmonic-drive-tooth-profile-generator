package main

import (
    "bytes"
    "fmt"
    "log"
    "math"
)

// https://www.youtube.com/watch?v=ctxKuCeRqaE
func (e Ellipse) TangentOnePointAbsolute(buf *bytes.Buffer, p1 Point, rotate float64, debug bool) float64 {
    if e.Width > e.Height {
        log.Fatal("width higher than height")
    }
    hp := math.Sqrt(math.Pow(e.Height/2.0,2)-math.Pow(e.Width/2.0,2))
  
    p1.X -= e.X
    p1.Y -= e.Y
    p1.Rotate(-rotate)

    p2 := Point{0, -hp}
    p3 := Point{0, hp}
    
    if debug {
        printDot(buf, Point{e.X*2, +e.Y*2}, "center")
        printDot(buf, Point{p2.X+e.X*2, p2.Y+e.Y*2}, "p2")
        printDot(buf, Point{p3.X+e.X*2, p3.Y+e.Y*2}, "p3")
    }
    
    printDot(buf, Point{p1.X+e.X*2, p1.Y+e.Y*2}, "p1")

    
    /*
    p2 := Point{0, -hp}
    p3 := Point{0, hp}
    p2.Rotate(rotate)
    p3.Rotate(rotate)
    p2.X += e.X
    p3.X += e.X
    p2.Y += e.Y
    p3.Y += e.Y
*/
    if debug {
        printDot(buf, p1, "p1 test")
        //printDot(buf, p2, "")
    }

    //fmt.Println(p1, p2, p3)

    // https://www.triangle-calculator.com/de/?what=vc
    a := math.Hypot(p3.X-p1.X, p3.Y-p1.Y)
    b := math.Hypot(p2.X-p1.X, p2.Y-p1.Y)
    c := math.Hypot(p2.X-p3.X, p2.Y-p3.Y)
    //alpha := rToD(math.Acos((math.Pow(b, 2) + math.Pow(c, 2) - math.Pow(a, 2)) / (2 * b * c)))
    gamma := rToD(math.Acos((math.Pow(a, 2) + math.Pow(b, 2) - math.Pow(c, 2)) / (2 * a * b)))

    tan := (360.0 - 2.0*gamma) / 4.0
    
    return tan
}

func main() {

    var b bytes.Buffer
    b.WriteString("<svg width=\"1200\" height=\"1000\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")
    printDot(&b, Point{350.0, 280.0}, "")

    e := Ellipse{Point{350.0, 280.0}, 250.0, 150.0 } // px, py, h, w
    teeth := 30 
    toothArc := e.Circumference() / float64(teeth)

    //for rotate := 63.0; rotate <= 117.0; rotate += 63.0 {
        
        p, pLast := Point{}, Point{}
        first := true
        total := 0.0
        counter := 0
        
        rotate := 0.0

        for i := 0.0; i <= 360.0; i += 0.0001 {
        
            p = e.PointInAbsoluteAngleRotated(i, rotate)
            if first {
                first = false
                pLast = p

                // ref := e.PointInAngleRotated(i, rotate)
                // fmt.Println(ref)
                // Winkel unterschied zu p vom Mittelpunkt ausrechnen, dann weiss ich genau, in welchem Sektor ich mich befinde.

                tan := e.TangentOnePointAbsolute(&b, p, rotate, true)
                
                // printDot ausbauen
                b.WriteString(fmt.Sprintf("<circle id=\"%f-%d\" cx=\"%f\" cy=\"%f\"  r=\"1\" fill=\"red\" /><text x=\"%f\" y=\"%f\" font-size=\"0.7em\" fill=\"black\" > t %.2f</text>\n", i, counter, p.X, p.Y,p.X+5, p.Y+5, tan))  
                counter++
                continue
            }

            total += p.DistanceTo(&pLast) / 2
            pLast = p
            if total >= (float64(counter) * toothArc) {

                tan := e.TangentOnePointAbsolute(&b, p, rotate, false)
                b.WriteString(fmt.Sprintf("<circle id=\"%f-%d\" cx=\"%f\" cy=\"%f\"  r=\"1\"/><text x=\"%f\" y=\"%f\" font-size=\"0.7em\" fill=\"black\" > t %.2f</text>\n", i, counter, p.X, p.Y,p.X+5, p.Y+5, tan))    

                counter++
            }

            if counter == teeth {
                fmt.Println("teeth complete")
                break
            }
        }
    //}

    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

}
