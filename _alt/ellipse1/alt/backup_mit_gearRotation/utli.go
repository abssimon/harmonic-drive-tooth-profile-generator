package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

const intScaleFactor = 10000000000000

// float64 to int for clipper
func scaleUp(value float64) int {
	return int(value * intScaleFactor)
}

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

func printDot(b *bytes.Buffer, p Point, id string) {
	b.WriteString(fmt.Sprintf("<circle id=\"%s\" cx=\"%.14f\" cy=\"%.14f\"  r=\"1\" />\n", id, p.X, p.Y))
}
