%

#1 = 1 (assign parameter #1 the value of 1)
o101 while [#1 LT 4]
G01 X95 Y50 Z75 A180 F750
G01 X-95 Y-50 Z-75 A-180 F750
G01 X0 Y0 Z0 A0 F750
#1 = [#1+1] (increment the test counter)
o101 endwhile


%

