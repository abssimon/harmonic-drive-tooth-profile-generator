package main

import (
    "bytes"
    "fmt"
    "log"
    "math"
    "os"
)
type Point struct {
    X, Y float64
}

func (p *Point) Rotate(a float64) {
    rad := dToR(a)
    x, y := p.X, p.Y
    p.X = x*math.Cos(rad) - y*math.Sin(rad)
    p.Y = x*math.Sin(rad) + y*math.Cos(rad)
}

/*
    A B
    D C
*/
type Rect struct {
    A *Point
    B *Point
    C *Point
    D *Point
}

func (r Rect) Rotate(a float64) {
    r.A.Rotate(a)
    r.B.Rotate(a)
    r.C.Rotate(a)
    r.D.Rotate(a)
}

// degree to radian
func dToR(deg float64) float64 {
    return deg * (math.Pi / 180.0)
}

// radian to degree
func rToD(rad float64) float64 {
    return rad * (180.0 / math.Pi)
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


func main() {

    var b bytes.Buffer
    b.WriteString("<svg width=\"1500\" height=\"1000\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")
    b.WriteString("<path d=\"M")
    
    r := Rect {
        &Point{-700, 353},
        &Point{700, 353},
        &Point{72, 436},
        &Point{-72, 436},
    }
    
    moveX := 780.0
    moveY := 147.0
    
    r.Rotate(20)
    
    b.WriteString(fmt.Sprintf("%.14f %.14f L", moveX+r.A.X, moveY+r.A.Y))           
    b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX+r.B.X, moveY+r.B.Y))  
    b.WriteString(fmt.Sprintf("%.14f %.14f ", moveX+r.C.X, moveY+r.C.Y))   
    b.WriteString(fmt.Sprintf("%.14f %.14f Z", moveX+r.D.X, moveY+r.D.Y))

    b.WriteString("\"/>\n")   
    b.WriteString("</g></svg>")


    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

}
