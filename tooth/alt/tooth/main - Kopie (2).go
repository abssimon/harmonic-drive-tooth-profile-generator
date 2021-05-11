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

	middle := 300.0
	left := 200.0

	var b bytes.Buffer
	b.WriteString("<svg width=\"700\" height=\"700\" xmlns=\"http://www.w3.org/2000/svg\">")
	b.WriteString("<line x1=\"0\" y1=\"300\" x2=\"700\" y2=\"300\" style=\"stroke:rgb(255,0,0);stroke-width:2\" />")
	b.WriteString("<g transform=\"translate(0,700)\"><g transform=\"scale(1,-1)\"><g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\"><circle cx=\"10\" cy=\"10\"  r=\"1\" />")

	radiusTop := 225.0
	topCenterX := -80.0
	topCenterY := -60.0

	for i := 0.0; i <= radiusTop; i += 5.0 {
		b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", i+topCenterX+left, getY(radiusTop, i, -1)+middle+topCenterY))
	}

	radiusBottom := 245.0
	bottomCenterX := 430.0
	bottomCenterY := 73.0

	for i := 0.0; i >= radiusBottom*-1; i -= 5.0 {
		b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", i+bottomCenterX+left, getY(radiusBottom, i, 1)+middle+bottomCenterY))
	}

	/** Tangente **/

	// https://www.weltderfertigung.de/suchen/lernen/mathematik/beruehrpunktberechnung-tangente-an-zwei-kreisen.php

	// Abstand der Mittelpunkte
	b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" style=\"stroke:rgb(255,0,0);stroke-width:2\" />\n", topCenterX+left, middle+topCenterY))
	b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" style=\"stroke:rgb(255,0,0);stroke-width:2\" />\n", bottomCenterX+left, middle+bottomCenterY))
	c1 := math.Sqrt(math.Pow(middle+topCenterY, 2) + math.Pow(math.Abs(topCenterX)+bottomCenterX, 2))
	fmt.Println("Abstand der Mittelpunkte", c1)

	// Verhaeltnis der beiden Kreise
	spx := radiusBottom / radiusTop
	fmt.Println("spx", spx)
	lVerhaeltnis := (100.0 / (spx + 1.0)) * spx
	fmt.Println("spx verhaeltnis", lVerhaeltnis)

	// Anteilige Laenge
	l1 := c1 / 100.0 * lVerhaeltnis //
	fmt.Println("l1", l1)
	l2 := c1 - l1
	fmt.Println("l2", l2)

	b.WriteString(fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" />\n", topCenterX+left, middle+topCenterY, bottomCenterX+left, middle+bottomCenterY))

	// Winkel

	beta2 := rToD(math.Asin(radiusBottom / l1)) // rToD((radiusBottom * math.Sin(dToR(90))) / l1)

	gamma2 := 180.0 - (90 + beta2)

	alpha2 := rToD(math.Tan((middle + topCenterY) / (math.Abs(topCenterX) + bottomCenterX)))

	alpha5 := gamma2 - alpha2
	fmt.Println("alpha5", alpha5)

	// ---

	// TODO weltderfertigung Beispiel 1zu1 nachbauen

	a5 := math.Sin(dToR(alpha5)) * radiusBottom
	b5 := math.Cos(dToR(alpha5)) * radiusBottom
	fmt.Println(a5)
	fmt.Println(b5)

	// T5 Punkt
	t5x := bottomCenterX - a5
	t5y := bottomCenterY + b5

	fmt.Println("t5x", t5x)
	fmt.Println("t5y", middle+t5y)

	b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" style=\"stroke:rgb(255,0,0);stroke-width:2\" />\n", t5x+left, middle+t5y))

	/*

	   alpha := rToD(math.Asin(radiusBottom/l1)) // rToD((radiusBottom * math.Sin(dToR(90))) / l1)
	       fmt.Println("alpha", alpha)

	       gamma := 180.0 - (90 + alpha)
	       fmt.Println("gamma", gamma)

	       alpha2 := rToD(math.Tan((middle + topCenterY) / (math.Abs(topCenterX) + bottomCenterX)))
	       fmt.Println("alpha2", alpha2)

	       alpha3 := 90 - (alpha2 + gamma)
	       fmt.Println("alpha3", alpha3)

	       //

	       a3 := math.Sin(dToR(alpha3)) * radiusBottom
	       b3 := math.Cos(dToR(alpha3)) * radiusBottom
	       fmt.Println(a3)
	       fmt.Println(b3)

	       // T5 Punkt
	       t5x := bottomCenterX-a3
	       t5y := bottomCenterY+b3

	       fmt.Println("t5x", t5x)
	       fmt.Println("t5y", middle+t5y)

	       b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" style=\"stroke:rgb(255,0,0);stroke-width:2\" />\n", t5x+left, middle+t5y))


	*/

	// --------------

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
