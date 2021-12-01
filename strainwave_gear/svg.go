package main

import (
    "fmt"
    "log"
    "os"
)



// one html file ... because svg objects local dont work, see 
// https://support.mozilla.org/en-US/questions/1279538
func aniTop(b *LineWriter) {
    b.WriteLine(`<!DOCTYPE html>`)
    b.WriteLine(`<html>`)
    b.WriteLine(`<head>`)
    b.WriteLine(`<meta charset="UTF-8">`)
    b.WriteLine(`<style>`)
    b.WriteLine(`    .flex {`)
    b.WriteLine(`        display: none`)
    b.WriteLine(`    }`)
    b.WriteLine(`</style>`)
    b.WriteLine(`<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>`)
    b.WriteLine(`<script>`)
    b.WriteLine(`    $(document).ready(function() {`)
    b.WriteLine(`        var counter = 0;`)
    b.WriteLine(`        setInterval(function(){`)
    b.WriteLine(`            document.getElementById("gear"+counter).style.display = "block";`)
    b.WriteLine(`            if (counter === 0) {`)
    b.WriteLine(`                document.getElementById("gear99").style.display = "none";`)
    b.WriteLine(`            } else {`)
    b.WriteLine(`                document.getElementById("gear"+(counter-1)).style.display = "none";`)
    b.WriteLine(`            }`)
    b.WriteLine(`            counter++`)
    b.WriteLine(`            if (counter === 100) {`)
    b.WriteLine(`                counter=0`)
    b.WriteLine(`            }`)
    b.WriteLine(`        }, 40);`)
    b.WriteLine(`    });`)
    b.WriteLine(``)
    b.WriteLine(`</script>`)
    b.WriteLine(`</head>`)
    b.WriteLine(`<body>`)
}

func aniBottom(b *LineWriter) {
    b.WriteLine(`</body>`)
    b.WriteLine(`</html>`)
}

func svg(mode string, gears []Gear) { 

    b := NewLineWriter()
    x := 500.0
    y := 500.0
  
    if mode == "ani" {
         aniTop(b)
    }
    
    b.WriteLine(`<svg width="1200" height="1200" xmlns="http://www.w3.org/2000/svg"><g stroke-width="0.5" fill="none">`)

    for i, gear := range gears {
            
        color := "blue"
        css := `class="flex"`
        if mode == "ani" {
            if i == 100 {
                color = "red"
                css = ""
            }
        } else {
            if i == 0 {
                color = "red"
                css = ""
            }
        }
         
              
        b.WriteString(fmt.Sprintf("<path id=\"gear%d\" stroke=\"%s\" %s d=\"M", i, color, css))
        
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
        
        b.WriteLine(" Z\"/>")
   
    }
    
    printDot(b, Point{x, y}, "black")

    if mode == "both" {
        b.WriteLine(fmt.Sprintf("<ellipse cx=\"%f\" cy=\"%f\" rx=\"%f\" ry=\"%f\" style=\"stroke:grey;stroke-width:0.5\" />", x, y, 4.2 * 95, 4.035 * 95))
    }
    
    b.WriteLine("</g></svg>")
    
    fName := "test.svg"
    if mode == "ani" {
        aniBottom(b)
        fName = "test.html"
    }

    err := writeSvg(fName, b.Bytes())
    if err != nil {
        log.Fatal(err)
    }
}

var dotCounter int

func printDot(b *LineWriter, p Point, id string) {
    b.WriteLine(fmt.Sprintf("<circle id=\"%d\" fill=\"%s\" cx=\"%.14f\" cy=\"%.14f\"  r=\"0.8\" />", dotCounter, id, p.X, p.Y)) 
    dotCounter++
}

func writeSvg(fName string, data []byte) error {
    f, err := os.OpenFile(fName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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
