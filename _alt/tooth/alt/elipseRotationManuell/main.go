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

func main() {

    var b bytes.Buffer
    b.WriteString("<svg width=\"700\" height=\"700\" xmlns=\"http://www.w3.org/2000/svg\">")
    b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")

    teeth := 72
    
    elipse := Ellipse{
        Point{350.0, 350.0}, 
        250.0,
        210.0,
    }
    toothArc := elipse.Circumference() / float64(teeth)
    
    colors := []string{"red", "yellow", "green", "blue", "grey", "HotPink", "Orange", "LawnGreen"}
    ccounter := 0 
    //rotate := 20.0
for rotate := 0.0; rotate <= 120.0; rotate += 60.0 { 
    
    p_first, p_last := Point{}, Point{}
    total := 0.0
    counter := 0
    
    // hmm, simulation fuer einen zahn?
    
    // find the right start point ...
    start := 360.0 - rotate
    for {
        p := elipse.PointInAngleRotated(start, rotate)
        if p.Y > elipse.X {
            start -= 0.0001
        } else {
            start += 0.0001
        }   
        if math.Abs(p.Y - elipse.X) < 0.001 {
            fmt.Println("break")
            break
        }
    }
    
    for i := start; i <= 360.0; i += 0.1 { 
        p := elipse.PointInAngleRotated(i, rotate)

        if p_first == (Point{}) {
            p_first, p_last = p, p
            b.WriteString(fmt.Sprintf("<circle id=\"%f\" cx=\"%f\" cy=\"%f\"  r=\"5\" fill=\"%s\" />\n", i, p.X, p.Y, colors[ccounter]))
            counter++
            continue
        }

        total += (math.Sqrt(math.Pow(p.X-p_last.X, 2) + math.Pow(p.Y-p_last.Y, 2))) / 2
        p_last = p

        if total >= (float64(counter) * toothArc) {
            b.WriteString(fmt.Sprintf("<circle id=\"%f-%d\" cx=\"%.14f\" cy=\"%.14f\" r=\"5\" fill=\"%s\" />\n", rotate, counter, p.X, p.Y, colors[ccounter]))
            counter++
        }
        if counter == teeth { 
            break
        }
    }
    for i := 0.0; i < start; i += 0.1 { 
        p := elipse.PointInAngleRotated(i, rotate)
        
        if p_first == (Point{}) {
            p_first, p_last = p, p
            b.WriteString(fmt.Sprintf("<circle id=\"%f\" cx=\"%f\" cy=\"%f\"  r=\"5\" fill=\"%s\" />\n", i, p.X, p.Y, colors[ccounter]))
            counter++
            continue
        }
        
        total += (math.Sqrt(math.Pow(p.X-p_last.X, 2) + math.Pow(p.Y-p_last.Y, 2))) / 2
        p_last = p
        
        if total >= (float64(counter) * toothArc) {
            b.WriteString(fmt.Sprintf("<circle id=\"%f-%d\" cx=\"%.14f\" cy=\"%.14f\" r=\"5\" fill=\"%s\" />\n", rotate, counter, p.X, p.Y, colors[ccounter]))
            counter++
        }
        if counter == teeth {
            break
        }
    }
ccounter++    
}    
    
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

    err := writeSvg(b.Bytes())
    if err != nil {
        log.Fatal(err)
    }

}
func printDot(b *bytes.Buffer, p Point) {
    b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"1\" />\n", p.X, p.Y))
}
