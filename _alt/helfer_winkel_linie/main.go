package main

import (
    "fmt"
    "math"
)

// degree to radian
func dToR(deg float64) float64 {
    return deg * (math.Pi / 180.0)
}

// radian to degree
func rToD(rad float64) float64 {
    return rad * (180.0 / math.Pi)
}
// https://www.massmatics.de/merkzettel/#!714:Winkel_zwischen_Geraden_&_x-Achse
func main() { 
    // <circle xmlns="http://www.w3.org/2000/svg" id="205" fill="black" cx="525.71366898563929" cy="276.44912644106262" r="0.8"/>
    // <circle xmlns="http://www.w3.org/2000/svg" id="204" fill="black" cx="525.71366898563917" cy="288.74587135086523" r="0.8"/>
/*
    cx1 := 525.71366898563929
    cy1 := 276.44912644106262
    
    cx2 := 536.56728476470539
    cy2 := 282.22917487487030
*/
    cx1 := 534.32302306239023
    cy1 := 285.82171758598162
    
    cx2 := 540.79312013617027
    cy2 := 289.26734116609913

 //   <circle xmlns="http://www.w3.org/2000/svg" id="399" fill="black" cx="534.32302306239023" cy="285.82171758598162" r="0.8"/>
 //   <circle xmlns="http://www.w3.org/2000/svg" id="401" fill="black" cx="540.79312013617027" cy="289.26734116609913" r="0.8"/>
    // <circle xmlns="http://www.w3.org/2000/svg" id="205" fill="black" cx="525.71366898563929" cy="276.44912644106262" r="0.8"/>
    // <circle xmlns="http://www.w3.org/2000/svg" id="206" fill="black" cx="536.56728476470539" cy="282.22917487487030" r="0.8"/>
/*
    cx2 = 969.14817422635542
    cy2 = 287.89993441560125
    
    cx1 = 969.16665743140413
    cy1 = 295.23029159456138   
*/    
    // atan((y2- y1) / (x2 - x1))
    fmt.Printf("%.14f\n", rToD(math.Atan((cy2 - cy1) / (cx2 - cx1))))
    //fmt.Println(math.Pi/2.0)
    

}
