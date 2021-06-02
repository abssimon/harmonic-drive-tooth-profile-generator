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

	/*
	   ellipse := Ellipse{
	       Point{350.0, 280.0},
	       250.0,
	       150.0,
	   }
	*/
	w := 3.0
	h := 4.0

	bb := 20.0
	bb *= math.Pi / 180.0
	cb, sb := math.Cos(bb), math.Sin(bb)
	d := math.Atan2(-w*sb, h*cb)

	//first := true
	for t := 0.0; t <= 360.0; t += 3.0 { // .001 is a +-0.00218px precision

		printDot(&b, Point{
			(w*math.Cos(dToR(t)+d)*cb-h*math.Sin(dToR(t)+d)*sb)*20.0 + 300.0,
			(w*math.Cos(dToR(t)+d)*sb+h*math.Sin(dToR(t)+d)*cb)*20.0 + 300.0,
		}, fmt.Sprintf("%f:%f", t, d))

		d := math.Atan2(-ellipse.Width*math.Sin(dToR(rotate)), ellipse.Height*math.Cos(dToR(rotate)))
		fmt.Println(i, rToD(math.Atan2(ellipse.Height*math.Cos(dToR(i)+d), -ellipse.Width*math.Sin(dToR(i)+d))))

		/*
		   p := ellipse.PointInAbsoluteAngleRotated(i, rotate)

		   if first {
		       first = false
		       b.WriteString(fmt.Sprintf("<circle id=\"%f\" cx=\"%f\" cy=\"%f\"  r=\"1\" fill=\"red\" />\n", i, p.X, p.Y))


		       continue
		   }

		   b.WriteString(fmt.Sprintf("<circle id=\"%f\" cx=\"%.14f\" cy=\"%.14f\" r=\"1\" />\n", i, p.X, p.Y))
		*/
		//d := math.Atan2(-ellipse.Width * math.Sin(dToR(rotate)), ellipse.Height * math.Cos(dToR(rotate)))
		//fmt.Println(i, rToD(math.Atan2(ellipse.Height * math.Cos(dToR(i)+d), -ellipse.Width * math.Sin(dToR(i)+d))))
	}

	b.WriteString("</g></svg>")

	err := writeSvg(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}

}
