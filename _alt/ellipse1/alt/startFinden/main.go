package main

import (
	"fmt"
	"math"
)

func main() {
	e := Ellipse{Point{600.0, 500.0}, 400.0, 375.0} // px, py, h, w

	teeth := 100
	//toothWidth  := e.Circumference() / float64(teeth)
	gearRotation := 0.0

	counter := 0
	for rotate := 0.0; rotate <= 360.0; rotate += 90.0 {
		counter = 0

		// absolute angles and elipse angles differ, find the right ones
		start := 360.0 - rotate + gearRotation
		fmt.Println("\nInitial mit", start)
		p := Point{}
		for {
			p = e.PointByAngleRotated(start, rotate)
			counter++

			if rToD(math.Atan2(p.Y/e.Height, p.X/e.Width)) > gearRotation {
				//if p.Y > 0.0 {
				start -= 0.0001 //    0.1
			} else {
				start += 0.0001 //  0.1
			}
			//fmt.Println("   ",start, rToD(math.Atan2(p.Y/e.Height, p.X/e.Width)), ">", gearRotation)

			if math.Abs(rToD(math.Atan2(p.Y/e.Height, p.X/e.Width))-gearRotation) < 0.001 {
				//if math.Abs(p.Y) < 0.001 {
				//if counter > 20 {
				break
			}
		}

		teethDiff := 2.0
		ratio := float64(teeth) / teethDiff
		perDegree := 360.0 / ratio / 360.0
		gearRotation += (perDegree * 90.0) // 90 see loop step

		// "start", "counter", "gearRotation", "rToD(math.Atan2(p.Y/e.Height, p.X/e.Width))"

		fmt.Printf("= start: %f\t counter: %d \t gearRotation:%f \t Atan2:%f\n", start, counter, gearRotation, rToD(math.Atan2(p.Y/e.Height, p.X/e.Width)))

	}

}
