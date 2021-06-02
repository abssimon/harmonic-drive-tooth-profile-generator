package main

import (
    "bytes"
    "fmt"
    "log"
)

type STooth struct {
    c1 Circle
    c2 Circle
    c3 Circle
    c4 Circle
}

func (t STooth) rotate() {

}

func (t STooth) scale() {

}

func (t STooth) moveX() {

}

func (t STooth) moveY() {

}

func main() {
    // tipCenter Point{700, 353}
    // bottomCenter Point{72, 436}
    // tipRadius  260
    // bottomRadius 501
    // tipLimitPoint Point{52, 67}
    // bottomLimitPoint Point{170, 169}

    circle1 := newCircle(-700, 353, 501) // anders
    circle2 := newCircle(72, 436, 260)  
    circle3 := newCircle(-72, 436, 260) 
    circle4 := newCircle(700, 353, 501) // anders

    // calculate
    t1, _ := circle1.InnerTangentTo(circle2)
    circle1.StartPoint = &Point{circle1.X + 170, circle1.Y - 169}
    circle1.StopPoint = t1.p2
    
    circle2.StartPoint = &Point{circle2.X - 52, circle2.Y + 67}
    circle2.StopPoint = t1.p1
    
    _, t2 := circle3.InnerTangentTo(circle4)
    circle3.StartPoint = t2.p2
    circle3.StopPoint = &Point{circle3.X + 52, circle3.Y + 67}
    
    circle4.StartPoint = t2.p1
    circle4.StopPoint = &Point{circle4.X - 170, circle4.Y - 169}

    


    // draw  
    
    rotate := 20.0
    circle1.Rotate(rotate)
    circle2.Rotate(rotate)
    circle3.Rotate(rotate)
    circle4.Rotate(rotate)
    
    moveX := 780.0
    moveY := 147.0
    
    var b bytes.Buffer
    b.WriteString("<svg width=\"1500\" height=\"1000\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<line x1=\"0\" y1=\"47\" x2=\"1500\" y2=\"47\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")
    b.WriteString("<path d=\"M")

    points := circle1.Coordinates()
    b.WriteString(fmt.Sprintf("%.14f %.14f L", moveX+points[0].X, moveY+points[0].Y))      
    for _, point := range points[1:] {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX+point.X, moveY+point.Y))
    }

    points = circle2.Coordinates()
    for i := len(points) - 1; i >= 0; i-- {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX+points[i].X, moveY+points[i].Y))
    }

    points = circle3.Coordinates()
    for i := len(points) - 1; i >= 0; i-- {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX+points[i].X, moveY+points[i].Y))
    }

    points = circle4.Coordinates()
    for _, point := range points {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX+point.X, moveY+point.Y))
    }

    b.WriteString("\"/>\n")

    // Debug
    printDot(&b, t1.p1, moveX, moveY)
    printDot(&b, t1.p2, moveX, moveY)

    printDot(&b, circle1.Point, moveX, moveY)
    printDot(&b, circle2.Point, moveX, moveY)
    printDot(&b, circle3.Point, moveX, moveY)
    printDot(&b, circle4.Point, moveX, moveY)

    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

}
