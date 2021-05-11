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

type Point struct {
    X, Y float64
}

type Circle struct {
    Point
    Radius float64
}

func (c Circle) PointInAngle(a float64) Point {
    x := math.Cos(dToR(a))*c.Radius + c.X
    y := math.Sin(dToR(a))*c.Radius + c.Y
    return Point{x, y}
}

func newCircle(x, y, r float64) *Circle {
    return &Circle{Point{x, y}, r}
}

type Ellipse struct {
    Point
    Height float64
    Width  float64
}

func (e Ellipse) PointInAngle(a float64) Point {
    da := dToR(a)
    x := e.Width*math.Cos(da) + e.X
    y := e.Height*math.Sin(da) + e.Y
    return Point{x, y}
}

// Rotated Elypse, get point in an absolute angel, not relativ

// https://math.stackexchange.com/questions/2645689/what-is-the-parametric-equation-of-a-rotated-ellipse-given-the-angle-of-rotatio
func (e Ellipse) PointInAngleRotated(a float64, r float64) Point {
    da := dToR(a)
    dr := dToR(r)
    x := e.Width*math.Cos(da)*math.Cos(dr) - e.Height*math.Sin(da)*math.Sin(dr) + e.X
    y := e.Width*math.Cos(da)*math.Sin(dr) + e.Height*math.Sin(da)*math.Cos(dr) + e.Y
    return Point{x, y}
}

func (e Ellipse) Circumference() float64 {
    a := e.Height / 2.0
    b := e.Width / 2.0
    t := (a - b) / (a + b)
    return (a + b) * math.Pi * (1.0 + 3.0*t*t/(10.0+math.Sqrt(4.0-3.0*t*t)))
}

func newEllipseByHeightAndCirc(x, y, r float64) {
    // TODO
}

type FlexSpline struct {
    Ellipse
    Arclen      float64 // to rotate teeth precisely
    Teeth       int
    Coordinates []Point // to rotate ellipse
}

// https://math.stackexchange.com/questions/2645689/what-is-the-parametric-equation-of-a-rotated-ellipse-given-the-angle-of-rotatio
func (f *FlexSpline) Rotate() {
    // wo rotieren?
}

type Cache struct {
    Point
    Angle float64
}

func main() {

    var b bytes.Buffer
    b.WriteString("<svg width=\"700\" height=\"700\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")

    //teeth := 72
    

    elipse := Ellipse{
        Point{350.0, 350.0}, // needs to be 0 for rotation, add this later... but do not move dots!!! ????
        250.0,
        210.0,
        // Coordinaten []Points // type FlexSpline?
    }
    //toothArc := elipse.Circumference() / float64(teeth)

    rotate := 20.0
    

//for rotate := 0.0; rotate <= 25.0; rotate += 4.0 {
    
    // calculate tooth points
    //p, p_first, p_last := Point{}, Point{}, Point{}
    //total := 0.0
    //counter := 0
    
    cache := []Cache{}
    for i := 0.0; i <= 360.0; i += 0.001 { // .001 is a +-0.00218px precision    // TODO - Schleife anpassen!!!
        p := elipse.PointInAngleRotated(i, rotate)
        cache = append(cache, Cache{p, i})
    }
    
    for i, v := range cache {
        if v.Angle >= 360.0 - rotate {
        
        }
    }
    
    fmt.Println(len(cache))
    
    /*
    for i := 0.0; i <= 360.0; i += 0.1 { // .001 is a +-0.00218px precision    // TODO - Schleife anpassen!!!
        p = elipse.PointInAngleRotated(i, rotate)

        if p_first == (Point{}) {
            p_first, p_last = p, p
            b.WriteString(fmt.Sprintf("<circle id=\"%f\" cx=\"%f\" cy=\"%f\"  r=\"5\" fill=\"red\" />\n", i, p.X, p.Y))
            counter++
            continue
        }

        dist := (math.Sqrt(math.Pow(p.X-p_last.X, 2) + math.Pow(p.Y-p_last.Y, 2))) / 2 // its fit only with "/2", why?
        total += dist
        p_last = p

        fmt.Println(i, p, dist, total > (float64(counter) * toothArc))

        if total >= (float64(counter) * toothArc) {
            b.WriteString(fmt.Sprintf("<circle id=\"%f-%d\" cx=\"%.14f\" cy=\"%.14f\" r=\"5\" />\n", i, counter, p.X, p.Y))
            counter++
        }

        if counter == teeth { // avoid double points with high resolution
            break
        }
    }
    */
    //fmt.Println("Umfang", elipse.Circumference())
    //fmt.Println("Zaehne", teeth, "vs", counter)
    //fmt.Println("Zahnabstand", toothArc)
    //fmt.Println("Umfang summiert", total)
//}    

/*
    // rigid spline always follow...
    circle := newCircle(350.0, 350.0, 250.0)
    for i := 0.0; i <= 360.0; i += (360.0 / float64(teeth)) {
        p := circle.PointInAngle(i)
        b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", p.X, p.Y))
    }
*/
    printDot(&b, Point{350.0, 350.0})

    b.WriteString("</g></svg>")

    // toolthickness
    // https://gearsolutions.com/departments/tooth-tips/determining-tooth-thickness-of-various-gear-types/

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

}
func printDot(b *bytes.Buffer, p Point) {
    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"1\" />\n", p.X, p.Y))
}
