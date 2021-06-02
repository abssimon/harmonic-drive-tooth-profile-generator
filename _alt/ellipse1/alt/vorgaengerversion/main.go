package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
)

func main() {

	var b bytes.Buffer
	b.WriteString("<svg width=\"700\" height=\"700\" xmlns=\"http://www.w3.org/2000/svg\">")
	b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")

	ellipse := Ellipse{
		Point{350.0, 280.0},
		250.0,
		240.0,
	}
	teeth := 200 // (200 - 202)/200 = -0.01 oder 1/100
	toothArc := ellipse.Circumference() / float64(teeth)

	circle := newCircle(350.0, 280.0, 250.0)

	// bei einer umdrehung 1/100 weiter.
	// 3,6 Grad pro Umdrehung
	// oder 2 Zaehne
	// Welche strecke

	// Die Drehung der Ellipse
	movement := 0.0
	for rotate := 63.0; rotate <= 117.0; rotate += 9.0 {

		p, p_first, p_last := Point{}, Point{}, Point{}
		total := 0.0
		counter := 0

		// TODO: Zaehne in Rotation versetzen
		for i := 0.0; i <= 360.0; i += 0.0001 { // .001 is a +-0.00218px precision
			p = ellipse.PointInAbsoluteAngleRotated(i, rotate)

			if p_first == (Point{}) {
				p_first, p_last = p, p
				b.WriteString(fmt.Sprintf("<circle id=\"%f\" cx=\"%f\" cy=\"%f\"  r=\"1\" fill=\"red\" />\n", i, p.X, p.Y))

				counter++
				continue
			}

			dist := (math.Sqrt(math.Pow(p.X-p_last.X, 2) + math.Pow(p.Y-p_last.Y, 2))) / 2
			total += dist
			p_last = p

			if total >= (float64(counter)*toothArc)+movement {
				b.WriteString(fmt.Sprintf("<circle id=\"%f-%d\" cx=\"%.14f\" cy=\"%.14f\" r=\"1\" />\n", i, counter, p.X, p.Y))
				counter++
			}

			if counter == teeth {
				break
			}
		}

		cp := circle.PointInAngle(rotate - 90.0)
		b.WriteString(fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" stroke=\"green\" />\n", circle.X, circle.Y, cp.X, cp.Y))

		// first draw is plus zero
		// two teeth per revolution,
		movement += (2.0 / 360.0) * 9.0 * toothArc
		fmt.Println(movement)

	}

	//b.WriteString(fmt.Sprintf("<circle cx=\"%.14f\" cy=\"%.14f\"  r=\"%.14f\" />\n", ellipse.X, ellipse.Y, ellipse.Height))

	// rigid spline always follow...
	circle = newCircle(350.0, 280.0, 252.5)
	for i := 0.0; i <= 360.0; i += (360.0 / float64(teeth+2)) {
		p := circle.PointInAngle(i)
		b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\" fill=\"yellow\"  r=\"1\" />\n", p.X, p.Y))
	}
	printDot(&b, Point{350.0, 280.0})

	fmt.Println("", ellipse.Circumference())

	b.WriteString("</g></svg>")

	err := writeSvg(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}

}
