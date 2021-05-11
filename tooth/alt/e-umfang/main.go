package main

import(
    "fmt"
    "math"
)

func main() {
    modul := 0.5 //cm

    // Rigid Gear
    rTeeth := 102
    
    rDiameter := modul * float64(rTeeth) // circle trough the middle of the teeth
    rCircum  := 2 * math.Pi * (rDiameter/2)
  
    fmt.Println("Rigid Diameter:", rDiameter, "mm")
    fmt.Println("Rigid Circumference :", rCircum, "mm")

    // Flex Gear
    fTeeth := 100
    
    fDiameter := modul * float64(fTeeth) // circle trough the middle of the teeth
    fCircum  := 2 * math.Pi * (fDiameter/2)
  
    fmt.Println("Flex Circle Diameter:", fDiameter, "mm")
    fmt.Println("Flex Circle Circumference:", fCircum, "mm")

    // Eine Elipse finden mit der Hoehe des Kreises und dem Umfang des vom Flex Gear

    // umfang

    a := 4.0 // cm radius
    b := 3.5 // cm radius

    t := (a - b) / (a + b)
    umfang := (a + b) * math.Pi * (1.0 + 3.0 * t * t / (10.0 + math.Sqrt(4.0 - 3.0 * t * t)))
    
    fmt.Println(umfang) // 23.588132121002868
 
}