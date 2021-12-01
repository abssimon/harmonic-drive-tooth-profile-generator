package main

import (
	"flag"
	"fmt"
	"log"
	"math"
)

// Teeth mashing is pretty hard. The only working solution is found here
// https://commons.wikimedia.org/wiki/File:HarmonicDriveAni.gif
func toothPosition(mode string, ellipseCoordinates []Point, nPoints float64, offset float64) []Point {

	// small coordinate changes
	sum := make([]float64, len(ellipseCoordinates))
	for i, v := range Diff(ellipseCoordinates) {
		sum[i+1] = math.Hypot(v.X, v.Y)
	}
	cumLen := CumSum(sum)
	circumference := cumLen[len(cumLen)-1]
	if mode == "both" {
		fmt.Println("flex_circumference is ", circumference/math.Pi)
	}

	// location on curcumference, with offset
	pos := LinSpace(0.0, 1.0, int(nPoints)+1)
	pos = pos[:len(pos)-1]
	for i, v := range pos {
		pos[i] = math.Mod(v+offset, 1) * circumference
	}

	coordX := make([]float64, len(ellipseCoordinates))
	coordY := make([]float64, len(ellipseCoordinates))
	for i, v := range ellipseCoordinates {
		coordX[i] = v.X
		coordY[i] = v.Y
	}

	// predict X for pos
	x, err := interp1(cumLen, coordX, pos)
	if err != nil {
		panic(err)
	}
	y, err := interp1(cumLen, coordY, pos)
	if err != nil {
		panic(err)
	}

	r := make([]Point, len(x))
	for i, v := range x {
		r[i] = Point{v, y[i]}
	}

	return r
}

func main() {

	m := flag.String("mode", "", "flag mode (both, flex, rigid, ani) is optional")
	j := flag.String("config", "", "flag config (*.json) is optional")
	flag.Parse()

	cf := "conf"
	mode := "both"
	if *j != "" {
		cf = *j
	}
	if *m == "flex" {
		mode = "flex"
	}
	if *m == "rigid" {
		mode = "rigid"
	}
	if *m == "ani" {
		mode = "ani"
	}

	conf, err := getConfig(cf + ".json")
	if err != nil {
		log.Fatalf("Config konnte nicht gelesen werden: %v\n", err)
	}

	rigidTheets := conf.RigidTheets
	flexTheets := rigidTheets - 2.0

	gears := []Gear{}
	gear := Gear{}

	if mode == "flex" {

		// draw undeformed flexgear for production
		c2 := &Circle{
			&Point{0.0, 0.0},
			conf.FlexCircumference / 2, // get this mode=both
			0.0,
			math.Pi * 2.0,
		}
		angle := LinSpace(0.0, math.Pi*2.0, int(flexTheets)+1)
		angle = angle[:len(angle)-1]

		for _, a := range angle {
			gear.Tooths = append(gear.Tooths, NewTooth(conf, -math.Pi/2.0+a, 1.56, c2.PointInAngle(a)))
		}
		gears = append(gears, gear)

	} else {

		if mode != "rigid" {

			// draw flexgear
			nFrames := 1
			if mode == "ani" {
				nFrames = 100
			}
			frameAngles := LinSpace(0.0, -math.Pi, nFrames+1)
			frameAngles = frameAngles[:len(frameAngles)-1]

			// sample points for prediction
			e := Ellipse{Point{0, 0}, conf.DiameterH, conf.DiameterV}
			ellipseCoordinates := e.Coordinates(1000)

			for _, angleWaveGen := range frameAngles {

				// position prediction on ellipse
				angleFlexTeeth := angleWaveGen * (flexTheets - rigidTheets) / flexTheets
				offset := (-angleWaveGen + angleFlexTeeth) / 2.0 / math.Pi
				pos := toothPosition(mode, ellipseCoordinates, flexTheets, offset)

				gear = Gear{}
				for _, p := range pos {
					gear.Tooths = append(gear.Tooths, NewTooth(conf, math.Pi+e.TangentByPoint(p), conf.BottomStopFlex, p))
				}
				gear.Angle = angleWaveGen
				gears = append(gears, gear)
			}
		}

		// draw rigid gear
		c := &Circle{&Point{0.0, 0.0}, conf.DiameterH / 2, 0.0, math.Pi * 2.0}
		angle := LinSpace(0.0, math.Pi*2.0, int(rigidTheets)+1)
		angle = angle[:len(angle)-1]

		gear = Gear{}
		for _, a := range angle {
			gear.Tooths = append(gear.Tooths, NewTooth(conf, -math.Pi/2.0+a, conf.BottomStopRigid, c.PointInAngle(a)))
		}
		gears = append(gears, gear)

	}

	svg(mode, gears)
}
