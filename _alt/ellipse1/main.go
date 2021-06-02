package main

import (
    "bytes"
    "fmt"
    "log"
    "math"
)

// grad weg und atan2 (halbkreis mit rad)
// rotiere alle punkte zum schluss (-PointByAngleRotated - TangentByPoint)
// offset berechnen (wie wiki)
// distance unabhängi methode
// line weg

/*
func linspace(start, end float64, num int) []float64 {
    result := make([]float64, num)
    step := (end - start) / float64(num-1)
    for i := range result {
        result[i] = start + float64(i)*step
    }
    return result
}
*/

func main() {

    b := Buffer()
    e := Ellipse{Point{600.0, 500.0}, 300.0, 337.0} // px, py, h, w // 391.0
    teeth := 40 
    toothWidth := e.Circumference() / float64(teeth)
    
    b.WriteString(fmt.Sprintf("<line x1=\"%.14f\" y1=\"%.14f\" x2=\"%.14f\" y2=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />", e.X-600, e.Y, e.X+600, e.Y ))
    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n", e.X, e.Y, e.Width+35))
    printDot(b, e.Point, "black")
    
    circle := newCircle(e.X, e.Y, e.Width+35)
    for i := 0.0; i <= 360.0; i += 360.0 / (float64(teeth)+2.0) {        
        p := circle.PointInAngle(i)
        b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n", p.X, p.Y))
    }
    // 180 Grad = 3.5294117647058822
        
    // Hyperbel einzeichnen
    // https://rechneronline.de/trigonometrie/erweiterte-hyperbelfunktionen.php

    //step := 10.0
    rotate := 20.0
    gcounter := 0
    //for rotate := 0.0; rotate <= 180.0; rotate += step {
        

        p, pLast := Point{}, Point{}
        first := true
        total := 0.0
        tcounter := 0
        
        start := 0.0;
        start -= rotate * (3.5294117647058822 / 180.0) * 55
        
        // loop start-360
        for i := start; i <= 360.0; i += 0.001 {
        
            p = e.PointByAngleRotated(i, rotate)
            if first {
                //printDot(b, Point{p.X + e.X, p.Y + e.Y}, "#1f1e1e")
                
                b.WriteString(fmt.Sprintf("<path id=\"gear%d\" stroke=\"#1f1e1e\" d=\"M", gcounter))
                gcounter++
                
                drawTooth(b, e, i, p, rotate, first)
                tcounter++
                
                first = false
                pLast = p
                
                continue
            }

            // set teeth by distance
            total += p.DistanceTo(&pLast) / 2.0
            pLast = p
            
            if total >= (float64(tcounter) * toothWidth) {
                drawTooth(b, e, i, p, rotate, first)
                tcounter++
            }

            if tcounter == teeth {
                break
            }
        }
        

        // start-360
        for i := 0.0; i < start; i += 0.001 {

            p = e.PointByAngleRotated(i, rotate)
            if first {
                //printDot(b, Point{p.X + e.X, p.Y + e.Y}, "#1f1e1e")
                
                b.WriteString(fmt.Sprintf("<path id=\"gear%d\" stroke=\"#1f1e1e\" d=\"M", gcounter))
                gcounter++
                
                drawTooth(b, e, i, p, rotate, first)
                tcounter++
                
                first = false
                pLast = p
                
                continue
            }

            // set teeth by distance
            total += p.DistanceTo(&pLast) / 2.0
            pLast = p
            
            if total >= (float64(tcounter) * toothWidth) {
                drawTooth(b, e, i, p, rotate, first)
                tcounter++
            }

            if tcounter == teeth {
                break
            }
        }

        b.WriteString("Z\"/>\n")
        
    //}

    b.WriteString(`
    </g></svg>
    </body>
    </html>`)

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }
}

func drawTooth(b *bytes.Buffer, e Ellipse, i float64, p Point, rotate float64, first bool) {
    tScale := 16.0 
    format := "%.8f %.8f " // "%.14f %.14f "

    tooth := newSTooth(
        // tipCenter
        Point{2.333 * tScale, 1.176 * tScale}, 
        // tipRadius
        0.866*tScale,       
        // tipLimit - point in relation to center
        Point{0.173 * tScale, 0.223 * tScale}, 
        // bottomCenter
        Point{0.24 * tScale, 1.453 * tScale},  
        // bottomRadius
        1.67*tScale,  
        // bottomLimit
        Point{0.566 * tScale, 0.563 * tScale}, 
    )

    // tan calculation via a non rotated ellipse
    cp := p.CopyRotated(-rotate)
    tan := e.TangentByPoint(cp, i)
    tooth.rotate(180 + rToD(tan) + rotate)

    points := tooth.C4.Coordinates()
    for i := len(points) - 1; i >= 0; i-- {
        b.WriteString(fmt.Sprintf(format, p.X+points[i].X+e.X, p.Y+points[i].Y+e.Y))
        // first pair with L
        if first {
            b.WriteString("L ")
            first = false
        }
    }

    points = tooth.C3.Coordinates()
    for _, point := range points {
        b.WriteString(fmt.Sprintf(format, p.X+point.X+e.X, p.Y+point.Y+e.Y))
    }

    points = tooth.C2.Coordinates()
    for _, point := range points {
        b.WriteString(fmt.Sprintf(format, p.X+point.X+e.X, p.Y+point.Y+e.Y))
    }

    points = tooth.C1.Coordinates()
    for i := len(points) - 1; i >= 0; i-- {
        b.WriteString(fmt.Sprintf(format, p.X+points[i].X+e.X, p.Y+points[i].Y+e.Y))
    }
}

func Buffer() *bytes.Buffer {
     var b bytes.Buffer
     
     b.WriteString(`<!DOCTYPE html>
 <html>
 <head>
   <style>
   /*
       path {
         display: none
       }
   */
   </style>
   <script src="https://code.jquery.com/jquery-3.5.0.js"></script>
 <script>

   /* 
     $(document).ready(function() {
  
        var counter = 0;
        setInterval(function(){
        
            document.getElementById("gear"+counter).style.display = "block";
            if (counter === 0) {
                document.getElementById("gear18").style.display = "none";
            } else {
                document.getElementById("gear"+(counter-1)).style.display = "none";
            }

            counter++
            if (counter === 19) {
                counter=0
            }
            
        }, 1500);      
     });
   */  
 </script>
 </head>
 <body>
 <button id="button" onclick="hideSVG('gear0')">1</button>
 <svg width="1200" height="1000" xmlns="http://www.w3.org/2000/svg">
 <g stroke-width="0.5" fill="none">
`)

    return &b
}