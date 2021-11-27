package main

import (
	"fmt"
	"math"
	"os"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) DistanceTo(q *Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func samplesEllipse(diameterH float64, diameterV float64, nPoints int) []Point {
	r := []Point{}
	max := 2.0 * math.Pi
	for i := 0.0; i <= max; i += max / float64(nPoints-1) { // hmmm, hier koennte einer weniger kommen, float ungenauigkeit...
		r = append(r, Point{math.Cos(i) * diameterH / 2.0, math.Sin(i) * diameterV / 2.0})
	}
	return r
}

func equidistantSamplesEllipse(diameterH float64, diameterV float64, nPoints float64, offset float64) {

	pathXY := samplesEllipse(diameterH, diameterV, 100)

	// distance between the points
	d := diff(pathXY)
	// pow 2
	for i, v := range d {
		d[i] = Point{math.Pow(v.X, 2), math.Pow(v.Y, 2)}
	}

	sum := make([]float64, len(d)+1)
	sum[0] = 0.0 // add the starting point
	for i, v := range d {
		sum[i+1] = math.Sqrt(v.X + v.Y) // +sqrt
	}

	cumLen := make([]float64, len(sum))
	cumLen = CumSum(cumLen, sum) // cumulative sum
	circumference := cumLen[len(cumLen)-1]

	finalStepLocs := make([]float64, 320)
	max := 1.0
	c := 0
	for i := 0.0; i < max; i += max / 320.0 {
		finalStepLocs[c] = (i + offset) * circumference
		c++
	}

	pathX := make([]float64, len(pathXY))
	pathY := make([]float64, len(pathXY))
	for i, v := range pathXY {
		pathX[i] = v.X
		pathY[i] = v.Y
	}

	x, err := interp1(cumLen, pathX, finalStepLocs)
	if err != nil {
		panic(err)
	}
	y, err := interp1(cumLen, pathY, finalStepLocs)
	if err != nil {
		panic(err)
	}

	for i, v := range x {
		fmt.Println(v, y[i])
	}

	// STIMMT !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! :=))
	// npoint war bei 100 - wieder anpassen.  original script
}

func main() {
	nTeethOutGear := 42.0
	nTeethFlex := nTeethOutGear - 2.0
	modul := 0.1

	frameAngles := []float64{0.0, -0.0314, -0.0628, -0.0942} //, -0.1257, -0.1571, -0.1885, -0.2199, -0.2513, -0.2827, -0.3142, -0.3456, -0.3770, -0.4084, -0.4398, -0.4712, -0.5027, -0.5341, -0.5655, -0.5969, -0.6283, -0.6597, -0.6912, -0.7226, -0.7540, -0.7854, -0.8168, -0.8482, -0.8796, -0.9111, -0.9425, -0.9739, -1.0053, -1.0367, -1.0681, -1.0996, -1.1310, -1.1624, -1.1938, -1.2252, -1.2566, -1.2881, -1.3195, -1.3509, -1.3823, -1.4137, -1.4451, -1.4765, -1.5080, -1.5394, -1.5708, -1.6022, -1.6336, -1.6650, -1.6965, -1.7279, -1.7593, -1.7907, -1.8221, -1.8535, -1.8850, -1.9164, -1.9478, -1.9792, -2.0106, -2.0420, -2.0735, -2.1049, -2.1363, -2.1677, -2.1991, -2.2305, -2.2619, -2.2934, -2.3248, -2.3562, -2.3876, -2.4190, -2.4504, -2.4819, -2.5133, -2.5447, -2.5761, -2.6075, -2.6389, -2.6704, -2.7018, -2.7332, -2.7646, -2.7960, -2.8274, -2.8588, -2.8903, -2.9217, -2.9531, -2.9845, -3.0159, -3.0473, -3.0788, -3.1102}

	for _, angleWaveGen := range frameAngles {

		angleFlexTeeth := angleWaveGen * (nTeethFlex - nTeethOutGear) / nTeethFlex
		effectiveDiameter := modul * nTeethOutGear
		deformedDiameter := effectiveDiameter * 0.9022
		offsetOnCircumference := (-angleWaveGen + angleFlexTeeth) / 2.0 / math.Pi

		equidistantSamplesEllipse(effectiveDiameter, deformedDiameter, nTeethFlex*8, offsetOnCircumference)

	}

}
