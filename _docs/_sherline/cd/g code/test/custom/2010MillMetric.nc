%

#1 = 1 (assign parameter #1 the value of 1)
o101 while [#1 LT 5]
G01 X4 Y3 Z3.25 A180 F100
G01 X-4 Y-3 Z-3.25 A-180 F100
G01 X0 Y0 Z0 A0 F100
#1 = [#1+1] (increment the test counter)
o101 endwhile


%

