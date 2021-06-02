package main

import (
    "fmt"
    "math"
    "sort"
)


// See https://www.npmjs.com/package/interp1

type Zip struct {
    X, V float64
}

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

// Interpolates values linearly in one dimension.
func interp1(x, v, xqs []float64) ([]float64, error) {
    if len(x) != len(v) {
        return nil, fmt.Errorf("Arrays of sample points xs and corresponding values vs have to have equal length.: %d vs. %d\n", len(x), len(v))
    }

    // combine x and v
    p := []Zip{} // make len()
    for i := range x {
        p = append(p, Zip{x[i], v[i]})
    }

    // sort asc
    sort.Slice(p, func(i, j int) bool {
        return p[i].X < p[j].X
    })

    // check for double x values
    for i, v := range p[1:] {
        if p[i].X == v.X {
            return nil, fmt.Errorf("Two sample points have equal value %f. This is not allowed.", v.X)
        }
    }

    // split
    sortedX, sortedV := []float64{}, []float64{}
    for _, v := range p {
        sortedX = append(sortedX, v.X)
        sortedV = append(sortedV, v.V)
    }

    // interpolate
    r := []float64{}
    for _, xq := range xqs {
        // Determine index of range of query value.
        index := binaryFindIndex(sortedX, xq)

        // Check if value lies in interpolation range.
        if index == -1.0 {
            return nil, fmt.Errorf("Query value %f lies outside of range. Extrapolation is not supported.", xq)
        }

        r = append(r, interpolate(sortedV, index))
    }

    return r, nil
}