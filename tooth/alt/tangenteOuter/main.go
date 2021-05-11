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
    return math.Abs(a-b) <= float64EqualityThreshold
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
    b.WriteString("<svg width=\"1000\" height=\"1000\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<line x1=\"0\" y1=\"300\" x2=\"700\" y2=\"300\" style=\"stroke:rgb(255,0,0);stroke-width:2\" />")
    b.WriteString("<g transform=\"translate(0,1000)\"><g transform=\"scale(1,-1)\"><g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\"><circle cx=\"60.000000\" cy=\"100.0\" r=\"1\"/>")

    radiusTop := 50.0
    topCenterX := 726.0
    topCenterY := 330.0
    b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"%f\" />\n", topCenterX, topCenterY, radiusTop))

    for i := 0.0; i <= radiusTop; i += 2.0 {
       // b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", i+topCenterX, getY(radiusTop, i, -1)+topCenterY))
    }

    radiusBottom := 257.0
    bottomCenterX := 849.0
    bottomCenterY := 520.0
    b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"%f\" />\n", bottomCenterX, bottomCenterY, radiusBottom))

    for i := 0.0; i >= radiusBottom*-1; i -= 2.0 {
       // b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", i+bottomCenterX, getY(radiusBottom, i, 1)+bottomCenterY))
    }


    // Tangente
    // https://www.weltderfertigung.de/suchen/lernen/mathematik/beruehrpunktberechnung-tangente-an-zwei-kreisen.php (excel)
    distCenterToCenter := math.Sqrt(math.Pow(topCenterY-bottomCenterY, 2) + math.Pow(bottomCenterX-topCenterX, 2)) // distCenterToCenter

    fmt.Println(distCenterToCenter)

    // Outer 
    radiusDiff := math.Abs(radiusTop-radiusBottom)
    diffRatio := distCenterToCenter / radiusDiff
    distTanIntersection := diffRatio * radiusTop
    
    alpha1 := rToD(math.Asin(radiusTop / distTanIntersection))
    alpha2 := rToD(math.Asin((topCenterY - bottomCenterY) / distCenterToCenter)) - alpha1
    alpha3 := math.Asin((topCenterY - bottomCenterY) / distCenterToCenter)
    alpha4 := 90 - (rToD(math.Asin((topCenterY - bottomCenterY) / distCenterToCenter)) + alpha1) 
    
    c1 := math.Sqrt(math.Pow(distTanIntersection,2) - math.Pow(radiusTop,2))
    a1 := math.Sin(dToR(alpha2)) * c1
    b1 := math.Cos(dToR(alpha2)) * c1
    
    c2 := distCenterToCenter + distTanIntersection
    b2 := math.Cos(alpha3) * c2
    a2 := math.Sin(alpha3) * c2
            
    b3 := math.Sqrt(math.Pow(c2, 2) - math.Pow(radiusBottom, 2)) 
    b4 := math.Cos(dToR(alpha2)) * b3
    a4 := math.Sin(dToR(alpha2)) * b3   
    b5 := math.Cos(dToR(alpha4)) * b3
    a5 := math.Sin(dToR(alpha4)) * b3
    b6 := math.Cos(dToR(alpha4)) * c1
    a6 := math.Sin(dToR(alpha4)) * c1

    ttx1 := bottomCenterX - b2 + b1
    tty1 := bottomCenterY + a2 - a1
    ttx2 := bottomCenterX - b2 + b4
    tty2 := bottomCenterY + a2 - a4
    tbx1 := bottomCenterX - b2 + a5 
    tby1 := bottomCenterY + a2 - b5    
    tbx2 := bottomCenterX - b2 + a6 
    tby2 := bottomCenterY + a2 - b6
    
    b.WriteString(fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" />\n", tbx1, tby1, tbx2, tby2))
    b.WriteString(fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" />\n", ttx1, tty1, ttx2, tty2))
    
    
    
    
    
/*

    // Inner Tangente
    sizeRatio := radiusBottom / radiusTop   
    distCenterToCenter1 := distCenterToCenter / 100 * (100.0 / (sizeRatio + 1.0) * sizeRatio)
    //distCenterToCenter2 := distCenterToCenter - distCenterToCenter1

    // angles
    alpha1 := rToD(math.Asin(radiusBottom / distCenterToCenter1))
    gamma1 := 180.0 - (alpha1 + 90.0)
    alpha2 := rToD(math.Asin((topCenterY - bottomCenterY) / distCenterToCenter))
    alpha3 := 90 - (alpha2 + gamma1)
    alpha4 := gamma1 - alpha2
    
    tbx3 := math.Sin(dToR(alpha3)) * radiusBottom
    tby3 := math.Cos(dToR(alpha3)) * radiusBottom
    tbx4 := math.Sin(dToR(alpha4)) * radiusBottom
    tby4 := math.Cos(dToR(alpha4)) * radiusBottom

    ttx3 := math.Sin(dToR(alpha3)) * radiusTop
    tty3 := math.Cos(dToR(alpha3)) * radiusTop
    ttx4 := math.Sin(dToR(alpha4)) * radiusTop
    tty4 := math.Cos(dToR(alpha4)) * radiusTop

    b.WriteString(fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" />\n", bottomCenterX-tbx3, bottomCenterY+tby3, topCenterX+ttx3, topCenterY-tty3))
    b.WriteString(fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" />\n", bottomCenterX-tby4, bottomCenterY-tbx4, topCenterX+tty4, topCenterY+ttx4))


*/


    b.WriteString("</g></g></g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

}

// https://www.mathepower.com/dreieck.php
func getY(radius float64, x float64, flip float64) float64 {
    // calculate angles first
    beta := rToD(math.Asin(x * math.Sin(dToR(90)) / radius))
    gamma := (beta * -1.0) - 90.0 + 180.0

    // calculate y
    y := math.Sqrt(math.Pow(radius, 2)-2*radius*x*math.Cos(dToR(gamma))+math.Pow(x, 2)) * flip // -1 flip it
    return y
}
