%

#1 = 1 (assign parameter #1 the value of 1)
o101 while [#1 LT 5]
G01 X175 Z20 F1500
G01 X-175 Z-20 F1500
G01 X0 Z0 F1500
#1 = [#1+1] (increment the test counter)
o101 endwhile


%






