package main

import (
    "bytes"
    "fmt"
    "log"
)

func main() {
       
    // In Kreisform:
    // - Den Zahn skalieren (was ist 1???)
    // - Den Zahn auf 0 bringen mit Mittelpunkt
    // - Zahn rotieren um x
    // - Zahn auf den Mittelpunkt im Kreis setzen

    // Zahnedesign
    // - Vorlage abkupfern
    // - Flexpline rotieren (2 Zähne pro durchlauf und Elipse rotieren) - 100x und übereinanderlegen, alles weiße ist das rigid gear

    var b bytes.Buffer
    b.WriteString("<svg width=\"1500\" height=\"1000\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<line x1=\"0\" y1=\"47\" x2=\"1500\" y2=\"47\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")

    moveX := 780.0
    moveY := 47.0
    
    // Höhe 691
    // Breite 644

    // 0 basiert
    // move
    // scale - modulo   // simpel, wie move..  moveX + (point.X / scale) - vorher Höhe auf 1 bringen... alles durch 64.4
    // rotate
   
    circle1 := newCircle(-700, 353, 501)
    circle2 := newCircle(72, 436, 260)
    circle3 := newCircle(-72, 436, 260)
    circle4 := newCircle(700, 353, 501)

    l1, l2 := circle1.InnerTangentTo(circle2)
    b.WriteString("<path d=\"M")
        
    points := circle1.CoordinatesBetween(Point{circle1.X + 170, circle1.Y - 169}, l1.p2)
    first := true
    for _, point := range points {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX + point.X , moveY + point.Y))
        if first {
            first = false
            b.WriteString("L")
        }
    }
    
    points = circle2.CoordinatesBetween(Point{20, 503}, l1.p1) 
    for i := len(points) - 1; i >= 0; i-- {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX + points[i].X, moveY + points[i].Y))
    }

    _, l2 = circle3.InnerTangentTo(circle4)
    points = circle3.CoordinatesBetween(l2.p2, Point{-20, 503})
    for i := len(points) - 1; i >= 0; i-- {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX + points[i].X, moveY + points[i].Y))
    }

    points = circle4.CoordinatesBetween(l2.p1, Point{circle4.X - 170, circle4.Y - 169})
    for _, point := range points {
        b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX + point.X, moveY + point.Y))
    }

    b.WriteString("\"/>\n")

    // Debug
    printDot(&b, l1.p1, moveX, moveY)
    printDot(&b, l1.p2, moveX, moveY)
   
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
