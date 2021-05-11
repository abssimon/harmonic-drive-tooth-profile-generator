package main

import (
    "bytes"
    "fmt"
    "log"
)

type STooth struct {
    C1   *Circle
    C2   *Circle
    C3   *Circle
    C4   *Circle
    Tan1 Line
    Tan2 Line
}

func (t STooth) rotate(a float64) {
    t.C1.Rotate(a)
    t.C2.Rotate(a)
    t.C3.Rotate(a)
    t.C4.Rotate(a)
}

func (t STooth) scale() {

}

func (t STooth) moveX() {

}

func (t STooth) moveY() {

}

func newSTooth(tipCenter Point, tipRadius float64, tipLimit Point, bottomCenter Point, bottomRadius float64, bottomLimit Point) STooth {

    tooth := STooth{}
    tooth.C1 = newCircle(-tipCenter.X, tipCenter.Y, bottomRadius)
    tooth.C2 = newCircle(bottomCenter.X, bottomCenter.Y, tipRadius)
    tooth.C3 = newCircle(-bottomCenter.X, bottomCenter.Y, tipRadius)
    tooth.C4 = newCircle(tipCenter.X, tipCenter.Y, bottomRadius)

    tooth.Tan1, _ = tooth.C1.InnerTangentTo(tooth.C2)
    tooth.C1.StartPoint = &Point{tooth.C1.X + bottomLimit.X, tooth.C1.Y - bottomLimit.Y}
    tooth.C1.StopPoint = tooth.Tan1.p2

    tooth.C2.StartPoint = &Point{tooth.C2.X - tipLimit.X, tooth.C2.Y + tipLimit.Y}
    tooth.C2.StopPoint = tooth.Tan1.p1

    _, tooth.Tan2 = tooth.C3.InnerTangentTo(tooth.C4)
    tooth.C3.StartPoint = tooth.Tan2.p2
    tooth.C3.StopPoint = &Point{tooth.C3.X + tipLimit.X, tooth.C3.Y + tipLimit.Y}

    tooth.C4.StartPoint = tooth.Tan2.p1
    tooth.C4.StopPoint = &Point{tooth.C4.X - bottomLimit.X, tooth.C4.Y - bottomLimit.Y}

    return tooth
}

func main() {
    
    var b bytes.Buffer
    b.WriteString("<svg width=\"1500\" height=\"2000\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<line x1=\"0\" y1=\"47\" x2=\"1500\" y2=\"47\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")
                

    
    scale := 42.0

    
    
    // Circles

    elipse := newEllipse(750.0, 900.0, 440.0*1.4, 500.0*1.4)
    tooths := 200
    
    // 37.80000000000000
    
    // 223.20000000000044 0
    // 264.60000000000070 unvollstandig
    // 37.80000000000000 0
    // 84.59999999999992 verkuemmert
    
    84.59999999999992
    
    for i := 0.0; i <= 360.0; i += (360.0/float64(tooths)) {
        p := elipse.PointInAngle(i) // fangt rechts an...
                
        //if i > 190 && i < 350 {
            
            tooth := newSTooth(
                Point{700 / scale, 353 / scale}, // tipCenter
                260.0/scale,                     // tipRadius
                Point{52 / scale, 67 / scale},   // tipLimit - point in relation to center
                Point{72 / scale, 436 / scale},  // bottomCenter
                501.0/scale,                     // bottomRadius
                Point{170 / scale, 169 / scale}, // bottomLimit
            )
            tooth.rotate(i-90)
            moveX := p.X
            moveY := p.Y

            b.WriteString(fmt.Sprintf("<path id=\"%.14f\" d=\"M", i))

            // tooth.MoveX = 780.0
            // tooth.MoveY = 147.0
            // tooth.Coordinates()

            points := tooth.C1.Coordinates()
            b.WriteString(fmt.Sprintf("%.14f %.14f L", moveX+points[0].X, moveY+points[0].Y))
            for _, point := range points[1:] {
                b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX+point.X, moveY+point.Y))
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

            b.WriteString("\"/>\n")
            
            
            
        //} else {
        //    b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\" r=\"1\" />\n", p.X, p.Y))
        //}
        
    }
    
    printDot(&b, &Point{750.0, 900.0}, 0, 0)
    
    /*
    circle := newCircle(750.0, 750.0, 500.0)

    for i := 0.0; i <= 360.0; i += (360.0/202) {
        p := circle.PointInAngle(i)
        b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", p.X, p.Y))
    }
    */
    
    // Debug
    /*
    printDot(&b, tooth.Tan1.p1, moveX, moveY)
    printDot(&b, tooth.Tan1.p2, moveX, moveY)

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
