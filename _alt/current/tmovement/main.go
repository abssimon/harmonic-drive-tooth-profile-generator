package main

import (
    "bytes"
    "fmt"
    "log"
    //"math"
)

func main() {

    var b bytes.Buffer
    
    b.WriteString("<svg width=\"1200\" height=\"1000\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<g stroke-width=\"0.5\" fill=\"none\">")

    e := Ellipse{Point{600.0, 500.0}, 303.0, 400.0} // px, py, h, w 
    teeth := 400
    toothWidth := e.Circumference() / float64(teeth)
    
    b.WriteString(fmt.Sprintf("<line x1=\"%.14f\" y1=\"%.14f\" x2=\"%.14f\" y2=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />", e.X-600, e.Y, e.X+600, e.Y ))
    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n", e.X, e.Y, e.Width+40))
    printDot(&b, e.Point, "black")
    
    r1, r2 := Point{}, Point{}
    
    c := newCircle(600.0, 500.0, 400.0)
    for i := 0.0; i <= 360.0; i += (360.0 / 102.0) {
        p := c.PointInAngle(i)
        printDot(&b, Point{p.X + e.X, p.Y + e.Y}, "blue")

        if i == 0.0 {
            r1 = p
        }
        if i > 3.0 && i < 4.0 {
            r2 = p
        }
    }  
    

    step := 1.0 // 181.0 
    first1 := true 
    pLast1 := Point{}
    for rotate := 0.0; rotate < 15.0; rotate += step { // TODO bei <= das = raus

        p, pLast := Point{}, Point{}
        first := true
        total := 0.0
        counter := 0
        start := 0.0
        
   
        p = e.PointByAngleRotated(0, rotate)
        if first1 {
            pLast1 = p
            first1 = false            
        }
        if !first1 {
            // first1 resetten wenn p2 erreicht
            fmt.Printf("%f %f %f %f\n", rotate, p.DistanceTo(&pLast1), r1.DistanceTo(&r2), HyperbelScale(p.DistanceTo(&pLast1), r1.DistanceTo(&r2)))
        }
        

        // draw 360 degrees gear from any rotation
        for i := start; i <= 360.0; i += 0.001 {

            p = e.PointByAngleRotated(i, rotate)
            if first {
                printDot(&b, Point{p.X + e.X, p.Y + e.Y}, "red") // 
                first = false
                pLast = p
                counter++
                continue
            }

            // set teeth by distance
            total += p.DistanceTo(&pLast) / 2
            pLast = p
            if total >= (float64(counter) * toothWidth) {
                printDot(&b, Point{p.X + e.X, p.Y + e.Y}, "#7d7c7c")
                counter++
            }

            if counter == teeth {
                break
            }
        }

        for i := 0.0; i < start; i += 0.001 {

            p = e.PointByAngleRotated(i, rotate)
            if first {
                printDot(&b, Point{p.X + e.X, p.Y + e.Y}, "yellow")
                first = false
                pLast = p
                counter++
                continue
            }

            total += p.DistanceTo(&pLast) / 2
            pLast = p
            if total >= (float64(counter) * toothWidth) {
                printDot(&b, Point{p.X + e.X, p.Y + e.Y}, "#7d7c7c")
                counter++
            }

            if counter == teeth {
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

