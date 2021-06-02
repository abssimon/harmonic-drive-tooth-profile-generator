package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
)

func main() {
	// Todo
	// - Intersect zweier Kreise Punkte berechnen
	// - ggf. runden
	// - Modulo (von Zahnmitte zu Zahnmitte) - alles von Modulo aus, Abstaende etc.
	// - Teethheight?

	// In Kreisform:
	// - Die Anzahl der Zähne
	// - 360° / 180 Zähne = 180 Unterteilungen
	//

	var b bytes.Buffer
	b.WriteString("<svg width=\"1200\" height=\"1000\" xmlns=\"http://www.w3.org/2000/svg\">")
	//b.WriteString("<line x1=\"368\" y1=\"0\" x2=\"368\" y2=\"1000\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />")
	//b.WriteString("<line x1=\"786\" y1=\"0\" x2=\"786\" y2=\"1000\" style=\"stroke:rgb(255,0,0);stroke-width:0.5\" />")
	b.WriteString("<g transform=\"translate(0,1000)\"><g transform=\"scale(1,-1)\"><g stroke=\"black\" stroke-width=\"0.5\" fill=\"none\">")

	circle1 := Circle{
		Point:  Point{-140.0, 463.0},
		Radius: 266.0,
	}
	circle2 := Circle{
		Point:  Point{388.0, 723.0},
		Radius: 207.0,
	}
	circle3 := Circle{
		Point:  Point{338.0, 723.0},
		Radius: 207.0,
	}
	circle4 := Circle{
		Point:  Point{866.0, 463.0},
		Radius: 266.0,
	}
	circle5 := Circle{
		Point:  Point{697.0, 463.0},
		Radius: 266.0,
	}
	/*
		circle5 := Circle{
			Point:  Point{786.0, 330.0},
			Radius: 100.0,
		}
	*/

	l1, l2 := circle1.InnerTangentTo(circle2)
	b.WriteString("<path d=\"M")
	points := circle1.CoordinatesBetween(Point{-90.0, 400.0}, l1.p2)
	first := true
	for _, point := range points {
		b.WriteString(fmt.Sprintf("%.14f %.14f ", point.X, point.Y))
		if first {
			first = false
			b.WriteString("L")
		}
	}

	points = circle2.CoordinatesBetween(Point{379.0, 743.0}, l1.p1)
	for i := len(points) - 1; i >= 0; i-- {
		b.WriteString(fmt.Sprintf("%.14f %.14f ", points[i].X, points[i].Y))
	}

	_, l2 = circle3.InnerTangentTo(circle4)
	points = circle3.CoordinatesBetween(l2.p2, Point{347.0, 743.0})
	for i := len(points) - 1; i >= 0; i-- {
		b.WriteString(fmt.Sprintf("%.14f %.14f ", points[i].X, points[i].Y))
	}

	//_, l2b := circle5.OuterTangentTo(circle4)

	points = circle4.CoordinatesBetween(l2.p1, Point{781.0, 211.0})
	for _, point := range points {
		b.WriteString(fmt.Sprintf("%.14f %.14f ", point.X, point.Y))
	}
	points = circle5.CoordinatesBetween(Point{781.0, 211.0}, Point{776.0, 463.0})
	for _, point := range points {
		b.WriteString(fmt.Sprintf("%.14f %.14f ", point.X, point.Y))
	}

	/*
		points = circle5.CoordinatesBetween(l2b.p1, Point{796.0, 330.0})
		for _, point := range points {
			b.WriteString(fmt.Sprintf("%.14f %.14f ", point.X, point.Y))
		}
	*/
	b.WriteString("\"/>\n")

	// printDot(&b, Point{781.0, 211.0}) Intersectpoint

	printDot(&b, circle1.Point)
	printDot(&b, circle2.Point)
	printDot(&b, circle3.Point)
	printDot(&b, circle4.Point)
	printDot(&b, circle5.Point)
	//printDot(&b, l2b.p1)
	//printDot(&b, l2b.p2)
	b.WriteString("</g></g></g></svg>")

	err := writeSvg(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}

}
