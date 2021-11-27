%

#1 = 1 (assign parameter #1 the value of 1)
o101 while [#1 LT 5]
G01 X7 Z1 F100
G01 X-7 Z-1 F100
G01 X0 Z0 F100
#1 = [#1+1] (increment the test counter)
o101 endwhile


%





