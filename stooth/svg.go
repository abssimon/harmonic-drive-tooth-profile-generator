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

func svg(e Ellipse, gear []Point) {

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

    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n", e.X, e.Y, e.Width+5))
    printDot(&b, e.Point, "black")
    b.WriteString(fmt.Sprintf(`<g transform="rotate(-45 600 500)"><ellipse rx="%.14f" ry="%.14f" cx="600" cy="500" stroke="#cccccc" stroke-width="0.5" fill="none" /></g>`, e.Height-10.0, e.Width-10.0))
    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }
}