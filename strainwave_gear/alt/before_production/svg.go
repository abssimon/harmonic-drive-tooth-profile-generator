package main

import (
    "bytes"
    "fmt"
    "log"
    "os"
)

func printDot(b *bytes.Buffer, p Point, id string) {
    b.WriteString(fmt.Sprintf("<circle fill=\"%s\" cx=\"%.14f\" cy=\"%.14f\"  r=\"0.8\" />\n", id, p.X, p.Y)) // 0.1
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

func svg(gears [][]Point) { // e Ellipse,

    var b bytes.Buffer
    b.WriteString(`<svg width="1500" height="1500" xmlns="http://www.w3.org/2000/svg"><g stroke-width="0.5" fill="none">`)

    colors := []string{"red", "blue", "#1f1e1e"}

    for i, gear := range gears {
        
    
        b.WriteString(fmt.Sprintf("<path id=\"gear%d\" stroke=\"%s\" d=\"M", i, colors[i]))
    
        first := true
        for _, point := range gear {
            b.WriteString(fmt.Sprintf("%.2f %.2f ", point.X+750.0, point.Y+650.0)) // "%.14f %.14f "
            if first {
                b.WriteString("L ")
                first = false
            }
        }
        b.WriteString("Z\"/>\n")
    }
    /*
       for _, p := range gear {
           //printDot(&b, Point{p.X*6000.0-8100.0, p.Y*6000.0-8100.0}, "black")
           //printDot(&b, Point{p.X*200.0+750.0, p.Y*200.0+650.0}, "black")
           printDot(&b, Point{p.X+750.0, p.Y+650.0}, "black")
       }
    */
    //b.WriteString(fmt.Sprintf("<line x1=\"%.14f\" y1=\"%.14f\" x2=\"%.14f\" y2=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />", e.X-600, e.Y, e.X+600, e.Y))
    //b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"%.14f\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />\n", e.X, e.Y, e.Width+35))
    printDot(&b, Point{750.0, 650.0}, "black")
    //b.WriteString(`<g transform="rotate(-45 750 650)"><ellipse rx="285" ry="322" cx="600" cy="500" stroke="#cccccc" stroke-width="0.5" fill="none" /></g>`)
    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }
}
