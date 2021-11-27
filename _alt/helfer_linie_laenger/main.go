package main

import (
    "fmt"
    "math"
)

func main() {
    // <circle xmlns="http://www.w3.org/2000/svg" id="1" fill="black" cx="1185.71161726138143" cy="650.00000000000000" r="0.8"/>
    // <circle xmlns="http://www.w3.org/2000/svg" id="0" fill="black" cx="1175.16918593522178" cy="643.67014398979245" r="0.8"/>
   
    cx2 := 1185.71161726138143
    cy2 := 650.00000000000000
    
    cx1 := 1175.16918593522178
    cy1 := 643.67014398979245
     
    // Nach innen
    fmt.Printf("<line stroke='black' x1='%f' y1='%f' x2='%f' y2='%f' />\n", cx2, cy2, cx1 + (math.Abs(cx1 - cx2) * 7.0), cy1 + (math.Abs(cy1 - cy2) * 7.0))
    
    // Nach aussen
    fmt.Printf("<line stroke='black' x1='%f' y1='%f' x2='%f' y2='%f' />\n", cx1, cy1, cx2 - (math.Abs(cx1 - cx2) * 7.0), cy2 - (math.Abs(cy1 - cy2) * 7.0))
     
    // <circle xmlns="http://www.w3.org/2000/svg" id="1" fill="black" cx="1185.71161726138143" cy="650.00000000000000" r="0.8"/>
    // <circle xmlns="http://www.w3.org/2000/svg" id="2" fill="black" cx="1175.16918593522178" cy="656.32985601020755" r="0.8"/>
    
    cx2 = 1185.71161726138143
    cy2 = 650.00000000000000
    
    cx1 = 1175.16918593522178
    cy1 = 656.32985601020755
    
    // Nach innen
    fmt.Printf("<line stroke='black' x1='%f' y1='%f' x2='%f' y2='%f' />\n", cx1, cy1, cx1 - (math.Abs(cx1 - cx2) * 6.0), cy1 + (math.Abs(cy1 - cy2) * 6.0))
    
    // Nach aussen
    fmt.Printf("<line stroke='black' x1='%f' y1='%f' x2='%f' y2='%f' />\n", cx2, cy2, cx2 + (math.Abs(cx1 - cx2) * 6.0), cy2 - (math.Abs(cy1 - cy2) * 6.0))
}
