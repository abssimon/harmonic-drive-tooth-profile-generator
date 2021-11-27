package main

import (
    "bytes"
    "fmt"
    "log"
    "os"
)

var dotCounter int

func printDot(b *bytes.Buffer, p Point, id string) {
    b.WriteString(fmt.Sprintf("<circle id=\"%d\" fill=\"%s\" cx=\"%.14f\" cy=\"%.14f\"  r=\"0.8\" />\n", dotCounter, id, p.X, p.Y)) // 0.1
    dotCounter++
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

func svg(gears []Gear) { 

    var b bytes.Buffer
    x := 750.0
    y := 650.0
    // x = 0
    // y = 0
    
    b.WriteString(`<svg width="1500" height="1500" xmlns="http://www.w3.org/2000/svg"><g stroke-width="0.5" fill="none">`)

    colors := []string{"red", "blue", "#1f1e1e"}

    for i, gear := range gears {
        
        b.WriteString(fmt.Sprintf("<path id=\"gear%d\" stroke=\"%s\" d=\"M", i, colors[i]))
        
        first := true
        for _, t := range gear.Tooths {
      
            if first {
                first = false 
            } else {
                b.WriteString(" L")
            }
        
            p1 := t.C4.PointInAngle(t.C4.Start)
            p2 := t.C4.PointInAngle(t.C4.Stop)
            
            // scale up and rotate gear if needed
            pu1 := Point{p1.X + t.pos.X * 200.0, p1.Y + t.pos.Y * 200.0}
            pu2 := Point{p2.X + t.pos.X * 200.0, p2.Y + t.pos.Y * 200.0}
            pu1.Rotate(gear.Angle)
            pu2.Rotate(gear.Angle)          
           
            b.WriteString(fmt.Sprintf(" %.3f %.3f A %.3f %.3f 0 0 0 %.3f %.3f", pu2.X + x, pu2.Y + y, t.C4.Radius, t.C4.Radius, pu1.X + x, pu1.Y + y))
           
            p1 = t.C3.PointInAngle(t.C3.Start)
            p2 = t.C3.PointInAngle(t.C3.Stop)
            
            pu1 = Point{p1.X + t.pos.X * 200.0, p1.Y + t.pos.Y * 200.0}
            pu2 = Point{p2.X + t.pos.X * 200.0, p2.Y + t.pos.Y * 200.0}
            pu1.Rotate(gear.Angle)
            pu2.Rotate(gear.Angle) 
            
            b.WriteString(fmt.Sprintf(" L %.3f %.3f A %.3f %.3f 0 0 1 %.3f %.3f", pu1.X + x, pu1.Y + y, t.C3.Radius, t.C3.Radius, pu2.X + x, pu2.Y + y))
            
            p1 = t.C2.PointInAngle(t.C2.Start)
            p2 = t.C2.PointInAngle(t.C2.Stop)
            
            pu1 = Point{p1.X + t.pos.X * 200.0, p1.Y + t.pos.Y * 200.0}
            pu2 = Point{p2.X + t.pos.X * 200.0, p2.Y + t.pos.Y * 200.0}
            pu1.Rotate(gear.Angle)
            pu2.Rotate(gear.Angle)  
            
            b.WriteString(fmt.Sprintf(" L %.3f %.3f A %.3f %.3f 0 0 1 %.3f %.3f", pu1.X + x, pu1.Y + y, t.C2.Radius, t.C2.Radius, pu2.X + x, pu2.Y + y))

            p1 = t.C1.PointInAngle(t.C1.Start)
            p2 = t.C1.PointInAngle(t.C1.Stop)
            
            pu1 = Point{p1.X + t.pos.X * 200.0, p1.Y + t.pos.Y * 200.0}
            pu2 = Point{p2.X + t.pos.X * 200.0, p2.Y + t.pos.Y * 200.0}
            pu1.Rotate(gear.Angle)
            pu2.Rotate(gear.Angle) 
            
            b.WriteString(fmt.Sprintf(" L %.3f %.3f A %.3f %.3f 0 0 0 %.3f %.3f", pu2.X + x, pu2.Y + y, t.C1.Radius, t.C1.Radius, pu1.X + x, pu1.Y + y))
        

        }
        
        b.WriteString(" Z\"/>\n")
   
    }
    
    printDot(&b, Point{x, y}, "black")

    b.WriteString(fmt.Sprintf("<ellipse cx=\"%f\" cy=\"%f\" rx=\"%f\" ry=\"%f\" style=\"stroke:grey;stroke-width:0.5\" />", x, y, 4.2 * 95, 4.035 * 95))
    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }
}
