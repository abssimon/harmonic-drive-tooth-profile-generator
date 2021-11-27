%

#1 = 1 (assign parameter #1 the value of 1)
o101 while [#1 LT 20]
G01 X2.5 Z1 F100
G01 X-2.5 Z-1 F100
G01 X0 Z0 F100
#1 = [#1+1] (increment the test counter)
o101 endwhile


%





