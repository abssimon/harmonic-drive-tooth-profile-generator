package main

import (
    "bytes"
    "fmt"
    "log"
)

func main() {

    var b bytes.Buffer
    b.WriteString("<svg width=\"1200\" height=\"1000\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")
    

    e := Ellipse{Point{350.0, 280.0}, 250.0, 150.0 } // px, py, h, w
    printDot(&b, e.Point, "")
    
    teeth := 30 
    toothArc := e.Circumference() / float64(teeth)

    for rotate := 63.0; rotate <= 117.0; rotate += 9.0 {
        
        p, pLast := Point{}, Point{}
        first := true
        total := 0.0
        counter := 0
        
        // rotate := 45.0
        
        // anfangspunkt finden. rotiere nach oben 
        // bis der erste punkt vom alten übereinstimmt. - winkel vom mittelpunkt
        // dann entsprechend rotieren 354-0 0-354

        for i := 0.0; i <= 360.0; i += 0.0001 {
        
            p = e.PointByAngleRotated(i, rotate)
    
            if first {
                first = false
                pLast = p

                // cp := p.CopyRotated(-rotate)
                // tan := e.TangentByPoint(cp, i)
                
                b.WriteString(fmt.Sprintf("<circle id=\"%f-%d\" cx=\"%f\" cy=\"%f\" r=\"1\" fill=\"red\" />\n", // <text x=\"%f\" y=\"%f\" font-size=\"0.7em\" fill=\"black\" > t %.2f</text>
                    i, counter, p.X + e.X, p.Y + e.Y)) // , p.X + e.X + 5, p.Y + e.Y + 5, rToD(tan)))  
                counter++
                continue
            }

            total += p.DistanceTo(&pLast) / 2
            pLast = p
            if total >= (float64(counter) * toothArc) {

                // cp := p.CopyRotated(-rotate)
                // tan := e.TangentByPoint(cp, i)
                b.WriteString(fmt.Sprintf("<circle id=\"%f-%d\" cx=\"%f\" cy=\"%f\" r=\"1\"/>\n", // <text x=\"%f\" y=\"%f\" font-size=\"0.7em\" fill=\"black\" > t %.2f</text>
                    i, counter, p.X + e.X, p.Y + e.Y)) // , p.X + e.X + 5, p.Y + e.Y + 5, rToD(tan)))    
                counter++
            }

            if counter == teeth {
                fmt.Println("teeth complete")
                break
            }
        }
    }

    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

}
