package main


func diff(p []Point) []Point {
    n := make([]Point, len(p))
    for i, v := range p[1:] {
        n[i] = Point{v.X - p[i].X, v.Y - p[i].Y}
    }
    return n
}
