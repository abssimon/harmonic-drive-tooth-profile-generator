package main

import (
    "fmt"
    "math"
    "sort"
)

// https://github.com/cpmech/gosl/tree/main/utl
func LinSpace(start, stop float64, num int) (res []float64) {
    if num <= 0 {
        return []float64{}
    }
    if num == 1 {
        return []float64{start}
    }
    step := (stop - start) / float64(num-1)
    res = make([]float64, num)
    res[0] = start
    for i := 1; i < num; i++ {
        res[i] = start + float64(i)*step
    }
    res[num-1] = stop
    return
}

func CumSum(s []float64) (dst []float64) {
    if len(s) == 0 {
        return []float64{}
    }
    dst = make([]float64, len(s))
    dst[0] = s[0]
    for i, v := range s[1:] {
        dst[i+1] = dst[i] + v
    }
    return 
}

func Diff(p []Point) (n []Point) {
    if len(p) == 0 {
        return []Point{}
    }
    n = make([]Point, len(p)-1)
    for i, v := range p[1:] {
        n[i] = Point{v.X - p[i].X, v.Y - p[i].Y}
    }
    return 
}

// See https://www.npmjs.com/package/interp1

// Finds the index of range in which a query value is included in a sorted
// array with binary search.
func binaryFindIndex(xs []float64, xq float64) float64 {
    // Special case of only one element in array.
    if len(xs) == 1 && xs[0] == xq {
        return 0.0
    }

    // Determine bounds.
    lower := 0
    upper := len(xs) - 1

    // Find index of range.
    for lower < upper {
        // Determine test range.
        mid := math.Floor(float64(lower+upper) / 2.0)
        prev := xs[int(mid)]
        next := xs[int(mid)+1]

        if xq < prev {
            // Query value is below range.
            upper = int(mid)
        } else if xq > next {
            // Query value is above range.
            lower = int(mid) + 1
        } else {
            // Query value is in range.
            return mid + (xq-prev)/(next-prev)
        }
    }

    // Range not found.
    return -1.0
}

// Interpolates a value linear.
func interpolate(vs []float64, index float64) float64 {
    prev := math.Floor(index)
    next := math.Ceil(index)
    lambda := index - prev
    return (1-lambda)*vs[int(prev)] + lambda*vs[int(next)]
}

type zip struct {
    X, V float64
}

// Interpolates values linearly in one dimension.
func interp1(x, v, xqs []float64) ([]float64, error) {
    if len(x) != len(v) {
        return nil, fmt.Errorf("Arrays of sample points xs and corresponding values vs have to have equal length.: %d vs. %d\n", len(x), len(v))
    }

    // sort x and v together
    p := make([]zip, len(x)) 
    for i := range x {
        p[i] = zip{x[i], v[i]}
    }
    sort.Slice(p, func(i, j int) bool {
        return p[i].X < p[j].X
    })

    // check for double x values
    for i, v := range p[1:] {
        if p[i].X == v.X {
            return nil, fmt.Errorf("Two sample points have equal value %f. This is not allowed.", v.X)
        }
    }

    // unzip
    sortedX := make([]float64, len(p))
    sortedV := make([]float64, len(p))
    for i, v := range p {
        sortedX[i] = v.X
        sortedV[i] = v.V
    }

    // interpolate
    r := make([]float64, len(xqs))
    for i, xq := range xqs {
        // Determine index of range of query value.
        index := binaryFindIndex(sortedX, xq)

        // Check if value lies in interpolation range.
        if index == -1.0 {
            return nil, fmt.Errorf("Query value %f lies outside of range. Extrapolation is not supported.", xq)
        }

        r[i] = interpolate(sortedV, index)
    }

    return r, nil
}
