package main

import (
    "bytes"
    "fmt"
    "os"
    "log"
)

func printDot(b *bytes.Buffer, p Point, id string) {
    b.WriteString(fmt.Sprintf("<circle stroke=\"%s\" cx=\"%.14f\" cy=\"%.14f\"  r=\"1\" />\n", id, p.X, p.Y))
}

func writeSvg(data []byte) error {
    f, err := os.OpenFile("test.svg", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.Write(data)
    if err != nil {
        return err
    }

    return nil
}

func svg(e Ellipse, t Tooth) {

    var b bytes.Buffer
    b.WriteString(`<svg width="1200" height="1000" xmlns="http://www.w3.org/2000/svg"><g stroke-width="0.5" fill="none">`)

    b.WriteString("<path stroke=\"blue\" d=\"M")  

    p1 := t.C4.PointInAngle(t.C4.Start)
    p2 := t.C4.PointInAngle(t.C4.Stop)
    b.WriteString(fmt.Sprintf(" %.2f %.2f A %.2f %.2f 0 0 0 %.2f %.2f", e.X + p2.X, e.Y + p2.Y, t.C4.Radius, t.C4.Radius, e.X + p1.X, e.Y + p1.Y))

    p1 = t.C3.PointInAngle(t.C3.Start)
    p2 = t.C3.PointInAngle(t.C3.Stop)
    b.WriteString(fmt.Sprintf("L %.2f %.2f A %.2f %.2f 0 0 1 %.2f %.2f", e.X + p1.X, e.Y + p1.Y, t.C3.Radius, t.C3.Radius, e.X + p2.X, e.Y + p2.Y))

    p1 = t.C2.PointInAngle(t.C2.Start)
    p2 = t.C2.PointInAngle(t.C2.Stop)
    b.WriteString(fmt.Sprintf("L %.2f %.2f A %.2f %.2f 0 0 1 %.2f %.2f", e.X + p1.X, e.Y + p1.Y, t.C2.Radius, t.C2.Radius, e.X + p2.X, e.Y + p2.Y))
    
    p1 = t.C1.PointInAngle(t.C1.Start)
    p2 = t.C1.PointInAngle(t.C1.Stop)
    b.WriteString(fmt.Sprintf("L %.2f %.2f A %.2f %.2f 0 0 0 %.2f %.2f", e.X + p2.X, e.Y + p2.Y, t.C1.Radius, t.C1.Radius, e.X + p1.X, e.Y + p1.Y ))


    /*    
    M 85 350 = x,y Startpunkt
                                        +--- x-Endpunkt
                                        |
             gegen Uhrzeigersinn ---+   |   +--- y-Endpunkt
                                    |   |   |
    <path d="M 85 350 A 150 180 0 0 0  280 79" stroke="red" fill="none"/>
                         |   |  | |
     1 Radius x-Achse ---+   |  | +--- 4 kurzer / langer Weg
                             |  |
         2 Radius y-Achse ---+  +--- 3 Rotation x
    */     
    
    b.WriteString("\"/>\n")

    printDot(&b, e.Point, "black")

    printDot(&b, Point{t.C1.Point.X + e.X, t.C1.Point.Y + e.Y}, "black")
    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n", t.C1.Point.X + e.X, t.C1.Point.Y + e.Y, t.C1.Radius))
    printDot(&b, Point{t.C2.Point.X + e.X, t.C2.Point.Y + e.Y}, "black")
    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n", t.C2.Point.X + e.X, t.C2.Point.Y + e.Y, t.C2.Radius))
    printDot(&b, Point{t.C3.Point.X + e.X, t.C3.Point.Y + e.Y}, "black")
    printDot(&b, Point{t.C4.Point.X + e.X, t.C4.Point.Y + e.Y}, "black")
    
    tan := t.C1.PointInAngle(t.C1.Stop)    
    printDot(&b, Point{tan.X + e.X, tan.Y + e.Y}, "white")
    
    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }
}