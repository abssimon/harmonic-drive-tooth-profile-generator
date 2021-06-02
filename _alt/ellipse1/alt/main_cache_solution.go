package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	//"io/ioutil"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func angleCache() (map[string]float64, error) {
	f, err := os.Open("angles.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	c := make(map[string]float64)
	r := csv.NewReader(f)
	r.Comma = ','
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		s, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			return nil, err
		}
		c[line[0]] = s
	}
	return c, nil
}

func main() {

	ac, err := angleCache()
	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer
	//var w bytes.Buffer
	b.WriteString("<svg width=\"1200\" height=\"1000\" xmlns=\"http://www.w3.org/2000/svg\">")
	b.WriteString("<g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")

	e := Ellipse{Point{350.0, 280.0}, 250.0, 150.0} // px, py, h, w
	printDot(&b, e.Point, "")

	teeth := 30
	toothArc := e.Circumference() / float64(teeth)

	//for rotate := 63.0; rotate <= 117.0; rotate += 63.0 {

	p, pLast := Point{}, Point{}
	first := true
	total := 0.0
	counter := 0

	//rotate := 45.0
	for i := 0.0; i <= 360.0; i += 0.001 {

		p = e.PointInAngle(i)
		//w.WriteString(fmt.Sprintf("%.6f,%.6f\n", math.Atan2(p.Y, p.X), dToR(i)))

		if first {
			first = false
			pLast = p

			//j:= math.Atan2(p.Y, p.X)
			tan := e.Tangent(i)
			b.WriteString(fmt.Sprintf("<circle id=\"%f-%d\" cx=\"%f\" cy=\"%f\"  r=\"1\" fill=\"red\" /><text x=\"%f\" y=\"%f\" font-size=\"0.7em\" fill=\"black\" > t %.2f</text>\n",
				i, counter, p.X+e.X, p.Y+e.Y, p.X+e.X+5, p.Y+e.Y+5, rToD(tan)))
			counter++
			continue
		}

		total += p.DistanceTo(&pLast) / 2
		pLast = p
		if total >= (float64(counter) * toothArc) {

			check := fmt.Sprintf("%.6f\n", math.Atan2(p.Y, p.X))
			if _, ok := ac[check]; ok {
				fmt.Println("Hit")
			} else {
				fmt.Println("nope", check, "from atan2:", p, "// i:", dToR(i))
			}

			tan := e.Tangent(i)
			b.WriteString(fmt.Sprintf("<circle id=\"%f-%d\" cx=\"%f\" cy=\"%f\"  r=\"1\"/><text x=\"%f\" y=\"%f\" font-size=\"0.7em\" fill=\"black\" > t %.2f</text>\n",
				i, counter, p.X+e.X, p.Y+e.Y, p.X+e.X+5, p.Y+e.Y+5, rToD(tan)))
			counter++
		}

		if counter == teeth {
			fmt.Println("teeth complete")
			break
		}
	}
	//}

	b.WriteString("</g></svg>")

	err = writeSvg(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	/*
	   err = ioutil.WriteFile("angles.txt", w.Bytes(), 0644)
	   if err != nil {
	       log.Fatal(err)
	   }
	*/
}
