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
	b.WriteString("<svg width=\"700\" height=\"700\" xmlns=\"http://www.w3.org/2000/svg\">")
	b.WriteString("<line x1=\"0\" y1=\"300\" x2=\"700\" y2=\"300\" style=\"stroke:rgb(255,0,0);stroke-width:2\" />")
	b.WriteString("<g transform=\"translate(0,700)\"><g transform=\"scale(1,-1)\"><g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\"><circle cx=\"60.000000\" cy=\"100.0\" r=\"1\"/>")

	radiusTop := 15.0
	topCenterX := 60.0
	topCenterY := 100.0

	for i := 0.0; i <= radiusTop; i += 2.0 {
		b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", i+topCenterX, getY(radiusTop, i, -1)+topCenterY))
	}

	radiusBottom := 50.0
	bottomCenterX := 130.0
	bottomCenterY := 40.0

	for i := 0.0; i >= radiusBottom*-1; i -= 2.0 {
		b.WriteString(fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\"  r=\"1\" />\n", i+bottomCenterX, getY(radiusBottom, i, 1)+bottomCenterY))
	}

	// Tangente
	// https://www.weltderfertigung.de/suchen/lernen/mathematik/beruehrpunktberechnung-tangente-an-zwei-kreisen.php

	c1 := math.Sqrt(math.Pow(topCenterY-bottomCenterY, 2) + math.Pow(bottomCenterX-topCenterX, 2))
	fmt.Println(topCenterY, bottomCenterY, bottomCenterX, topCenterX)
	v := radiusBottom / radiusTop
	l1 := c1 / 100 * (100.0 / (v + 1.0) * v)
	//l2 := c1 - l1

	// angles
	a0 := rToD(math.Asin(radiusBottom / l1))
	a1 := 180.0 - (a0 + 90.0)
	a2 := rToD(math.Asin((topCenterY - bottomCenterY) / c1))
	a3 := 90 - (a2 + a1)
	a4 := a1 - a2

	sa3 := math.Sin(dToR(a3)) * radiusBottom
	sb3 := math.Cos(dToR(a3)) * radiusBottom
	sa5 := math.Sin(dToR(a4)) * radiusBottom
	sb5 := math.Cos(dToR(a4)) * radiusBottom

	osa5 := math.Sin(dToR(a3)) * radiusTop
	osb5 := math.Cos(dToR(a3)) * radiusTop
	osa7 := math.Sin(dToR(a4)) * radiusTop
	osb7 := math.Cos(dToR(a4)) * radiusTop

	b.WriteString(fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" />\n", bottomCenterX-sa3, bottomCenterY+sb3, topCenterX+osa5, topCenterY-osb5))
	b.WriteString(fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" />\n", bottomCenterX-sb5, bottomCenterY-sa5, topCenterX+osb7, topCenterY+osa7))

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
