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
		150.0,
	}
	teeth := 200 // (200 - 202)/200 = -0.01 oder 1/100
	toothArc := ellipse.Circumference() / float64(teeth)

	// Die Drehung der Ellipse
	//for rotate := 63.0; rotate <= 117.0; rotate += 9.0 {
	rotate := 30.0

	p, pFirst, pLast := Point{}, Point{}, Point{}
	total := 0.0
	counter := 0

	// TODO: Zaehne in Rotation versetzen
	for j := 0; j <= 3600000; j++ { // .001 is a +-0.00218px precision

		i := float64(j) * 0.0001

		p = ellipse.PointInAbsoluteAngleRotated(i, rotate)

		if pFirst == (Point{}) {
			pFirst, pLast = p, p
			b.WriteString(fmt.Sprintf("<circle id=\"%f\" cx=\"%f\" cy=\"%f\"  r=\"1\" fill=\"red\" />\n", i, p.X, p.Y))
			counter++

			d := math.Atan2(-ellipse.Width*math.Sin(dToR(rotate)), ellipse.Height*math.Cos(dToR(rotate)))
			fmt.Println(counter, i, rToD(math.Atan2(ellipse.Height*math.Cos(dToR(i)+d), -ellipse.Width*math.Sin(dToR(i)+d))))
			continue
		}

		dist := (math.Sqrt(math.Pow(p.X-pLast.X, 2) + math.Pow(p.Y-pLast.Y, 2))) / 2
		total += dist
		pLast = p

		if total >= (float64(counter) * toothArc) {

			b.WriteString(fmt.Sprintf("<circle id=\"%f-%d\" cx=\"%.14f\" cy=\"%.14f\" r=\"1\" />\n", i, counter, p.X, p.Y))
			counter++

			d := math.Atan2(-ellipse.Width*math.Sin(dToR(rotate)), ellipse.Height*math.Cos(dToR(rotate)))
			fmt.Println(counter, i, rToD(math.Atan2(ellipse.Height*math.Cos(dToR(i)+d), -ellipse.Width*math.Sin(dToR(i)+d))))

		}

		if counter == teeth {
			fmt.Println("Exit")
			break
		}
	}

	//}

	printDot(&b, Point{350.0, 280.0})

	fmt.Println("", ellipse.Circumference())

	b.WriteString("</g></svg>")

	err := writeSvg(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}

}
