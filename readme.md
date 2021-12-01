# Strain Wave/Harmonic Drive - Gear Generator

With this generator, you can develop and simulate tooth profiles for a harmonic drive. These profiles can be imported into CAD programs for further use.


### Examples:

- [50 Tooth Profile](https://retwin.com/github/test1.html)
- [30 Tooth Profile](https://retwin.com/github/test2.html)
- [50 Tooth Profile (small)](https://retwin.com/github/test3.html) (a slotted plastic version of this profile can be seen [here](https://www.youtube.com/watch?v=iLJkPBIP0VU))
 
### General Usage

The programm can be called with two flags. mode and config, for example

    go run . -mode=ani -config=conf_50_small

##### -mode

- mode=both - is default and will generate a flex and a rigid gear (in "test.svg") 
- mode=ani - will generate the gear animation (in "test.html")
- mode=flex - will generate a round "undeformed" flexgear (in "test.svg")
- mode=rigid - will generate the a rigid gear (in "test.svg")

##### -config

The default file is "conf.json", but you can specifiy here a different one.

### Gear Definition 

A single gear tooth is defined by two circles (which are mirrored). A tip and a bottom circle with a center point (x, y) and a radius. The circles are connected with a tangent at one end, and at the other end, they stop in a certain direction.

[![N|Solid](https://retwin.com/github/teeth_circles.jpg)](https://retwin.com/github/teeth_circles.jpg)

Referring to this, there are variables in the config file

	"tip_center_x": -0.0,   // both tip points are together, so a round tip...
	"tip_center_y": 1.265,
	"tip_radius": 0.506,
	"tip_stop": 1.56,       // = 180 degree, stop right over the center

	"bottom_center_x": 1.305,
	"bottom_center_y": 0.985,
	"bottom_radius": 0.506,
	"bottom_stop_flex": 1.56,
	"bottom_stop_rigid": 0.95, // in this case, the stop for the rigid gear is different than in flex

Scale factor will bring the teeth to the correct size

	"scale": 10.03
	
The number of teeths are defined by

	"rigid_theets": 102.0  // flex theets are always two less...
 	
The shape of the ellipse is defined by

	"diameter_h": 4.2,
	"diameter_v": 4.035,

For example, when you decrease the teeth height, you can increase diameter_v. So more teeths will have contact during rotation. You can check this in detail with --mode=both. Note, when you change the diameter, you also need to update "flex_circumference" 

	"flex_circumference": 4.117906474475928

When you run --mode=both, the circumference is always shown, copy the new value to your config file. The right value is very important for -mode=flex. 






