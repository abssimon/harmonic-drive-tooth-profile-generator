package main

type STooth struct {
    C1   *Circle
    C2   *Circle
    C3   *Circle
    C4   *Circle
    Tan1 Line
    Tan2 Line
}

func (t STooth) rotate(a float64) {
    t.C1.Rotate(a)
    t.C2.Rotate(a)
    t.C3.Rotate(a)
    t.C4.Rotate(a)
}

func (t STooth) scale() {

}

func (t STooth) moveX() {

}

func (t STooth) moveY() {

}

func newSTooth(tipCenter Point, tipRadius float64, tipLimit Point, bottomCenter Point, bottomRadius float64, bottomLimit Point) STooth {

    tooth := STooth{}
    tooth.C1 = newCircle(-tipCenter.X, tipCenter.Y, bottomRadius)
    tooth.C2 = newCircle(bottomCenter.X, bottomCenter.Y, tipRadius)
    tooth.C3 = newCircle(-bottomCenter.X, bottomCenter.Y, tipRadius)
    tooth.C4 = newCircle(tipCenter.X, tipCenter.Y, bottomRadius)

    // C1 - from BottomLimit to Tagent
    tooth.Tan1, _ = tooth.C1.InnerTangentTo(tooth.C2)
    tooth.C1.StartPoint = &Point{tooth.C1.X + bottomLimit.X, tooth.C1.Y - bottomLimit.Y}
    tooth.C1.StopPoint = tooth.Tan1.p2

    // C2 - from Tangent to Toplimit...
    tooth.C2.StartPoint = &Point{tooth.C2.X - tipLimit.X, tooth.C2.Y + tipLimit.Y}
    tooth.C2.StopPoint = tooth.Tan1.p1

    _, tooth.Tan2 = tooth.C3.InnerTangentTo(tooth.C4)
    tooth.C3.StartPoint = tooth.Tan2.p2
    tooth.C3.StopPoint = &Point{tooth.C3.X + tipLimit.X, tooth.C3.Y + tipLimit.Y}

    tooth.C4.StartPoint = tooth.Tan2.p1
    tooth.C4.StopPoint = &Point{tooth.C4.X - bottomLimit.X, tooth.C4.Y - bottomLimit.Y}

    return tooth
}