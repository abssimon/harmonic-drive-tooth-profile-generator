%

#1 = 1 (assign parameter #1 the value of 1)
o101 while [#1 LT 20]

G01 X3 Y2 Z2.5 A180 F100
G01 X-3 Y-2 Z-2.5 A-180 F100
G01 X0 Y0 Z0 A0 F100

#1 = [#1+1] (increment the test counter)
o101 endwhile


%
