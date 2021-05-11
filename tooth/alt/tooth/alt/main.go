package main

import (
    "bytes"
    "fmt"
    "log"
    "math"
    "os"
)

type Parameter struct {
    z1 int     // Number of flexible gear teeth 240
    z2 int     // Number of rigid gear teeth 242
    i  int     // Transmission ratio 120
    m  float32 // Modulus 0.25
    pa int     // (mm) Arc radius of the tooth profile of a convex tooth
    la int     // (mm) Lateral offset of a convex tooth profile
    ea int     // (mm) Longitudinal offset of a convex tooth profile
    pf int     // (mm) Arc radius of the tooth profile of a concave tooth
    lf int     // (mm) Lateral offset of concave tooth profile
    ef int     // (mm) Longitudinal offset of concave tooth profile
    Sa int     // (mm) Tooth thickness
    om float32 // Tangent angle
    t  int     // (mm) Wall thickness of the flexible gear
}

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
    /*
        p := Parameter{
            z1: 240,
            z2: 242,
            i:  120,
            m:  0.25, // falsch in relation
            pa: 237,
            la: 66,
            ea: 108,
            pf: 293,
            lf: 0, //121 + (p.m / 2.0), // math.Pi
            ef: -114,
            Sa: 395,
            om: 10.0,
            t:  56,
        }
        fmt.Println(p)
        fmt.Println(math.Pi * p.m / 2.0)
    */
    radius := 5.0
    //x := 3.0

    var b bytes.Buffer

    b.WriteString("<svg width=\"600\" height=\"600\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<line x1=\"0\" y1=\"300\" x2=\"600\" y2=\"300\" style=\"stroke:rgb(255,0,0);stroke-width:2\" />")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.1\" fill=\"none\"><circle id=\"pointA\" cy=\"0\" cx=\"0\" r=\"1\" />")
    for i := 0.0; i <= radius; i += 0.1 {
        if almostEqual(i, 3.0) {
            fmt.Println("aa")
            b.WriteString(fmt.Sprintf("<circle cy=\"%f\" cx=\"%f\" r=\"1\" style=\"stroke:rgb(255,0,0);stroke-width:2\" />\n", getY(radius, i)*50.0, i*50.0))
        } else {
            fmt.Println(i)
            b.WriteString(fmt.Sprintf("<circle cy=\"%f\" cx=\"%f\" r=\"1\" />\n", getY(radius, i)*50.0, i*50.0))
        }
    }
    b.WriteString("</g></svg>")

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

    // TODO
    // Kreis verschieben....

}

func getY(radius float64, x float64) float64 {
    // calculate angles first
    beta := rToD(math.Asin(x * math.Sin(dToR(90)) / radius))
    gamma := (beta * -1.0) - 90.0 + 180.0

    // calculate y
    y := math.Sqrt(math.Pow(radius, 2) - 2*radius*x*math.Cos(dToR(gamma)) + math.Pow(x, 2))
    return y
}

func getYUntenLinksObenRechts(radius float64, x float64) float64 {
    // calculate angles first
    beta := rToD(math.Asin(x * math.Sin(dToR(90)) / radius))
    gamma := (beta * -1.0) - 90.0 + 180.0

    // calculate y
    y := math.Sqrt(math.Pow(radius, 2) - 2*radius*x*math.Cos(dToR(gamma)) + math.Pow(x, 2))
    return y
}
