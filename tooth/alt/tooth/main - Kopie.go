package main

import (
    "bytes"
    "fmt"
    "log"
    "math"
    "os"
)

// degree to radian
func dToR(deg float64) float64 {
    return deg * (math.Pi / 180.0)
}

// radian to degree
func rToD(rad float64) float64 {
    return rad * (180.0 / math.Pi)
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
    return math.Abs(a - b) <= float64EqualityThreshold
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

    scale := 50.0
    middle := 300.0

    var b bytes.Buffer
    b.WriteString("<svg width=\"600\" height=\"700\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<line x1=\"0\" y1=\"300\" x2=\"600\" y2=\"300\" style=\"stroke:rgb(255,0,0);stroke-width:2\" />")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\"><circle id=\"pointA\" cy=\"0\" cx=\"0\" r=\"1\" />")


    radiusTop := 4.5
    topCenterX := -80.0
    topCenterY := -60.0
    
    for i := 0.0; i <= radiusTop; i += 0.1 {
        b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", (i*scale)+topCenterX, (getY(radiusTop, i, -1)*scale)+middle+topCenterY))
    }

    radiusBottom := 4.9
    bottomCenterX := 430.0
    bottomCenterY := 73.0
    
    for i := 0.0; i >= radiusBottom *-1; i -= 0.1 {
        b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", (i*scale)+bottomCenterX, (getY(radiusBottom, i, 1)*scale)+middle+bottomCenterY))
    }
    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

    // TODO

}


func getY(radius float64, x float64, flip float64) float64 {
    // calculate angles first
    beta := rToD(math.Asin(x * math.Sin(dToR(90)) / radius))
    gamma := (beta * -1.0) - 90.0 + 180.0

    // calculate y
    y := math.Sqrt(math.Pow(radius, 2) - 2*radius*x*math.Cos(dToR(gamma)) + math.Pow(x, 2)) * flip // -1 flip it
    return y
}