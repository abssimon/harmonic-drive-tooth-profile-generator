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

func svg(e Ellipse, gear []Point, t Tooth) {

    var b bytes.Buffer
    b.WriteString(`<svg width="1200" height="1000" xmlns="http://www.w3.org/2000/svg"><g stroke-width="0.5" fill="none">`)



    b.WriteString("<path stroke=\"#1f1e1e\" d=\"M")   
    first := true
    for _, point := range gear {
        b.WriteString(fmt.Sprintf("%.8f %.8f ", point.X+e.X, point.Y+e.Y)) // "%.14f %.14f "
        if first {
            b.WriteString("L ")
            first = false
        }
    }
    b.WriteString("Z\"/>\n")

    //b.WriteString(fmt.Sprintf("<line x1=\"%.14f\" y1=\"%.14f\" x2=\"%.14f\" y2=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />", e.X-600, e.Y, e.X+600, e.Y))
    //b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n", e.X, e.Y, e.Width+35))
    printDot(&b, e.Point, "black")
    


    printDot(&b, Point{t.C1.Point.X + e.X, t.C1.Point.Y + e.Y}, "black")
    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n", t.C1.Point.X + e.X, t.C1.Point.Y + e.Y, t.C1.Radius))
    printDot(&b, Point{t.C2.Point.X + e.X, t.C2.Point.Y + e.Y}, "black")
    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n", t.C2.Point.X + e.X, t.C2.Point.Y + e.Y, t.C2.Radius))
    printDot(&b, Point{t.C3.Point.X + e.X, t.C3.Point.Y + e.Y}, "black")
    printDot(&b, Point{t.C4.Point.X + e.X, t.C4.Point.Y + e.Y}, "black")
    
    tan := t.C1.PointInAngle(t.C1.Stop)
    
    printDot(&b, Point{tan.X + e.X, tan.Y + e.Y}, "white")
    
    //b.WriteString(`<g transform="rotate(-45 600 500)"><ellipse rx="285" ry="322" cx="600" cy="500" stroke="#cccccc" stroke-width="0.5" fill="none" /></g>`)
    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }
}