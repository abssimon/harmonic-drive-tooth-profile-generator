package main

// CumSum finds the cumulative sum of the first i elements in
// s and puts them in place into the ith element of the
// destination dst.
// It panics if the argument lengths do not match.
//
// At the return of the function, dst[i] = s[i] + s[i-1] + s[i-2] + ...
func CumSum(dst, s []float64) []float64 {
    if len(dst) != len(s) {
        panic("CumSum arg lenght must be the same")
    }
    if len(dst) == 0 {
        return dst
    }
    
    dst[0] = s[0]
    for i, v := range s[1:] {
        dst[i+1] = dst[i] + v
    }
    return dst
}