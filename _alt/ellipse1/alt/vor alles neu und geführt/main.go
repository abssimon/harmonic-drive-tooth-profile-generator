package main

import (
    "bytes"
    "fmt"
    "log"
    "math"
)

func main() {

    b := Buffer()
    e := Ellipse{Point{600.0, 500.0}, 303.0, 337.0} // px, py, h, w // 391.0
    
    b.WriteString(fmt.Sprintf("<line x1=\"%.14f\" y1=\"%.14f\" x2=\"%.14f\" y2=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />", e.X-600, e.Y, e.X+600, e.Y ))
    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n", e.X, e.Y, e.Width+40))
    printDot(b, e.Point, "black")
    
    teeth := 40 
    teethRigid := 42 
    toothWidth := e.Circumference() / float64(teeth)
    gearRotation := 0.0

    
    step := 30.0
    rotate := 0.0
    //step := 30.0 // 181.0  
    //for rotate := 0.0; rotate <= 180.0; rotate += step {

        p, pLast := Point{}, Point{}
        first := true
        total := 0.0
        counter := 0


        fmt.Println(360.0, " - ", rotate, " - ", gearRotation)        
        start := 360.0 - rotate - gearRotation
        for {
            p := e.PointByAngleRotated(start, rotate)
            a2 := rToD(math.Atan2(p.Y, p.X))

            if a2 > -gearRotation {
                start -= 0.0001
            } else {
                start += 0.0001
            }
            
            if math.Abs(a2 - -gearRotation) < 0.001 {
                break
            }
        }
        
        degreePerTooth := 360.0 / float64(teethRigid)
        degreePerEllipseRotation := degreePerTooth * 2.0
        perDegree := degreePerEllipseRotation / 360
        gearRotation += (perDegree * step) // 90 see loop step
    

        // loop start-360
        for i := start; i <= 360.0; i += 0.001 {

            p = e.PointByAngleRotated(i, rotate)
            if first {
                printDot(b, Point{p.X + e.X, p.Y + e.Y}, "#1f1e1e")
                
                b.WriteString("<path stroke=\"#1f1e1e\" d=\"M")
                drawTooth(b, e, i, p, rotate, first)
                
                first = false
                pLast = p
                counter++
                continue
            }

            // set teeth by distance
            total += p.DistanceTo(&pLast) / 2
            pLast = p
            if total >= (float64(counter) * toothWidth) {
                drawTooth(b, e, i, p, rotate, first)
                counter++
            }

            if counter == teeth {
                break
            }
        }

        // start-360
        for i := 0.0; i < start; i += 0.001 {

            p = e.PointByAngleRotated(i, rotate)
            if first {
                printDot(b, Point{p.X + e.X, p.Y + e.Y}, "#1f1e1e")
                
                b.WriteString("<path stroke=\"#1f1e1e\" d=\"M")
                drawTooth(b, e, i, p, rotate, first)
                
                first = false
                pLast = p
                counter++
                continue
            }

            total += p.DistanceTo(&pLast) / 2
            pLast = p
            if total >= (float64(counter) * toothWidth) {
                drawTooth(b, e, i, p, rotate, first)
                counter++
            }

            if counter == teeth {
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
    tooth.rotate(180)
    tooth.rotate(rToD(tan) + rotate)

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
   .highlight {
     background: white;
   }
   </style>
   <script src="https://code.jquery.com/jquery-3.5.0.js"></script>
 <script>
 function hideSVG(id) {
   var style = document.getElementById(id).style.display;
   if(style === "none")
     document.getElementById(id).style.display = "block";
   else
     document.getElementById(id).style.display = "none";
 }
 
 $(document).ready(function() {
 $( "button" ).click(function() {
   $( this ).toggleClass( "highlight" );
 });
 });
 
 </script>
 </head>
 <body>
 <button id="button" onclick="hideSVG('gear1')">1</button>
 <button onclick="hideSVG('gear2')">2</button>
 <button onclick="hideSVG('gear3')">3</button>
 <button onclick="hideSVG('gear4')">4</button>
 <button onclick="hideSVG('gear5')">5</button>
 <button onclick="hideSVG('gear6')">6</button>
 <button onclick="hideSVG('gear7')">7</button><br>
 <svg width="1200" height="1000" xmlns="http://www.w3.org/2000/svg">
 <g stroke-width="0.5" fill="none">
`)

    return &b
}